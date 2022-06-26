package de.faust.compiler60.semantics;

import de.faust.compiler60.CompilationException;
import de.faust.compiler60.StdLibrary;
import de.faust.compiler60.ast.*;
import de.faust.compiler60.types.IntType;
import de.faust.compiler60.types.VoidType;

import java.util.*;

public class NameAnalysis extends ExpressionMutateAstVisitor {

  public final Deque<Map<String, Symbol>> symbolTable = new ArrayDeque<>();

  private String currentFunction;
  private Symbol returnVarSymbol;

  public final void performNameAnalysis(Program program) {
    Map<String, Symbol> stdlib = new HashMap<>();
    StdLibrary.addToSymbolTable(stdlib);
    symbolTable.addFirst(stdlib);
    visit(program.block);
    symbolTable.removeFirst();
    assert symbolTable.isEmpty();
  }

  @Override
  public Expression visit(BlockStatement block) {
    symbolTable.addFirst(new HashMap<>(block.declarations.size()));
    // add declarations to symbol table
    for (Declaration declaration : block.declarations) {
      Symbol s = new Symbol(declaration);
      if (symbolTable.getFirst().putIfAbsent(s.getName(), s) != null) {
        throw new CompilationException(s + " already defined", declaration);
      }
    }
    // assign symbols to identifiers
    // and go into procedure bodies
    for (Declaration declaration : block.declarations) {
      declaration.accept(this);
    }
    for (Statement statement : block.statements) {
      statement.accept(this);
    }
    symbolTable.removeFirst();
    return null;
  }

  @Override
  public Expression visit(Identifier identifier) {
    for (Map<String, Symbol> symbolScope : symbolTable) {
      Symbol symbol = symbolScope.get(identifier.name);
      if (symbol == null) {
        continue;
      }
      identifier.symbol = symbol;
      // check if is variable access or procedure call without args
      if (symbol.declaration instanceof ProcedureDeclaration) {
        return new ProcedureExpression(identifier.lineNumber, identifier, List.of());
      }
      return identifier;
    }
    throw new CompilationException("Unknown " + identifier, identifier);
  }

  @Override
  public Expression visit(ProcedureDeclaration procedureDeclaration) {
    if (symbolTable.size() > 2) {
      throw new CompilationException("Nested functions not supported", procedureDeclaration);
    }
    visit(procedureDeclaration.identifier);

    // check if identifiers in parameters are also declared
    Map<String, Declaration> paramsDecls = new HashMap<>();
    for (Identifier parameter : procedureDeclaration.parameters) {
      Optional<Declaration> decl = procedureDeclaration.procedureBody.declarations
              .stream()
              .filter(d -> d.identifier.name.equals(parameter.name))
              .findAny();
      if (decl.isEmpty()) {
        throw new CompilationException("Parameter " + parameter.name + " is in parameter list but not declared", parameter);
      }
      if (paramsDecls.putIfAbsent(parameter.name, decl.get()) != null) {
        throw new CompilationException("Duplicate parameter " + parameter.name + " in parameter list", parameter);
      }
    }
    IntVarDeclaration returnVarDecl = null;
    // add declaration and symbol for return value
    if (procedureDeclaration.returnType == IntType.INSTANCE) {
      currentFunction = procedureDeclaration.identifier.name;
      Identifier returnVarIdentifier = new Identifier(procedureDeclaration.lineNumber, currentFunction + "$returnVar");
      returnVarDecl = new IntVarDeclaration(procedureDeclaration.lineNumber, returnVarIdentifier);
      Optional<Declaration> dupDecl = procedureDeclaration.procedureBody.declarations
              .stream()
              .filter(decl -> decl.identifier.name.equals(currentFunction))
              .findAny();
      if (dupDecl.isPresent()) {
        throw new CompilationException(procedureDeclaration.identifier.name + " already defined " +
                "as implicit return variable", dupDecl.get());
      }
      returnVarSymbol = new Symbol(returnVarDecl);
      returnVarIdentifier.symbol = returnVarSymbol;
      procedureDeclaration.returnVariable = returnVarSymbol;
    } else if (procedureDeclaration.returnType != VoidType.INSTANCE) {
      throw new IllegalStateException();
    }

    visit(procedureDeclaration.procedureBody);

    if (returnVarDecl != null) {
      currentFunction = null;
      returnVarSymbol = null;
      // add returnVar for codegen
      procedureDeclaration.procedureBody.declarations.add(returnVarDecl);
    }

    // assign symbols to parameters
    procedureDeclaration.parameters.forEach(p ->
            p.symbol = paramsDecls.get(p.name).identifier.symbol);

    return null;
  }

  @Override
  public Expression visit(AssignmentStatement assignmentStatement) {
    super.visit(assignmentStatement);
    // check for assignment to return var
    if (assignmentStatement.target instanceof ProcedureExpression procedureExpression) {
      if (procedureExpression.procedureName.name.equals(currentFunction)) {
        Identifier returnVarId = new Identifier(procedureExpression.lineNumber, currentFunction);
        assignmentStatement.target = returnVarId;
        returnVarId.symbol = returnVarSymbol;
      } else {
        throw new CompilationException("Illegal assignment to function " + procedureExpression.procedureName.name, assignmentStatement);
      }
    }
    return null;
  }

  @Override
  public Expression visit(BinaryExpression binaryExpression) {
    binaryExpression.left = binaryExpression.left.accept(this);
    binaryExpression.right = binaryExpression.right.accept(this);
    // replace 0-1 with -1
    if (binaryExpression.op == BinaryExpression.Op.SUBTRACT) {
      if (binaryExpression.left instanceof IntConstant lConst && lConst.value == 0) {
        if (binaryExpression.right instanceof IntConstant rConst && Math.abs(rConst.value) == 1) {
          return new IntConstant(binaryExpression.lineNumber, -rConst.value);
        }
      }
    }
    return binaryExpression;
  }
}
