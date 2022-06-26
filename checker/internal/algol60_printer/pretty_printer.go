package algol60_printer

import (
	"checker/internal/algol60_parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"strconv"
	"strings"
)

var DefaultIndent = strings.Repeat(" ", 2)

// TODO also use pwnyNames

type PrettyPrintVisitor struct {
	writer strings.Builder

	currIndent string
}

func (v *PrettyPrintVisitor) GetOutput() string {
	return v.writer.String()
}

func (v *PrettyPrintVisitor) join(nodes interface{}, separator string, doIndent bool) {
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
		if doIndent {
			v.writer.WriteString(v.currIndent)
		}
		nodesV.Index(i).Interface().(antlr.ParseTree).Accept(v)
		v.writer.WriteString(separator)
	}
	if doIndent {
		v.writer.WriteString(v.currIndent)
	}
	nodesV.Index(last).Interface().(antlr.ParseTree).Accept(v)
}

func (v *PrettyPrintVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *PrettyPrintVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitProgram(ctx *algol60_parser.ProgramContext) interface{} {
	ctx.Block().Accept(v)
	v.writer.WriteByte('\n')
	return nil
}

func (v *PrettyPrintVisitor) VisitBlock(ctx *algol60_parser.BlockContext) interface{} {
	v.writer.WriteString("'BEGIN'\n")
	prevIndent := v.currIndent
	v.currIndent += DefaultIndent
	ctx.Block_contents().Accept(v)
	v.currIndent = prevIndent
	v.writer.WriteString(prevIndent)
	v.writer.WriteString("'END'")
	return nil
}

func (v *PrettyPrintVisitor) VisitBlock_contents(ctx *algol60_parser.Block_contentsContext) interface{} {
	if len(ctx.GetDecls()) > 0 {
		v.join(ctx.GetDecls(), ";\n", true)
		v.writer.WriteString(";\n")
	}
	stmts := ctx.GetStmts()
	if len(stmts) > 0 {
		v.join(stmts, ";\n", true)
		if stmts[len(stmts)-1].GetStart().GetTokenType() == algol60_parser.Algol60V2ParserCOMMENT {
			// if ends with comment need additional semicolon
			v.writer.WriteByte(';')
		}
		v.writer.WriteByte('\n')
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitDeclaration(ctx *algol60_parser.DeclarationContext) interface{} {
	if ctx.COMMENT() == nil {
		v.VisitChildren(ctx)
		return nil
	}
	v.writer.WriteString(ctx.COMMENT().GetText())
	return nil
}

func (v *PrettyPrintVisitor) VisitType_declaration(ctx *algol60_parser.Type_declarationContext) interface{} {
	v.writer.WriteString("'INTEGER' ")
	v.join(ctx.GetVars(), ", ", false)
	return nil
}

func (v *PrettyPrintVisitor) VisitArray_declaration(ctx *algol60_parser.Array_declarationContext) interface{} {
	v.writer.WriteString("'INTEGER' 'ARRAY' ")
	v.join(ctx.GetSegments(), ", ", false)
	return nil
}

func (v *PrettyPrintVisitor) VisitArray_segment(ctx *algol60_parser.Array_segmentContext) interface{} {
	v.join(ctx.GetVars(), ", ", false)
	v.writer.WriteByte('[')
	v.join(ctx.GetSizes(), ", ", false)
	v.writer.WriteByte(']')
	return nil
}

func (v *PrettyPrintVisitor) VisitString_declaration(ctx *algol60_parser.String_declarationContext) interface{} {
	v.writer.WriteString("'STRING' ")
	v.join(ctx.GetSegments(), ", ", false)
	return nil
}

func (v *PrettyPrintVisitor) VisitString_segment(ctx *algol60_parser.String_segmentContext) interface{} {
	v.join(ctx.GetVars(), ", ", false)
	v.writer.WriteByte('[')
	ctx.GetSize().Accept(v)
	v.writer.WriteByte(']')
	return nil
}

func (v *PrettyPrintVisitor) VisitProcedure_declaration(ctx *algol60_parser.Procedure_declarationContext) interface{} {
	// first line
	if ctx.GetReturnType() != nil {
		v.writer.WriteString("'INTEGER' ")
	}
	v.writer.WriteString("'PROCEDURE' ")
	ctx.Identifier().Accept(v)
	if ctx.GetParams() != nil {
		v.writer.WriteByte('(')
		ctx.GetParams().Accept(v)
		v.writer.WriteByte(')')
	}
	v.writer.WriteString(";\n")
	// indent decls
	prevIndent := v.currIndent
	v.currIndent += DefaultIndent
	v.join(ctx.GetDecls(), ";\n", true)
	if len(ctx.GetDecls()) > 0 {
		v.writer.WriteString(";\n")
	}
	// statement
	switch ctx.Statement().GetStart().GetTokenType() {
	case algol60_parser.Algol60V2ParserBEGIN,
		algol60_parser.Algol60V2ParserFOR,
		algol60_parser.Algol60V2ParserIF:
		// do not indent wrapping statements (block, for, if)
		v.currIndent = prevIndent
	}
	v.writer.WriteString(v.currIndent)
	ctx.Statement().Accept(v)
	v.currIndent = prevIndent
	return nil
}

func (v *PrettyPrintVisitor) VisitFormal_parameter_list(ctx *algol60_parser.Formal_parameter_listContext) interface{} {
	ctx.Identifier().Accept(v)
	if ctx.Parameter_delimiter() != nil {
		ctx.Parameter_delimiter().Accept(v)
		ctx.Formal_parameter_list().Accept(v)
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitParameter_delimiter(ctx *algol60_parser.Parameter_delimiterContext) interface{} {
	if ctx.IDENTIFIER() == nil {
		v.writer.WriteString(", ")
	} else {
		v.writer.WriteString(") ")
		v.writer.WriteString(ctx.IDENTIFIER().GetText())
		v.writer.WriteString(" : (")
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitStatement(ctx *algol60_parser.StatementContext) interface{} {
	if ctx.COMMENT() == nil {
		v.VisitChildren(ctx)
		return nil
	}
	v.writer.WriteString(ctx.COMMENT().GetText())
	return nil
}

func (v *PrettyPrintVisitor) VisitUnconditional_statement(ctx *algol60_parser.Unconditional_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PrettyPrintVisitor) VisitAssignment_statement(ctx *algol60_parser.Assignment_statementContext) interface{} {
	ctx.Left_part().Accept(v)
	v.writer.WriteString(" := ")
	ctx.Expression().Accept(v)
	return nil
}

func (v *PrettyPrintVisitor) VisitLeft_part(ctx *algol60_parser.Left_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PrettyPrintVisitor) VisitProcedure_statement(ctx *algol60_parser.Procedure_statementContext) interface{} {
	ctx.Identifier().Accept(v)
	if ctx.Actual_parameter_list() != nil {
		v.writer.WriteString("(")
		ctx.Actual_parameter_list().Accept(v)
		v.writer.WriteString(")")
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitActual_parameter_list(ctx *algol60_parser.Actual_parameter_listContext) interface{} {
	ctx.Expression().Accept(v)
	if ctx.Parameter_delimiter() != nil {
		ctx.Parameter_delimiter().Accept(v)
		ctx.Actual_parameter_list().Accept(v)
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitConditional_statement(ctx *algol60_parser.Conditional_statementContext) interface{} {
	v.writer.WriteString("'IF' ")
	ctx.Boolean_expression().Accept(v)
	v.writer.WriteString(" 'THEN'")
	prevIndent := v.currIndent
	switch ctx.GetThenBlock().GetStart().GetTokenType() {
	case algol60_parser.Algol60V2ParserBEGIN:
		// do not indent block statements but use 'BEGIN' and 'END' as sort of brackets
		v.writer.WriteByte(' ')
		ctx.GetThenBlock().Accept(v)
		if ctx.GetElseBlock() != nil {
			v.writer.WriteString(" 'ELSE'")
		}
	default:
		v.writer.WriteByte('\n')
		v.currIndent += DefaultIndent
		v.writer.WriteString(v.currIndent)
		ctx.GetThenBlock().Accept(v)
		v.currIndent = prevIndent
		if ctx.GetElseBlock() != nil {
			v.writer.WriteByte('\n')
			v.writer.WriteString(prevIndent)
			v.writer.WriteString("'ELSE'")
		}
	}
	if ctx.GetElseBlock() == nil {
		return nil
	}
	switch ctx.GetElseBlock().GetStart().GetTokenType() {
	case algol60_parser.Algol60V2ParserBEGIN:
		// do not indent block statements but use 'BEGIN' and 'END' as sort of brackets
		v.writer.WriteByte(' ')
		ctx.GetElseBlock().Accept(v)
	default:
		v.writer.WriteByte('\n')
		v.currIndent += DefaultIndent
		v.writer.WriteString(v.currIndent)
		ctx.GetElseBlock().Accept(v)
		v.currIndent = prevIndent
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitFor_statement(ctx *algol60_parser.For_statementContext) interface{} {
	v.writer.WriteString("'FOR' ")
	ctx.GetLoopVar().Accept(v)
	v.writer.WriteString(" := ")
	ctx.For_list().Accept(v)
	v.writer.WriteString(" 'DO'")
	switch ctx.GetBody().GetStart().GetTokenType() {
	case algol60_parser.Algol60V2ParserBEGIN:
		// do not indent block statements but use 'BEGIN' and 'END' as sort of brackets
		v.writer.WriteByte(' ')
		ctx.GetBody().Accept(v)
	default:
		v.writer.WriteByte('\n')
		prevIndent := v.currIndent
		v.currIndent += DefaultIndent
		v.writer.WriteString(v.currIndent)
		ctx.GetBody().Accept(v)
		v.currIndent = prevIndent
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitFor_list(ctx *algol60_parser.For_listContext) interface{} {
	v.join(ctx.GetElems(), ", ", false)
	return nil
}

func (v *PrettyPrintVisitor) VisitFor_list_element(ctx *algol60_parser.For_list_elementContext) interface{} {
	ctx.GetFirstExpr().Accept(v)
	switch {
	case ctx.GetStep() != nil:
		v.writer.WriteString(" 'STEP' ")
		if ctx.GetStepNeg() != nil {
			v.writer.WriteByte('-')
		}
		ctx.GetStep().Accept(v)
		v.writer.WriteString(" 'UNTIL' ")
		ctx.GetUntil().Accept(v)
	case ctx.GetWhileCond() != nil:
		v.writer.WriteString(" 'WHILE' ")
		ctx.GetWhileCond().Accept(v)
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitExpression(ctx *algol60_parser.ExpressionContext) interface{} {
	if ctx.GetCond() == nil {
		ctx.Simple_expression().Accept(v)
	} else {
		v.writer.WriteString("'IF' ")
		ctx.GetCond().Accept(v)
		v.writer.WriteString(" 'THEN' ")
		ctx.GetTrueCase().Accept(v)
		v.writer.WriteString(" 'ELSE' ")
		ctx.GetFalseCase().Accept(v)
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitSimple_expression(ctx *algol60_parser.Simple_expressionContext) interface{} {
	ops := ctx.GetOperators()
	operands := ctx.GetOperands()
	if len(ops) == len(operands) {
		// prefix operator
		v.writer.WriteString(ops[0].GetText())
		ops = ops[1:]
	}
	operands[0].Accept(v)
	for i, operand := range operands[1:] {
		v.writer.WriteByte(' ')
		v.writer.WriteString(ops[i].GetText())
		v.writer.WriteByte(' ')
		operand.Accept(v)
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitTerm(ctx *algol60_parser.TermContext) interface{} {
	ops := ctx.GetOperators()
	operands := ctx.GetOperands()
	operands[0].Accept(v)
	for i, operand := range operands[1:] {
		v.writer.WriteByte(' ')
		v.writer.WriteString(ops[i].GetText())
		v.writer.WriteByte(' ')
		operand.Accept(v)
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitFactor(ctx *algol60_parser.FactorContext) interface{} {
	switch {
	case ctx.Number() != nil:
		ctx.Number().Accept(v)
	case ctx.QUOTED_STRING() != nil:
		v.writer.WriteString(ctx.QUOTED_STRING().GetText())
	case ctx.Identifier() != nil:
		ctx.Identifier().Accept(v)
	case ctx.Function_designator() != nil:
		ctx.Function_designator().Accept(v)
	case ctx.Subscripted_variable() != nil:
		ctx.Subscripted_variable().Accept(v)
	case ctx.Expression() != nil:
		v.writer.WriteByte('(')
		ctx.Expression().Accept(v)
		v.writer.WriteByte(')')
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitFunction_designator(ctx *algol60_parser.Function_designatorContext) interface{} {
	ctx.Identifier().Accept(v)
	// this always has at least one argument
	// as a function call with zero arguments would match the identifier in the "factor" rule
	v.writer.WriteByte('(')
	ctx.GetArgs().Accept(v)
	v.writer.WriteByte(')')
	return nil
}

func (v *PrettyPrintVisitor) VisitBoolean_expression(ctx *algol60_parser.Boolean_expressionContext) interface{} {
	if ctx.GetCond() == nil {
		ctx.Implication().Accept(v)
	} else {
		v.writer.WriteString("'IF' ")
		ctx.GetCond().Accept(v)
		v.writer.WriteString(" 'THEN' ")
		ctx.GetThenValue().Accept(v)
		v.writer.WriteString(" 'ELSE' ")
		ctx.GetElseValue().Accept(v)
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitImplication(ctx *algol60_parser.ImplicationContext) interface{} {
	operands := ctx.GetOperands()
	operands[0].Accept(v)
	for _, operand := range operands[1:] {
		v.writer.WriteString(" ⊃ ")
		operand.Accept(v)
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitBoolean_term(ctx *algol60_parser.Boolean_termContext) interface{} {
	operands := ctx.GetOperands()
	operands[0].Accept(v)
	for _, operand := range operands[1:] {
		v.writer.WriteString(" ∨ ")
		operand.Accept(v)
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitBoolean_factor(ctx *algol60_parser.Boolean_factorContext) interface{} {
	operands := ctx.GetOperands()
	operands[0].Accept(v)
	for _, operand := range operands[1:] {
		v.writer.WriteString(" ∧ ")
		operand.Accept(v)
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitBoolean_secondary(ctx *algol60_parser.Boolean_secondaryContext) interface{} {
	if ctx.GetOperator() != nil {
		v.writer.WriteString("¬")
	}
	ctx.GetOperand().Accept(v)
	return nil
}

func (v *PrettyPrintVisitor) VisitBoolean_primary(ctx *algol60_parser.Boolean_primaryContext) interface{} {
	if ctx.GetLeft() != nil {
		ctx.GetLeft().Accept(v)
		v.writer.WriteByte(' ')
		v.writer.WriteString(ctx.GetOperator().GetText())
		v.writer.WriteByte(' ')
		ctx.GetRight().Accept(v)
	} else {
		v.writer.WriteByte('(')
		ctx.Boolean_expression().Accept(v)
		v.writer.WriteByte(')')
	}
	return nil
}

func (v *PrettyPrintVisitor) VisitSubscripted_variable(ctx *algol60_parser.Subscripted_variableContext) interface{} {
	ctx.GetArrayVar().Accept(v)
	v.writer.WriteByte('[')
	v.join(ctx.GetIdxExprs(), ", ", false)
	v.writer.WriteByte(']')
	return nil
}

func (v *PrettyPrintVisitor) VisitIdentifier(ctx *algol60_parser.IdentifierContext) interface{} {
	v.writer.WriteString(ctx.IDENTIFIER().GetText())
	return nil
}

func (v *PrettyPrintVisitor) VisitNumber(ctx *algol60_parser.NumberContext) interface{} {
	numStr := strconv.FormatInt(ctx.GetValue(), 10)
	v.writer.WriteString(numStr)
	return nil
}

func (v *PrettyPrintVisitor) VisitTerminal(_ antlr.TerminalNode) interface{} {
	panic("not implemented")
}

func (v *PrettyPrintVisitor) VisitErrorNode(_ antlr.ErrorNode) interface{} {
	panic("not implemented")
}
