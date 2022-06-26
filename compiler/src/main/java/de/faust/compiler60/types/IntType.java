package de.faust.compiler60.types;

public final class IntType extends Type {

  public static final IntType INSTANCE = new IntType();

  private IntType() {
  }

  @Override
  public boolean isReturnType() {
    return true;
  }

  @Override
  public boolean isParameterType() {
    return true;
  }

  @Override
  public boolean canBeAssigned() {
    return true;
  }

  @Override
  public boolean equals(Object o) {
    return o == this;
  }

  @Override
  public String toString() {
    return "'INTEGER'";
  }
}
