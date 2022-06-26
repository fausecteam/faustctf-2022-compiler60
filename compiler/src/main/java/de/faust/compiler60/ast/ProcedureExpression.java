package de.faust.compiler60.ast;

import java.util.List;
import java.util.Objects;

public class ProcedureExpression extends Expression {

  public final Identifier procedureName;

  public final List<Expression> arguments;

  public ProcedureExpression(int lineNumber, Identifier procedureName, List<Expression> arguments) {
    super(lineNumber);
    this.procedureName = procedureName;
    this.arguments = arguments;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), procedureName, arguments);
  }
}
