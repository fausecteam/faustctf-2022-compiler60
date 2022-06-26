// Code generated from Algol60V2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package algol60_parser // Algol60V2
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseAlgol60V2Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAlgol60V2Visitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitBlock_contents(ctx *Block_contentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitType_declaration(ctx *Type_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitArray_declaration(ctx *Array_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitArray_segment(ctx *Array_segmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitString_declaration(ctx *String_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitString_segment(ctx *String_segmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitProcedure_declaration(ctx *Procedure_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitFormal_parameter_list(ctx *Formal_parameter_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitParameter_delimiter(ctx *Parameter_delimiterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitUnconditional_statement(ctx *Unconditional_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitAssignment_statement(ctx *Assignment_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitLeft_part(ctx *Left_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitProcedure_statement(ctx *Procedure_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitActual_parameter_list(ctx *Actual_parameter_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitConditional_statement(ctx *Conditional_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitFor_statement(ctx *For_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitFor_list(ctx *For_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitFor_list_element(ctx *For_list_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitSimple_expression(ctx *Simple_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitFunction_designator(ctx *Function_designatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitBoolean_expression(ctx *Boolean_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitImplication(ctx *ImplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitBoolean_term(ctx *Boolean_termContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitBoolean_factor(ctx *Boolean_factorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitBoolean_secondary(ctx *Boolean_secondaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitBoolean_primary(ctx *Boolean_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitSubscripted_variable(ctx *Subscripted_variableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlgol60V2Visitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}
