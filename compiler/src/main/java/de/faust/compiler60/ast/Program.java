package de.faust.compiler60.ast;

public class Program {
  public final int programHash;
  
  public final BlockStatement block;

  public Program(BlockStatement block) {
    this.block = block;
    this.programHash = block.hashCode();
  }
}
