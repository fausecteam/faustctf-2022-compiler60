package de.faust.compiler60;

import java.io.IOException;
import java.nio.ByteBuffer;
import java.nio.ByteOrder;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Base64;
import java.util.HexFormat;

public class SignedElf {
  // 2 ticks
  public static final long EXPIRES_AFTER = 2 * 3 * 60 * 1000; 
  
  public static final Path SIGN_KEY_PATH = Path.of("/app/sign_key");
  public static final byte[] DEFAULT_KEY = "foobar".getBytes(StandardCharsets.UTF_8);

  private final byte[] binary;
  private final long expiry;
  private final byte[] signature;

  public SignedElf(byte[] binary) throws IOException {
    this.binary = binary;
    this.expiry = System.currentTimeMillis() + EXPIRES_AFTER;
    signature = signBinary();
  }

  private static byte[] getSignKey() {
    if (!Files.isReadable(SIGN_KEY_PATH)) {
      return DEFAULT_KEY;
    }
    try {
      return Files.readAllBytes(SIGN_KEY_PATH);
    } catch (IOException e) {
      return DEFAULT_KEY;
    }
  }
  
  private static byte[] bytesFromLong(long l) {
    return ByteBuffer
            .allocate(Long.BYTES)
            .order(ByteOrder.BIG_ENDIAN)
            .putLong(l)
            .array();
  }

  private byte[] signBinary() throws IOException {
    ReadElf readElf = new ReadElf(ByteBuffer.wrap(binary));
    int rodataOff = Math.toIntExact(readElf.mRodataOffset);
    int rodataSize = readElf.mRodataSize;
    try {
      MessageDigest messageDigest = MessageDigest.getInstance("SHA3-256");

      // SHA3 does not need HMAC
      messageDigest.update(getSignKey());

      // sign expiry
      messageDigest.update(bytesFromLong(expiry));

      // sign binary except .rodata section (if any)
      if (rodataSize == 0) {
        messageDigest.update(binary);
      } else {
        messageDigest.update(binary, 0, rodataOff);
        messageDigest.update(binary, rodataOff + rodataSize, binary.length - rodataOff - rodataSize);
      }

      return messageDigest.digest();
    } catch (NoSuchAlgorithmException e) {
      throw new IOException(e);
    }
  }

  public String getAsJSON() {
    String binaryBase64 = Base64.getEncoder().encodeToString(binary);
    String signatureHex = HexFormat.of().formatHex(signature);
    return String.format(
            "{\"expiry\": %d, \"binary\": \"%s\", \"signature\": \"%s\"}",
            expiry, binaryBase64, signatureHex);
  }
}
