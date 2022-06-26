package de.faust.compiler60;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.params.provider.Arguments;

import java.io.IOException;
import java.io.UncheckedIOException;
import java.net.URISyntaxException;
import java.net.URL;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.stream.Stream;

public class TestUtils {
  public static Stream<Path> walkAlgolFiles(URL url) throws URISyntaxException, IOException {
    Path path = Path.of(url.toURI());
    Stream<Path> walk = Files.walk(path);
    return walk
            .filter(Files::isRegularFile)   // is a file
            .filter(p -> p.getFileName().toString().endsWith(".a60"));
  }

  public static List<Arguments> getAlgolSourcesFromURL(URL url) throws URISyntaxException, IOException {
    return walkAlgolFiles(url)
            .map(p -> {
              String fileName = p.subpath(p.getNameCount() - 3, p.getNameCount()).toString();
              try {
                return Arguments.of(
                        Files.readString(p),
                        fileName
                );
              } catch (IOException e) {
                throw new UncheckedIOException(e);
              }
            })
            .toList();
  }

  public static List<Arguments> getAlgolSourcesAndOutputFromUrl(URL url) throws URISyntaxException, IOException {
    return walkAlgolFiles(url)
            .map(p -> {
              Path outputFile = p.resolveSibling(p.getFileName().toString().replace(".a60", ".outp"));
              Assertions.assertTrue(Files.exists(outputFile), ".outp for " + p);
              String fileName = p.subpath(p.getNameCount() - 3, p.getNameCount()).toString();
              try {
                return Arguments.of(
                        Files.readString(p),
                        Files.readAllBytes(outputFile),
                        fileName
                );
              } catch (IOException e) {
                throw new UncheckedIOException(e);
              }
            })
            .toList();
  }

  public static int getLineFromFileName(String fileName) {
    assert fileName.endsWith(".a60");
    int idx = fileName.lastIndexOf("-L");
    if (idx == -1) {
      return -1;
    }
    String line = fileName.substring(idx + 2, fileName.length() - ".a60".length());
    return Integer.parseInt(line);
  }
}
