package de.faust.compiler60.ir;

public abstract class AssignmentInstruction extends IRInstruction {
  public IRVariable target;

  public AssignmentInstruction(IRVariable target) {
    this.target = target;
  }
}
