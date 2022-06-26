package de.faust.compiler60.ast;

import de.faust.compiler60.semantics.Symbol;
import de.faust.compiler60.types.ProcedureType;
import de.faust.compiler60.types.Type;

import java.util.List;
import java.util.Objects;

public class ProcedureDeclaration extends Declaration {

  public final List<Identifier> parameters;

  public final Type returnType;

  public final BlockStatement procedureBody;

  public Symbol returnVariable;

  public ProcedureType type;

  public ProcedureDeclaration(int lineNumber, Identifier identifier,
                              List<Identifier> parameters, Type returnType, BlockStatement procedureBody) {
    super(lineNumber, identifier);
    this.parameters = parameters;
    this.returnType = returnType;
    this.procedureBody = procedureBody;
  }

  @Override
  public ProcedureType getType() {
    return type;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), parameters, returnType, procedureBody);
  }
}
