package de.faust.compiler60.types;

public class StringType extends Type implements Comparable<StringType> {
  public final long size;

  public StringType(long size) {
    this.size = size;
  }

  public boolean canBeAssigned() {
    return true;
  }

  @Override
  public int compareTo(StringType o) {
    return (int) (this.size - o.size);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;
    StringType stringType = (StringType) o;
    return stringType.size == this.size;
  }

  @Override
  public String toString() {
    return "'STRING'[" + size + "]";
  }
}
