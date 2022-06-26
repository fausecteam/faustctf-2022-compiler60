package de.faust.compiler60.ir;

public abstract class IRVisitor {
  public void visit(IRProgram program) {
    for (IRProcedure procedure : program.procedures) {
      visit(procedure);
    }
  }

  public void visit(IRProcedure procedure) {
    visit(procedure.body);
  }

  public void visit(IRInstructionList instrList) {
    for (IRInstruction instruction : instrList.instructions) {
      instruction.accept(this);
    }
  }

  public void visit(AssignmentInstruction assignmentInstruction) {
  }

  public void visit(BinaryInstruction binaryInstruction) {
    visit((AssignmentInstruction) binaryInstruction);
  }

  public void visit(BinaryInstruction.Add add) {
    visit((BinaryInstruction) add);
  }

  public void visit(BinaryInstruction.Sub sub) {
    visit((BinaryInstruction) sub);
  }

  public void visit(BinaryInstruction.Mul mul) {
    visit((BinaryInstruction) mul);
  }

  public void visit(BinaryInstruction.Div div) {
    visit((BinaryInstruction) div);
  }

  public void visit(Mov mov) {
    visit((AssignmentInstruction) mov);
  }

  public void visit(Load load) {
    visit((AssignmentInstruction) load);
  }

  public void visit(Store store) {
  }

  public void visit(Call call) {
    visit((AssignmentInstruction) call);
  }

  public void visit(Label label) {
  }

  public void visit(Jump jump) {
  }

  public void visit(Jump.UnconditionalJump unconditionalJump) {
    visit((Jump) unconditionalJump);
  }

  public void visit(Jump.ConditionalJump conditionalJump) {
    visit((Jump) conditionalJump);
  }
}
