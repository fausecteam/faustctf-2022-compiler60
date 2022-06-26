package de.faust.compiler60.ast;

import java.util.Objects;

public class BooleanNotExpression extends BooleanExpression {
  public final BooleanExpression toInvertExpression;

  public BooleanNotExpression(int lineNumber, BooleanExpression toInvertExpression) {
    super(lineNumber);
    this.toInvertExpression = toInvertExpression;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), toInvertExpression);
  }
}
