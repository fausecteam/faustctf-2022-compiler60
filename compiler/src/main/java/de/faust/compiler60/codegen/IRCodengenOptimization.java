package de.faust.compiler60.codegen;

import de.faust.compiler60.ir.*;

/*
Some optimizations to make code generation better, like swapping constants right for add/mul

All constants except (-1, 0, 1) are put into .rodata and therefor not removed/folded
 */
public class IRCodengenOptimization extends IRVisitor {
  IRInstructionList currentInstructionList;

  @Override
  public void visit(IRInstructionList instrList) {
    currentInstructionList = instrList;
    super.visit(instrList);
  }

  @Override
  public void visit(BinaryInstruction.Add add) {
    if (add.left instanceof IRIntConstant lConst) {
      if (lConst.value == 0) {
        // replace 0 + x with mov x
        currentInstructionList.replace(new Mov(add.target, add.right), add);
        return;
      }
      // swap constant right
      add.left = add.right;
      add.right = lConst;
      return;
    }
    if (add.right instanceof IRIntConstant rConst && rConst.value == 0) {
      // replace x + 0 with mov x
      currentInstructionList.replace(new Mov(add.target, add.left), add);
    }
    if (add.target.equals(add.right)) {
      // swap target left
      add.right = add.left;
      add.left = add.target;
    }
  }

  @Override
  public void visit(BinaryInstruction.Sub sub) {
    if (sub.right instanceof IRIntConstant rConst) {
      if (rConst.value == 0) {
        // replace x - 0 with mov x
        currentInstructionList.replace(new Mov(sub.target, sub.left), sub);
        return;
      }
      if (Math.abs(rConst.value) == 1) {
        // replace x - (+/-)1 with x +  (-/+)1
        currentInstructionList.replace(new BinaryInstruction.Add(sub.target, sub.left, new IRIntConstant(-rConst.value)), sub);
      }
    }
  }

  @Override
  public void visit(BinaryInstruction.Mul mul) {
    IRIntConstant c;
    IROperand other;
    if (mul.left instanceof IRIntConstant lConst) {
      c = lConst;
      other = mul.right;
    } else if (mul.right instanceof IRIntConstant rConst) {
      c = rConst;
      other = mul.left;
    } else {
      if (mul.target.equals(mul.right)) {
        // swap target left
        mul.right = mul.left;
        mul.left = mul.target;
      }
      return;
    }
    if (c.value == 0) {
      // replace 0 * x with mov 0
      currentInstructionList.replace(new Mov(mul.target, IRIntConstant.ZERO), mul);
      return;
    }
    if (c.value == 1) {
      // replace 1 * x with mov x
      currentInstructionList.replace(new Mov(mul.target, other), mul);
      return;
    }
    if (c.value == -1) {
      // replace -1 * x with sub 0, x
      currentInstructionList.replace(new BinaryInstruction.Sub(mul.target, IRIntConstant.ZERO, other), mul);
      return;
    }
    if (c == mul.left) {
      // swap constant right
      mul.left = mul.right;
      mul.right = c;
    }
  }

  @Override
  public void visit(BinaryInstruction.Div div) {
    if (div.left instanceof IRIntConstant lConst && lConst.value == 0) {
      // replace 0 / x with mov 0
      currentInstructionList.replace(new Mov(div.target, IRIntConstant.ZERO), div);
      return;
    }
    if (div.right instanceof IRIntConstant rConst) {
      if (rConst.value == 0) {
        // replace division by zero with mov 0
        currentInstructionList.replace(new Mov(div.target, IRIntConstant.ZERO), div);
        return;
      }
      if (rConst.value == 1) {
        // replace x / 1 with mov x
        currentInstructionList.replace(new Mov(div.target, div.left), div);
        return;
      }
      if (rConst.value == -1) {
        // replace x / -1 with sub 0, x
        currentInstructionList.replace(new BinaryInstruction.Sub(div.target, IRIntConstant.ZERO, div.left), div);
      }
    }
  }

  @Override
  public void visit(Jump.ConditionalJump conditionalJump) {
    if (conditionalJump.left instanceof IRIntConstant lConst) {
      // swap constant right
      conditionalJump.left = conditionalJump.right;
      conditionalJump.right = lConst;
      conditionalJump.operator = conditionalJump.operator.swapSides();
    }
  }
}
