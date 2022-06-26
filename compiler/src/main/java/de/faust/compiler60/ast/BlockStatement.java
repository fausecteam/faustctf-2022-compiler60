package de.faust.compiler60.ast;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class BlockStatement extends Statement {

  public final List<Declaration> declarations;

  public final List<Statement> statements;

  public BlockStatement(int lineNumber, List<Declaration> declarations, List<Statement> statements) {
    super(lineNumber);
    this.declarations = new ArrayList<>(declarations);
    this.statements = new ArrayList<>(statements);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), declarations, statements);
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }
}
