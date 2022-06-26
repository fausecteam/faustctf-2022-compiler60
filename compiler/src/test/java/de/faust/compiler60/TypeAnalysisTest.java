package de.faust.compiler60;

import de.faust.compiler60.semantics.NameAnalysis;
import de.faust.compiler60.semantics.TypeAnalysis;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

public class TypeAnalysisTest {

  @ParameterizedTest(name = "{1}")
  @MethodSource
  void testTypeAnalysisPrograms(String sourceCode, String fileName) {
    Algol60Compiler compiler = new Algol60Compiler(sourceCode);

    new NameAnalysis().performNameAnalysis(compiler.program);
    new TypeAnalysis().performTypeAnalysis(compiler.program);
  }

  static List<Arguments> testTypeAnalysisPrograms() throws URISyntaxException, IOException {
    return TestUtils.getAlgolSourcesFromURL(TypeAnalysisTest.class.getResource("/programs"));
  }

  @ParameterizedTest(name = "{1}")
  @MethodSource
  void testTypeAnalysisFailing(String sourceCode, String fileName) {
    Algol60Compiler compiler = new Algol60Compiler(sourceCode);

    new NameAnalysis().performNameAnalysis(compiler.program);

    CompilationException compilationException = assertThrows(CompilationException.class, () -> {
      new TypeAnalysis().performTypeAnalysis(compiler.program);
    });

    assertEquals(TestUtils.getLineFromFileName(fileName), compilationException.lineNumber);
  }

  static List<Arguments> testTypeAnalysisFailing() throws IOException, URISyntaxException {
    return TestUtils.getAlgolSourcesFromURL(TypeAnalysisTest.class.getResource("/types"));
  }
}
