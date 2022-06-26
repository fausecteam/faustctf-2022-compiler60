package de.faust.compiler60.ir.gen;

import de.faust.compiler60.ast.*;
import de.faust.compiler60.types.IntType;

import java.util.ArrayList;
import java.util.List;
import java.util.ListIterator;

public class ForToWhileLoop extends ASTVisitor<BlockStatement> {

  public void transform(Program program) {
    visit(program.block);
  }

  @Override
  public BlockStatement visit(BlockStatement block) {
    for (Declaration declaration : block.declarations) {
      declaration.accept(this);
    }
    for (ListIterator<Statement> iterator = block.statements.listIterator(); iterator.hasNext(); ) {
      Statement statement = iterator.next();
      if (statement instanceof ForStatement forStatement) {
        iterator.set(visit(forStatement));
      } else {
        statement.accept(this);
      }
    }
    return null;
  }

  @Override
  public BlockStatement visit(ConditionalStatement conditionalStatement) {
    conditionalStatement.condition.accept(this);
    if (conditionalStatement.thenBlock instanceof ForStatement forStatement) {
      conditionalStatement.thenBlock = visit(forStatement);
    } else {
      conditionalStatement.thenBlock.accept(this);
    }
    if (conditionalStatement.elseBlock != null) {
      if (conditionalStatement.elseBlock instanceof ForStatement forStatement) {
        conditionalStatement.elseBlock = visit(forStatement);
      } else {
        conditionalStatement.elseBlock.accept(this);
      }
    }
    return null;
  }

  @Override
  public BlockStatement visit(ForStatement forStatement) {
    Statement body = forStatement.body;
    if (body instanceof ForStatement bodyFor) {
      body = visit(bodyFor);
    } else {
      body.accept(this);
    }
    List<Statement> statements = new ArrayList<>();

    for (ForStatement.ForListElement forListElement : forStatement.forList) {
      if (forListElement instanceof ForStatement.SingleValue singleValue) {
        // assign single value to loop var
        var startValueAssign = new AssignmentStatement(-1, forStatement.loopVar, singleValue.value);
        statements.add(startValueAssign);
        statements.add(body);
      } else if (forListElement instanceof ForStatement.StepUntil stepUntil) {
        // assign start value to loop var
        var startValueAssign = new AssignmentStatement(-1, forStatement.loopVar, stepUntil.startValue);
        statements.add(startValueAssign);

        boolean posStep = stepUntil.step.value > 0;
        // construct while condition: (loopVar - until) <= 0
        BooleanExpression whileCond = new ComparisonExpression(-1, forStatement.loopVar, stepUntil.until,
                posStep ? ComparisonExpression.Op.LE : ComparisonExpression.Op.GE);
        var reducedForBody = new BlockStatement(-1, List.of(), new ArrayList<>());
        var reducedFor = new ForStatement(-1, null, List.of(new ForStatement.While(null, whileCond)), reducedForBody);

        // add loop body and increment
        reducedForBody.statements.add(body);
        var increment = new AssignmentStatement(-1, forStatement.loopVar, add(forStatement.loopVar, stepUntil.step));
        reducedForBody.statements.add(increment);

        statements.add(reducedFor);
      } else if (forListElement instanceof ForStatement.While whileElem) {

        var reducedForBody = new BlockStatement(-1, List.of(), new ArrayList<>());
        var reducedFor = new ForStatement(-1, null, List.of(new ForStatement.While(null, whileElem.whileCond)), reducedForBody);

        // update loopVar before while
        var update = new AssignmentStatement(-1, forStatement.loopVar, whileElem.updateExpr);
        statements.add(update);
        // add loop body and update loopVar in body before condition
        reducedForBody.statements.add(body);
        reducedForBody.statements.add(update);

        statements.add(reducedFor);
      }
    }

    return new BlockStatement(forStatement.lineNumber, List.of(), statements);
  }

  static BinaryExpression add(Expression a, Expression b) {
    assert a.getType() == IntType.INSTANCE;
    assert b.getType() == IntType.INSTANCE;

    BinaryExpression binExpr = new BinaryExpression(-1, a, b, BinaryExpression.Op.ADD);
    binExpr.setType(IntType.INSTANCE);
    return binExpr;
  }
}
