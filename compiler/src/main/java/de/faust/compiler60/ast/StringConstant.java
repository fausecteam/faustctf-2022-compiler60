package de.faust.compiler60.ast;

import de.faust.compiler60.types.StringType;
import de.faust.compiler60.types.Type;

import java.nio.charset.StandardCharsets;
import java.util.Objects;
import java.util.regex.Pattern;

public class StringConstant extends Expression {
  public final String value;
  public final int size;

  public StringConstant(int lineNumber, String value) {
    super(lineNumber);
    this.value = value;
    this.size = getSize();
  }

  private int getSize() {
    // +1 for terminating null byte
    int size = value.getBytes(StandardCharsets.UTF_8).length + 1;
    // subtract escape sequences
    size -= Pattern.compile("\\\\(.)").matcher(value).results().mapToInt(match ->
            switch (match.group(1)) {
              // hex escape: len(\xFF)=4 -> 1 
              case "x" -> 3;
              // octal escape len(\111)=4 -> 1
              case "0", "1", "2", "3", "4", "5", "6", "7" -> 3;
              // other cases like \\, \n, \r...
              default -> 1;
            }).sum();
    return size;
  }

  @Override
  public StringType getType() {
    return new StringType(size);
  }

  @Override
  public void setType(Type type) {
    throw new UnsupportedOperationException();
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), value);
  }

  @Override
  public String toString() {
    return "String[" + value + "]";
  }
}
