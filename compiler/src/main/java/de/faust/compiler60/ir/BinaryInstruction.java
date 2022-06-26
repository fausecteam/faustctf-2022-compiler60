package de.faust.compiler60.ir;

public abstract class BinaryInstruction extends AssignmentInstruction {
  public IROperand left;
  public IROperand right;

  public BinaryInstruction(IRVariable target, IROperand left, IROperand right) {
    super(target);
    this.left = left;
    this.right = right;
  }

  @Override
  public String toString() {
    return String.format("%s = %s %s, %s", target, mnemonic(), left, right);
  }

  public static class Add extends BinaryInstruction {
    public Add(IRVariable target, IROperand left, IROperand right) {
      super(target, left, right);
    }

    @Override
    public String mnemonic() {
      return "add";
    }

    @Override
    public void accept(IRVisitor visitor) {
      visitor.visit(this);
    }
  }

  public static class Sub extends BinaryInstruction {
    public Sub(IRVariable target, IROperand left, IROperand right) {
      super(target, left, right);
    }

    @Override
    public String mnemonic() {
      return "sub";
    }

    @Override
    public void accept(IRVisitor visitor) {
      visitor.visit(this);
    }
  }

  public static class Mul extends BinaryInstruction {
    public Mul(IRVariable target, IROperand left, IROperand right) {
      super(target, left, right);
    }

    @Override
    public String mnemonic() {
      return "mul";
    }


    @Override
    public void accept(IRVisitor visitor) {
      visitor.visit(this);
    }
  }

  public static class Div extends BinaryInstruction {
    public Div(IRVariable target, IROperand left, IROperand right) {
      super(target, left, right);
    }

    @Override
    public String mnemonic() {
      return "div";
    }

    @Override
    public void accept(IRVisitor visitor) {
      visitor.visit(this);
    }
  }
}
