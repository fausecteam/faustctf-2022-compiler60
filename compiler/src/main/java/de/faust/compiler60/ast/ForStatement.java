package de.faust.compiler60.ast;

import java.util.List;
import java.util.Objects;

public class ForStatement extends Statement {
  public final Identifier loopVar;

  public final List<ForListElement> forList;

  public final Statement body;

  public ForStatement(int lineNumber, Identifier loopVar, List<ForListElement> forList, Statement body) {
    super(lineNumber);
    this.loopVar = loopVar;
    this.forList = forList;
    this.body = body;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), loopVar, forList, body);
  }

  public abstract static class ForListElement {
  }

  public static class SingleValue extends ForListElement {
    public Expression value;

    public SingleValue(Expression value) {
      this.value = value;
    }

    @Override
    public int hashCode() {
      return Objects.hash(this.getClass().getSimpleName(), value);
    }
  }

  public static class StepUntil extends ForListElement {
    public Expression startValue;
    public IntConstant step;
    public Expression until;

    public StepUntil(Expression startValue, IntConstant step, Expression until) {
      this.startValue = startValue;
      this.step = step;
      this.until = until;
    }

    @Override
    public int hashCode() {
      return Objects.hash(this.getClass().getSimpleName(), startValue, until);
    }
  }

  public static class While extends ForListElement {
    public Expression updateExpr;
    public final BooleanExpression whileCond;

    public While(Expression updateExpr, BooleanExpression whileCond) {
      this.updateExpr = updateExpr;
      this.whileCond = whileCond;
    }

    @Override
    public int hashCode() {
      return Objects.hash(this.getClass().getSimpleName(), updateExpr, whileCond);
    }
  }
}
