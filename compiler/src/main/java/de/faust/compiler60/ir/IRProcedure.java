package de.faust.compiler60.ir;

import de.faust.compiler60.ast.ProcedureDeclaration;

import java.util.ArrayList;
import java.util.List;

public class IRProcedure {
  public final String name;

  public final List<IRVariable> parameters;
  public final List<IRVariable> localVariables;

  public final IRInstructionList body;

  public IROperand returnValue;

  public IRProcedure(ProcedureDeclaration decl) {
    name = decl.identifier.name;
    parameters = new ArrayList<>(decl.parameters.size());
    localVariables = new ArrayList<>();
    body = new IRInstructionList();
  }
}
