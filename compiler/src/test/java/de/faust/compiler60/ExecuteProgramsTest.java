package de.faust.compiler60;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.io.IOException;
import java.net.URISyntaxException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.attribute.PosixFilePermissions;
import java.util.List;
import java.util.concurrent.TimeUnit;

import static org.junit.jupiter.api.Assertions.*;

public class ExecuteProgramsTest {
  @ParameterizedTest(name = "{2}")
  @MethodSource
  void executePrograms(String sourceCode, byte[] expectedOutput, String fileName) throws IOException, InterruptedException {
    Algol60Compiler compiler = new Algol60Compiler(sourceCode);

    byte[] binary = compiler.compile();

    var perms = PosixFilePermissions.asFileAttribute(PosixFilePermissions.fromString("rwxrwx---"));
    Path binaryFile = Files.createTempFile("compiler60-" + Path.of(fileName).getFileName().toString(), ".exe", perms);

    try {
      Files.write(binaryFile, binary);

      assertEquals(binary.length, Files.size(binaryFile));
      assertTrue(Files.isExecutable(binaryFile));

      Process process = new ProcessBuilder()
              .redirectInput(ProcessBuilder.Redirect.PIPE)
              .redirectError(ProcessBuilder.Redirect.DISCARD)
              .redirectOutput(ProcessBuilder.Redirect.PIPE)
              .command("timeout", "-s", "KILL", "2",
                      binaryFile.toString())
              .start();

      try {
        process.getOutputStream().close();

        byte[] stdout = process.getInputStream().readNBytes(1 + expectedOutput.length);

        assertTrue(process.waitFor(100, TimeUnit.MILLISECONDS));
        int expectedExitCode = 0;
        if (fileName.contains("illegal-index")) {
          expectedExitCode = 138;
        }
        assertEquals(expectedExitCode, process.exitValue());
        assertArrayEquals(expectedOutput, stdout);
        assertEquals(-1, process.getInputStream().read());
      } finally {
        process.destroyForcibly();
      }
    } finally {
      Files.deleteIfExists(binaryFile);
    }
  }

  static List<Arguments> executePrograms() throws URISyntaxException, IOException {
    return TestUtils.getAlgolSourcesAndOutputFromUrl(NameAnalysisTest.class.getResource("/programs"));
  }
}
