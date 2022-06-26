package de.faust.compiler60;

import de.faust.compiler60.antlr.Algol60V2Lexer;
import de.faust.compiler60.antlr.Algol60V2Parser;
import de.faust.compiler60.ast.Program;
import de.faust.compiler60.codegen.BinaryBuilder;
import de.faust.compiler60.codegen.CodegenX64;
import de.faust.compiler60.ir.IRProgram;
import de.faust.compiler60.ir.gen.IRGenerator;
import de.faust.compiler60.semantics.NameAnalysis;
import de.faust.compiler60.semantics.TypeAnalysis;
import org.antlr.v4.runtime.*;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

public class Algol60Compiler {

  public final Program program;

  public Algol60Compiler(String sourceCode) {
    Algol60V2Lexer lexer = getLexer(sourceCode);

    Algol60V2Parser parser = new Algol60V2Parser(new CommonTokenStream(lexer));

    parser.removeErrorListeners();
    lexer.removeErrorListeners();
    BaseErrorListener listener = new BaseErrorListener() {
      @Override
      public void syntaxError(Recognizer<?, ?> recognizer, Object offendingSymbol, int line, int charPositionInLine, String msg, RecognitionException e) {
        throw new CompilationException(msg + " @ line " + line + ":" + charPositionInLine, line);
      }
    };
    parser.addErrorListener(listener);
    lexer.addErrorListener(listener);

    program = parser.program().ast;
  }

  public byte[] compile() {
    new NameAnalysis().performNameAnalysis(program);
    new TypeAnalysis().performTypeAnalysis(program);

    IRProgram irProgram = new IRGenerator().generateIR(program);

    String assembly = new CodegenX64().generateAssembly(irProgram);

    return BinaryBuilder.createBinary(assembly);
  }

  private static Algol60V2Lexer getLexer(String sourceCode) {
    return new Algol60V2Lexer(CharStreams.fromString(sourceCode));
  }

  public static void main(String[] args) throws IOException {

    // for debugging prints ir and assembly
    String code = """
            'BEGIN'
              'INTEGER' 'PROCEDURE' even(i);
                  'INTEGER' i;
                even := 'IF' i > 1 'THEN' odd(i-1) 'ELSE' 'IF' i = 0 'THEN' 1 'ELSE' 0
              ;

              'INTEGER' 'PROCEDURE' odd(i);
                  'INTEGER' i;
                odd := 'IF' i > 1 'THEN' even(i-1) 'ELSE' 'IF' i = 1 'THEN' 1 'ELSE' 0
              ;

              'PROCEDURE' test(i);
                  'INTEGER' i;
              'BEGIN'
                 outinteger(i);
                 outstring(" is even: ");
                 outinteger(even(i));
                 outstring(" odd: ");
                 outinteger(odd(i));
                 outchar(10);
              'END';

              test(0);
              test(1);

              test(33);
              test(27);
              test(24);

              test(256);
              test(257);
            'END'
            """;

    Algol60Compiler compiler = new Algol60Compiler(code);
    new NameAnalysis().performNameAnalysis(compiler.program);
    new TypeAnalysis().performTypeAnalysis(compiler.program);

    IRProgram irProgram = new IRGenerator().generateIR(compiler.program);

    String assembly = new CodegenX64().generateAssembly(irProgram);
    System.out.println("Intermediate representation:");
    System.out.println(irProgram.procedures.get(0).body);
    System.out.println("x64 assembly:");
    System.out.println(assembly);

    byte[] binary = BinaryBuilder.createBinary(assembly);

    Files.write(Path.of(".", "a.out"), binary);
  }
}
