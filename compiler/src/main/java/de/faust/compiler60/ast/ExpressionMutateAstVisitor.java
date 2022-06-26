package de.faust.compiler60.ast;

import java.util.List;

public abstract class ExpressionMutateAstVisitor extends ASTVisitor<Expression> {
  @Override
  public Expression visit(AssignmentStatement assignmentStatement) {
    assignmentStatement.target = assignmentStatement.target.accept(this);
    assignmentStatement.expression = assignmentStatement.expression.accept(this);
    return null;
  }

  @Override
  public Expression visit(ForStatement forStatement) {
    visit(forStatement.loopVar);
    for (ForStatement.ForListElement forListElement : forStatement.forList) {
      if (forListElement instanceof ForStatement.SingleValue singleValue) {
        singleValue.value = singleValue.value.accept(this);
      } else if (forListElement instanceof ForStatement.StepUntil stepUntil) {
        stepUntil.startValue = stepUntil.startValue.accept(this);
        visit(stepUntil.step);
        stepUntil.until = stepUntil.until.accept(this);
      } else if (forListElement instanceof ForStatement.While w) {
        w.updateExpr = w.updateExpr.accept(this);
        w.whileCond.accept(this);
      }
    }
    forStatement.body.accept(this);
    return null;
  }

  @Override
  public Expression visit(TernaryExpression ternaryExpression) {
    ternaryExpression.cond.accept(this);
    ternaryExpression.trueCase = ternaryExpression.trueCase.accept(this);
    ternaryExpression.falseCase = ternaryExpression.falseCase.accept(this);
    return ternaryExpression;
  }

  @Override
  public Expression visit(BinaryExpression binaryExpression) {
    binaryExpression.left = binaryExpression.left.accept(this);
    binaryExpression.right = binaryExpression.right.accept(this);
    return binaryExpression;
  }

  @Override
  public Expression visit(ProcedureExpression procedureExpression) {
    visit(procedureExpression.procedureName);
    List<Expression> arguments = procedureExpression.arguments;
    for (int i = 0; i < arguments.size(); i++) {
      Expression argument = arguments.get(i);
      arguments.set(i, argument.accept(this));
    }
    return procedureExpression;
  }

  @Override
  public Expression visit(ComparisonExpression comparisonExpression) {
    comparisonExpression.left = comparisonExpression.left.accept(this);
    comparisonExpression.right = comparisonExpression.right.accept(this);
    return null;
  }

  @Override
  public Expression visit(SubscriptedVariable subscriptedVariable) {
    visit(subscriptedVariable.arrayVar);
    List<Expression> idxExprs = subscriptedVariable.idxExprs;
    for (int i = 0; i < idxExprs.size(); i++) {
      Expression idxExpr = idxExprs.get(i);
      idxExprs.set(i, idxExpr.accept(this));
    }
    return subscriptedVariable;
  }

  @Override
  public Expression visit(IntConstant intConstant) {
    return intConstant;
  }

  @Override
  public Expression visit(StringConstant stringConstant) {
    return stringConstant;
  }

  @Override
  public Expression visit(Identifier identifier) {
    return identifier;
  }
}
