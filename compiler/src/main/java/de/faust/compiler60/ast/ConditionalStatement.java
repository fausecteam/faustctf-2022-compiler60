package de.faust.compiler60.ast;

import java.util.Objects;

public class ConditionalStatement extends Statement {

  public final BooleanExpression condition;
  public Statement thenBlock;
  public Statement elseBlock;

  public ConditionalStatement(int lineNumber, BooleanExpression condition, Statement thenBlock, Statement elseBlock) {
    super(lineNumber);
    this.condition = condition;
    this.thenBlock = thenBlock;
    this.elseBlock = elseBlock;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), condition, thenBlock, elseBlock);
  }
}
