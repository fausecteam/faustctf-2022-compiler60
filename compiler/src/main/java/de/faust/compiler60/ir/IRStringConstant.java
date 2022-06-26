package de.faust.compiler60.ir;

import de.faust.compiler60.ast.StringConstant;
import de.faust.compiler60.types.StringType;

public class IRStringConstant extends IROperand {
  public String value;
  public StringType type;

  public IRStringConstant(StringConstant astStringConstant) {
    this.value = astStringConstant.value;
    this.type = astStringConstant.getType();
  }

  @Override
  public StringType getType() {
    return type;
  }

  @Override
  public String toString() {
    return "$\"" + value + '"';
  }
}
