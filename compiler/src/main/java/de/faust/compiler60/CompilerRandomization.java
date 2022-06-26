package de.faust.compiler60;

import de.faust.compiler60.ast.Program;
import de.faust.compiler60.ir.IRProgram;

import java.io.IOException;
import java.io.PrintWriter;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.*;
import java.util.function.BiConsumer;

public final class CompilerRandomization {
  private CompilerRandomization() {
  }

  public static final Path TEAM_NUM_PATH = Path.of("/etc/team-num");
  public static final String DEFAULT_TEAM = "foobar";

  private static String getTeamNum() {
    if (!Files.isReadable(TEAM_NUM_PATH)) {
      return DEFAULT_TEAM;
    }
    try {
      return Files.readString(TEAM_NUM_PATH);
    } catch (IOException e) {
      return DEFAULT_TEAM;
    }
  }

  public static void setupRandomization(Program astProgram, IRProgram irProgram) {
    
    // The Compiler must still be deterministic.
    // Checker will check this. 
    long compilationSeed = ((long)getTeamNum().hashCode() << 32L) + astProgram.programHash;
    
    irProgram.compilerRandomization = new Random(compilationSeed);
    
    randomizeDeclarationOrder(irProgram);
  }
  
  public static void randomizeDeclarationOrder(IRProgram irProgram) {
    Collections.shuffle(irProgram.globals, irProgram.compilerRandomization);
    Collections.shuffle(irProgram.procedures, irProgram.compilerRandomization);
  }
  
  public static void randomNopPadding(IRProgram irProgram, PrintWriter writer) {
    int count = irProgram.compilerRandomization.nextInt(10);
    writer.print("nop\n".repeat(count));
  }
  
  public static <K, V> void randomConstantsOrder(IRProgram irProgram, Map<K, V> constantsMap, BiConsumer<K, V> forEach) {
    List<K> keys = new ArrayList<>(constantsMap.keySet());
    Collections.shuffle(keys, irProgram.compilerRandomization);
    keys.forEach(k -> forEach.accept(k, constantsMap.get(k)));
  }
}
