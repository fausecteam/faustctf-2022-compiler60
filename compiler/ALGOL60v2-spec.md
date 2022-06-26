
# <center>ALGOL60v2 Specification</center>

Based on the Modified Report on the Algorithmic Language ALGOL 60.
> Modified Report on the Algorithmic Language ALGOL 60. The Computer Journal, Vol. 19, No. 4, Nov. 1976, pp. 364—379.

## Contents

- [Formalism for syntactic description](#formalism-for-syntactic-description)
- [Basic symbols, identifiers, numbers, and strings](#basic-symbols-identifiers-numbers-and-strings)
    - [Basic concepts](#basic-concepts)
    - [Identifiers](#identifiers)
    - [Numbers](#numbers)
    - [Strings](#strings)
    - [Values and types](#values-and-types)
- [Expressions](#expressions)
    - [Variables](#variables)
    - [Function designators](#function-designators)
    - [Standard functions](#standard-functions)
    - [Arithmetic expressions](#arithmetic-expressions)
    - [Boolean expressions](#boolean-expressions)
- [Statements](#statements)
    - [Blocks](#blocks)
    - [Assignment statements](#assignment-statements)
    - [Conditional statements](#conditional-statements)
    - [For statements](#for-statements)
    - [Procedure statements](#procedure-statements)
- [Declarations](#declarations)
    - [Type declarations](#type-declarations)
    - [Array declarations](#array-declarations)
    - [String declarations](#string-declarations)
    - [Procedure declarations](#procedure-declarations)

### Formalism for syntactic description

The syntactic rules of the language will be described using the original BNF syntax. The original report was the first
to use that syntax. Now follows the syntax definition from the original report.

Sequences of characters enclosed in the bracket `<>` represent metalinguistic variables whose values are sequences of
symbols. The marks `::=` and `|` (the latter with the meaning of 'or') are metalinguistic connectives. Any mark in a
formula, which is not a variable or a connective, denotes itself (or the class of marks which are similar to it).
Juxtaposition of marks and/or variables in a formula signifies juxtaposition of the sequences denoted.

## Basic symbols, identifiers, numbers, and strings

### Basic concepts

Source code of the language is to be interpreted as UTF-8 string. All whitespace characters (space, tabulator, newline
and carriage return) in it are ignored, except when enclosed in double quotes `"`.

All keywords are in uppercase and enclosed in single quotes `'`. The remaining symbols are built up from the following
basic symbols:

```
<letter> ::= a | b | c | d | e | f | g | h | i | j | k | l |
             m | n | o | p | q | r | s | t | u | v | w | x | y | z | A |
             B | C | D | E | F | G | H | I | J | K | L | M | N | O | P |
             Q | R | S | T | U | V | W | X | Y | Z

<digit> ::= 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
```

### Identifiers

#### Syntax

`<identifier> ::= <letter> | <identifier>  <letter> | <identifier> <digit>`

#### Semantics

Identifiers have no inherent meaning, but serve for the identification of simple variables, arrays, strings and
procedures. They may be chosen freely. Identifiers also act as formal parameters of procedures, in which capacity they
represent a simple variable.

The same identifier cannot be used to denote two different entities except when these entities have disjoint scopes as
defined by the declarations of the program. This rule applies also to the formal parameters of procedures.

### Numbers

#### Syntax

`<number> ::= <digit> | <number> <digit>`

#### Semantics

Decimal numbers have their conventional meaning.

### Strings

#### Syntax

`<open string> ::= <any sequence of characters not containing " > | <empty>`

`<string> ::= " <open string> "`

#### Semantics

In order to enable the language to handle sequences of characters string literals enclosed in double quotes `"` are
introduced. Characters in a string might be any valid unicode codepoint. The type of a string literal will be a string
with the size of the byte length of it's UTF-8 representation plus a terminating null byte.

The following escape sequences are defined:

```
\"     Represents one " character.
\\     Represents one \ character.
\r     Represents a carriage return; A byte with value 13.
\n     Represents a newline; A byte with value 10.
\t     Represents a horizontal tab; A byte with value 9.
\ <0-8> <0-8> <0-8>
       A byte with the numeric value of the 3 octal digits.
\x <hex-digit> <hex-digit>
       A byte with the numeric value of the 2 hex digits (upper or lowercase).
```

### Values and types

Certain syntactic units are said to possess values. These values will in general change during the execution of the
program. The values of expressions and their constituents are defined in Section [Expressions](#expressions).

The INTEGER type denotes a 64-bit signed number.

The STRING type denotes a sequence of bytes (8-bit) of fixed size.

The INTEGER ARRAY type denotes a sequence of numbers with INTEGER type of fixed size. The type might be organized in
multiple dimensions.

## Expressions

### Variables

#### Syntax

`<subscript list> ::= <expression> |
<subscript list> , <expression>`

`<subscripted variable> ::= <identifier> [ <subscript list> ]`

`<variable> ::= <identifier> | <subscripted variable>`

#### Semantics

A variable is a designation given to a single value. This value may be used in expressions for forming other values and
may be changed at will by means of assignment statements. The type of the value of a particular variable is defined in
the declaration for the variable itself or for the corresponding array identifier .

#### Subscripts

Subscripted variable designate values which are components of multidimensional arrays or one-dimensional strings. Each
expression of the subscript list occupies one subscript position of the subscripted variable and is called a subscript.
The complete list of subscripts is enclosed in the subscript brackets [ ]. The array component referred to by a
subscripted variable is specified by the actual numerical value of its subscripts. The count of subscripts must match
the dimension of the variable type.

For an integer array the subscript expression will be of integer type. The value of the subscripted integer array is
defined only if the value of the subscript expression is within the subscript bounds of the array.

When the subscripted variable is a string the single subscript refers to the byte-position in the UTF-8 representation
of the string. The value of a subscript expression of a string variable is the unsigned byte value at this position
converted to an integer. An access with an index beyond the first terminating null byte of a string is illegal and will
result in the termination of the program with exit code 138. For an assignment to a subscripted string even
the index of the terminating null byte is illegal.

#### Initial values of variables

The value of a integer or integer array variable, not declared in the top-most block, is undefined from entry into the
block in which it is declared until an assignment is made to it. The value of a variable declared in the top-most block
is zero.

The value of a string variable is a string of length zero from entry into the block in which it is declared until an
assignment is made to it.

### Function designators

#### Syntax

```
<parameter delimiter> ::= , | ) <identifier> : (

<actual parameter list> ::= <expression> | <actual parameter list> <parameter delimiter> <expression>

<function designator> ::= <identifier> | <identifier> ( <actual parameter list> )
```

#### Semantics

Function designators define integer or string values which result through the application of given sets of rules defined
by a procedure declaration to fixed sets of actual parameters. The rules governing specification of actual parameters
and execution are given in Section [Procedure statements](#procedure-statements).

### Standard functions

Certain standard functions are declared in the environmental block with the following procedure declarations. Defining
strings as parameters or return value of procedures is only allowed for standard functions and not possible in a
program:

```
'PROCEDURE' exit(code);
  'INTEGER' code;
  'COMMENT' Exits the execution with the given code.;
  ...
;
'PROCEDURE' outchar(c);
  'INTEGER' c;
  'COMMENT' Writes the lower 8-bit of c to stdout;
  writechar(1, c)
;
'PROCEDURE' outinteger(i);
  'INTEGER' i;
  'COMMENT' Writes the decimal representation of i to stdout;
  ...
;
'PROCEDURE' outstring(s);
  'STRING' s[1024];
  'COMMENT' Writes the bytes of s to stdout until a terminating null byte (which is not written);
  writestring(1, s);
;
'STRING'[21] 'PROCEDURE' integer2string(i);
  'INTEGER' i;
  'COMMENT' Returns the decimal representation of i as string;
  ...
;
'INTEGER' 'PROCEDURE' openRW(passwd);
  'STRING' passwd[1024];
  'COMMENT' Opens or creates the file corresponding to the passwd in read-write mode and returns the 
  file descriptor number. The passwd is hashed using SHA1 to get the filename. If the file exist and you do not 
  have write permissions  (cause you are not the creator) a negative value is returned (see man open for details).;
  ...
;
'INTEGER' 'PROCEDURE' openRO(passwd);
  'STRING' passwd[1024];
  'COMMENT' Opens the file corresponding to the passwd in read-only mode and returns the file descriptor number.
   The passwd is hashed using SHA1 to get the filename. If the file does not exist a negative value is returned 
   (see man open for details).;
  ...
;
'INTEGER' 'PROCEDURE' openWOConfidential(filename);
  'STRING' filename[41];
  'COMMENT' Opens the file with name filename in write-only mode and returns the file descriptor number.
   You will only be able to read the file contents later, if the filename is a hex-encoded SHA1-hash of some known value.
   If the file does already exist and you are not the creator a negative value is returned 
   (see man open for details).;  
  ...
;
'INTEGER' 'PROCEDURE' readchar(fd);
  'INTEGER' fd;
  'COMMENT' Read one byte from the file descriptor fd and return it unsigned. If EOF is reached -1 is returned.;
  ...
;
'STRING'[128] 'PROCEDURE' readstring(fd);
  'INTEGER' fd;
  'COMMENT' Reads bytes from the the file descriptor fd until a null-byte or 127 bytes are read or EOF is reached
    or an error occurs. The resulting string is null terminated.;
  ...
;
'PROCEDURE' writechar(fd, c);
  'INTEGER' fd, c;
  'COMMENT' Writes the lower 8-bit of c to the file descriptor fd;
  ...
;
'PROECEDURE' writestring(fd, s);
  'INTEGER' fd;
  'STRING' s[128];
  'COMMENT' Writes the bytes of s to the file descriptor fd until a terminating null byte (which is not written);
  ...
;
```

### Arithmetic expressions

#### Syntax

```
<adding operator> ::= + | -

<multiplying operator> ::= × | ÷

<factor> ::= <number> | <string> | <variable>
        <function designator> | ( <expression> )

<term> ::= <factor> | <term> <multiplying operator> <factor>

<simple expression> ::= <term> | 
        <adding operator> <term> |
        <simple expression> <adding operator> <term>

<expression> ::= <simple expression> |
        'IF' <Boolean expression> 'THEN' <simple expression> 'ELSE' <expression>
```

#### Semantics

An expression is a rule for computing a numerical value or a string. In the case of simple expressions this value is
obtained by executing the indicated arithmetic operations on the actual numerical values of the primaries of the
expression, as explained the next [Section](#operators-and-types) below. The actual numerical value of a primary is
obvious in the case of numbers. For variables it is the current value (assigned last in the dynamic sense), and for
function designators it is the value arising from the computing rules defining the procedure when applied to the current
values of the procedure parameters given in the expression. Finally, for expression enclosed in parenthesis the value
must through a recursive analysis be expressed in terms of the values of primaries of the other three kinds.

In the more general expressions, which include if clauses, one out of several simple expressions is selected on the
basis of the actual values of the Boolean expressions. This selection is made as follows: The Boolean expressions of the
if clauses are evaluated one by one in sequence from left to right until one having the value true is found. The value
of the expression is then the value if the first expression following this Boolean. If none of the Boolean expressions
has the value true, then the value of the expression is the value of expression following the final else.

The order of evaluation of primaries within an expression is left to right. (WTF in the original it was undefined
behavior)

In evaluating an expression, it is understood that all the factors within that expression are evaluated, except those
within any expression that is governed by an if clause but not selected by it.

#### Operators and types

The type of an expression of the form  
`'IF' B 'THEN' T 'ELSE' E`    
is the type of `T` and `E`, which must both be integer type or both be string type. In the case of strings
(e.g. `T: STRING[a]` and `E: STRING[b]`) the resulting type will be the one with the larger size
(e.g. `STRING[max(a,b)]`).

Apart from the Boolean expressions of clauses, the constituents of simple expressions must be of integer type or string
type. The meaning of the basic operators and the types of the expressions to which they lead are given by the following
rules:

- The operators +, -, and × have the conventional meaning (addition, subtraction, and multiplication) for integer
  operands. The type of the expression will be integer.
- The operator + for string operands (e.g. `STRING[a]` and `STRING[b]`) concatenates the two strings. The bytes of the
  UTF-8 representation of the first string are copied to the destination stopping at a null byte, which is not copied.
  Into the following positions in the destination the second string is copied until a null byte was copied. The type of
  the expression will be a string with the sum of the string sizes minus one (e.g. `STRING[a+b-1]`).
- The operator ÷ denotes integer division. The operation is undefined if the divisor has the value zero or for (-2^63)
  ÷(-1). The fractional result will be rounded to the next integer towards zero.

#### Precedence of operators

```
1st:   + -
2nd:   × ÷
```

### Boolean expressions

#### Syntax

```
<relational operator> ::= < | ≤ | = | ≥ | > | ≠

<Boolean primary> ::= ( <Boolean expression> ) |
        <simple expression> <relational operator> <simple expression>

<Boolean secondary> ::= <Boolean primary> |
        ¬ <Boolean primary>

<Boolean factor> ::= <Boolean secondary> |
        <Boolean factor> ∧ <Boolean secondary>

<Boolean term> ::= <Boolean factor> |
        <Boolean term> ∨ <Boolean factor>

<implication> ::= <Boolean term> |
        <implication> ⊃ <Boolean term>

<Boolean expression> ::= <implication> | 
        'IF' <Boolean expression> 'THEN' <implication> 'ELSE' <Boolean expression>
```

#### Semantics

A Boolean expression is a rule for computing a logical value.

The relational operators <, ≤, =, ≥, > and ≠ have their conventional meaning (less than, less than or equal to, equal
to, greater than or equal to, greater than, not equal to). The expressions must be of integer type. Relations take on
the value true whenever the corresponding relation is satisfied for the expressions involved, otherwise false. The order
of evaluation is left to right.

The logical operators ¬ (not), ∧ (and), ∨ (or) and ⊃ (implies) have their conventional meaning. They are evaluated in
order left to right until the logical value of the whole expression can no longer change even if the expression was not
fully evaluated (short-circuit semantic).

#### Precedence of operators

```
1st:   < ≤ = ≥ > ≠
2nd:   ¬
3rd:   ∧
4th:   ∨
5th:   ⊃
```

## Statements

The units of operation within the language are called statements. They will normally be executed consecutively as
written. However, this sequence of operations may be shortened by conditional statements, which may cause certain
statements to be skipped, and lengthened by for statements which cause certain statements to be repeated.

### Blocks

#### Syntax

```
<unconditional statement> ::= <for statement> | <block> | 
        <assignment statement> | <procedure statement> 

<statement> ::= 'COMMENT' <any sequence of characters not containing ;>  |
        <unconditional statement> | <conditional statement>

<compound tail> ::= <statement> | <statement> ; <compound tail>

<block head> ::= <declaration> | <block head> ; <declaration>

<block contents> ::= <block head> ; <compound tail>

<block> ::= 'BEGIN' block contents 'END'

<program> ::= <block>
```

#### Semantics

Every block automatically introduces a new level of nomenclature. This is realised as follows: Any identifier occurring
within the block may through a suitable declaration be specified to be local to the block in question. This means (a)
that the entity specified by this identifier inside the block has no existence outside it and (b) that any entity
represented by this identifier outside the block is completely inaccessible inside the block.

Identifiers occurring within a block and not being declared to this block will be non-local to it, i.e. will represent
the same entity inside the block and in the level immediately outside it. Since a statement of a block may again itself
be a block the concepts local and non-local to a block must be understood recursively. Thus an identifier which is
non-local to a block A, may or may not be non-local to the block B in which A is one statement.

### Assignment statements

#### Syntax

`<left part> ::= <identifier> | <subscripted variable>`

`<assignment statement> ::= <left part> := <expression>`

#### Semantics

Assignment statements serve for assigning the value of an expression to one destination. Assignment to a procedure
identifier may only occur within the body of a procedure defining the value of the function designator denoted by that
identifier. If assignment is made to a subscripted variable, the values of all the subscripts must lie within the
appropriate subscript bounds. Otherwise the action of the program becomes undefined. The process will in the general
case be understood to take place in three steps as follows:

- Any subscript expressions occurring in the destinations are evaluated in sequence from left to right.
- The expression of the statement is evaluated.
- The value of the expression is assigned to the destination, with any subscript expressions having values as evaluated
  in the first step.

An assignment to a variable of string type (e.g. `STRING[a]`) without subscript is only allowed when the expression is a
combination of standard functions, concatenations using `+`, conditional expressions and strings in double quotes `"`.
The string type of the expression (e.g. `STRING[b]`) must not have a larger size than the variable (e.g. `a ≥ b` must
hold). The assignment will copy the expression as bytes in UTF-8 representation until a null byte is copied.

The type of the arithmetic expression must not differ from that associated with the destination. The type associated
with a procedure identifier is given by the declarator which appears as the first symbol of the corresponding procedure
declaration.

### Conditional statements

#### Syntax

```
<else clause> ::= 'ELSE' <statement> | <empty>

<conditional statement> ::= 'IF' <Boolean expression> 'THEN'
        <unconditional statement> <else clause>
```

#### Semantics

An if statement is of the form  
`'IF' B 'THEN' Su 'ELSE' S`  
where `B` is a Boolean expression, `Su` is an unconditional statement `S` is a statement. In execution, `B` is
evaluated; if the result is true, `Su` is executed and `S` is skipped; if the result is false, `Su` is not executed
and `S` is executed.

### For statements

#### Syntax

```
<signed number> ::= - <number> | <number>

<for list element> ::= <expression> | 
        <expression> 'STEP' <signed number> 'UNTIL' <expression> |
        <expression> 'WHILE' <Boolean expression>

<for list> ::= <for list element> | <for list> ,
        <for list element>

<for statement> ::= 'FOR' <identifier> := <for list> 'DO' <statement>
```

#### Semantics

A for clause causes the statement `S` which it precedes to be repeatedly executed zero or more times. In addition it
performs a sequence of assignments to its controlled variable (the variable after 'FOR'). The controlled variable must
be of integer type.

#### The for list elements

If the for list contain more than one element then

`'FOR' V := X, Y 'DO' S`

where `X` is a for list element, and `Y` is a for list (which may consist of one element or more), is equivalent to

```
'BEGIN'
  'FOR' V := X 'DO' S; 
  'FOR' V := Y 'DO' S
'END'
```

A for list element `X` can be one of three types, they are executed as follows:

-   ##### Expression element

    If `X` is an arithmetic expression then

    `'FOR' V := X 'DO' S`

    is equivalent to

        'BEGIN'
          V := X; 
          S
        'END'

-   ##### Step-until element

    If `A` and `B` are expressions and `x` a (possibly negative) number then

    `'FOR' V := A 'STEP' x 'UNTIL' B 'DO' S`

    is equivalent to

        'BEGIN'
          V := A;
          Lstart:
          'IF' (V - B) × sign(x) ≤ 0 'THEN'
          'BEGIN'
            S; 
            V := V + x;
            <goto Lstart>
          'END'
        'END'

    Note that `B` is evaluated each iteration.

-   ##### While element

    If `E` is an expression and `F` is a Boolean expression then

    `'FOR' V := E 'WHILE' F 'DO' S`

    is equivalent to

        'BEGIN'
          Lstart:
          V := E;
          'IF' F 'THEN'
          'BEGIN'
            S; 
            <goto Lstart>
          'END'
        'END'

Upon exit from the for statement by exhaustion of the for list, the controlled variable retains the last value assigned
to it.

### Procedure statements

#### Syntax

```
<parameter delimiter> ::= , | ) <identifier> : (

<actual parameter list> ::= <expression> | <actual parameter list> <parameter delimiter> <expression>

<procedure statement> ::= <identifier> | <identifier> ( <actual parameter list> )
```

#### Semantics

A procedure statement serves to invoke (call for) the execution of a procedure body. The effect of this execution will
be equivalent to the effect of performing the following operations on the program at the time of execution of the
procedure statement:

- All formal parameters of the procedure heading are assigned the values of the corresponding actual parameters, these
  assignments being considered as being performed explicitly before entering the procedure body. The effect is as though
  an additional block embracing the procedure body were created in which these assignments were made to variables local
  to this fictitious block with types as given in the corresponding specifications.
- The procedure body is executed. When the end of the procedure body is reached, execution returns to the statements
  following the procedure statement.

The correspondence between the actual parameters of the procedure statement and formal parameters of the procedure
heading is established as follows: The actual parameter list of the procedure statement must have the same number of
entries as the formal parameter list of the procedure declaration heading. The correspondence is obtained by taking the
entries of these two lists in the same order.

All parameter delimiters (even the elaborate one with a colon `:`) are understood to be equivalent. No correspondence
between the parameter delimiters used in a procedure statement and those used in the procedure heading is expected
beyond their number being the same. Thus the information conveyed by using the elaborate ones is entirely optional.

## Declarations

Declarations serve to define certain properties of the simple variables, arrays, strings and procedures used in the
program, and to associate them with identifiers. A declaration of an identifier is valid for one block. Outside this
block the particular identifier may be used for other purposes.

Dynamically this implies the following: at the time of an entry into a block all identifiers declared for the block
assume the significance implied by the nature of the declarations given. If these identifiers had already been defined
by other declarations outside they are for the time being given a new significance. Identifiers which are not declared
for the block, on the other hand, retain their old meaning.

At the time of an exit from a block all identifiers which are declared for the block lose their local significance.

Apart from standard functions in the environmental block, each identifier appearing in a program must be explicitly
declared within a program.

No identifier may be declared more than once in any one block head.

#### Syntax

```
<declaration> ::= 'COMMENT' <any sequence of characters not containing ;>  |
        <type declaration> | <array declaration> |
        <string declaration> | <procedure declaration>
```

### Type declarations

#### Syntax

`<type list> ::= <identifier> | <identifier> , <type list>`

`<type declaration> ::= 'INTEGER' <type list>`

#### Semantics

Type declarations serve to declare certain identifiers to represent simple variables of integer type.

### Array declarations

#### Syntax

```
<dimension list> ::= <number> |
        <dimension list> , <number> 

<array segment> ::= <identifier> [ <dimension list> ] |
        <identifier> , <array segment>

<array list> ::= <array segment> |
        <array list> , <array segment>

<array declaration> ::= 'INTEGER' 'ARRAY' <array list>
```

#### Semantics

An array declaration declares one or several identifiers to represent multidimensional integer arrays of subscripted
variables and gives the dimensions of the arrays and the bounds of the subscripts.

The subscript bounds for any array are given in the first subscript brackets following the identifier of this array in
the form of a dimension list. Each item of this list gives the upper bound as a numeric constant.

The dimensions are given as the number of entries in the dimension lists.

### String declarations

#### Syntax

```
<string segment> ::= <identifier> [ <number> ] |
        <identifier> , <string segment>

<string list> ::= <string segment> |
        <string list> , <string segment>

<string declaration> ::= 'STRING' <string list>
```

#### Semantics

A string declaration declares one or several identifiers to represent one-dimensional strings and specifies the sizes of
the string variables.

The sizes for any string are given in the first subscript brackets following the identifier of this string in the form
of a numeric constant. The sizes have to include the space for a terminating null byte so a size of zero is forbidden.

### Procedure declarations

#### Syntax

```
<formal parameter list> ::= <identifier> |
        <formal parameter list> <parameter delimiter> <identifier>

<procedure heading> ::= <identifier> | 
        <identifier> ( <formal parameter list> )

<procedure body> ::= <statement> | <declaration> ; <procedure body>

<procedure declaration> ::= 'PROCEDURE' <procedure heading> ; <procedure body> |
        'INTEGER' 'PROCEDURE' <procedure heading> ; <procedure body>
```

#### Semantics

A procedure declaration serves to define the procedure associated with a procedure identifier. The principal constituent
of a procedure declaration is a statement or more generally a block, the procedure body, which through the use of
procedure statements and/or function designators may be activated from other parts of the program. Associated with the
body is a heading, which specifies certain identifiers occurring within the body to represent formal parameters. Formal
parameters in the procedure body will, whenever the procedure is activated, be assigned the values of the actual
parameters. The procedure body always acts like a block, whether it has the form of one or not. Consequently, the scope
of any declaration can never extend beyond the procedure body. In addition, if the identifier of a formal parameter is
declared anew within a block inside the procedure body, it is hereby given a local significance and actual parameters
which correspond to it are inaccessible throughout the scope of this inner local variable.

No identifier may appear more than once in any one formal parameter list and if the procedure has a return type a formal
parameter list must not contain the procedure identifier of the same procedure heading. Every identifier in the formal
parameter list must be declared in the specification part of the procedure body. Like declarations in blocks the same
identifier may not be declared more than once in the same scope.

Unlike standard functions the corresponding declaration of an identifier in the formal parameter list must be of integer
type.

#### Values of function designators

For a procedure identifier to be used in a function designator it must have a return type declared as the very first
symbol of the procedure declaration. The value of a function designator is the value of an implicit virtual variable
after execution at the end of the procedure body. This virtual variable can be assigned to using the procedure
identifier as a destination in an assignment statement. Any occurrence of the procedure identifier within the body of
the procedure other than as a destination in an assignment statement denotes activation of the procedure.

If a function designator is used as a procedure statement, then the resulting value is discarded, but such a statement
may be used, if desired, for the purpose of invoking side effects.

Only standard functions can declare a return type other than none or the integer type.
