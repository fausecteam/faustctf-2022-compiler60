package de.faust.compiler60.codegen;

import de.faust.compiler60.CompilerRandomization;
import de.faust.compiler60.ir.IRIntConstant;
import de.faust.compiler60.ir.IRProgram;
import de.faust.compiler60.ir.IRStringConstant;

import java.io.PrintWriter;
import java.util.HashMap;
import java.util.Map;

/**
 * All constants except (-1, 0, 1) are put into .rodata
 */
public class ConstantPool {
  protected Map<String, Long> constants64Bit = new HashMap<>();
  protected Map<String, String> strings = new HashMap<>();

  public boolean willBePTR(IRIntConstant constant) {
    long value = constant.value;
    return value != -1 && value != 0 && value != 1;
  }

  public String useStringConstant(IRStringConstant constant) {
    String name = strings.computeIfAbsent(constant.value, x -> ".L$S" + strings.size());
    return String.format("QWORD PTR [rip + %s]", name);
  }

  public String useIntConstant(IRIntConstant constant) {
    long value = constant.value;
    if (!willBePTR(constant)) {
      return Long.toString(value);
    }
    String name = ".L$C" + constants64Bit.size();
    constants64Bit.put(name, constant.value);
    return String.format("QWORD PTR [rip + %s]", name);
  }

  public void generateRodata(IRProgram program, PrintWriter writer) {
    writer.println(".section .rodata");
    CompilerRandomization.randomConstantsOrder(program, constants64Bit,
            (name, value) -> writer.printf("""
                    %s=.
                    .quad %d
                    """, name, value));

    writer.println(".section .strings, \"aS\", @progbits");
    CompilerRandomization.randomConstantsOrder(program, strings, 
            (value, name) -> writer.printf("""
                    %s=.
                    .asciz "%s"
                    """, name, value));
  }
}
