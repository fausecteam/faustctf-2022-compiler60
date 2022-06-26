package de.faust.compiler60.ast;

import java.util.Objects;

public class ComparisonExpression extends BooleanExpression {

  public Expression left;
  public Expression right;

  public final Op op;

  public ComparisonExpression(int lineNumber, Expression left, Expression right, Op op) {
    super(lineNumber);
    this.left = left;
    this.right = right;
    this.op = op;
  }

  @Override
  public <T> T accept(ASTVisitor<T> visitor) {
    return visitor.visit(this);
  }

  @Override
  public int hashCode() {
    return Objects.hash(this.getClass().getSimpleName(), left, right);
  }

  public enum Op {
    LT,   // less than
    LE,   // less equal
    GE,   // greater equal
    GT,   // greater than
    EQ,   // equal
    NEQ,  // not equal
    // unsigned operations
    U_B, // unsigned below
    U_BE,// unsigned below equal
    U_AE,// unsigned above equal
    U_A;  // unsigned above

    public Op invert() {
      return switch (this) {
        case LT -> GE;
        case LE -> GT;
        case GE -> LT;
        case GT -> LE;
        case EQ -> NEQ;
        case NEQ -> EQ;
        case U_B -> U_AE;
        case U_BE -> U_A;
        case U_AE -> U_B;
        case U_A -> U_BE;
      };
    }

    public Op swapSides() {
      return switch (this) {
        case LT -> GT;
        case LE -> GE;
        case GE -> LE;
        case GT -> LT;
        case EQ, NEQ -> this;
        case U_B -> U_A;
        case U_BE -> U_AE;
        case U_AE -> U_BE;
        case U_A -> U_B;
      };
    }
  }
}
