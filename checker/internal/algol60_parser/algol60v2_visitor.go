// Code generated from Algol60V2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package algol60_parser // Algol60V2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by Algol60V2Parser.
type Algol60V2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Algol60V2Parser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#block_contents.
	VisitBlock_contents(ctx *Block_contentsContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#type_declaration.
	VisitType_declaration(ctx *Type_declarationContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#array_declaration.
	VisitArray_declaration(ctx *Array_declarationContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#array_segment.
	VisitArray_segment(ctx *Array_segmentContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#string_declaration.
	VisitString_declaration(ctx *String_declarationContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#string_segment.
	VisitString_segment(ctx *String_segmentContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#procedure_declaration.
	VisitProcedure_declaration(ctx *Procedure_declarationContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#formal_parameter_list.
	VisitFormal_parameter_list(ctx *Formal_parameter_listContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#parameter_delimiter.
	VisitParameter_delimiter(ctx *Parameter_delimiterContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#unconditional_statement.
	VisitUnconditional_statement(ctx *Unconditional_statementContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#assignment_statement.
	VisitAssignment_statement(ctx *Assignment_statementContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#left_part.
	VisitLeft_part(ctx *Left_partContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#procedure_statement.
	VisitProcedure_statement(ctx *Procedure_statementContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#actual_parameter_list.
	VisitActual_parameter_list(ctx *Actual_parameter_listContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#conditional_statement.
	VisitConditional_statement(ctx *Conditional_statementContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#for_statement.
	VisitFor_statement(ctx *For_statementContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#for_list.
	VisitFor_list(ctx *For_listContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#for_list_element.
	VisitFor_list_element(ctx *For_list_elementContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#simple_expression.
	VisitSimple_expression(ctx *Simple_expressionContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#function_designator.
	VisitFunction_designator(ctx *Function_designatorContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#boolean_expression.
	VisitBoolean_expression(ctx *Boolean_expressionContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#implication.
	VisitImplication(ctx *ImplicationContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#boolean_term.
	VisitBoolean_term(ctx *Boolean_termContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#boolean_factor.
	VisitBoolean_factor(ctx *Boolean_factorContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#boolean_secondary.
	VisitBoolean_secondary(ctx *Boolean_secondaryContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#boolean_primary.
	VisitBoolean_primary(ctx *Boolean_primaryContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#subscripted_variable.
	VisitSubscripted_variable(ctx *Subscripted_variableContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by Algol60V2Parser#number.
	VisitNumber(ctx *NumberContext) interface{}
}
