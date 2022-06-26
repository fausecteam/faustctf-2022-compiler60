package de.faust.compiler60.ir;

import java.util.Arrays;
import java.util.stream.Collectors;

public class Call extends AssignmentInstruction {
  public String name;
  public IROperand[] args;

  public Call(IRVariable target, String name, IROperand[] args) {
    super(target);
    this.name = name;
    this.args = args;
  }

  @Override
  public String mnemonic() {
    return "call";
  }

  @Override
  public void accept(IRVisitor visitor) {
    visitor.visit(this);
  }

  @Override
  public String toString() {
    return String.format("%s = call %s(%s)", target, name, Arrays.stream(args).map(IROperand::toString).collect(Collectors.joining(", ")));
  }
}
