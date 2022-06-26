package de.faust.compiler60.ir;

import de.faust.compiler60.ast.ComparisonExpression;

public abstract class Jump extends IRInstruction {
  public Label target;

  public Jump(Label target) {
    this.target = target;
  }

  public static class UnconditionalJump extends Jump {
    public UnconditionalJump(Label target) {
      super(target);
    }

    @Override
    public String mnemonic() {
      return "jmp";
    }

    @Override
    public void accept(IRVisitor visitor) {
      visitor.visit(this);
    }

    @Override
    public String toString() {
      return "jmp " + target.mnemonic();
    }
  }

  public static class ConditionalJump extends Jump {
    public IROperand left;
    public IROperand right;

    public ComparisonExpression.Op operator;

    public ConditionalJump(Label target, IROperand left, IROperand right, ComparisonExpression.Op operator) {
      super(target);
      this.left = left;
      this.right = right;
      this.operator = operator;
    }

    @Override
    public String mnemonic() {
      return switch (operator) {
        case LT -> "jlt";
        case LE -> "jle";
        case GE -> "jge";
        case GT -> "jgt";
        case EQ -> "jeq";
        case NEQ -> "jne";
        case U_B -> "jb";
        case U_BE -> "jbe";
        case U_AE -> "jae";
        case U_A -> "ja";
      };
    }

    @Override
    public void accept(IRVisitor visitor) {
      visitor.visit(this);
    }

    @Override
    public String toString() {
      return String.format("%s %s, %s, %s", mnemonic(), left, right, target.mnemonic());
    }
  }
}
