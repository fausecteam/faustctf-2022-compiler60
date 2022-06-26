package algol60_printer

import (
	"checker/internal/algol60_parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// TODO randomize keywords

type RandomPrintVisitor struct {
	// main source of randomization
	Rand *rand.Rand
	// line on which elems must remain on same line - see GetNewErrorLine()
	ErrorLine int
	// only randomize things which do not affect compiler output
	DeterministicCompilation bool

	newErrorLine int
	onErrorLine  bool
	// string builder for resulting code
	writer strings.Builder

	// for randomizing identifiers
	seenIdentifier        map[string]bool
	identifierReplacement map[string]string
}

func (v *RandomPrintVisitor) GetOutput() string {
	return v.writer.String()
}

// GetNewErrorLine the line number in the generated source which corresponds to ErrorLine in the ParseTree
func (v *RandomPrintVisitor) GetNewErrorLine() int {
	return v.newErrorLine
}

func (v *RandomPrintVisitor) sometimes() bool {
	return v.Rand.Int31n(4) == 0
}

var regexLastLineOnlyWhitespace = regexp.MustCompile("\\n[ \\r\\t]*$")

func (v *RandomPrintVisitor) randSpaceForCtx(ctx antlr.ParserRuleContext) {
	v.randSpaceForToken(ctx.GetStart())
}

func (v *RandomPrintVisitor) randSpaceForToken(ctx antlr.Token) {
	line := ctx.GetLine()
	if line == v.ErrorLine {
		output := v.GetOutput()
		currLine := 1 + strings.Count(output, "\n")
		if !v.onErrorLine {
			// first elem on ErrorLine
			v.onErrorLine = true
			if !regexLastLineOnlyWhitespace.MatchString(output) {
				// if last line has non whitespace place ErrorLine on new line
				v.writer.WriteByte('\n')
				currLine++
			}
			v.newErrorLine = currLine
		} else if v.newErrorLine != currLine {
			panic("newErrorLine line changed")
		}
	} else if v.onErrorLine {
		v.onErrorLine = false
		// make sure elem after ErrorLine comes in a new line
		v.writer.WriteByte('\n')
	}
	v.randSpace()
}

func (v *RandomPrintVisitor) randSpace() {
	alphabet := " \t\r\n"
	if v.onErrorLine {
		// no \n to keep line in ErrorLine
		alphabet = " \t\r"
	}
	if !v.sometimes() {
		return
	}
	l := int32(len(alphabet))
	for i := v.Rand.Int31n(3); i > 0; i-- {
		v.writer.WriteByte(alphabet[v.Rand.Int31n(l)])
	}
}

func (v *RandomPrintVisitor) writeSeparator(sep byte) {
	v.randSpace()
	v.writer.WriteByte(sep)
}

const alphabetLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphabetAlphaNum = alphabetLetters + "0123456789"

func (v *RandomPrintVisitor) randIdentifier() string {
	if v.Rand.Int31n(2) != 0 {
		return pwnyNames[v.Rand.Int31n(int32(len(pwnyNames)))]
	}
	// always make at least 5 characters
	// to avoid hitting x86 register names and triggering vulnerability
	size := 5 + v.Rand.Int31n(7)
	result := make([]byte, size)
	result[0] = alphabetLetters[v.Rand.Int31n(int32(len(alphabetLetters)))]
	l := int32(len(alphabetAlphaNum))
	for i := int32(1); i < size; i++ {
		result[i] = alphabetAlphaNum[v.Rand.Int31n(l)]
	}
	if isPredefinedName[string(result)] {
		// try again
		return v.randIdentifier()
	}
	return string(result)
}

func (v *RandomPrintVisitor) writeSpaceableString(str string) {
	for _, r := range str {
		v.randSpace()
		v.writer.WriteRune(r)
	}
}

func (v *RandomPrintVisitor) join(nodes interface{}, separator byte) {
	// use heavy reflection here because we cannot cast the arrays to []antlr.ParseTree:
	// https://www.reddit.com/r/golang/comments/crzqdp/casting_between_slices_of_interfaces/
	nodesV := reflect.ValueOf(nodes)
	parseTreeType := reflect.TypeOf((*antlr.ParseTree)(nil)).Elem()
	if !nodesV.Type().Elem().Implements(parseTreeType) {
		panic("Wrong type")
	}
	last := nodesV.Len() - 1
	if last == -1 {
		return
	}
	for i := 0; i < last; i++ {
		nodesV.Index(i).Interface().(antlr.ParseTree).Accept(v)
		v.writeSeparator(separator)
	}
	nodesV.Index(last).Interface().(antlr.ParseTree).Accept(v)
}

func (v *RandomPrintVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *RandomPrintVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}

func (v *RandomPrintVisitor) VisitProgram(ctx *algol60_parser.ProgramContext) interface{} {
	v.randSpaceForCtx(ctx)
	ctx.Block().Accept(v)
	return nil
}

func (v *RandomPrintVisitor) VisitBlock(ctx *algol60_parser.BlockContext) interface{} {
	v.randSpaceForCtx(ctx)
	v.writer.WriteString("'BEGIN'")
	ctx.Block_contents().Accept(v)
	v.randSpace()
	v.writer.WriteString("'END'")
	v.randSpace()
	return nil
}

func (v *RandomPrintVisitor) VisitBlock_contents(ctx *algol60_parser.Block_contentsContext) interface{} {
	v.randSpaceForCtx(ctx)
	decls := ctx.GetDecls()
	if len(decls) > 0 {
		if len(decls) > 1 && !v.DeterministicCompilation && v.ErrorLine == -1 && v.sometimes() {
			v.Rand.Shuffle(len(decls), func(i, j int) {
				decls[i], decls[j] = decls[j], decls[i]
			})
		}
		v.join(decls, ';')
		v.writeSeparator(';')
	}
	stmts := ctx.GetStmts()
	if len(stmts) > 0 {
		v.join(stmts, ';')
		if stmts[len(stmts)-1].GetStart().GetTokenType() == algol60_parser.Algol60V2ParserCOMMENT {
			// if ends with comment need additional semicolon
			v.writer.WriteByte(';')
		}
	}
	v.randSpace()
	return nil
}

func (v *RandomPrintVisitor) VisitDeclaration(ctx *algol60_parser.DeclarationContext) interface{} {
	v.randSpaceForCtx(ctx)
	if ctx.COMMENT() != nil {
		v.writer.WriteString(ctx.COMMENT().GetText())
		return nil
	}
	v.VisitChildren(ctx)
	return nil
}

func (v *RandomPrintVisitor) VisitType_declaration(ctx *algol60_parser.Type_declarationContext) interface{} {
	v.randSpaceForCtx(ctx)
	v.writer.WriteString("'INTEGER'")
	vars := ctx.GetVars()
	if v.ErrorLine == -1 && v.sometimes() {
		v.Rand.Shuffle(len(vars), func(i, j int) {
			vars[i], vars[j] = vars[j], vars[i]
		})
	}
	v.join(vars, ',')
	return nil
}

func (v *RandomPrintVisitor) VisitArray_declaration(ctx *algol60_parser.Array_declarationContext) interface{} {
	v.randSpaceForCtx(ctx)
	v.writer.WriteString("'INTEGER' 'ARRAY'")
	segs := ctx.GetSegments()
	if v.ErrorLine == -1 && v.sometimes() {
		v.Rand.Shuffle(len(segs), func(i, j int) {
			segs[i], segs[j] = segs[j], segs[i]
		})
	}
	v.join(segs, ',')
	return nil
}

func (v *RandomPrintVisitor) VisitArray_segment(ctx *algol60_parser.Array_segmentContext) interface{} {
	v.randSpaceForCtx(ctx)
	vars := ctx.GetVars()
	if v.ErrorLine == -1 && v.sometimes() {
		v.Rand.Shuffle(len(vars), func(i, j int) {
			vars[i], vars[j] = vars[j], vars[i]
		})
	}
	v.join(vars, ',')
	v.writeSeparator('[')
	v.join(ctx.GetSizes(), ',')
	v.writeSeparator(']')
	return nil
}

func (v *RandomPrintVisitor) VisitString_declaration(ctx *algol60_parser.String_declarationContext) interface{} {
	v.randSpaceForCtx(ctx)
	v.writer.WriteString("'STRING'")
	segs := ctx.GetSegments()
	if v.ErrorLine == -1 && v.sometimes() {
		v.Rand.Shuffle(len(segs), func(i, j int) {
			segs[i], segs[j] = segs[j], segs[i]
		})
	}
	v.join(segs, ',')
	return nil
}

func (v *RandomPrintVisitor) VisitString_segment(ctx *algol60_parser.String_segmentContext) interface{} {
	v.randSpaceForCtx(ctx)
	vars := ctx.GetVars()
	if v.ErrorLine == -1 && v.sometimes() {
		v.Rand.Shuffle(len(vars), func(i, j int) {
			vars[i], vars[j] = vars[j], vars[i]
		})
	}
	v.join(vars, ',')
	v.writeSeparator('[')
	ctx.GetSize().Accept(v)
	v.writeSeparator(']')
	return nil
}

func (v *RandomPrintVisitor) VisitProcedure_declaration(ctx *algol60_parser.Procedure_declarationContext) interface{} {
	v.randSpaceForCtx(ctx)
	// first line
	if ctx.GetReturnType() != nil {
		v.writer.WriteString("'INTEGER'")
		v.randSpace()
	}
	v.writer.WriteString("'PROCEDURE'")
	ctx.Identifier().Accept(v)
	if ctx.GetParams() != nil {
		v.writeSeparator('(')
		ctx.GetParams().Accept(v)
		v.writeSeparator(')')
	}
	v.writeSeparator(';')
	// decls
	decls := ctx.GetDecls()
	if len(decls) > 0 {
		if len(decls) > 1 && v.ErrorLine == -1 && v.sometimes() {
			v.Rand.Shuffle(len(decls), func(i, j int) {
				decls[i], decls[j] = decls[j], decls[i]
			})
		}
		v.join(decls, ';')
		v.writeSeparator(';')
	}
	// statement
	ctx.Statement().Accept(v)
	return nil
}

func (v *RandomPrintVisitor) VisitFormal_parameter_list(ctx *algol60_parser.Formal_parameter_listContext) interface{} {
	v.randSpaceForCtx(ctx)
	ctx.Identifier().Accept(v)
	if ctx.Parameter_delimiter() != nil {
		ctx.Parameter_delimiter().Accept(v)
		ctx.Formal_parameter_list().Accept(v)
	}
	return nil
}

func (v *RandomPrintVisitor) VisitParameter_delimiter(ctx *algol60_parser.Parameter_delimiterContext) interface{} {
	v.randSpaceForCtx(ctx)
	var namedParameterDelimiter string
	if ctx.IDENTIFIER() != nil {
		namedParameterDelimiter = ctx.IDENTIFIER().GetText()
	}
	if v.sometimes() {
		namedParameterDelimiter = v.randIdentifier()
	}
	if len(namedParameterDelimiter) == 0 {
		v.writeSeparator(',')
	} else {
		v.writeSeparator(')')
		v.writeSpaceableString(namedParameterDelimiter)
		v.writeSpaceableString(":(")
	}
	return nil
}

func (v *RandomPrintVisitor) VisitStatement(ctx *algol60_parser.StatementContext) interface{} {
	if ctx.COMMENT() != nil {
		v.writer.WriteString(ctx.COMMENT().GetText())
		return nil
	}
	if !v.DeterministicCompilation && v.sometimes() {
		v.writer.WriteString("'BEGIN'")
		v.VisitChildren(ctx)
		v.writer.WriteString("'END'")
		return nil
	}
	v.VisitChildren(ctx)
	return nil
}

func (v *RandomPrintVisitor) VisitUnconditional_statement(ctx *algol60_parser.Unconditional_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *RandomPrintVisitor) VisitAssignment_statement(ctx *algol60_parser.Assignment_statementContext) interface{} {
	v.randSpaceForCtx(ctx)
	ctx.Left_part().Accept(v)
	v.writeSpaceableString(":=")
	ctx.Expression().Accept(v)
	return nil
}

func (v *RandomPrintVisitor) VisitLeft_part(ctx *algol60_parser.Left_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *RandomPrintVisitor) VisitProcedure_statement(ctx *algol60_parser.Procedure_statementContext) interface{} {
	v.randSpaceForCtx(ctx)
	ctx.Identifier().Accept(v)
	if ctx.Actual_parameter_list() != nil {
		v.writeSeparator('(')
		ctx.Actual_parameter_list().Accept(v)
		v.writeSeparator(')')
	}
	return nil
}

func (v *RandomPrintVisitor) VisitActual_parameter_list(ctx *algol60_parser.Actual_parameter_listContext) interface{} {
	v.randSpaceForCtx(ctx)
	ctx.Expression().Accept(v)
	if ctx.Parameter_delimiter() != nil {
		ctx.Parameter_delimiter().Accept(v)
		ctx.Actual_parameter_list().Accept(v)
	}
	return nil
}

func (v *RandomPrintVisitor) VisitConditional_statement(ctx *algol60_parser.Conditional_statementContext) interface{} {
	v.randSpaceForCtx(ctx)
	v.writer.WriteString("'IF'")
	ctx.Boolean_expression().Accept(v)
	v.randSpace()
	v.writer.WriteString("'THEN'")
	ctx.GetThenBlock().Accept(v)
	if ctx.GetElseBlock() != nil {
		v.randSpace()
		v.writer.WriteString("'ELSE'")
		ctx.GetElseBlock().Accept(v)
	}
	return nil
}

func (v *RandomPrintVisitor) VisitFor_statement(ctx *algol60_parser.For_statementContext) interface{} {
	v.randSpaceForCtx(ctx)
	v.writer.WriteString("'FOR'")
	ctx.GetLoopVar().Accept(v)
	v.writeSpaceableString(":=")
	ctx.For_list().Accept(v)
	v.writer.WriteString("'DO'")
	ctx.GetBody().Accept(v)
	return nil
}

func (v *RandomPrintVisitor) VisitFor_list(ctx *algol60_parser.For_listContext) interface{} {
	v.randSpaceForCtx(ctx)
	v.join(ctx.GetElems(), ',')
	return nil
}

func (v *RandomPrintVisitor) VisitFor_list_element(ctx *algol60_parser.For_list_elementContext) interface{} {
	v.randSpaceForCtx(ctx)
	ctx.GetFirstExpr().Accept(v)
	switch {
	case ctx.GetStep() != nil:
		v.randSpace()
		v.writer.WriteString("'STEP'")
		if ctx.GetStepNeg() != nil {
			v.writeSeparator('-')
		}
		ctx.GetStep().Accept(v)
		v.randSpace()
		v.writer.WriteString("'UNTIL'")
		ctx.GetUntil().Accept(v)
	case ctx.GetWhileCond() != nil:
		v.randSpace()
		v.writer.WriteString("'WHILE'")
		ctx.GetWhileCond().Accept(v)
	}
	return nil
}

func (v *RandomPrintVisitor) VisitExpression(ctx *algol60_parser.ExpressionContext) interface{} {
	if v.sometimes() && v.sometimes() {
		v.writeSeparator('(')
		v.VisitExpression(ctx)
		v.writeSeparator(')')
		return nil
	}
	v.randSpaceForCtx(ctx)
	if ctx.GetCond() == nil {
		ctx.Simple_expression().Accept(v)
	} else {
		v.writer.WriteString("'IF'")
		ctx.GetCond().Accept(v)
		v.randSpace()
		v.writer.WriteString("'THEN'")
		ctx.GetTrueCase().Accept(v)
		v.randSpace()
		v.writer.WriteString("'ELSE'")
		ctx.GetFalseCase().Accept(v)
	}
	return nil
}

func (v *RandomPrintVisitor) VisitSimple_expression(ctx *algol60_parser.Simple_expressionContext) interface{} {
	v.randSpaceForCtx(ctx)
	ops := ctx.GetOperators()
	operands := ctx.GetOperands()
	if len(ops) == len(operands) {
		// prefix operator
		v.writer.WriteString(ops[0].GetText())
		ops = ops[1:]
	} else if v.Rand.Int31n(4) == 0 {
		// unnecessary + prefix
		v.writeSeparator('+')
	}
	operands[0].Accept(v)
	for i, operand := range operands[1:] {
		v.randSpaceForToken(ops[i])
		v.writer.WriteString(ops[i].GetText())
		operand.Accept(v)
	}
	return nil
}

func (v *RandomPrintVisitor) VisitTerm(ctx *algol60_parser.TermContext) interface{} {
	v.randSpaceForCtx(ctx)
	ops := ctx.GetOperators()
	operands := ctx.GetOperands()
	operands[0].Accept(v)
	for i, operand := range operands[1:] {
		v.randSpaceForToken(ops[i])
		v.writer.WriteString(ops[i].GetText())
		operand.Accept(v)
	}
	return nil
}

func (v *RandomPrintVisitor) VisitFactor(ctx *algol60_parser.FactorContext) interface{} {
	v.randSpaceForCtx(ctx)
	switch {
	case ctx.Number() != nil:
		ctx.Number().Accept(v)
	case ctx.QUOTED_STRING() != nil:
		// TODO add '" + "' to string constants
		v.writer.WriteString(ctx.QUOTED_STRING().GetText())
	case ctx.Identifier() != nil:
		ctx.Identifier().Accept(v)
	case ctx.Function_designator() != nil:
		ctx.Function_designator().Accept(v)
	case ctx.Subscripted_variable() != nil:
		ctx.Subscripted_variable().Accept(v)
	case ctx.Expression() != nil:
		v.writeSeparator('(')
		ctx.Expression().Accept(v)
		v.writeSeparator(')')
	}
	return nil
}

func (v *RandomPrintVisitor) VisitFunction_designator(ctx *algol60_parser.Function_designatorContext) interface{} {
	v.randSpaceForCtx(ctx)
	ctx.Identifier().Accept(v)
	// this always has at least one argument
	// as a function call with zero arguments would match the identifier in the "factor" rule
	v.writeSeparator('(')
	ctx.GetArgs().Accept(v)
	v.writeSeparator(')')
	return nil
}

func (v *RandomPrintVisitor) VisitBoolean_expression(ctx *algol60_parser.Boolean_expressionContext) interface{} {
	if v.sometimes() && v.sometimes() {
		v.writeSeparator('(')
		v.VisitBoolean_expression(ctx)
		v.writeSeparator(')')
		return nil
	}
	v.randSpaceForCtx(ctx)
	if ctx.GetCond() == nil {
		ctx.Implication().Accept(v)
	} else {
		v.writer.WriteString("'IF'")
		ctx.GetCond().Accept(v)
		v.randSpace()
		v.writer.WriteString("'THEN'")
		ctx.GetThenValue().Accept(v)
		v.randSpace()
		v.writer.WriteString("'ELSE'")
		ctx.GetElseValue().Accept(v)
	}
	return nil
}

func (v *RandomPrintVisitor) VisitImplication(ctx *algol60_parser.ImplicationContext) interface{} {
	v.randSpaceForCtx(ctx)
	ops := ctx.GetOperators()
	operands := ctx.GetOperands()
	operands[0].Accept(v)
	for i, operand := range operands[1:] {
		v.randSpaceForToken(ops[i])
		v.writer.WriteString("⊃")
		operand.Accept(v)
	}
	return nil
}

func (v *RandomPrintVisitor) VisitBoolean_term(ctx *algol60_parser.Boolean_termContext) interface{} {
	v.randSpaceForCtx(ctx)
	ops := ctx.GetOperators()
	operands := ctx.GetOperands()
	operands[0].Accept(v)
	for i, operand := range operands[1:] {
		v.randSpaceForToken(ops[i])
		v.writer.WriteString("∨")
		operand.Accept(v)
	}
	return nil
}

func (v *RandomPrintVisitor) VisitBoolean_factor(ctx *algol60_parser.Boolean_factorContext) interface{} {
	v.randSpaceForCtx(ctx)
	ops := ctx.GetOperators()
	operands := ctx.GetOperands()
	operands[0].Accept(v)
	for i, operand := range operands[1:] {
		v.randSpaceForToken(ops[i])
		v.writer.WriteString("∧")
		operand.Accept(v)
	}
	return nil
}

func (v *RandomPrintVisitor) VisitBoolean_secondary(ctx *algol60_parser.Boolean_secondaryContext) interface{} {
	v.randSpaceForCtx(ctx)
	if ctx.GetOperator() != nil {
		v.randSpaceForToken(ctx.GetOperator())
		v.writer.WriteString("¬")
	}
	ctx.GetOperand().Accept(v)
	return nil
}

func (v *RandomPrintVisitor) VisitBoolean_primary(ctx *algol60_parser.Boolean_primaryContext) interface{} {
	v.randSpaceForCtx(ctx)
	if ctx.GetLeft() != nil {
		ctx.GetLeft().Accept(v)
		v.randSpaceForToken(ctx.GetOperator())
		v.writer.WriteString(ctx.GetOperator().GetText())
		ctx.GetRight().Accept(v)
	} else {
		v.writeSeparator('(')
		ctx.Boolean_expression().Accept(v)
		v.writeSeparator(')')
	}
	return nil
}

func (v *RandomPrintVisitor) VisitSubscripted_variable(ctx *algol60_parser.Subscripted_variableContext) interface{} {
	v.randSpaceForCtx(ctx)
	ctx.GetArrayVar().Accept(v)
	v.writeSeparator('[')
	v.join(ctx.GetIdxExprs(), ',')
	v.writeSeparator(']')
	return nil
}

func (v *RandomPrintVisitor) VisitIdentifier(ctx *algol60_parser.IdentifierContext) interface{} {
	v.randSpaceForCtx(ctx)
	ident := ctx.IDENTIFIER().GetText()
	if v.DeterministicCompilation || isPredefinedName[ident] {
		v.writeSpaceableString(ident)
		return nil
	}

	if v.seenIdentifier == nil {
		v.seenIdentifier = make(map[string]bool)
		v.identifierReplacement = make(map[string]string)
	}
	if !v.seenIdentifier[ident] {
		v.seenIdentifier[ident] = true
		v.identifierReplacement[ident] = ident
		if v.sometimes() {
			newIdent := v.randIdentifier()
			for v.seenIdentifier[newIdent] {
				newIdent = v.randIdentifier()
			}
			v.identifierReplacement[ident] = newIdent
			// if newIdent actually gets used later, use old name to avoid collisions
			v.seenIdentifier[newIdent] = true
			v.identifierReplacement[newIdent] = ident
		}
	}
	ident = v.identifierReplacement[ident]
	v.writeSpaceableString(ident)
	return nil
}

func (v *RandomPrintVisitor) VisitNumber(ctx *algol60_parser.NumberContext) interface{} {
	v.randSpaceForCtx(ctx)
	_, canReplaceNumber := ctx.GetParent().(*algol60_parser.FactorContext)
	if canReplaceNumber && !v.DeterministicCompilation && v.sometimes() {
		val := ctx.GetValue()
		addend := v.Rand.Int63()
		if v.Rand.Int31n(2) != 0 {
			val -= addend
			v.writeSpaceableString(fmt.Sprintf("(%d + %d)", val, addend))
		} else {
			val += addend
			v.writeSpaceableString(fmt.Sprintf("(%d - %d)", val, addend))
		}
		return nil
	}
	numStr := strconv.FormatInt(ctx.GetValue(), 10)
	v.writeSpaceableString(numStr)
	return nil
}

func (v *RandomPrintVisitor) VisitTerminal(_ antlr.TerminalNode) interface{} {
	panic("not implemented")
}

func (v *RandomPrintVisitor) VisitErrorNode(_ antlr.ErrorNode) interface{} {
	panic("not implemented")
}
