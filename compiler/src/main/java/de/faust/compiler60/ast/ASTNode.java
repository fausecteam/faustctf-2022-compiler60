package de.faust.compiler60.ast;

public abstract class ASTNode {

  public final int lineNumber;

  public ASTNode(int lineNumber) {
    this.lineNumber = lineNumber;
  }

  /* Double dispatch magic */
  public abstract <T> T accept(ASTVisitor<T> visitor);

  @Override
  public abstract int hashCode();
}
