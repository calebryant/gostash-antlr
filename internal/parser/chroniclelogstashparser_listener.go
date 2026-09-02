// Code generated from ChronicleLogstashParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ChronicleLogstashParser

import "github.com/antlr4-go/antlr/v4"

// ChronicleLogstashParserListener is a complete listener for a parse tree produced by ChronicleLogstashParser.
type ChronicleLogstashParserListener interface {
	antlr.ParseTreeListener

	// EnterFilterblock is called when entering the filterblock production.
	EnterFilterblock(c *FilterblockContext)

	// EnterConditionalblock is called when entering the conditionalblock production.
	EnterConditionalblock(c *ConditionalblockContext)

	// EnterFor_var is called when entering the for_var production.
	EnterFor_var(c *For_varContext)

	// EnterFor_iterable is called when entering the for_iterable production.
	EnterFor_iterable(c *For_iterableContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterUnary_expression is called when entering the unary_expression production.
	EnterUnary_expression(c *Unary_expressionContext)

	// EnterBinary_expression is called when entering the binary_expression production.
	EnterBinary_expression(c *Binary_expressionContext)

	// EnterExpression_val is called when entering the expression_val production.
	EnterExpression_val(c *Expression_valContext)

	// EnterSigned_number is called when entering the signed_number production.
	EnterSigned_number(c *Signed_numberContext)

	// EnterParen_list is called when entering the paren_list production.
	EnterParen_list(c *Paren_listContext)

	// EnterParen_value is called when entering the paren_value production.
	EnterParen_value(c *Paren_valueContext)

	// EnterMath_statement is called when entering the math_statement production.
	EnterMath_statement(c *Math_statementContext)

	// EnterMath_expression is called when entering the math_expression production.
	EnterMath_expression(c *Math_expressionContext)

	// EnterBoolean_op is called when entering the boolean_op production.
	EnterBoolean_op(c *Boolean_opContext)

	// EnterBoolean_eval is called when entering the boolean_eval production.
	EnterBoolean_eval(c *Boolean_evalContext)

	// EnterPlugin is called when entering the plugin production.
	EnterPlugin(c *PluginContext)

	// EnterKeyvalue is called when entering the keyvalue production.
	EnterKeyvalue(c *KeyvalueContext)

	// EnterKv_lvalue is called when entering the kv_lvalue production.
	EnterKv_lvalue(c *Kv_lvalueContext)

	// EnterKv_rvalue is called when entering the kv_rvalue production.
	EnterKv_rvalue(c *Kv_rvalueContext)

	// EnterHash is called when entering the hash production.
	EnterHash(c *HashContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterIf_list is called when entering the if_list production.
	EnterIf_list(c *If_listContext)

	// EnterList_value is called when entering the list_value production.
	EnterList_value(c *List_valueContext)

	// ExitFilterblock is called when exiting the filterblock production.
	ExitFilterblock(c *FilterblockContext)

	// ExitConditionalblock is called when exiting the conditionalblock production.
	ExitConditionalblock(c *ConditionalblockContext)

	// ExitFor_var is called when exiting the for_var production.
	ExitFor_var(c *For_varContext)

	// ExitFor_iterable is called when exiting the for_iterable production.
	ExitFor_iterable(c *For_iterableContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitUnary_expression is called when exiting the unary_expression production.
	ExitUnary_expression(c *Unary_expressionContext)

	// ExitBinary_expression is called when exiting the binary_expression production.
	ExitBinary_expression(c *Binary_expressionContext)

	// ExitExpression_val is called when exiting the expression_val production.
	ExitExpression_val(c *Expression_valContext)

	// ExitSigned_number is called when exiting the signed_number production.
	ExitSigned_number(c *Signed_numberContext)

	// ExitParen_list is called when exiting the paren_list production.
	ExitParen_list(c *Paren_listContext)

	// ExitParen_value is called when exiting the paren_value production.
	ExitParen_value(c *Paren_valueContext)

	// ExitMath_statement is called when exiting the math_statement production.
	ExitMath_statement(c *Math_statementContext)

	// ExitMath_expression is called when exiting the math_expression production.
	ExitMath_expression(c *Math_expressionContext)

	// ExitBoolean_op is called when exiting the boolean_op production.
	ExitBoolean_op(c *Boolean_opContext)

	// ExitBoolean_eval is called when exiting the boolean_eval production.
	ExitBoolean_eval(c *Boolean_evalContext)

	// ExitPlugin is called when exiting the plugin production.
	ExitPlugin(c *PluginContext)

	// ExitKeyvalue is called when exiting the keyvalue production.
	ExitKeyvalue(c *KeyvalueContext)

	// ExitKv_lvalue is called when exiting the kv_lvalue production.
	ExitKv_lvalue(c *Kv_lvalueContext)

	// ExitKv_rvalue is called when exiting the kv_rvalue production.
	ExitKv_rvalue(c *Kv_rvalueContext)

	// ExitHash is called when exiting the hash production.
	ExitHash(c *HashContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitIf_list is called when exiting the if_list production.
	ExitIf_list(c *If_listContext)

	// ExitList_value is called when exiting the list_value production.
	ExitList_value(c *List_valueContext)
}
