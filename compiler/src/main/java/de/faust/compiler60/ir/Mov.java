package de.faust.compiler60.ir;

public class Mov extends AssignmentInstruction {
  public IROperand source;

  public Mov(IRVariable target, IROperand source) {
    super(target);
    this.source = source;
  }

  @Override
  public String mnemonic() {
    return "mov";
  }

  @Override
  public void accept(IRVisitor visitor) {
    visitor.visit(this);
  }

  @Override
  public String toString() {
    return String.format("%s = mov %s", target, source);
  }
}
