package de.faust.compiler60.ast;

import de.faust.compiler60.types.IntType;
import de.faust.compiler60.types.Type;

public class IntConstant extends Expression {

  public final long value;

  public IntConstant(int lineNumber, long value) {
    super(lineNumber);
    this.value = value;
  }

  @Override
  public IntType getType() {
    return IntType.INSTANCE;
  }

  @Override
  public void setType(Type type) {
    throw new UnsupportedOperationException();
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    // This must not contain the actual int value.
    // Also int constants will be put into .rodata which is not included in signature.   
    // Compilation must be deterministic when changing only int constants.
    // Checker will check this. 
    return this.getClass().getSimpleName().hashCode();
  }

  @Override
  public String toString() {
    return "Int[" + value + "]";
  }
}
