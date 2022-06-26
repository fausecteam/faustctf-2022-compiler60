package de.faust.compiler60.types;

import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;

public class ProcedureType extends Type {
  public final Type returnType;

  public final List<Type> parameterTypes;

  public ProcedureType(Type returnType, Type... parameterTypes) {
    this(returnType, List.of(parameterTypes));
  }

  public ProcedureType(Type returnType, List<Type> parameterTypes) {
    this.returnType = returnType;
    this.parameterTypes = parameterTypes;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;
    ProcedureType that = (ProcedureType) o;
    return returnType.equals(that.returnType)
            && parameterTypes.equals(that.parameterTypes);
  }

  @Override
  public String toString() {
    return returnType + "(" + parameterTypes.stream().map(Objects::toString).collect(Collectors.joining(",")) + ")";
  }
}
