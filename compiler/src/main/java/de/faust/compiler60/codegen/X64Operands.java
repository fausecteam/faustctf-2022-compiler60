package de.faust.compiler60.codegen;

import de.faust.compiler60.ir.*;
import de.faust.compiler60.types.ArrayType;
import de.faust.compiler60.types.IntType;
import de.faust.compiler60.types.StringType;
import de.faust.compiler60.types.Type;

import java.io.PrintWriter;
import java.util.ArrayDeque;
import java.util.Arrays;
import java.util.Deque;
import java.util.HashMap;

import static de.faust.compiler60.codegen.X64Operands.HardwareRegister.CALLEE_SAVE;
import static de.faust.compiler60.codegen.X64Operands.HardwareRegister.CALLER_SAVE;

public class X64Operands {
  public final ConstantPool constantPool;

  public final CodegenX64 codegen;
  public final PrintWriter writer;

  public HashMap<IRVariable, HardwareLocation> variableLocations;

  public X64Operands(CodegenX64 codegen) {
    this.writer = codegen.writer;
    this.codegen = codegen;
    variableLocations = new HashMap<>();
    constantPool = new ConstantPool();
  }

  public Deque<HardwareRegister> getRegistersForFunction(IRProcedure procedure) {
    boolean useCalleeSaveRegsFirst = procedure.body.instructions.stream()
            .anyMatch(instruction -> instruction instanceof Call);
    Deque<X64Operands.HardwareRegister> regs = new ArrayDeque<>(CALLEE_SAVE.length + CALLER_SAVE.length);
    if (useCalleeSaveRegsFirst) {
      regs.addAll(Arrays.asList(CALLEE_SAVE));
      regs.addAll(Arrays.asList(CALLER_SAVE));
    } else {
      regs.addAll(Arrays.asList(CALLER_SAVE));
      regs.addAll(Arrays.asList(CALLEE_SAVE));
    }
    return regs;
  }

  public int getSize(Type t) {
    if (t == IntType.INSTANCE) {
      return 8;
    }
    if (t instanceof ArrayType arrayType) {
      assert arrayType.sizes.size() == 1;
      return Math.toIntExact(arrayType.sizes.get(0) * 8);
    }
    if (t instanceof StringType stringType) {
      // round up to 8 bytes
      return Math.toIntExact(stringType.size + 7) / 8 * 8;
    }
    throw new IllegalStateException();
  }

  public String getTargetOperand(IRVariable target) {
    assert target.type == IntType.INSTANCE;
    HardwareLocation loc = variableLocations.get(target);
    if (loc instanceof HardwareRegister reg) {
      if (reg == HardwareRegister.RAX) throw new UnsupportedOperationException();
      codegen.targetReg = reg;
      codegen.targetRegOrRax = reg;
      return reg.registerName;
    }
    codegen.targetReg = null;
    codegen.targetRegOrRax = HardwareRegister.RAX;
    if (loc instanceof GlobalSymbol symbol) {
      return String.format("QWORD PTR [rip + %s]", symbol.name);
    }
    if (loc instanceof StackOffset stackOffset) {
      return String.format("QWORD PTR [rsp %+d]", stackOffset.currentSPOffset(codegen.currentSPToBPOffset));
    }
    throw new UnsupportedOperationException();
  }

  // move codegen.targetRegOrRax to targetOperand
  public void movToTarget(String targetOperand) {
    if (codegen.targetReg != null) {
      assert codegen.targetReg == codegen.targetRegOrRax;
      return; // result already is in targetReg
    }
    assert codegen.targetRegOrRax.equals(HardwareRegister.RAX);
    // move result from RAX to target memory location
    writer.printf("mov %s, rax\n", targetOperand);
  }

  public void mov(HardwareRegister regInto, IROperand operand) {
    if (operand instanceof IRIntConstant c) {
      if (c.value == 0) {
        String regName = regInto.getZeroExtendName();
        writer.printf("xor %s, %s\n", regName, regName);
        return;
      }
      if (c.value == 1) {
        String regName = regInto.getZeroExtendName();
        writer.printf("mov %s, 1\n", regName);
        return;
      }
    } else if (regInto == getRegister(operand)) {
      // no need to write something like "mov rcx, rcx"
      return;
    }
    writer.printf("mov %s, %s\n", regInto, getOperand(operand, OperandMode.IMM32_OR_REG_OR_MEM));
  }

  public HardwareRegister getRegister(IROperand operand) {
    return operand instanceof IRVariable var && variableLocations.get(var) instanceof HardwareRegister reg ? reg : null;
  }

  public boolean isRegister(IROperand operand) {
    return getRegister(operand) != null;
  }

  public boolean isImm32(IROperand operand) {
    return operand instanceof IRIntConstant c && !constantPool.willBePTR(c);
  }

  enum OperandMode {
    IMM32(true, false, false, false),
    IMM32_OR_REG(true, true, false, false),             // e.g. x in add PTR[], x
    IMM32_OR_REG_OR_MEM(true, true, true, false),       // e.g. x in add rdx, x
    IMM32_OR_REG_OR_MEM_NO_RAX(true, true, true, true), // e.g. x in add rax, x
    REG(false, true, false, false),                     // e.g. x in lea rax, [x + y]
    REG_OR_MEM(false, true, true, false)                // e.g. y in imul rax, y
    ;

    final boolean imm32;
    final boolean reg;
    final boolean mem;
    final boolean noRax;

    OperandMode(boolean imm32, boolean reg, boolean mem, boolean noRax) {
      this.imm32 = imm32;
      this.reg = reg;
      this.mem = mem;
      this.noRax = noRax;
    }
  }

  protected String getOperand(IROperand operand, OperandMode mode) {
    if (operand instanceof IRIntConstant c) {
      if (constantPool.willBePTR(c)) {
        if (mode.mem) {
          return constantPool.useIntConstant(c);
        }
        if (mode.reg) {
          mov(HardwareRegister.RAX, c);
          return HardwareRegister.RAX.registerName;
        }
      } else {
        if (mode.imm32) {
          return constantPool.useIntConstant(c);
        }
        if (mode.reg && !mode.noRax) {
          mov(HardwareRegister.RAX, c);
          return HardwareRegister.RAX.registerName;
        }
      }
      throw new UnsupportedOperationException();
    }
    if (operand instanceof IRStringConstant) {
      // must use movStringAddr()
      throw new UnsupportedOperationException();
    }
    HardwareLocation loc = variableLocations.get((IRVariable) operand);
    if (loc instanceof GlobalSymbol symbol) {
      if (mode.mem) {
        return String.format("QWORD PTR [rip + %s]", symbol.name);
      }
      if (mode.reg) {
        mov(HardwareRegister.RAX, operand);
        return HardwareRegister.RAX.registerName;
      }
    }
    if (loc instanceof StackOffset stackOffset) {
      if (mode.mem) {
        return String.format("QWORD PTR [rsp %+d]", stackOffset.currentSPOffset(codegen.currentSPToBPOffset));
      }
      if (mode.reg) {
        mov(HardwareRegister.RAX, operand);
        return HardwareRegister.RAX.registerName;
      }
    }
    if (loc instanceof HardwareRegister reg) {
      if (mode.reg) {
        return reg.registerName;
      }
    }
    throw new UnsupportedOperationException();
  }

  public String getStringPTR(IRVariable variable) {
    assert variable.getType() instanceof StringType;
    HardwareLocation loc = variableLocations.get(variable);
    if (loc instanceof GlobalSymbol symbol) {
      return String.format("[rip + %s]", symbol.name);
    }
    if (loc instanceof StackOffset stackOffset) {
      return String.format("[rsp %+d]", stackOffset.currentSPOffset(codegen.currentSPToBPOffset));
    }
    throw new UnsupportedOperationException();
  }

  public void movStringAddr(HardwareRegister regInto, IROperand operand) {
    assert operand.getType() instanceof StringType;
    if (operand instanceof IRStringConstant s) {
      String stringAddr = constantPool.useStringConstant(s);
      writer.printf("lea %s, %s\n", regInto, stringAddr);
      return;
    }
    writer.printf("lea %s, %s\n", regInto, getStringPTR((IRVariable) operand));
  }

  public String getArrayPTR(IRVariable array, IROperand index) {
    ArrayType arrayType = (ArrayType) array.type;
    assert arrayType.sizes.size() == 1;
    
    // prevent accessing indexes out of bounds
    long size = arrayType.sizes.get(0);
    if (size > Integer.MAX_VALUE) {
      writer.printf("mov rax, %d\n", size);
      writer.printf("cmp rax, %s\n", getOperand(index, OperandMode.IMM32_OR_REG_OR_MEM_NO_RAX));
      writer.printf("jbe .%s\n", Label.ILLEGAL_INDEX.mnemonic());
    } else {
      writer.printf("cmp %s, %d\n", getOperand(index, OperandMode.REG_OR_MEM), size);
      writer.printf("jae .%s\n", Label.ILLEGAL_INDEX.mnemonic());
    }
    
    HardwareLocation arrayLoc = variableLocations.get(array);
    if (index instanceof IRIntConstant c && !constantPool.willBePTR(c)) {
      // use assembler to calculate index
      if (arrayLoc instanceof StackOffset stackOffset) {
        return String.format("QWORD PTR [rsp + (8 * %+d %+d)]",
                c.value, stackOffset.currentSPOffset(codegen.currentSPToBPOffset));
      } else if (arrayLoc instanceof GlobalSymbol symbol) {
        return String.format("QWORD PTR [rip + (8 * %+d + %s)]", c.value, symbol.name);
      }
    }
    if (arrayLoc instanceof StackOffset stackOffset) {
      return String.format("QWORD PTR[rsp + 8 * %s %+d]",
              getOperand(index, OperandMode.REG), stackOffset.currentSPOffset(codegen.currentSPToBPOffset));
    }
    if (arrayLoc instanceof GlobalSymbol symbol) {
      writer.printf("lea rax, [rip + %s]\n", symbol.name);
      if (isRegister(index)) {
        return String.format("QWORD PTR[rax + 8 * %s]", getOperand(index, OperandMode.REG));
      }
      if (codegen.targetReg != null) {
        mov(codegen.targetReg, index);
        return String.format("QWORD PTR[rax + 8 * %s]", codegen.targetReg);
      }
      writer.printf("shr rax, 3\n");
      writer.printf("add rax, %s\n", getOperand(index, OperandMode.IMM32_OR_REG_OR_MEM_NO_RAX));
      return "QWORD PTR[8 * rax]";
    }
    throw new UnsupportedOperationException();
  }

  interface HardwareLocation {
  }

  enum HardwareRegister implements HardwareLocation {
    RAX("rax"),
    RCX("rcx"),
    RDX("rdx"),
    RBX("rbx"),
    //RSP("rsp"),
    RBP("rbp"),
    RSI("rsi"),
    RDI("rdi"),
    R8("r8"),
    R9("r9"),
    R10("r10"),
    R11("r11"),
    R12("r12"),
    R13("r13"),
    R14("r14"),
    R15("r15");
    // do not use RAX here because of division (and as temporary register)
    static final HardwareRegister[] CALLER_SAVE = new HardwareRegister[]{RCX, RSI, RDI, R8, R9, R10, R11, RDX};
    static final HardwareRegister[] CALLEE_SAVE = new HardwareRegister[]{RBX, RBP, R12, R13, R14, R15};
    static final HardwareRegister[] CALLING_CONVENTION = new HardwareRegister[]{RDI, RSI, RDX, RCX, R8, R9};

    final String registerName;

    HardwareRegister(String registerName) {
      this.registerName = registerName;
    }

    String getZeroExtendName() {
      if (Character.isDigit(registerName.charAt(1))) {
        // here the d extension does not save bytes
        // xor r10d, r10d
        // xor r10, r10
        // has same encoded length
        return registerName;
      }
      // use 32-bit registers
      return "e" + registerName.substring(1);
    }

    @Override
    public String toString() {
      return registerName;
    }
  }

  record GlobalSymbol(String name) implements HardwareLocation {
  }

  record StackOffset(int offsetFromBP) implements HardwareLocation {
    int currentSPOffset(int currentSPToBPOffset) {
      return currentSPToBPOffset + this.offsetFromBP;
    }
  }
}
