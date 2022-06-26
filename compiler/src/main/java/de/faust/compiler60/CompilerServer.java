package de.faust.compiler60;

import com.sun.net.httpserver.*;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.charset.CodingErrorAction;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.concurrent.Executors;

import static java.net.HttpURLConnection.*;

public class CompilerServer {
  static {
    System.setProperty("sun.net.httpserver.maxReqTime", "3");
    System.setProperty("sun.net.httpserver.maxRspTime", "3");
  }

  public static final int PORT = 6061;

  public static final int MAX_SOURCE_SIZE = 16 * 1024;

  public static final Path GUI_PATH = Path.of("/app/gui.html");

  public static final String CORS_ORIGIN = "Access-Control-Allow-Origin";
  public static final String CONTENT_TYPE = "Content-Type";
  public static final String CACHE_CONTROL = "Cache-Control";
  public static final String CACHE_2_HOURS = "max-age=7200";
  public static final String TEXT_PLAIN = "text/plain";
  public static final String JSON = "application/json;charset=UTF-8";

  static class CompilerHandler implements HttpHandler {
    @Override
    public void handle(HttpExchange exchange) throws IOException {
      byte[] reqBody = exchange.getRequestBody().readNBytes(MAX_SOURCE_SIZE + 1);
      if (reqBody.length > MAX_SOURCE_SIZE) {
        throw new CompilationException("Source to large: larger than " + MAX_SOURCE_SIZE + " bytes");
      }
      // make sure to throw on utf-8 error
      String sourceCode = StandardCharsets.UTF_8.newDecoder()
              .onUnmappableCharacter(CodingErrorAction.REPORT)
              .onMalformedInput(CodingErrorAction.REPORT)
              .decode(ByteBuffer.wrap(reqBody))
              .toString();

      byte[] binary = new Algol60Compiler(sourceCode).compile();

      SignedElf signedElf = new SignedElf(binary);

      byte[] resp = signedElf.getAsJSON().getBytes(StandardCharsets.UTF_8);
      exchange.getResponseHeaders().add(CONTENT_TYPE, JSON);
      exchange.sendResponseHeaders(HTTP_OK, resp.length);
      exchange.getResponseBody().write(resp);
    }
  }
  
  static class ContentTypeFilter extends Filter {
    final String type;

    public ContentTypeFilter(String type) {
      this.type = type;
    }

    @Override
    public void doFilter(HttpExchange exchange, Chain chain) throws IOException {
      if (!type.equals(exchange.getRequestHeaders().getFirst(CONTENT_TYPE))) {
        byte[] body = ("Expected " + CONTENT_TYPE + ": " + type).getBytes(StandardCharsets.UTF_8);
        exchange.sendResponseHeaders(HTTP_UNSUPPORTED_TYPE, body.length);
        exchange.getResponseBody().write(body);
        return;
      }
      chain.doFilter(exchange);
    }

    @Override
    public String description() {
      return "Content-Type method must be " + type;
    }
  }

  static class CORSAwareRequestMethodFilter extends Filter {
    final String method;
    final String corsConfig;

    public CORSAwareRequestMethodFilter(String method, String corsConfig) {
      this.method = method;
      this.corsConfig = corsConfig;
    }

    @Override
    public void doFilter(HttpExchange exchange, Chain chain) throws IOException {
      if (exchange.getRequestMethod().equals("OPTIONS")) {
        exchange.getResponseHeaders().add(CORS_ORIGIN, this.corsConfig);
        exchange.sendResponseHeaders(HTTP_NO_CONTENT, 0);
        return;
      }
      if (exchange.getRequestMethod().equals(this.method)) {
        exchange.getResponseHeaders().add(CORS_ORIGIN, this.corsConfig);
        chain.doFilter(exchange);
        return;
      }

      byte[] body = ("Only " + method + " allowed").getBytes(StandardCharsets.UTF_8);
      exchange.sendResponseHeaders(HTTP_BAD_METHOD, body.length);
      exchange.getResponseBody().write(body);
    }

    @Override
    public String description() {
      return "Request method must be " + method;
    }
  }

  static class ExactPathFilter extends Filter {
    @Override
    public void doFilter(HttpExchange exchange, Chain chain) throws IOException {
      if (!exchange.getHttpContext().getPath().equals(exchange.getRequestURI().getPath())) {
        exchange.sendResponseHeaders(HTTP_NOT_FOUND, -1);
        return;
      }
      chain.doFilter(exchange);
    }

    @Override
    public String description() {
      return "Path must match HttpContext";
    }
  }
  
  static class ExceptionHandler extends Filter {
    @Override
    public void doFilter(HttpExchange exchange, Chain chain) throws IOException {
      try (exchange) {
        try {
          chain.doFilter(exchange);
        } catch (CompilationException compilationException) {
          String body;
          if (compilationException.lineNumber != -1) {
            body = String.format("""
                    {
                      "error": "%s",
                      "line": %d
                    }
                    """, convertToJSONString(compilationException.getMessage()), compilationException.lineNumber);
          } else {
            body = String.format("""
                    {
                      "error": "%s"
                    }
                    """, convertToJSONString(compilationException.getMessage()));
          }
          byte[] bodyBytes = body.getBytes(StandardCharsets.UTF_8);
          exchange.getResponseHeaders().add(CONTENT_TYPE, JSON);
          exchange.sendResponseHeaders(HTTP_BAD_REQUEST, bodyBytes.length);
          exchange.getResponseBody().write(bodyBytes);
        } catch (Exception e) {
          e.printStackTrace();
          byte[] body = String.format("""
                          {
                            "error": "Uncaught Exception: %s"
                          }
                          """, convertToJSONString(e.toString()))
                  .getBytes(StandardCharsets.UTF_8);
          exchange.getResponseHeaders().add(CONTENT_TYPE, JSON);
          exchange.sendResponseHeaders(HTTP_INTERNAL_ERROR, body.length);
          exchange.getResponseBody().write(body);
        }
      }
    }

    @Override
    public String description() {
      return "Catches exceptions";
    }
  }

  private static String convertToJSONString(String in) {
    return in.replace("\n", "\\n")
            .replace("\"", "\\\"");
  }

  private static class GuiServeHandler implements HttpHandler {
    @Override
    public void handle(HttpExchange exchange) throws IOException {
      byte[] body;
      try {
        body = Files.readAllBytes(GUI_PATH);
      } catch (IOException e) {
        exchange.sendResponseHeaders(HTTP_NOT_FOUND, -1);
        return;
      }
      exchange.getResponseHeaders().add(CACHE_CONTROL, CACHE_2_HOURS);
      exchange.sendResponseHeaders(HTTP_OK, body.length);
      exchange.getResponseBody().write(body);
    }
  }

  public static void main(String[] args) throws IOException {
    HttpServer server = HttpServer.create(new InetSocketAddress(PORT), 0);
    server.setExecutor(Executors.newFixedThreadPool(8));

    HttpContext compileCtx = server.createContext("/compile", new CompilerHandler());
    compileCtx.getFilters().add(new ExceptionHandler());
    compileCtx.getFilters().add(new ExactPathFilter());
    compileCtx.getFilters().add(new CORSAwareRequestMethodFilter("POST", "*"));
    compileCtx.getFilters().add(new ContentTypeFilter(TEXT_PLAIN));

    server.createContext("/", new GuiServeHandler());

    server.start();
    System.out.println("Server started on port " + PORT);
  }
}
