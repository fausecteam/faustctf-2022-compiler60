package de.faust.compiler60.ir;

import de.faust.compiler60.types.Type;

public class IRVariable extends IROperand {
  public final String name;

  public Type type;

  public final boolean isGlobal;

  public IRVariable(String name, Type type, boolean isGlobal) {
    this.name = name;
    this.type = type;
    this.isGlobal = isGlobal;
  }

  @Override
  public Type getType() {
    return type;
  }

  @Override
  public String toString() {
    return name;
  }
}
