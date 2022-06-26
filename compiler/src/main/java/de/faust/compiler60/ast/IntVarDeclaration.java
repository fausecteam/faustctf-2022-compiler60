package de.faust.compiler60.ast;

import de.faust.compiler60.types.IntType;

public class IntVarDeclaration extends Declaration {
  public IntVarDeclaration(int lineNumber, Identifier identifier) {
    super(lineNumber, identifier);
  }

  @Override
  public IntType getType() {
    return IntType.INSTANCE;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return this.getClass().getSimpleName().hashCode();
  }
}
