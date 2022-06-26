package de.faust.compiler60;

import de.faust.compiler60.ast.*;
import de.faust.compiler60.semantics.NameAnalysis;
import de.faust.compiler60.types.IntType;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

public class NameAnalysisTest {
  @Test
  void testNameAnalysisReturnVar() {
    String code = """
            'BEGIN'
              'INTEGER' 'PROCEDURE' foo(a);
                'INTEGER' a;
                foo := a + 1
              ;
              
              foo(1)
            'END'
            """;

    Algol60Compiler compiler = new Algol60Compiler(code);

    new NameAnalysis().performNameAnalysis(compiler.program);

    BlockStatement block = compiler.program.block;
    assertEquals(1, block.declarations.size());
    assertInstanceOf(ProcedureDeclaration.class, block.declarations.get(0));
    ProcedureDeclaration fooDecl = (ProcedureDeclaration) block.declarations.get(0);
    assertSame(IntType.INSTANCE, fooDecl.returnType);
    assertSame(fooDecl, fooDecl.identifier.symbol.declaration);

    assertEquals(2, fooDecl.procedureBody.declarations.size());
    assertSame(fooDecl.parameters.get(0).symbol, fooDecl.procedureBody.declarations.get(0).identifier.symbol);

    assertEquals(fooDecl.returnVariable, fooDecl.procedureBody.declarations.get(1).identifier.symbol);
    assertEquals("foo$returnVar", fooDecl.returnVariable.getName());

    assertEquals(1, fooDecl.procedureBody.statements.size());
    assertInstanceOf(AssignmentStatement.class, fooDecl.procedureBody.statements.get(0));
    AssignmentStatement statement = (AssignmentStatement) fooDecl.procedureBody.statements.get(0);

    assertSame(fooDecl.returnVariable, ((Identifier) statement.target).symbol);
    BinaryExpression binaryExpression = (BinaryExpression) statement.expression;
    assertSame(fooDecl.parameters.get(0).symbol, ((Identifier) binaryExpression.left).symbol);

    assertSame(fooDecl.identifier.symbol, ((ProcedureStatement) block.statements.get(0)).procedureExpression.procedureName.symbol);
  }

  @ParameterizedTest(name = "{1}")
  @MethodSource
  void testNameAnalysisPrograms(String sourceCode, String fileName) {
    Algol60Compiler compiler = new Algol60Compiler(sourceCode);

    new NameAnalysis().performNameAnalysis(compiler.program);
  }

  static List<Arguments> testNameAnalysisPrograms() throws URISyntaxException, IOException {
    return TestUtils.getAlgolSourcesFromURL(NameAnalysisTest.class.getResource("/programs"));
  }

  @ParameterizedTest(name = "{1}")
  @MethodSource
  void testNameAnalysisFailing(String sourceCode, String fileName) {
    Algol60Compiler compiler = new Algol60Compiler(sourceCode);

    CompilationException compilationException = assertThrows(CompilationException.class, () -> {
      new NameAnalysis().performNameAnalysis(compiler.program);
    });

    assertEquals(TestUtils.getLineFromFileName(fileName), compilationException.lineNumber);
  }

  static List<Arguments> testNameAnalysisFailing() throws IOException, URISyntaxException {
    return TestUtils.getAlgolSourcesFromURL(NameAnalysisTest.class.getResource("/names"));
  }
}
