package de.faust.compiler60.ir.gen;

import de.faust.compiler60.ast.*;
import de.faust.compiler60.semantics.Symbol;
import de.faust.compiler60.types.ArrayType;
import de.faust.compiler60.types.IntType;
import de.faust.compiler60.types.StringType;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class MultiDimArrayFlatten extends ASTVisitor<Void> {

  Map<Symbol, ArrayType> oldArrayTypes;

  public void transform(Program program) {
    oldArrayTypes = new HashMap<>();
    visit(program.block);
  }

  @Override
  public Void visit(ArrayDeclaration arrayDeclaration) {
    oldArrayTypes.putIfAbsent(arrayDeclaration.identifier.symbol, arrayDeclaration.type);
    List<Long> sizes = arrayDeclaration.type.sizes;
    if (sizes.size() == 1) {
      return null;
    }

    long flattenedLength = sizes.stream().mapToLong(Long::longValue).reduce(1L, (left, right) -> left * right);
    arrayDeclaration.type = new ArrayType(List.of(flattenedLength));
    return null;
  }

  @Override
  public Void visit(SubscriptedVariable subscriptedVariable) {
    if (subscriptedVariable.arrayVar.getType() instanceof StringType) {
      // string subscripts must have one dimension as TypeAnalysis would fail 
      return super.visit(subscriptedVariable);
    }

    ArrayType oldType = oldArrayTypes.computeIfAbsent(subscriptedVariable.arrayVar.symbol,
            sym -> (ArrayType) sym.getType());
    if (oldType.sizes.size() == 1) {
      return super.visit(subscriptedVariable);
    }

    List<Long> sizes = oldType.sizes;
    // first do the constant indexes
    long constantPart = 0;
    long currDimension = 1;
    for (int i = sizes.size() - 1; i >= 0; i--) {
      Expression idxExpr = subscriptedVariable.idxExprs.get(i);
      if (idxExpr instanceof IntConstant c) {
        constantPart += currDimension * c.value;
      }
      currDimension *= sizes.get(i);
    }
    // then do the non-constant expressions
    Expression resultIdxExpr = constantPart == 0 ? null : new IntConstant(-1, constantPart);
    currDimension = 1;
    for (int i = sizes.size() - 1; i >= 0; i--) {
      Expression idxExpr = subscriptedVariable.idxExprs.get(i);
      if (!(idxExpr instanceof IntConstant)) {
        var scaled = currDimension == 1 ? idxExpr : mul(new IntConstant(-1, currDimension), idxExpr);
        resultIdxExpr = resultIdxExpr == null ? scaled :
                add(scaled, resultIdxExpr);
      }
      currDimension *= sizes.get(i);
    }
    if (resultIdxExpr == null)
      resultIdxExpr = new IntConstant(-1, 0);

    subscriptedVariable.idxExprs.clear();
    subscriptedVariable.idxExprs.add(resultIdxExpr);
    return super.visit(subscriptedVariable);
  }

  static BinaryExpression mul(Expression a, Expression b) {
    assert a.getType() == IntType.INSTANCE;
    assert b.getType() == IntType.INSTANCE;

    BinaryExpression binExpr = new BinaryExpression(-1, a, b, BinaryExpression.Op.MULTIPLY);
    binExpr.setType(IntType.INSTANCE);
    return binExpr;
  }

  static BinaryExpression add(Expression a, Expression b) {
    assert a.getType() == IntType.INSTANCE;
    assert b.getType() == IntType.INSTANCE;

    BinaryExpression binExpr = new BinaryExpression(-1, a, b, BinaryExpression.Op.ADD);
    binExpr.setType(IntType.INSTANCE);
    return binExpr;
  }
}
