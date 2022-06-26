package de.faust.compiler60.semantics;

import de.faust.compiler60.CompilationException;
import de.faust.compiler60.ast.*;
import de.faust.compiler60.types.*;

import java.util.List;

public class TypeAnalysis extends ASTVisitor<Type> {

  public void performTypeAnalysis(Program program) {
    visit(program.block);
  }

  private Type visitExpression(Expression e) {
    if (e instanceof Identifier
            || e instanceof IntConstant
            || e instanceof StringConstant) {
      return e.getType();
    }
    Type t = e.accept(this);
    e.setType(t);
    return t;
  }

  @Override
  public Type visit(Identifier identifier) {
    return identifier.getType();
  }

  @Override
  public Type visit(IntConstant intConstant) {
    return intConstant.getType();
  }

  @Override
  public Type visit(StringConstant stringConstant) {
    return stringConstant.getType();
  }

  @Override
  public Type visit(BlockStatement block) {
    // use two passes for ProcedureTypes
    // 1. create ProcedureTypes
    for (Declaration declaration : block.declarations) {
      declaration.accept(this);
    }
    // 2. descent into procedure bodies
    for (Declaration declaration : block.declarations) {
      if (declaration instanceof ProcedureDeclaration procedureDeclaration) {
        visit(procedureDeclaration.procedureBody);
      }
    }
    for (Statement statement : block.statements) {
      statement.accept(this);
    }
    return null;
  }

  @Override
  public Type visit(StringDeclaration stringDeclaration) {
    if (stringDeclaration.getType().size == 0) {
      throw new CompilationException("String declaration of size 0", stringDeclaration);
    }
    return null;
  }

  @Override
  public Type visit(ProcedureDeclaration procedureDeclaration) {
    List<Type> params = procedureDeclaration.parameters.stream()
            .map(e -> {
              if (!e.getType().isParameterType()) {
                throw new CompilationException("Not a parameter type " + e.getType(), e);
              }
              return e.getType();
            })
            .toList();

    if (!procedureDeclaration.returnType.isReturnType()) {
      throw new CompilationException("Not a return type " + procedureDeclaration.returnType, procedureDeclaration);
    }
    if (procedureDeclaration.returnVariable != null &&
            !procedureDeclaration.returnType.equals(procedureDeclaration.returnVariable.getType())) {
      throw new IllegalStateException();
    }
    procedureDeclaration.type = new ProcedureType(procedureDeclaration.returnType, params);
    return null;
  }

  @Override
  public Type visit(AssignmentStatement assignmentStatement) {
    Type lValueType = visitExpression(assignmentStatement.target);
    Type rValueType = visitExpression(assignmentStatement.expression);

    if (!rValueType.canBeAssigned()) {
      throw new CompilationException("Illegal rvalue type for assignment: " + rValueType, assignmentStatement.expression);
    }
    if (!lValueType.canBeAssigned()) {
      throw new CompilationException("Illegal lvalue type for assignment: " + lValueType, assignmentStatement.target);
    }

    if (lValueType instanceof StringType lStringT && rValueType instanceof StringType rStringT) {
      if (lStringT.size < rStringT.size) {
        throw new CompilationException("Illegal string assignment: Left side (" + lStringT
                + ") is smaller than right side (" + rStringT + ")", assignmentStatement);
      }
    } else if (rValueType != lValueType) {
      throw new CompilationException("Incompatible assignment types. Found: " + rValueType
              + " required: " + lValueType, assignmentStatement);
    }
    return null;
  }

  @Override
  public Type visit(ForStatement forStatement) {
    Type t;
    if ((t = visitExpression(forStatement.loopVar)) != IntType.INSTANCE) {
      throw new CompilationException("Variable in for loop must be of type INTEGER but is " + t, forStatement.loopVar);
    }
    for (ForStatement.ForListElement forListElement : forStatement.forList) {
      if (forListElement instanceof ForStatement.SingleValue singleValue) {
        if ((t = visitExpression(singleValue.value)) != IntType.INSTANCE) {
          throw new CompilationException("Expression for loop variable must be of type INTEGER but is " + t, singleValue.value);
        }
      } else if (forListElement instanceof ForStatement.StepUntil stepUntil) {
        if ((t = visitExpression(stepUntil.startValue)) != IntType.INSTANCE) {
          throw new CompilationException("Step start value in for loop must be of type INTEGER but is " + t, stepUntil.startValue);
        }
        // stepUntil.step instanceof IntConstant;
        if ((t = visitExpression(stepUntil.until)) != IntType.INSTANCE) {
          throw new CompilationException("Until value in for loop must be of type INTEGER but is " + t, stepUntil.until);
        }
      } else if (forListElement instanceof ForStatement.While w) {
        if ((t = visitExpression(w.updateExpr)) != IntType.INSTANCE) {
          throw new CompilationException("While update expression must be of type INTEGER but is " + t, w.updateExpr);
        }
        w.whileCond.accept(this);
      }
    }
    forStatement.body.accept(this);
    return null;
  }

  @Override
  public Type visit(TernaryExpression ternaryExpression) {
    ternaryExpression.cond.accept(this);
    Type trueType = visitExpression(ternaryExpression.trueCase);
    Type falseType = visitExpression(ternaryExpression.falseCase);
    if (!trueType.canBeAssigned()) {
      throw new CompilationException("Illegal type in ternary expression: " + trueType, ternaryExpression.trueCase);
    }
    if (!falseType.canBeAssigned()) {
      throw new CompilationException("Illegal type in ternary expression: " + falseType, ternaryExpression.falseCase);
    }
    if (trueType instanceof StringType tStringT && falseType instanceof StringType fStringT) {
      return tStringT.size > fStringT.size ? tStringT : fStringT;
    } else if (!trueType.equals(falseType)) {
      throw new CompilationException("Types in ternary expression do not match: " + trueType + " != " + falseType, ternaryExpression);
    }
    assert trueType == IntType.INSTANCE;
    return trueType;
  }

  @Override
  public Type visit(BinaryExpression binaryExpression) {
    Type leftType = visitExpression(binaryExpression.left);
    Type rightType = visitExpression(binaryExpression.right);
    if (!leftType.canBeAssigned()) {
      throw new CompilationException("Illegal type in binary expression: " + leftType, binaryExpression.left);
    }
    if (!rightType.canBeAssigned()) {
      throw new CompilationException("Illegal type in binary expression: " + rightType, binaryExpression.right);
    }
    if (leftType instanceof StringType lStringT && rightType instanceof StringType rStringT) {
      if (binaryExpression.op != BinaryExpression.Op.ADD) {
        throw new CompilationException("Illegal operator for strings in binary expression: " + binaryExpression.op, binaryExpression);
      }
      return new StringType(lStringT.size + rStringT.size - 1);
    } else if (!leftType.equals(rightType)) {
      throw new CompilationException("Types in binary expression do not match: " + leftType + " != " + rightType, binaryExpression);
    }
    assert leftType == IntType.INSTANCE;
    return leftType;
  }

  @Override
  public Type visit(ProcedureExpression procedureExpression) {
    Type nameType = visit(procedureExpression.procedureName);
    if (!(nameType instanceof ProcedureType procedureType)) {
      throw new CompilationException("Expected procedure call but " + procedureExpression.procedureName + " is not a procedure",
              procedureExpression.procedureName);
    }
    if (procedureType.parameterTypes.size() != procedureExpression.arguments.size()) {
      throw new CompilationException("Procedure call has " + procedureExpression.arguments.size()
              + " arguments, but expected to have " + procedureType.parameterTypes.size(), procedureExpression);
    }
    List<Expression> arguments = procedureExpression.arguments;
    for (int i = 0; i < arguments.size(); i++) {
      Type argType = visitExpression(arguments.get(i));
      Type paramType = procedureType.parameterTypes.get(i);

      if (paramType instanceof StringType stringParam && argType instanceof StringType stringArg) {
        if (stringParam.size >= stringArg.size) {
          continue;
        }
      } else if (argType.equals(paramType)) {
        continue;
      }
      throw new CompilationException("Incompatible " + i + "-th argument types. Found: " + argType
              + " required: " + paramType, arguments.get(i));
    }
    return procedureType.returnType;
  }

  @Override
  public Type visit(ComparisonExpression comparisonExpression) {
    Type leftType = visitExpression(comparisonExpression.left);
    Type rightType = visitExpression(comparisonExpression.right);
    if (leftType != IntType.INSTANCE) {
      throw new CompilationException("Illegal type in compare expression: " + leftType, comparisonExpression.left);
    }
    if (rightType != IntType.INSTANCE) {
      throw new CompilationException("Illegal type in compare expression: " + rightType, comparisonExpression.right);
    }
    return null;
  }

  @Override
  public Type visit(SubscriptedVariable subscriptedVariable) {
    Type type = visitExpression(subscriptedVariable.arrayVar);

    if (type instanceof StringType) {
      if (subscriptedVariable.idxExprs.size() != 1) {
        throw new CompilationException("Subscripted string variable more than 1 dimension, found "
                + subscriptedVariable.idxExprs.size(), subscriptedVariable);
      }

      Expression idxExpr = subscriptedVariable.idxExprs.get(0);
      Type idxType = visitExpression(idxExpr);
      if (idxType != IntType.INSTANCE) {
        throw new CompilationException("Index expression in subscripted string variable not of integer type: " + idxType, idxExpr);
      }
      return IntType.INSTANCE;
    }

    if (!(type instanceof ArrayType arrayType)) {
      throw new CompilationException(subscriptedVariable.arrayVar +
              " expected to be array/string type but is " + type, subscriptedVariable.arrayVar);
    }

    if (arrayType.sizes.size() != subscriptedVariable.idxExprs.size()) {
      throw new CompilationException("Subscripted variable access has " + arrayType.sizes.size() +
              " dimensions, but found " + subscriptedVariable.idxExprs.size(), subscriptedVariable);
    }

    for (Expression idxExpr : subscriptedVariable.idxExprs) {
      Type idxType = visitExpression(idxExpr);
      if (idxType != IntType.INSTANCE) {
        throw new CompilationException("Index expression in subscripted variable not of integer type: " + idxType, idxExpr);
      }
    }
    return arrayType.getBaseType();
  }
}
