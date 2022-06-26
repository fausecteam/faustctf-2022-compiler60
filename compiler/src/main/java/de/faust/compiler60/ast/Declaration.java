package de.faust.compiler60.ast;

import de.faust.compiler60.types.Type;

public abstract class Declaration extends ASTNode {

  public final Identifier identifier;

  public Declaration(int lineNumber, Identifier identifier) {
    super(lineNumber);
    this.identifier = identifier;
  }

  public abstract Type getType();
}
