package de.faust.compiler60.ast;

import java.util.Objects;

public class TernaryExpression extends Expression {

  public final BooleanExpression cond;

  public Expression trueCase;
  public Expression falseCase;

  public TernaryExpression(int lineNumber, BooleanExpression cond, Expression trueCase, Expression falseCase) {
    super(lineNumber);
    this.cond = cond;
    this.trueCase = trueCase;
    this.falseCase = falseCase;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), cond, trueCase, falseCase);
  }
}
