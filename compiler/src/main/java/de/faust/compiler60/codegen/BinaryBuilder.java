package de.faust.compiler60.codegen;

import de.faust.compiler60.CompilationException;

import java.io.BufferedWriter;
import java.io.IOException;
import java.net.URISyntaxException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Path;
import java.util.concurrent.TimeUnit;

public class BinaryBuilder {
  public static final int MAX_ASM_SIZE = 64 * 1024;
  public static final int MAX_BINARY_SIZE = 64 * 1024;

  public static byte[] createBinary(String assembly) {
    try {
      Path buildWrapperPath = Path.of(BinaryBuilder.class.getResource("/assembleAndLink").toURI());
      Path stdLibPath = Path.of(BinaryBuilder.class.getResource("/stdlib.o").toURI());

      // yes, we go from java to c here cause java cannot use the memfd_create syscall
      Process buildWraper = new ProcessBuilder()
              .redirectInput(ProcessBuilder.Redirect.PIPE)
              .redirectError(ProcessBuilder.Redirect.PIPE)
              .redirectOutput(ProcessBuilder.Redirect.PIPE)
              .command(buildWrapperPath.toString(), stdLibPath.toString())
              .start();
      try {
        BufferedWriter writer = buildWraper.outputWriter(StandardCharsets.UTF_8);
        writer.write(assembly);
        writer.close();

        byte[] binary = buildWraper.getInputStream().readNBytes(MAX_BINARY_SIZE + 1);
        if (binary.length > MAX_BINARY_SIZE) {
          throw new CompilationException("Generated binary too large [> " + MAX_BINARY_SIZE + " bytes]");
        }
        if (!buildWraper.waitFor(500, TimeUnit.MILLISECONDS)) {
          throw new CompilationException("Build wrapper timeout");
        }
        if (buildWraper.exitValue() != 0) {
          String err = "";
          while (buildWraper.errorReader().ready()) {
            err += "\n" + buildWraper.errorReader().readLine();
          }
          throw new CompilationException("Build wrapper exit[" + buildWraper.exitValue() + "]: " + err);
        }

        return binary;

      } finally {
        buildWraper.destroyForcibly();
      }
    } catch (IOException | InterruptedException | URISyntaxException e) {
      throw new CompilationException(e.toString());
    }
  }
}
