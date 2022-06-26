package de.faust.compiler60.ast;

import java.util.Objects;

public class AssignmentStatement extends Statement {
  public Expression target;
  public Expression expression;

  public AssignmentStatement(int lineNumber, Expression target, Expression expression) {
    super(lineNumber);
    this.target = target;
    this.expression = expression;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), target, expression);
  }
}
