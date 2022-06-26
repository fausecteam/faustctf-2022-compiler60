package de.faust.compiler60.ir;

import de.faust.compiler60.types.IntType;
import de.faust.compiler60.types.Type;

public class IRIntConstant extends IROperand {
  public static final IRIntConstant ZERO = new IRIntConstant(0);

  public long value;

  public IRIntConstant(long value) {
    this.value = value;
  }

  @Override
  public Type getType() {
    return IntType.INSTANCE;
  }

  @Override
  public String toString() {
    return "$" + value;
  }
}
