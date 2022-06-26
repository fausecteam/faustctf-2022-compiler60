package de.faust.compiler60.ir;

public class Label extends IRInstruction {
  public static final Label ILLEGAL_INDEX = new Label(-1) {
    @Override
    public String mnemonic() {
      return "LIllegalIndex";
    }
  };
  
  public final int labelIdx;

  public Label(int labelIdx) {
    this.labelIdx = labelIdx;
  }

  @Override
  public String mnemonic() {
    return "L" + labelIdx;
  }

  @Override
  public void accept(IRVisitor visitor) {
    visitor.visit(this);
  }

  @Override
  public String toString() {
    return "L" + labelIdx + ":";
  }
}
