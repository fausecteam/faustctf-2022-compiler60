package de.faust.compiler60.ast;

import de.faust.compiler60.types.Type;

public abstract class Expression extends ASTNode {
  private Type type;

  public Expression(int lineNumber) {
    super(lineNumber);
  }

  public Type getType() {
    return type;
  }

  public void setType(Type type) {
    this.type = type;
  }
}
