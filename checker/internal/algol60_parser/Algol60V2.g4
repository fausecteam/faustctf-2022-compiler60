grammar Algol60V2;

@lexer::header {
import "regexp";
}

program : block EOF;

block : BEGIN block_contents END;

block_contents : (decls+=declaration ';')* ';'*
    stmts+=statement? ';'* (';' stmts+=statement)* ';'*
;

declaration : type_declaration
  | array_declaration
  | string_declaration
  | procedure_declaration
  | COMMENT
;
  
type_declaration : INTEGER vars+=identifier (',' vars+=identifier)* ;
  
array_declaration : INTEGER ARRAY segments+=array_segment (',' segments+=array_segment)* ;
  
array_segment : vars+=identifier (',' vars+=identifier)* '[' sizes+=number (',' sizes+=number)* ']' ;

string_declaration : STRING segments+=string_segment (',' segments+=string_segment)* ;

string_segment : vars+=identifier (',' vars+=identifier)* '[' size=number ']' ;

procedure_declaration : (returnType=INTEGER)? PROCEDURE identifier
    ( '(' params=formal_parameter_list ')')? ';' 
    (decls+=declaration ';')*
    statement
; 
  
formal_parameter_list : identifier (parameter_delimiter formal_parameter_list)? ;

parameter_delimiter
  : ',' | ')' IDENTIFIER ':' '('
;

statement : unconditional_statement
  | conditional_statement
  | COMMENT
;

unconditional_statement : for_statement
  | block
  | assignment_statement
  | procedure_statement
;

assignment_statement : (target=left_part ASSIGN) expression ;

left_part : identifier
  | subscripted_variable
;

procedure_statement : proc_name=identifier ( '(' args=actual_parameter_list ')' )? ;

actual_parameter_list : expression (parameter_delimiter actual_parameter_list)? ;

conditional_statement : IF boolean_expression
    THEN thenBlock=unconditional_statement 
    (ELSE elseBlock=statement)?
;

for_statement : FOR loopVar=identifier ASSIGN for_list DO body=statement ;
  
for_list : elems+=for_list_element (',' elems+=for_list_element)* ;

for_list_element : firstExpr=expression
    ((STEP stepNeg='-'? step=number UNTIL until=expression)
    | (WHILE whileCond=boolean_expression))?
;

expression : simple_expression
  | IF cond=boolean_expression THEN trueCase=simple_expression ELSE falseCase=expression
;

simple_expression : (operators+=( '+' | '-' ))? operands+=term
    ( operators+=( '+' | '-' ) operands+=term )*
;

term : operands+=factor
    (operators+=('×' | '÷')  operands+=factor)*
;

factor : number
  | QUOTED_STRING
  | identifier
  | function_designator
  | subscripted_variable
  | '(' expression ')'
;

function_designator : proc_name=identifier '(' args=actual_parameter_list ')' ;

boolean_expression : implication
  | IF cond=boolean_expression THEN thenValue=implication ELSE elseValue=boolean_expression
;

implication : operands+=boolean_term
    (operators+='⊃' operands+=boolean_term)*
;
    
boolean_term : operands+=boolean_factor
    (operators+='∨' operands+=boolean_factor)*
;

boolean_factor : operands+=boolean_secondary
    (operators+='∧' operands+=boolean_secondary)*
; 

boolean_secondary : (operator='¬')? operand=boolean_primary ;

boolean_primary : left=simple_expression operator=('<'|'≤'|'='|'≥'|'>'|'≠') right=simple_expression
  | '(' boolean_expression ')'
;

subscripted_variable : arrayVar=identifier '[' idxExprs+=expression (',' idxExprs+=expression)* ']' ;

identifier : IDENTIFIER ;

number returns [ int64 value ]
  : NUMBER
  {
    val, err := strconv.ParseInt($text, 10, 64)
    if err != nil {
      panic(err)
    }
    $value = val
  }
;


// ----- Lexer rules ------

// we do not support whitespaces in keywords like in the real grammar

ASSIGN: ':' [ \t\r\n]* '=';

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

COMMENT : '\'COMMENT\'' ~[;]*;

// use fixed grammar here
QUOTED_STRING: '"' (~["\\] | ('\\' .))* '"';

NUMBER: [0-9][0-9 \t\r\n]*
{
  this.SetText(regexp.MustCompile("[ \t\r\n]").
      ReplaceAllLiteralString(this.GetText(), ""))
};

IDENTIFIER: [a-zA-Z][a-zA-Z0-9 \t\r\n]*
{
  this.SetText(regexp.MustCompile("[ \t\r\n]").
  	  ReplaceAllLiteralString(this.GetText(), ""))
};

SKIP_WS: [ \t\r\n]+ -> skip;
