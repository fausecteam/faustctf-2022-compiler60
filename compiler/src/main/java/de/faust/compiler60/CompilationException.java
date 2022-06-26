package de.faust.compiler60;

import de.faust.compiler60.ast.ASTNode;

public class CompilationException extends RuntimeException {
  public final int lineNumber;

  public CompilationException(String message, int lineNumber) {
    super(message);
    this.lineNumber = lineNumber;
  }
  
  public CompilationException(String message) {
    super(message);
    lineNumber = -1;
  }

  public CompilationException(String message, ASTNode offendingNode) {
    super(message + " @ line " + offendingNode.lineNumber);
    lineNumber = offendingNode.lineNumber;
  }
}
