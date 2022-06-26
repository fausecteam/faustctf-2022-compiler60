package de.faust.compiler60.semantics;

import de.faust.compiler60.ast.Declaration;
import de.faust.compiler60.types.Type;

public class Symbol {
  public final Declaration declaration;

  public Symbol(Declaration declaration) {
    this.declaration = declaration;
  }

  public String getName() {
    return declaration.identifier.name;
  }

  public Type getType() {
    return declaration.getType();
  }

  @Override
  public String toString() {
    return "Symbol[" + getName() + "]";
  }
}
