package de.faust.compiler60.ir;

public class Load extends AssignmentInstruction {
  public IRVariable arrayVar;
  public IROperand index;

  public Load(IRVariable target, IRVariable arrayVar, IROperand index) {
    super(target);
    this.arrayVar = arrayVar;
    this.index = index;
  }

  @Override
  public String mnemonic() {
    return "load";
  }

  @Override
  public void accept(IRVisitor visitor) {
    visitor.visit(this);
  }

  @Override
  public String toString() {
    return String.format("%s = load %s[%s]", target, arrayVar, index);
  }
}
