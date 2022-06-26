package de.faust.compiler60.types;

public final class VoidType extends Type {

  public static final VoidType INSTANCE = new VoidType();

  private VoidType() {
  }

  @Override
  public boolean isReturnType() {
    return true;
  }

  @Override
  public boolean equals(Object o) {
    return o == this;
  }

  @Override
  public String toString() {
    return "'VOID'";
  }
}
