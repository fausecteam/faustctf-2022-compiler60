package de.faust.compiler60.ir;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

public class IRInstructionList {
  public final List<IRInstruction> instructions = new ArrayList<>();

  public int findIndex(IRInstruction toFind) {
    for (int i = 0; i < instructions.size(); i++) {
      if (instructions.get(i) == toFind) {
        return i;
      }
    }
    throw new IllegalArgumentException("not found");
  }

  public void insertBefore(IRInstruction toInsert, IRInstruction toFind) {
    if (toFind == null) {
      instructions.add(toInsert);
      return;
    }
    instructions.add(findIndex(toFind), toInsert);
  }

  public void replace(IRInstruction toInsert, IRInstruction toFind) {
    instructions.set(findIndex(toFind), toInsert);
  }

  @Override
  public String toString() {
    return instructions.stream().map(IRInstruction::toString).collect(Collectors.joining("\n"));
  }
}
