package de.faust.compiler60.ir;

import java.util.ArrayList;
import java.util.List;
import java.util.Random;

public class IRProgram {
  public final List<IRVariable> globals;
  public final List<IRProcedure> procedures;

  public Random compilerRandomization;

  public IRProgram() {
    globals = new ArrayList<>();
    procedures = new ArrayList<>();
  }
}
