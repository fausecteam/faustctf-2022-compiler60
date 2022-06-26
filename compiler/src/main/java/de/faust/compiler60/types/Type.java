package de.faust.compiler60.types;

public abstract class Type {

  public boolean isReturnType() {
    return false;
  }

  public boolean isParameterType() {
    return false;
  }

  public boolean canBeAssigned() {
    return false;
  }
}
