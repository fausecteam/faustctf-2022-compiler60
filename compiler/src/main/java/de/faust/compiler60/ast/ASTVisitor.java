package de.faust.compiler60.ast;

public abstract class ASTVisitor<T> {
  public T visit(BlockStatement block) {
    for (Declaration declaration : block.declarations) {
      declaration.accept(this);
    }
    for (Statement statement : block.statements) {
      statement.accept(this);
    }
    return null;
  }

  public T visit(IntVarDeclaration intVarDeclaration) {
    visit(intVarDeclaration.identifier);
    return null;
  }

  public T visit(ArrayDeclaration arrayDeclaration) {
    visit(arrayDeclaration.identifier);
    return null;
  }

  public T visit(StringDeclaration stringDeclaration) {
    visit(stringDeclaration.identifier);
    return null;
  }

  public T visit(ProcedureDeclaration procedureDeclaration) {
    visit(procedureDeclaration.identifier);
    for (Identifier parameter : procedureDeclaration.parameters) {
      visit(parameter);
    }
    visit(procedureDeclaration.procedureBody);
    return null;
  }

  public T visit(AssignmentStatement assignmentStatement) {
    assignmentStatement.target.accept(this);
    assignmentStatement.expression.accept(this);
    return null;
  }

  public T visit(ProcedureStatement procedureStatement) {
    visit(procedureStatement.procedureExpression);
    return null;
  }

  public T visit(ConditionalStatement conditionalStatement) {
    conditionalStatement.condition.accept(this);
    conditionalStatement.thenBlock.accept(this);
    if (conditionalStatement.elseBlock != null)
      conditionalStatement.elseBlock.accept(this);
    return null;
  }

  public T visit(ForStatement forStatement) {
    visit(forStatement.loopVar);
    for (ForStatement.ForListElement forListElement : forStatement.forList) {
      if (forListElement instanceof ForStatement.SingleValue singleValue) {
        singleValue.value.accept(this);
      } else if (forListElement instanceof ForStatement.StepUntil stepUntil) {
        stepUntil.startValue.accept(this);
        stepUntil.step.accept(this);
        stepUntil.until.accept(this);
      } else if (forListElement instanceof ForStatement.While w) {
        w.updateExpr.accept(this);
        w.whileCond.accept(this);
      }
    }
    forStatement.body.accept(this);
    return null;
  }

  public T visit(TernaryExpression ternaryExpression) {
    ternaryExpression.cond.accept(this);
    ternaryExpression.trueCase.accept(this);
    ternaryExpression.falseCase.accept(this);
    return null;
  }

  public T visit(BinaryExpression binaryExpression) {
    binaryExpression.left.accept(this);
    binaryExpression.right.accept(this);
    return null;
  }

  public T visit(IntConstant intConstant) {
    return null;
  }

  public T visit(StringConstant stringConstant) {
    return null;
  }

  public T visit(ProcedureExpression procedureExpression) {
    visit(procedureExpression.procedureName);
    for (Expression argument : procedureExpression.arguments) {
      argument.accept(this);
    }
    return null;
  }

  public T visit(BinaryBooleanExpression binaryBooleanExpression) {
    binaryBooleanExpression.left.accept(this);
    binaryBooleanExpression.right.accept(this);
    return null;
  }

  public T visit(BooleanNotExpression booleanNotExpression) {
    booleanNotExpression.toInvertExpression.accept(this);
    return null;
  }

  public T visit(TernaryBooleanExpression ternaryBooleanExpression) {
    ternaryBooleanExpression.cond.accept(this);
    ternaryBooleanExpression.trueValue.accept(this);
    ternaryBooleanExpression.falseValue.accept(this);
    return null;
  }

  public T visit(ComparisonExpression comparisonExpression) {
    comparisonExpression.left.accept(this);
    comparisonExpression.right.accept(this);
    return null;
  }

  public T visit(SubscriptedVariable subscriptedVariable) {
    visit(subscriptedVariable.arrayVar);
    for (Expression idxExpr : subscriptedVariable.idxExprs) {
      idxExpr.accept(this);
    }
    return null;
  }

  public T visit(Identifier identifier) {
    return null;
  }
}
