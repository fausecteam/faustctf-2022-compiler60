package de.faust.compiler60.ast;

import java.util.Objects;

public class BinaryBooleanExpression extends BooleanExpression {
  public final BooleanExpression left;
  public final BooleanExpression right;

  public final BinaryBooleanExpression.Op op;

  public BinaryBooleanExpression(int lineNumber, BooleanExpression left, BooleanExpression right, BinaryBooleanExpression.Op op) {
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
    AND,
    OR
  }
}
