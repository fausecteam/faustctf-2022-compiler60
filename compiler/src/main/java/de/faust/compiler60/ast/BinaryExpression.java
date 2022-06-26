package de.faust.compiler60.ast;

import java.util.Objects;

public class BinaryExpression extends Expression {
  public Expression left;
  public Expression right;

  public final Op op;

  public BinaryExpression(int lineNumber, Expression left, Expression right, Op op) {
    super(lineNumber);
    this.left = left;
    this.right = right;
    this.op = op;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), left, right);
  }
  
  public enum Op {
    ADD,
    SUBTRACT,
    MULTIPLY,
    DIVIDE
  }
}
