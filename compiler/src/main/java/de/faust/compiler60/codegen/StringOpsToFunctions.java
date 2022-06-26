package de.faust.compiler60.codegen;

import de.faust.compiler60.StdLibrary;
import de.faust.compiler60.ir.*;
import de.faust.compiler60.types.StringType;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/*
Strings are filled with a null byte at function entry.

String concat and string assignment are implemented using $strcat and $strcpy functions.
So replace the usages with a function call.
Likewise with load and store of strings using the $strget and $strput functions.

Stdlib functions returning a string need the destination as first argument.             
 */
public class StringOpsToFunctions extends IRVisitor {
  IRInstructionList currentInstructionList;

  @Override
  public void visit(IRProcedure procedure) {
    super.visit(procedure);
    // add store for initial null byte of local strings
    for (IRVariable localVariable : procedure.localVariables) {
      if (localVariable.type instanceof StringType) {
        Store nullByte = new Store(localVariable, IRIntConstant.ZERO, IRIntConstant.ZERO);
        procedure.body.instructions.add(0, nullByte);
      }
    }
  }

  @Override
  public void visit(IRInstructionList instrList) {
    currentInstructionList = instrList;
    super.visit(instrList);
  }

  @Override
  public void visit(AssignmentInstruction assignmentInstruction) {
    if (assignmentInstruction.target.type instanceof StringType) {
      throw new IllegalStateException("Assignment to string should only be mov, add or call");
    }
  }

  @Override
  public void visit(Mov mov) {
    if (mov.target.type instanceof StringType) {
      Call strcpyCall = new Call(null, "$strcpy", new IROperand[]{mov.target, mov.source});
      currentInstructionList.replace(strcpyCall, mov);
    }
  }

  @Override
  public void visit(BinaryInstruction.Add add) {
    if (add.target.type instanceof StringType) {
      Call strcatCall = new Call(null, "$strcat", new IROperand[]{add.target, add.left, add.right});
      currentInstructionList.replace(strcatCall, add);
    }
  }

  @Override
  public void visit(Load load) {
    if (load.arrayVar.type instanceof StringType) {
      Call strgetCall = new Call(load.target, "$strget", new IROperand[]{load.arrayVar, load.index});
      currentInstructionList.replace(strgetCall, load);
    }
  }

  @Override
  public void visit(Store store) {
    if (store.arrayVar.type instanceof StringType) {
      Call strputCall = new Call(null, "$strput", new IROperand[]{store.arrayVar, store.index, store.expression});
      currentInstructionList.replace(strputCall, store);
    }
  }

  @Override
  public void visit(Call call) {
    if (call.target != null && call.target.type instanceof StringType) {
      if (!StdLibrary.stringReturningFunctions.contains(call.name)) {
        throw new IllegalStateException("Unknown function returning string: " + call.name);
      }
      List<IROperand> args = new ArrayList<>(Arrays.asList(call.args));
      args.add(0, call.target);
      call.args = args.toArray(IROperand[]::new);
      call.target = null;
    }
  }
}
