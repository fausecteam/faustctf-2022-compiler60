package de.faust.compiler60.ir;

public class Store extends IRInstruction {
  public IRVariable arrayVar;
  public IROperand index;
  public IROperand expression;

  public Store(IRVariable arrayVar, IROperand index, IROperand expression) {
    this.arrayVar = arrayVar;
    this.index = index;
    this.expression = expression;
  }

  @Override
  public String mnemonic() {
    return "store";
  }

  @Override
  public void accept(IRVisitor visitor) {
    visitor.visit(this);
  }

  @Override
  public String toString() {
    return String.format("%s[%s] = store %s", arrayVar, index, expression);
  }
}
