package de.faust.compiler60.ast;

import de.faust.compiler60.types.ArrayType;

import java.util.List;
import java.util.Objects;

public class ArrayDeclaration extends Declaration {

  public ArrayType type;

  public ArrayDeclaration(int lineNumber, Identifier identifier, List<Long> sizes) {
    super(lineNumber, identifier);
    type = new ArrayType(sizes);
  }

  @Override
  public ArrayType getType() {
    return type;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), type.sizes);
  }
}
