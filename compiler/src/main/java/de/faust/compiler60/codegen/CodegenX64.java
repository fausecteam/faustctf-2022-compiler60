package de.faust.compiler60.codegen;

import de.faust.compiler60.CompilationException;
import de.faust.compiler60.CompilerRandomization;
import de.faust.compiler60.ast.ComparisonExpression;
import de.faust.compiler60.ir.*;
import de.faust.compiler60.types.IntType;
import de.faust.compiler60.types.StringType;

import java.io.PrintWriter;
import java.io.StringWriter;
import java.util.*;

import static de.faust.compiler60.codegen.X64Operands.HardwareRegister.*;
import static de.faust.compiler60.codegen.X64Operands.OperandMode.*;

public class CodegenX64 extends IRVisitor {
  protected IRProgram program;

  protected StringBuffer stringBuffer;
  protected PrintWriter writer;

  protected X64Operands ops;

  protected Map<IRVariable, X64Operands.GlobalSymbol> globals;

  protected IRProcedure currentProcedure;
  // value to add to rsp to get to virtual rbp
  protected int currentSPToBPOffset;
  protected List<X64Operands.HardwareRegister> callerSaveRegs;

  protected X64Operands.HardwareRegister targetReg;
  protected X64Operands.HardwareRegister targetRegOrRax;

  public String generateAssembly(IRProgram program) {
    new StringOpsToFunctions().visit(program);
    new IRCodengenOptimization().visit(program);

    this.program = program;
    StringWriter sw = new StringWriter();
    stringBuffer = sw.getBuffer();
    writer = new PrintWriter(sw);
    writer.println(".intel_syntax noprefix");
    ops = new X64Operands(this);

    // add globals to bss
    writer.print("""
            .bss
            .align 8
            """);
    globals = new HashMap<>(program.globals.size());
    for (IRVariable global : program.globals) {
      writer.printf(".lcomm %s, %d\n", global.name, ops.getSize(global.type));
      globals.put(global, new X64Operands.GlobalSymbol(global.name));
    }
    writer.println();

    // visit functions
    writer.println(".text");
    addStringErrorHandlingTarget();
    super.visit(program);

    ops.constantPool.generateRodata(program, writer);

    checkOutputSize();
    stringBuffer = null;
    writer = null;
    ops = null;
    globals = null;
    return sw.toString();
  }

  @Override
  public void visit(IRProgram program) {
    throw new UnsupportedOperationException("Use generateAssembly()");
  }

  private void checkOutputSize() {
    if (stringBuffer.length() > BinaryBuilder.MAX_ASM_SIZE) {
      throw new CompilationException("Generated assembly too large [> " + BinaryBuilder.MAX_ASM_SIZE + " chars]");
    }
  }

  private void addStringErrorHandlingTarget() {
    // add exit(138) for illegal $strget and $strput
    visit(Label.ILLEGAL_INDEX);
    writer.println("mov rdi, 138");
    writer.println("call exit");
  }

  private void checkStringFunctionResult(Call call) {
    if ("$strget".equals(call.name) || "$strput".equals(call.name)) {
      // check return value - illegal on negative eax
      writer.println("test eax, eax");
      writer.printf("jl .%s\n", Label.ILLEGAL_INDEX.mnemonic());
    }
  }

  @Override
  public void visit(IRProcedure procedure) {
    writer.printf(".globl %s\n", procedure.name);
    writer.printf("%s=.\n", procedure.name);
    currentProcedure = procedure;
    ops.variableLocations = new HashMap<>(globals);

    callerSaveRegs = new ArrayList<>();
    List<X64Operands.HardwareRegister> calleeSaveRegs = new ArrayList<>();
    var availableRegisters = ops.getRegistersForFunction(procedure);

    // assign parameter locations
    int argOffset = 8; // saved return address
    List<IRVariable> parameters = procedure.parameters;
    for (int i = 0; i < parameters.size(); i++) {
      IRVariable parameter = parameters.get(i);
      if (i < CALLING_CONVENTION.length) {
        var reg = CALLING_CONVENTION[i];
        availableRegisters.remove(reg);
        ops.variableLocations.put(parameter, reg);
        callerSaveRegs.add(reg);
      } else {
        ops.variableLocations.put(parameter, new X64Operands.StackOffset(argOffset));
        argOffset += 8;
      }
    }

    // assign variable locations
    currentSPToBPOffset = 0;
    int stackFrameSize = 0;
    // first assign variables which might be in register
    for (IRVariable localVariable : procedure.localVariables) {
      if (localVariable.type != IntType.INSTANCE) {
        continue;
      }
      if (!availableRegisters.isEmpty()) {
        var reg = availableRegisters.removeFirst();
        boolean isCalleeSave = Arrays.asList(CALLEE_SAVE).contains(reg);
        if (isCalleeSave) {
          writer.printf("push %s\n", reg);
          // stack grows in negative direction from base pointer
          currentSPToBPOffset += 8;
          calleeSaveRegs.add(reg);
        } else {
          callerSaveRegs.add(reg);
        }
        ops.variableLocations.put(localVariable, reg);
      } else {
        int size = ops.getSize(localVariable.type);
        // stack grows in negative direction from base pointer
        currentSPToBPOffset += size;
        stackFrameSize += size;
        ops.variableLocations.put(localVariable, new X64Operands.StackOffset(-currentSPToBPOffset));
      }
    }
    // then assign arrays/strings which must be on stack
    for (IRVariable localVariable : procedure.localVariables) {
      if (localVariable.type == IntType.INSTANCE) {
        continue;
      }
      int size = ops.getSize(localVariable.type);
      // stack grows in negative direction from base pointer
      currentSPToBPOffset += size;
      stackFrameSize += size;
      ops.variableLocations.put(localVariable, new X64Operands.StackOffset(-currentSPToBPOffset));
    }

    if (stackFrameSize != 0)
      writer.printf("sub rsp, %d\n", stackFrameSize);

    // generate instructions
    visit(procedure.body);
    
    // mov result to RAX
    if (procedure.returnValue != null)
      ops.mov(RAX, procedure.returnValue);

    if (stackFrameSize != 0)
      writer.printf("add rsp, %d\n", stackFrameSize);
    currentSPToBPOffset -= stackFrameSize;

    for (int i = calleeSaveRegs.size() - 1; i >= 0; i--) {
      writer.printf("pop %s\n", calleeSaveRegs.get(i));
    }
    currentSPToBPOffset -= 8 * calleeSaveRegs.size();
    assert currentSPToBPOffset == 0;
    writer.println("ret");

    callerSaveRegs = null;
    currentProcedure = null;
    ops.variableLocations = null;

    CompilerRandomization.randomNopPadding(program, writer);
    writer.println();
    checkOutputSize();
  }

  @Override
  public void visit(IRInstructionList instrList) {
    super.visit(instrList);
  }

  @Override
  public void visit(Mov mov) {
    String target = ops.getTargetOperand(mov.target);
    if (targetReg != null) {
      ops.mov(targetReg, mov.source);
    } else {
      writer.printf("mov %s, %s\n", target, ops.getOperand(mov.source, IMM32_OR_REG));
    }
  }

  @Override
  public void visit(BinaryInstruction.Add add) {
    String target = ops.getTargetOperand(add.target);
    IROperand left = add.left;
    IROperand right = add.right;
    // check for assignment operator '+='
    if (add.target.equals(left)) {
      if (targetReg != null) {
        writer.printf("add %s, %s\n", target, ops.getOperand(right, IMM32_OR_REG_OR_MEM));
        return;
      }
      writer.printf("add %s, %s\n", target, ops.getOperand(right, IMM32_OR_REG));
      return;
    }
    // if left == right load operands only once
    if (left.equals(right)) {
      if (ops.isRegister(left)) {
        String reg = ops.getOperand(left, REG);
        writer.printf("lea %s, [%s+%s]\n", targetRegOrRax, reg, reg);
      } else {
        ops.mov(targetRegOrRax, left);
        writer.printf("add %s, %s\n", targetRegOrRax, targetRegOrRax);
      }
      ops.movToTarget(target);
      return;
    }
    // check lea possible
    if (targetReg != null && ops.isImm32(right)) {
      writer.printf("lea %s, [%s + %s]\n", target, ops.getOperand(left, REG), ops.getOperand(right, IMM32));
      return;
    } else if (ops.isRegister(left) && ops.isRegister(right)) {
      writer.printf("lea %s, [%s + %s]\n", targetRegOrRax, ops.getOperand(left, REG), ops.getOperand(right, REG));
      ops.movToTarget(target);
      return;
    }

    ops.mov(targetRegOrRax, left);
    var mode = targetRegOrRax == RAX ? IMM32_OR_REG_OR_MEM_NO_RAX : IMM32_OR_REG_OR_MEM;
    writer.printf("add %s, %s\n", targetRegOrRax, ops.getOperand(right, mode));
    ops.movToTarget(target);
  }

  @Override
  public void visit(BinaryInstruction.Sub sub) {
    String target = ops.getTargetOperand(sub.target);
    IROperand left = sub.left;
    IROperand right = sub.right;
    // check if can negate
    if (left instanceof IRIntConstant con && con.value == 0) {
      if (sub.target.equals(right)) {
        // check if target == right then use only neg
        // for memory targets this saves 2 mov instructions compared to the below case
        writer.printf("neg %s\n", target);
        return;
      }
      ops.mov(targetRegOrRax, right);
      writer.printf("neg %s\n", targetRegOrRax);
      ops.movToTarget(target);
      return;
    }
    // check for assignment operator '-='
    if (targetReg == null && sub.target.equals(left)) {
      // only do this on memory target as below will already do the same for register target
      writer.printf("sub %s, %s\n", target, ops.getOperand(right, IMM32_OR_REG));
      return;
    }

    if (targetReg != null && sub.target.equals(right)) {
      // when target == right and target is register use neg and add 
      writer.printf("neg %s\n", target);
      writer.printf("add %s, %s\n", target, ops.getOperand(left, IMM32_OR_REG_OR_MEM));
      return;
    }
    ops.mov(targetRegOrRax, left);
    var mode = targetRegOrRax == RAX ? IMM32_OR_REG_OR_MEM_NO_RAX : IMM32_OR_REG_OR_MEM;
    writer.printf("sub %s, %s\n", targetRegOrRax, ops.getOperand(right, mode));
    ops.movToTarget(target);
  }

  @Override
  public void visit(BinaryInstruction.Mul mul) {
    String target = ops.getTargetOperand(mul.target);
    IROperand left = mul.left;
    IROperand right = mul.right;
    // if left == right only load once
    if (left.equals(right)) {
      ops.mov(targetRegOrRax, left);
      writer.printf("imul %s, %s\n", targetRegOrRax, targetRegOrRax);
      ops.movToTarget(target);
      return;
    }

    ops.mov(targetRegOrRax, left);
    var mode = targetRegOrRax == RAX ? IMM32_OR_REG_OR_MEM_NO_RAX : IMM32_OR_REG_OR_MEM;
    writer.printf("imul %s, %s\n", targetRegOrRax, ops.getOperand(right, mode));
    ops.movToTarget(target);
  }

  @Override
  public void visit(BinaryInstruction.Div div) {
    String target = ops.getTargetOperand(div.target);
    ops.mov(RAX, div.left);
    boolean saveRdx = callerSaveRegs.contains(RDX);
    saveRdx = saveRdx && (!"rdx".equals(target) || div.target.equals(div.right));
    if (saveRdx) {
      writer.println("push rdx");
      currentSPToBPOffset += 8;
    }
    String right = ops.getOperand(div.right, REG_OR_MEM);
    if ("rdx".equals(right)) {
      right = "QWORD PTR[rsp]";
    }
    writer.printf("""
            cqo
            idiv %s
            """, right);
    if (saveRdx) {
      writer.println("pop rdx");
      currentSPToBPOffset -= 8;
    }
    writer.printf("mov %s, rax\n", target);
  }

  @Override
  public void visit(Label label) {
    writer.printf(".%s=.\n", label.mnemonic());
  }

  @Override
  public void visit(Jump.UnconditionalJump unconditionalJump) {
    writer.printf("jmp .%s\n", unconditionalJump.target.mnemonic());
  }

  @Override
  public void visit(Jump.ConditionalJump conditionalJump) {
    IROperand left = conditionalJump.left;
    IROperand right = conditionalJump.right;
    if (ops.isRegister(left)) {
      writer.printf("cmp %s, %s\n",
              ops.getOperand(left, REG),
              ops.getOperand(right, IMM32_OR_REG_OR_MEM));
    } else if (ops.isRegister(right)) {
      writer.printf("cmp %s, %s\n",
              ops.getOperand(left, IMM32_OR_REG_OR_MEM),
              ops.getOperand(right, REG));
    } else {
      ops.mov(RAX, left);
      writer.printf("cmp rax, %s\n", ops.getOperand(right, IMM32_OR_REG_OR_MEM));
    }
    ComparisonExpression.Op op = conditionalJump.operator;
    String jmpMemonic = switch (op) {
      case LT -> "jl";
      case LE -> "jle";
      case GE -> "jge";
      case GT -> "jg";
      case EQ -> "je";
      case NEQ -> "jne";
      case U_B -> "jb";
      case U_BE -> "jbe";
      case U_AE -> "jae";
      case U_A -> "ja";
    };

    writer.printf("%s .%s\n", jmpMemonic, conditionalJump.target.mnemonic());
  }

  @Override
  public void visit(Call call) {
    Map<X64Operands.HardwareRegister, X64Operands.StackOffset> alternativeLocation = new HashMap<>();
    // push caller save regs
    for (var callerSaveReg : callerSaveRegs) {
      writer.printf("push %s\n", callerSaveReg);
      currentSPToBPOffset += 8;
      alternativeLocation.put(callerSaveReg, new X64Operands.StackOffset(-currentSPToBPOffset));
    }
    int stackArgSize = 8 * Math.max(call.args.length - CALLING_CONVENTION.length, 0);
    // align stack - round to multiple of 16 + 8 cause of saved return address
    if ((currentSPToBPOffset + stackArgSize) % 16 != 8) {
      writer.println("sub rsp, 8");
      currentSPToBPOffset += 8;
      stackArgSize += 8;
    }
    // push stack arguments
    for (int i = call.args.length - 1; i >= CALLING_CONVENTION.length; i--) {
      IROperand arg = call.args[i];
      if (arg.getType() instanceof StringType) {
        throw new IllegalStateException("string parameter after 6th position?");
      } else {
        writer.printf("push %s\n", ops.getOperand(arg, IMM32_OR_REG_OR_MEM));
      }
      currentSPToBPOffset += 8;
    }
    // mov register arguments and keep track of overridden registers
    Set<X64Operands.HardwareRegister> clobberedRegisters = new HashSet<>();
    for (int i = 0; i < Math.min(call.args.length, CALLING_CONVENTION.length); i++) {
      var target = CALLING_CONVENTION[i];
      IROperand arg = call.args[i];
      if (arg instanceof IRVariable var && ops.variableLocations.get(var) instanceof X64Operands.HardwareRegister reg) {
        if (clobberedRegisters.contains(reg)) {
          ops.variableLocations.put(var, alternativeLocation.get(reg));
          ops.mov(target, var);
          ops.variableLocations.put(var, reg);
          clobberedRegisters.add(target);
          continue;
        }
      }
      if (arg.getType() instanceof StringType) {
        ops.movStringAddr(target, arg);
      } else {
        ops.mov(target, arg);
      }
      clobberedRegisters.add(target);
    }

    writer.printf("call %s\n", call.name);

    checkStringFunctionResult(call);

    // remove stack arguments
    if (stackArgSize != 0) {
      writer.printf("add rsp, %d\n", stackArgSize);
      currentSPToBPOffset -= stackArgSize;
    }

    for (int i = callerSaveRegs.size() - 1; i >= 0; i--) {
      var callerSaveReg = callerSaveRegs.get(i);
      writer.printf("pop %s\n", callerSaveReg);
      currentSPToBPOffset -= 8;
    }

    if (call.target != null) {
      writer.printf("mov %s, rax\n", ops.getTargetOperand(call.target));
    }
  }

  @Override
  public void visit(Load load) {
    String target = ops.getTargetOperand(load.target);
    String arrayPTR = ops.getArrayPTR(load.arrayVar, load.index);

    writer.printf("mov %s, %s\n", targetRegOrRax, arrayPTR);
    ops.movToTarget(target);
  }

  @Override
  public void visit(Store store) {
    targetReg = null;
    IROperand expression = store.expression;
    if (store.arrayVar.type instanceof StringType) {
      // initial null byte for string - other stores would use $strput
      if (!(store.expression instanceof IRIntConstant c) || c.value != 0) {
        throw new UnsupportedOperationException();
      }
      if (!(store.index instanceof IRIntConstant c2) || c2.value != 0) {
        throw new UnsupportedOperationException();
      }
      String arrayPtr = ops.getStringPTR(store.arrayVar);
      writer.printf("mov BYTE PTR %s, 0\n", arrayPtr);
      return;
    }

    if (ops.isRegister(expression) || ops.isImm32(expression)) {
      String arrayPTR = ops.getArrayPTR(store.arrayVar, store.index);
      writer.printf("mov %s, %s\n", arrayPTR, ops.getOperand(expression, IMM32_OR_REG));
      return;
    }
    // need one scratch register to hold expression value for mov to memory -> use RDX 
    // need one scratch register to hold PTR target address -> use RAX (in ops.getArrayPTR())
    boolean saveRdx = callerSaveRegs.contains(RDX);
    if (saveRdx) {
      writer.println("push rdx");
      currentSPToBPOffset += 8;
    }
    String arrayPTR = ops.getArrayPTR(store.arrayVar, store.index);
    ops.mov(RDX, expression);
    writer.printf("mov %s, rdx\n", arrayPTR);
    if (saveRdx) {
      writer.println("pop rdx");
      currentSPToBPOffset -= 8;
    }
  }
}
                            