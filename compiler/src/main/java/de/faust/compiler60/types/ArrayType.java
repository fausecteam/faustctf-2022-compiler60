package de.faust.compiler60.types;

import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;

public class ArrayType extends Type {

  public final List<Long> sizes;

  public ArrayType(List<Long> sizes) {
    this.sizes = Collections.unmodifiableList(sizes);
  }

  public Type getBaseType() {
    return IntType.INSTANCE;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;
    ArrayType arrayType = (ArrayType) o;
    return sizes.equals(arrayType.sizes);
  }

  @Override
  public String toString() {
    return getBaseType() + "[" + sizes.stream().map(Object::toString).collect(Collectors.joining("][")) + "]";
  }
}
