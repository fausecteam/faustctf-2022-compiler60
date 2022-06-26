package de.faust.compiler60.ast;

import java.util.Objects;

public class ProcedureStatement extends Statement {
  public final ProcedureExpression procedureExpression;

  public ProcedureStatement(int lineNumber, ProcedureExpression procedureExpression) {
    super(lineNumber);
    this.procedureExpression = procedureExpression;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), procedureExpression);
  }
}
