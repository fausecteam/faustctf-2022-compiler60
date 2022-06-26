package de.faust.compiler60.ir;

public abstract class IRInstruction {

  public abstract String mnemonic();

  public abstract void accept(IRVisitor visitor);

  @Override
  public abstract String toString();
}
