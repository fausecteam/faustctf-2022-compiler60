package de.faust.compiler60.ast;

import de.faust.compiler60.semantics.Symbol;
import de.faust.compiler60.types.Type;

import java.util.Objects;

public class Identifier extends Expression {
  public final String name;

  public Symbol symbol;

  public Identifier(int lineNumber, String name) {
    super(lineNumber);
    this.name = name;
  }

  @Override
  public Type getType() {
    return symbol.getType();
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
    return Objects.hash(this.getClass().getSimpleName(), name);
  }

  @Override
  public String toString() {
    return "Identifier[" + name + "]";
  }
}
