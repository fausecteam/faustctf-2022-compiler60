package de.faust.compiler60.ast;

import de.faust.compiler60.types.StringType;

import java.util.Objects;

public class StringDeclaration extends Declaration {
  public StringType type;

  public StringDeclaration(int lineNumber, Identifier identifier, long size) {
    super(lineNumber, identifier);
    this.type = new StringType(size);
  }

  @Override
  public StringType getType() {
    return type;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), type.size);
  }
}
