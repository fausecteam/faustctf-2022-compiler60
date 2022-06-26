package de.faust.compiler60;

import de.faust.compiler60.ast.Identifier;
import de.faust.compiler60.ast.ProcedureDeclaration;
import de.faust.compiler60.semantics.Symbol;
import de.faust.compiler60.types.IntType;
import de.faust.compiler60.types.ProcedureType;
import de.faust.compiler60.types.StringType;
import de.faust.compiler60.types.VoidType;

import java.util.List;
import java.util.Map;
import java.util.Set;

public class StdLibrary {
  public static final List<Symbol> stdLibraryFunction = List.of(
          createLibraryFunction("exit", new ProcedureType(VoidType.INSTANCE, IntType.INSTANCE)),

          createLibraryFunction("outinteger", new ProcedureType(VoidType.INSTANCE, IntType.INSTANCE)),
          createLibraryFunction("outchar", new ProcedureType(VoidType.INSTANCE, IntType.INSTANCE)),
          createLibraryFunction("outstring", new ProcedureType(VoidType.INSTANCE, new StringType(1024))),

          createLibraryFunction("integer2string", new ProcedureType(new StringType(21), IntType.INSTANCE)),

          createLibraryFunction("openRW", new ProcedureType(IntType.INSTANCE, new StringType(1024))),
          createLibraryFunction("openRO", new ProcedureType(IntType.INSTANCE, new StringType(1024))),
          createLibraryFunction("openWOConfidential", new ProcedureType(IntType.INSTANCE, new StringType(41))),

          createLibraryFunction("readchar", new ProcedureType(IntType.INSTANCE, IntType.INSTANCE)),
          createLibraryFunction("readstring", new ProcedureType(new StringType(128), IntType.INSTANCE)),
          createLibraryFunction("writechar", new ProcedureType(VoidType.INSTANCE, IntType.INSTANCE, IntType.INSTANCE)),
          createLibraryFunction("writestring", new ProcedureType(VoidType.INSTANCE, IntType.INSTANCE, new StringType(128)))
  );

  public static final Set<String> stringReturningFunctions = Set.of("$strcpy", "$strcat", "integer2string", "readstring");

  private static Symbol createLibraryFunction(String name, ProcedureType type) {
    ProcedureDeclaration declaration = new ProcedureDeclaration(-1, new Identifier(-1, name), null, type.returnType, null);
    declaration.type = type;
    return new Symbol(declaration);
  }

  public static void addToSymbolTable(Map<String, Symbol> symbolTable) {
    for (Symbol symbol : stdLibraryFunction) {
      symbolTable.put(symbol.getName(), symbol);
    }
  }
}
