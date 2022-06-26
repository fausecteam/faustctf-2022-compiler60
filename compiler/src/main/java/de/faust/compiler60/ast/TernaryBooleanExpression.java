package de.faust.compiler60.ast;

import java.util.Objects;

public class TernaryBooleanExpression extends BooleanExpression {
  public final BooleanExpression cond;

  public final BooleanExpression trueValue;
  public final BooleanExpression falseValue;

  public TernaryBooleanExpression(int lineNumber, BooleanExpression cond, BooleanExpression trueValue, BooleanExpression falseValue) {
    super(lineNumber);
    this.cond = cond;
    this.trueValue = trueValue;
    this.falseValue = falseValue;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), cond, trueValue, falseValue);
  }
}
