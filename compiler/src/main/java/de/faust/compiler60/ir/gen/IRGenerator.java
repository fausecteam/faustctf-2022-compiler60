package de.faust.compiler60.ir.gen;

import de.faust.compiler60.CompilerRandomization;
import de.faust.compiler60.ast.*;
import de.faust.compiler60.ir.*;
import de.faust.compiler60.semantics.Symbol;
import de.faust.compiler60.types.*;

import java.util.*;
import java.util.function.Function;
import java.util.stream.IntStream;

public class IRGenerator extends ASTVisitor<IRGenerator.SourceOperand> {

  protected IRProgram program;

  protected IRProcedure currentProcedure;

  protected Set<Symbol> globals;
  protected Map<Symbol, IRVariable> globalVars;
  protected Map<Symbol, IRVariable> localVars;

  protected Deque<IRVariable> virtualRegisterStack;
  protected TreeMap<StringType, IRVariable> virtualStringVariables;
  protected int labelCounter;

  protected IRInstruction addNextInstrBefore;

  protected Label conditionalJumpTarget;
  protected boolean conditionalJumpTargetIsTrue;

  public IRProgram generateIR(Program program) {
    // apply simplifications to ast
    new MultiDimArrayFlatten().transform(program);
    new ForToWhileLoop().transform(program);

    this.program = new IRProgram();
    visitTopLevelBlock(program.block);
    
    CompilerRandomization.setupRandomization(program, this.program);
    return this.program;
  }

  protected void visitTopLevelBlock(BlockStatement block) {
    // store symbols for globals, so we differentiate in getIRVar()
    globalVars = new HashMap<>();
    globals = new HashSet<>();
    for (Declaration declaration : block.declarations) {
      if (!(declaration instanceof ProcedureDeclaration)) {
        globals.add(declaration.identifier.symbol);
      }
    }
    // generate top level statements
    ProcedureDeclaration topLevelProcedure = new ProcedureDeclaration(-1, new Identifier(-1, "$top"), List.of(), VoidType.INSTANCE, block);
    visit(topLevelProcedure);
    
    // generate procedures
    for (Declaration declaration : block.declarations) {
      if (declaration instanceof ProcedureDeclaration procDecl) {
        visit(procDecl);
      }
    }
  }

  @Override
  public SourceOperand visit(ProcedureDeclaration procedureDeclaration) {
    currentProcedure = new IRProcedure(procedureDeclaration);
    program.procedures.add(currentProcedure);
    localVars = new HashMap<>();
    
    if (procedureDeclaration.returnType != VoidType.INSTANCE) {
      currentProcedure.returnValue = getIRVar(procedureDeclaration.returnVariable);
    }
    for (Identifier parameter : procedureDeclaration.parameters) {
      IRVariable paramVar = new IRVariable(parameter.name, parameter.getType(), false);
      currentProcedure.parameters.add(paramVar);
      localVars.put(parameter.symbol, paramVar);
    }
    virtualRegisterStack = new ArrayDeque<>();
    virtualStringVariables = new TreeMap<>();

    visit(procedureDeclaration.procedureBody);

    currentProcedure = null;
    localVars = null;
    virtualRegisterStack = null;
    assert addNextInstrBefore == null;
    assert conditionalJumpTarget == null;
    return null;
  }

  @Override
  public SourceOperand visit(BlockStatement block) {
    // Declarations are ignored cause getIRVar() creates IRVariables on use
    for (Statement statement : block.statements) {
      statement.accept(this);
    }
    return null;
  }

  @Override
  public SourceOperand visit(ForStatement forStatement) {
    assert forStatement.loopVar == null;
    assert forStatement.forList.size() == 1;
    assert ((ForStatement.While) forStatement.forList.get(0)).updateExpr == null;

    Label whileBody = createLabel();
    Label whileTailCond = createLabel();
    // jump to tail condition
    addInstr(new Jump.UnconditionalJump(whileTailCond));
    // generate while body
    addInstr(whileBody);
    forStatement.body.accept(this);
    // generate tail condition
    addInstr(whileTailCond);
    conditionalJumpTarget = whileBody;
    conditionalJumpTargetIsTrue = true;
    ((ForStatement.While) forStatement.forList.get(0)).whileCond.accept(this);

    conditionalJumpTarget = null;
    return null;
  }

  @Override
  public SourceOperand visit(ConditionalStatement conditionalStatement) {
    IRInstruction saveInstrBefore = addNextInstrBefore;

    conditionalJumpTarget = createLabel();
    conditionalJumpTargetIsTrue = false;

    addInstr(conditionalJumpTarget);

    // generate condition
    addNextInstrBefore = conditionalJumpTarget;
    conditionalStatement.condition.accept(this);

    // generate if body
    addNextInstrBefore = conditionalJumpTarget;
    conditionalJumpTarget = null;
    conditionalStatement.thenBlock.accept(this);

    if (conditionalStatement.elseBlock != null) {
      Label endElse = createLabel();
      addInstr(new Jump.UnconditionalJump(endElse));
      // generate else body
      addNextInstrBefore = saveInstrBefore;
      conditionalStatement.elseBlock.accept(this);

      addInstr(endElse);
    }
    addNextInstrBefore = saveInstrBefore;
    return null;
  }

  @Override
  public SourceOperand visit(BinaryBooleanExpression binaryBooleanExpression) {
    // create short circuit semantic of and / or
    IRInstruction saveInstrBefore = addNextInstrBefore;
    Label saveJmpTarget = conditionalJumpTarget;
    boolean saveJumpIsTrue = conditionalJumpTargetIsTrue;
    boolean isAnd = binaryBooleanExpression.op == BinaryBooleanExpression.Op.AND;
    if (isAnd == conditionalJumpTargetIsTrue) {
      // jump target needs to swap
      // for 'or' left side == true jumps over right side
      // for 'and' left side == false jumps over right side
      conditionalJumpTargetIsTrue = !isAnd;
      // create label to jump over the right side
      conditionalJumpTarget = createLabel();
      addInstr(conditionalJumpTarget);
      // add left side before label
      addNextInstrBefore = conditionalJumpTarget;
    }
    // create left side
    binaryBooleanExpression.left.accept(this);
    // restore condition for right side
    conditionalJumpTarget = saveJmpTarget;
    conditionalJumpTargetIsTrue = saveJumpIsTrue;
    // create right side
    binaryBooleanExpression.right.accept(this);
    // put following instructions after the label (if created)
    addNextInstrBefore = saveInstrBefore;
    return null;
  }

  @Override
  public SourceOperand visit(BooleanNotExpression booleanNotExpression) {
    conditionalJumpTargetIsTrue = !conditionalJumpTargetIsTrue;
    booleanNotExpression.toInvertExpression.accept(this);
    conditionalJumpTargetIsTrue = !conditionalJumpTargetIsTrue;
    return null;
  }

  @Override
  public SourceOperand visit(TernaryBooleanExpression ternaryBooleanExpression) {
    IRInstruction saveInstrBefore = addNextInstrBefore;
    Label saveJmpTarget = conditionalJumpTarget;
    boolean saveJumpIsTrue = conditionalJumpTargetIsTrue;

    Label beforeElse = createLabel();
    Label afterElse = createLabel();
    addInstr(beforeElse);
    addInstr(afterElse);

    // generate condition jumping to beforeElse on false 
    addNextInstrBefore = beforeElse;
    conditionalJumpTarget = beforeElse;
    conditionalJumpTargetIsTrue = false;
    ternaryBooleanExpression.cond.accept(this);
    // on true fall-though and do conditional jump with true case
    conditionalJumpTarget = saveJmpTarget;
    conditionalJumpTargetIsTrue = saveJumpIsTrue;
    ternaryBooleanExpression.trueValue.accept(this);
    // if conditional jump is not met need skip over else part to fall-through 
    addInstr(new Jump.UnconditionalJump(afterElse));

    // generate false case at beforeElse; f conditional jump is not met fall-through over afterElse
    addNextInstrBefore = afterElse;
    ternaryBooleanExpression.falseValue.accept(this);

    addNextInstrBefore = saveInstrBefore;
    return null;
  }

  @Override
  public SourceOperand visit(ComparisonExpression comparisonExpression) {
    ComparisonExpression.Op cmpOp = comparisonExpression.op;
    // check if we need to invert operator
    if (!conditionalJumpTargetIsTrue) {
      cmpOp = cmpOp.invert();
    }
    // descent into left side then right side
    SourceOperand lValue = comparisonExpression.left.accept(this);
    SourceOperand rValue = comparisonExpression.right.accept(this);

    // generate conditional jump
    var finalCmpOp = cmpOp;
    var substituteOperand = new CalculatedSourceOperand(null, List.of(lValue, rValue), ops ->
            new Jump.ConditionalJump(conditionalJumpTarget, ops[0], ops[1], finalCmpOp));
    // just call assignTo() with null target
    substituteOperand.assignTo(null, false);

    return null;
  }

  @Override
  public SourceOperand visit(AssignmentStatement assignmentStatement) {
    Expression target = assignmentStatement.target;
    if (target instanceof Identifier identifier) {
      SourceOperand sourceOperand = assignmentStatement.expression.accept(this);
      IRVariable targetVar = getIRVar(identifier.symbol);
      sourceOperand.assignTo(targetVar, true);
      return null;
    } else if (target instanceof SubscriptedVariable subscriptedVariable) {
      // assignment to array
      assert subscriptedVariable.idxExprs.size() == 1;
      // first get the array index value to assign
      SourceOperand arrayIdx = subscriptedVariable.idxExprs.get(0).accept(this);
      // then get the value to assign
      SourceOperand expression = assignmentStatement.expression.accept(this);

      IRVariable targetArrayVar = getIRVar(subscriptedVariable.arrayVar.symbol);

      var substituteOperand = new CalculatedSourceOperand(null, List.of(arrayIdx, expression), ops ->
              new Store(targetArrayVar, ops[0], ops[1]));
      // just call assignTo() with null target
      substituteOperand.assignTo(null, false);

      return null;
    }
    throw new IllegalStateException();
  }

  @Override
  public SourceOperand visit(ProcedureStatement procedureStatement) {
    var substituteOperand = visit(procedureStatement.procedureExpression);
    // call assignTo() with null target so no virtual register is used for the return value 
    substituteOperand.assignTo(null, false);

    return null;
  }

  @Override
  public SourceOperand visit(TernaryExpression ternaryExpression) {
    SourceOperand trueCase = ternaryExpression.trueCase.accept(this);
    SourceOperand falseCase = ternaryExpression.falseCase.accept(this);
    List<SourceOperand> ops = new ArrayList<>(List.of(trueCase, falseCase));
    return new CalculatedSourceOperand(ternaryExpression.getType(), ops, null) {
      @Override
      boolean checkCanUseTarget(IRVariable target) {
        // do both orders because the order does not matter here 
        boolean a = super.checkCanUseTarget(target);
        Collections.reverse(this.usedOperands);
        boolean b = super.checkCanUseTarget(target);
        return a | b;
      }

      @Override
      void assignTo(IRVariable target, boolean freshTarget) {
        IRInstruction saveInstrBefore = addNextInstrBefore;
        Label saveJmpTarget = conditionalJumpTarget;
        boolean saveJumpIsTrue = conditionalJumpTargetIsTrue;

        Label beforeElse = createLabel();
        Label afterElse = createLabel();
        addInstr(beforeElse);
        addInstr(afterElse);


        // generate condition
        // in theory, you cloud check if target is used in neither trueCase nor falseCase
        // and push target as virtual register for the condition, but that's a bit overkill
        addNextInstrBefore = beforeElse;
        conditionalJumpTarget = beforeElse;
        conditionalJumpTargetIsTrue = false;
        ternaryExpression.cond.accept(IRGenerator.this);
        // generate true case
        trueCase.assignTo(target, freshTarget);
        // jump over else part
        addInstr(new Jump.UnconditionalJump(afterElse));

        // generate false case
        addNextInstrBefore = afterElse;
        falseCase.assignTo(target, freshTarget);

        addNextInstrBefore = saveInstrBefore;
        conditionalJumpTarget = saveJmpTarget;
        conditionalJumpTargetIsTrue = saveJumpIsTrue;
      }
    };
  }

  @Override
  public SourceOperand visit(BinaryExpression binaryExpression) {
    SourceOperand leftValue = binaryExpression.left.accept(this);
    SourceOperand rightValue = binaryExpression.right.accept(this);

    Function<IROperand[], IRInstruction> create = switch (binaryExpression.op) {
      case ADD -> ops -> new BinaryInstruction.Add(null, ops[0], ops[1]);
      case SUBTRACT -> ops -> new BinaryInstruction.Sub(null, ops[0], ops[1]);
      case MULTIPLY -> ops -> new BinaryInstruction.Mul(null, ops[0], ops[1]);
      case DIVIDE -> ops -> new BinaryInstruction.Div(null, ops[0], ops[1]);
    };

    return new CalculatedSourceOperand(binaryExpression.getType(), List.of(leftValue, rightValue), create);
  }

  @Override
  public SourceOperand visit(SubscriptedVariable subscriptedVariable) {
    assert subscriptedVariable.idxExprs.size() == 1;

    SourceOperand arrayIdx = subscriptedVariable.idxExprs.get(0).accept(this);

    IRVariable targetArrayVar = getIRVar(subscriptedVariable.arrayVar.symbol);
    return new CalculatedSourceOperand(subscriptedVariable.getType(),
            List.of(arrayIdx), ops -> new Load(null, targetArrayVar, ops[0]));
  }

  @Override
  public SourceOperand visit(ProcedureExpression procedureExpression) {
    List<SourceOperand> args = procedureExpression.arguments.stream()
            .map(arg -> arg.accept(this))
            .toList();

    return new CalculatedSourceOperand(procedureExpression.getType(), args, ops ->
            new Call(null, procedureExpression.procedureName.name, ops));
  }

  @Override
  public SourceOperand visit(Identifier identifier) {
    if (identifier.getType() instanceof ProcedureType) {
      return null; // ignore procedure identifiers
    }
    // return IRVariable new/existing from this VariableSymbol
    return new DirectOperand(identifier.getType(), getIRVar(identifier.symbol));
  }

  @Override
  public SourceOperand visit(IntConstant intConstant) {
    return new DirectOperand(IntType.INSTANCE, new IRIntConstant(intConstant.value));
  }

  @Override
  public SourceOperand visit(StringConstant stringConstant) {
    return new DirectOperand(stringConstant.getType(), new IRStringConstant(stringConstant));
  }

  protected void addInstr(IRInstruction instr) {
    currentProcedure.body.insertBefore(instr, addNextInstrBefore);
  }

  protected Label createLabel() {
    return new Label(labelCounter++);
  }

  protected IRVariable getIRVar(Symbol symbol) {
    if (globals.contains(symbol)) {
      return globalVars.computeIfAbsent(symbol, sym -> {
        IRVariable var = new IRVariable(sym.getName(), sym.getType(), true);
        program.globals.add(var);
        return var;
      });
    }
    return localVars.computeIfAbsent(symbol, sym -> {
      IRVariable var = new IRVariable(sym.getName(), sym.getType(), false);
      currentProcedure.localVariables.add(var);
      return var;
    });
  }

  protected IRVariable popVirtualRegister(Type variableType) {
    if (variableType instanceof StringType stringType) {
      if (virtualStringVariables.isEmpty()) {
        String name = "%s" + currentProcedure.localVariables.size();
        IRVariable newStringVR = new IRVariable(name, stringType, false);
        currentProcedure.localVariables.add(newStringVR);
        return newStringVR;
      }
      Map.Entry<StringType, IRVariable> nextLargerEntry = virtualStringVariables.ceilingEntry(stringType);
      if (nextLargerEntry != null) {
        virtualStringVariables.remove(nextLargerEntry.getKey());
        return nextLargerEntry.getValue();
      }
      IRVariable largestStringVar = virtualStringVariables.pollLastEntry().getValue();
      // increase to necessary size
      largestStringVar.type = stringType;
      return largestStringVar;
    }
    if (variableType != IntType.INSTANCE)
      throw new IllegalArgumentException();
    // take from stack or create a new one
    if (!virtualRegisterStack.isEmpty()) {
      return virtualRegisterStack.pop();
    }
    String name = "%v" + currentProcedure.localVariables.size();
    IRVariable newVR = new IRVariable(name, IntType.INSTANCE, false);
    currentProcedure.localVariables.add(newVR);
    return newVR;
  }

  protected void pushVirtualRegister(IRVariable variable) {
    if (variable.type instanceof StringType stringType) {
      virtualStringVariables.put(stringType, variable);
    } else if (variable.type == IntType.INSTANCE) {
      virtualRegisterStack.push(variable);
    } else {
      throw new IllegalArgumentException();
    }
  }

  abstract class SourceOperand {
    final Type type;
    boolean canUseTarget;

    SourceOperand(Type type) {
      this.type = type;
    }

    abstract void assignTo(IRVariable target, boolean freshTarget);
  }

  private class DirectOperand extends SourceOperand {
    final IROperand operand;

    DirectOperand(Type type, IROperand operand) {
      super(type);
      this.operand = operand;
    }

    @Override
    void assignTo(IRVariable target, boolean freshTarget) {
      addInstr(new Mov(target, operand));
    }
  }

  private class CalculatedSourceOperand extends SourceOperand {
    final List<SourceOperand> usedOperands;
    final Function<IROperand[], IRInstruction> createInstruction;

    CalculatedSourceOperand(Type type, List<SourceOperand> usedOperands, Function<IROperand[], IRInstruction> createInstruction) {
      super(type);
      this.usedOperands = usedOperands;
      this.createInstruction = createInstruction;
    }

    boolean checkCanUseTarget(IRVariable target) {
      /// do a depth first traversal but in reverse
      // check that all operands are not directly the target
      for (int i = usedOperands.size() - 1; i >= 0; i--) {
        if (usedOperands.get(i) instanceof DirectOperand directOperand
                && directOperand.operand == target) {
          return true;
        }
      }
      // recurse for child nodes until usage of target is found
      for (int i = usedOperands.size() - 1; i >= 0; i--) {
        SourceOperand usedOperand = usedOperands.get(i);
        usedOperand.canUseTarget = true;
        if (usedOperand instanceof CalculatedSourceOperand calc
                && calc.checkCanUseTarget(target)) {
          return true;
        }
      }
      return false;
    }

    @Override
    void assignTo(IRVariable target, boolean freshTarget) {
      if (freshTarget && target != null && (!target.isGlobal || target.type instanceof StringType)) {
        // check how long old target value is still needed 
        checkCanUseTarget(target);
      }
      // find index of last CalculatedSourceOperand (which could modify globals)
      int lastCalculated = IntStream.range(0, usedOperands.size())
              .filter(i -> usedOperands.get(i) instanceof CalculatedSourceOperand).max().orElse(-1);
      IROperand[] operands = new IROperand[usedOperands.size()];
      IRVariable savedTarget = target;
      List<IRVariable> virtualRegisters = new ArrayList<>();
      for (int i = 0; i < usedOperands.size(); i++) {
        SourceOperand usedOperand = usedOperands.get(i);
        if (usedOperand instanceof DirectOperand directOperand) {
          // Prevent postponing the access of globals beyond a CalculatedSourceOperand,
          // which might be a ProcedureExpression and modify the global.
          // In the original ALGOL60 spec 3.3.3 this is actually not guaranteed.
          // e.g. x + inc_x() should be:
          // %v0 = mov x ; %v1 = call inc_x() ; %v0 = add %v0, %v1
          // instead of:
          // %v0 = call inc_x() ; %v0 = add x, %v0
          if (i >= lastCalculated || !(directOperand.operand instanceof IRVariable var) || !var.isGlobal) {
            operands[i] = directOperand.operand;
            continue;
          }
        }
        // for calculated operand and non-global target, 
        // try to use the target to store one operand saving one virtual register
        // only allowed if both integer type or both string type
        // string sizes do not matter cause string operations can only enlarge the size
        if (target != null && target.type.getClass() == usedOperand.type.getClass() && usedOperand.canUseTarget) {
          usedOperand.assignTo(target, false);
          operands[i] = target;
          target = null;
          continue;
        }
        // need to use next virtual register
        IRVariable virtualRegister = popVirtualRegister(usedOperand.type);
        virtualRegisters.add(virtualRegister);
        usedOperand.assignTo(virtualRegister, true);
        operands[i] = virtualRegister;
      }
      // after all operands are handled put back the virtual registers
      virtualRegisters.forEach(IRGenerator.this::pushVirtualRegister);
      // create the instructions with the generated operands
      IRInstruction instr = createInstruction.apply(operands);
      // possibly savedTarget==null for a ProcedureStatement or compare expression
      if (savedTarget != null)
        ((AssignmentInstruction) instr).target = savedTarget;

      addInstr(instr);
    }
  }
}
