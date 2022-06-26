package de.faust.compiler60.ast;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class SubscriptedVariable extends Expression {

  public final Identifier arrayVar;
  public final List<Expression> idxExprs;

  public SubscriptedVariable(int lineNumber, Identifier arrayVar, List<Expression> idxExprs) {
    super(lineNumber);
    this.arrayVar = arrayVar;
    this.idxExprs = new ArrayList<>(idxExprs);
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), arrayVar, idxExprs);
  }
}
