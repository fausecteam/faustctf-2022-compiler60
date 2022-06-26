grammar Algol60V2;

@parser::header {
import de.faust.compiler60.ast.*;
import de.faust.compiler60.types.*;
import java.util.stream.Collectors;
}

program returns [ Program ast ]
  : block EOF
  { $ast = new Program($block.ast); }
;

block returns [ BlockStatement ast ]
  : BEGIN block_contents END
  { $ast = $block_contents.ast; }
;

block_contents returns [ BlockStatement ast ]
  : ((decls+=declaration | COMMENT) ';')* ';'*
    (stmts+=statement|COMMENT)? ';'* (';' (stmts+=statement|COMMENT))* ';'*
  {
    var decls = $decls.stream()
        .flatMap(x -> x.ast.stream())
        .collect(Collectors.toList());

    $ast = new BlockStatement($ctx.getStart().getLine(),
      decls,
      $stmts.stream().map(x -> x.ast).collect(Collectors.toList()));
  }
;  

declaration returns [ List<Declaration> ast ]
  : type_declaration
  { $ast = $type_declaration.ast; }
  | array_declaration
  { $ast = $array_declaration.ast; }
  | string_declaration
  { $ast = $string_declaration.ast; }
  | procedure_declaration
  { $ast = List.of($procedure_declaration.ast); }
  ;
  
type_declaration returns [ List<Declaration> ast ]
  : INTEGER vars+=identifier (',' vars+=identifier)*
  {
    $ast = $vars.stream()
        .map(x -> new IntVarDeclaration(x.ast.lineNumber, x.ast))
        .collect(Collectors.toList());
  }
;
  
array_declaration returns [ List<Declaration> ast ]
  : INTEGER ARRAY segments+=array_segment (',' segments+=array_segment)*
  {
    $ast = $segments.stream()
        .flatMap(x -> x.ast.stream())
        .collect(Collectors.toList());
  }
 ;
  
array_segment returns [ List<Declaration> ast ]
  : vars+=identifier (',' vars+=identifier)* '[' sizes+=number (',' sizes+=number)* ']'
  {
    var sizes = $sizes.stream().map(x -> x.ast).collect(Collectors.toList());

    $ast = $vars.stream()
        .map(x -> new ArrayDeclaration(x.ast.lineNumber, x.ast, sizes))
        .collect(Collectors.toList());
  }
;

string_declaration returns [ List<Declaration> ast ]
  : STRING segments+=string_segment (',' segments+=string_segment)*
  {
    $ast = $segments.stream()
        .flatMap(x -> x.ast.stream())
        .collect(Collectors.toList());  
  }
;

string_segment returns [ List<Declaration> ast ]
  : vars+=identifier (',' vars+=identifier)* '[' size=number ']'
  {
    var size = $size.ast;

    $ast = $vars.stream()
        .map(x -> new StringDeclaration(x.ast.lineNumber, x.ast, size))
        .collect(Collectors.toList());
  }
;

procedure_declaration returns [ ProcedureDeclaration ast ]
  : (returnType=INTEGER)? PROCEDURE identifier
    ( '(' params=formal_parameter_list ')')? ';' 
    ((decls+=declaration|COMMENT) ';')*
    statement
  {
    var decls = $decls.stream()
        .flatMap(x -> x.ast.stream())
        .collect(Collectors.toList());
    var bodyLine = decls.size() > 0 ? decls.get(0).lineNumber : $statement.ast.lineNumber;
    var procedureBody = new BlockStatement(bodyLine, decls, List.of($statement.ast));

    $ast = new ProcedureDeclaration($ctx.getStart().getLine(),
        $identifier.ast,
        $params.ctx != null ? $params.ast : List.of(),
        $returnType != null ? IntType.INSTANCE : VoidType.INSTANCE,
        procedureBody);
  }
; 
  
formal_parameter_list returns [ List<Identifier> ast ]
  : identifier (parameter_delimiter formal_parameter_list)?
  {
    if ($formal_parameter_list.ctx == null) {
      $ast = new ArrayList<>();
    } else {
      $ast = $formal_parameter_list.ast;
    }
    $ast.add(0, $identifier.ast);
  }   
;

parameter_delimiter
  : ',' | ')' IDENTIFIER ':' '('
;

statement returns [ Statement ast ]
  : unconditional_statement
  { $ast = $unconditional_statement.ast; }
  | conditional_statement
  { $ast = $conditional_statement.ast; }
;

unconditional_statement returns [ Statement ast ]
  : for_statement
  { $ast = $for_statement.ast; }
  | block
  { $ast = $block.ast; }
  | assignment_statement
  { $ast = $assignment_statement.ast; }
  | procedure_statement
  { $ast = $procedure_statement.ast; }
;

assignment_statement returns [ Statement ast ]
  : target=left_part ASSIGN expression
  {
    $ast = new AssignmentStatement($ctx.getStart().getLine(), $target.ast, $expression.ast);
  }
;

left_part returns [ Expression ast ]
  : identifier
  { $ast = $identifier.ast; }
  | subscripted_variable
  { $ast = $subscripted_variable.ast; }
;

procedure_statement returns [ ProcedureStatement ast ]
  : proc_name=identifier ( '(' args=actual_parameter_list ')' )?
  {
    var procedureExpr = new ProcedureExpression(
      $ctx.getStart().getLine(),
      $proc_name.ast, 
      $args.ctx != null ? $args.ast : List.of());

    $ast = new ProcedureStatement(procedureExpr.lineNumber, procedureExpr);
  }  
;

actual_parameter_list returns [ List<Expression> ast ]
  : expression (parameter_delimiter actual_parameter_list)? 
  {
    if ($actual_parameter_list.ctx == null) {
      $ast = new ArrayList<>();
    } else {
      $ast = $actual_parameter_list.ast;
    }
    $ast.add(0, $expression.ast);
  }   
;

conditional_statement returns [ ConditionalStatement ast ]
  : IF boolean_expression 
    THEN thenBlock=unconditional_statement 
    (ELSE elseBlock=statement)?
  {
    $ast = new ConditionalStatement($ctx.getStart().getLine(), $boolean_expression.ast, $thenBlock.ast, 
      $elseBlock.ctx != null ? $elseBlock.ast : null);
  }
;

for_statement returns [ ForStatement ast ]
  : FOR loopVar=identifier ASSIGN for_list DO body=statement
  {
    $ast = new ForStatement($ctx.getStart().getLine(),
      $loopVar.ast, $for_list.ast, $body.ast);
  }
;
  
for_list returns [ List<ForStatement.ForListElement> ast ]
  : elems+=for_list_element (',' elems+=for_list_element)*
  { $ast = $elems.stream().map(x -> x.ast).collect(Collectors.toList()); }
;

for_list_element returns [ ForStatement.ForListElement ast ]
  : firstExpr=expression ((STEP stepNeg='-'? step=number UNTIL until=expression) | (WHILE whileCond=boolean_expression))?
  {
    if ($step.ctx != null) {
      var step = $stepNeg != null ? - $step.ast : $step.ast;
      $ast = new ForStatement.StepUntil($firstExpr.ast, new IntConstant($step.ctx.getStart().getLine(), step), $until.ast);
    } else if ($whileCond.ctx != null) {
      $ast = new ForStatement.While($firstExpr.ast, $whileCond.ast);
    } else {
      $ast = new ForStatement.SingleValue($firstExpr.ast);
    }
  }
;

expression returns [ Expression ast ]
  : simple_expression
  { $ast = $simple_expression.ast; }
  | IF cond=boolean_expression THEN trueCase=simple_expression ELSE falseCase=expression
  {
    $ast = new TernaryExpression($ctx.getStart().getLine(), $cond.ast, $trueCase.ast, $falseCase.ast);
  }
;

simple_expression returns [ Expression ast ]
  : (( '+' | operators+='-' ))? operands+=term
    ( operators+=( '+' | '-' ) operands+=term )*
  {
    if ($operators.size() == $operands.size()) {
      $ast = new IntConstant($operators.get(0).getLine(), 0);
    } else {
      assert $operators.size() == $operands.size() - 1;
      $ast = $operands.remove(0).ast;
    }

    for (int i = 0; i < $operands.size(); i++) {
      var op = $operators.get(i).getText().equals("+") ? BinaryExpression.Op.ADD : BinaryExpression.Op.SUBTRACT;
      $ast = new BinaryExpression($operators.get(i).getLine(), $ast, $operands.get(i).ast, op);
    }
  }
;

term returns [ Expression ast ]
  : operands+=factor
    (operators+=('×' | '÷')  operands+=factor)*
  {
    assert $operators.size() == $operands.size() - 1;
    $ast = $operands.get(0).ast;

    for (int i = 1; i < $operands.size(); i++) {
      var op = $operators.get(i-1).getText().equals("×") ? BinaryExpression.Op.MULTIPLY : BinaryExpression.Op.DIVIDE;
      $ast = new BinaryExpression($operators.get(i-1).getLine(), $ast, $operands.get(i).ast, op);
    }
  }
;

factor returns [ Expression ast ]
  : number
  { $ast = new IntConstant($ctx.getStart().getLine(), $number.ast); }
  | QUOTED_STRING
  { $ast = new StringConstant($ctx.getStart().getLine(), $text); }
  | identifier
  { $ast = $identifier.ast; }
  | function_designator
  { $ast = $function_designator.ast; }
  | subscripted_variable
  { $ast = $subscripted_variable.ast; }
  | '(' expression ')'
  { $ast = $expression.ast; }
;

function_designator returns [ ProcedureExpression ast ]
  : proc_name=identifier '(' args=actual_parameter_list ')'
  {
    $ast = new ProcedureExpression($ctx.getStart().getLine(),
          $proc_name.ast, $args.ast);
  }
;

boolean_expression returns [ BooleanExpression ast ]
  : implication
  { $ast = $implication.ast; }
  | IF cond=boolean_expression THEN thenValue=implication ELSE elseValue=boolean_expression
  {
    $ast = new TernaryBooleanExpression($ctx.getStart().getLine(),
      $cond.ast, $thenValue.ast, $elseValue.ast);
  }
;

implication returns [ BooleanExpression ast ]
  : operands+=boolean_term
    (operators+='⊃' operands+=boolean_term)*
  {
    $ast = $operands.get(0).ast;
    for (int i = 1; i < $operands.size(); i++) {
      $ast = new BooleanNotExpression($operators.get(i-1).getLine(), $ast);
      $ast = new BinaryBooleanExpression($ast.lineNumber, $ast, $operands.get(i).ast, BinaryBooleanExpression.Op.OR);
    }
  }
;
    
boolean_term returns [ BooleanExpression ast ]
  : operands+=boolean_factor
    (operators+='∨' operands+=boolean_factor)*
  {
    $ast = $operands.get(0).ast;
    for (int i = 1; i < $operands.size(); i++) {
      $ast = new BinaryBooleanExpression($operators.get(i-1).getLine(), $ast, $operands.get(i).ast, BinaryBooleanExpression.Op.OR);
    }
  }
;

boolean_factor returns [ BooleanExpression ast ]
  : operands+=boolean_secondary
    (operators+='∧' operands+=boolean_secondary)*
  {
    $ast = $operands.get(0).ast;
    for (int i = 1; i < $operands.size(); i++) {
      $ast = new BinaryBooleanExpression($operators.get(i-1).getLine(), $ast, $operands.get(i).ast, BinaryBooleanExpression.Op.AND);
    }
  }
; 

boolean_secondary returns [ BooleanExpression ast ]
  : (operator='¬')? operand=boolean_primary
  {
    $ast = $operand.ast;
    if ($operator != null) {
      $ast = new BooleanNotExpression($operator.getLine(), $ast);
    }
  }
;

boolean_primary returns [ BooleanExpression ast ]
  : left=simple_expression operator=('<'|'≤'|'='|'≥'|'>'|'≠') right=simple_expression
  {
    var op = switch($operator.getText()) {
      case "<" -> ComparisonExpression.Op.LT;
      case "≤" -> ComparisonExpression.Op.LE;
      case "≥" -> ComparisonExpression.Op.GE;
      case ">" -> ComparisonExpression.Op.GT;
      case "=" -> ComparisonExpression.Op.EQ;
      case "≠" -> ComparisonExpression.Op.NEQ;
      default -> throw new IllegalStateException();
    };
    $ast = new ComparisonExpression($operator.line, $left.ast, $right.ast, op);
  }
  | '(' boolean_expression ')'
  { $ast = $boolean_expression.ast; }
;

subscripted_variable returns [ SubscriptedVariable ast ]
  : arrayVar=identifier '[' idxExprs+=expression (',' idxExprs+=expression)*  ']'
  {
    var idxExprs = $idxExprs.stream().map(x -> x.ast).collect(Collectors.toList());

    $ast = new SubscriptedVariable($ctx.getStart().getLine(), $arrayVar.ast, idxExprs);
  }
;

identifier returns [ Identifier ast ]
  : IDENTIFIER
  {
    $ast = new Identifier($ctx.getStart().getLine(), $text);
  }
  ;

number returns [ long ast ]
  : NUMBER
  {
    $ast = Long.parseLong($text);
  }
;


// ----- Lexer rules ------

ASSIGN: ':' WS '=';

BEGIN: '\'BEGIN\'';
END: '\'END\'';
IF: '\'IF\'';
THEN: '\'THEN\'';
ELSE: '\'ELSE\'';
FOR: '\'FOR\'';
DO: '\'DO\'';
STEP: '\'STEP\'';
UNTIL: '\'UNTIL\'';
WHILE: '\'WHILE\'';
PROCEDURE: '\'PROCEDURE\'';

INTEGER: '\'INTEGER\'';
ARRAY: '\'ARRAY\'';
STRING: '\'STRING\'';

// accept above keywords even with spaces
KEYWORD: '\'' [A-Z \t\r\n]+ '\''
{
  switch (getText().replaceAll("[ \t\r\n]", "")) {
   case "'BEGIN'" -> {setType(BEGIN); yield true;}
   case "'END'" -> {setType(END); yield true;}
   case "'IF'" -> {setType(IF); yield true;}
   case "'THEN'" -> {setType(THEN); yield true;}
   case "'ELSE'" -> {setType(ELSE); yield true;}
   case "'FOR'" -> {setType(FOR); yield true;}
   case "'DO'" -> {setType(DO); yield true;}
   case "'STEP'" -> {setType(STEP); yield true;}
   case "'UNTIL'" -> {setType(UNTIL); yield true;}
   case "'WHILE'" -> {setType(WHILE); yield true;}
   case "'PROCEDURE'" -> {setType(PROCEDURE); yield true;}
   case "'INTEGER'" -> {setType(INTEGER); yield true;}
   case "'ARRAY'" -> {setType(ARRAY); yield true;}
   case "'STRING'" -> {setType(STRING); yield true;}
   default -> false;
 }
}?;

fragment WS: [ \t\r\n]*;
COMMENT : '\'' WS 'C' WS 'O' WS 'M' WS 'M' WS 'E' WS 'N' WS 'T' WS '\''
          ~[;]*;

QUOTED_STRING: '"' (~["] | '\\"')* '"'
{
  String t = getText();
  setText(t.substring(1, t.length() - 1));
};

NUMBER: [0-9][0-9 \t\r\n]*
{
  setText(getText().replaceAll("[ \t\r\n]", ""));
};

IDENTIFIER: [a-zA-Z][a-zA-Z0-9 \t\r\n]*
{
  setText(getText().replaceAll("[ \t\r\n]", ""));
};

SKIP_WS: [ \t\r\n]+ -> skip;
