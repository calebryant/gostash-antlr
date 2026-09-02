// Code generated from ChronicleLogstashParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ChronicleLogstashParser

import "github.com/antlr4-go/antlr/v4"

// BaseChronicleLogstashParserListener is a complete listener for a parse tree produced by ChronicleLogstashParser.
type BaseChronicleLogstashParserListener struct{}

var _ ChronicleLogstashParserListener = &BaseChronicleLogstashParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChronicleLogstashParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChronicleLogstashParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChronicleLogstashParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChronicleLogstashParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFilterblock is called when production filterblock is entered.
func (s *BaseChronicleLogstashParserListener) EnterFilterblock(ctx *FilterblockContext) {}

// ExitFilterblock is called when production filterblock is exited.
func (s *BaseChronicleLogstashParserListener) ExitFilterblock(ctx *FilterblockContext) {}

// EnterConditionalblock is called when production conditionalblock is entered.
func (s *BaseChronicleLogstashParserListener) EnterConditionalblock(ctx *ConditionalblockContext) {}

// ExitConditionalblock is called when production conditionalblock is exited.
func (s *BaseChronicleLogstashParserListener) ExitConditionalblock(ctx *ConditionalblockContext) {}

// EnterFor_var is called when production for_var is entered.
func (s *BaseChronicleLogstashParserListener) EnterFor_var(ctx *For_varContext) {}

// ExitFor_var is called when production for_var is exited.
func (s *BaseChronicleLogstashParserListener) ExitFor_var(ctx *For_varContext) {}

// EnterFor_iterable is called when production for_iterable is entered.
func (s *BaseChronicleLogstashParserListener) EnterFor_iterable(ctx *For_iterableContext) {}

// ExitFor_iterable is called when production for_iterable is exited.
func (s *BaseChronicleLogstashParserListener) ExitFor_iterable(ctx *For_iterableContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseChronicleLogstashParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseChronicleLogstashParserListener) ExitStatement(ctx *StatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseChronicleLogstashParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseChronicleLogstashParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterUnary_expression is called when production unary_expression is entered.
func (s *BaseChronicleLogstashParserListener) EnterUnary_expression(ctx *Unary_expressionContext) {}

// ExitUnary_expression is called when production unary_expression is exited.
func (s *BaseChronicleLogstashParserListener) ExitUnary_expression(ctx *Unary_expressionContext) {}

// EnterBinary_expression is called when production binary_expression is entered.
func (s *BaseChronicleLogstashParserListener) EnterBinary_expression(ctx *Binary_expressionContext) {}

// ExitBinary_expression is called when production binary_expression is exited.
func (s *BaseChronicleLogstashParserListener) ExitBinary_expression(ctx *Binary_expressionContext) {}

// EnterExpression_val is called when production expression_val is entered.
func (s *BaseChronicleLogstashParserListener) EnterExpression_val(ctx *Expression_valContext) {}

// ExitExpression_val is called when production expression_val is exited.
func (s *BaseChronicleLogstashParserListener) ExitExpression_val(ctx *Expression_valContext) {}

// EnterSigned_number is called when production signed_number is entered.
func (s *BaseChronicleLogstashParserListener) EnterSigned_number(ctx *Signed_numberContext) {}

// ExitSigned_number is called when production signed_number is exited.
func (s *BaseChronicleLogstashParserListener) ExitSigned_number(ctx *Signed_numberContext) {}

// EnterParen_list is called when production paren_list is entered.
func (s *BaseChronicleLogstashParserListener) EnterParen_list(ctx *Paren_listContext) {}

// ExitParen_list is called when production paren_list is exited.
func (s *BaseChronicleLogstashParserListener) ExitParen_list(ctx *Paren_listContext) {}

// EnterParen_value is called when production paren_value is entered.
func (s *BaseChronicleLogstashParserListener) EnterParen_value(ctx *Paren_valueContext) {}

// ExitParen_value is called when production paren_value is exited.
func (s *BaseChronicleLogstashParserListener) ExitParen_value(ctx *Paren_valueContext) {}

// EnterMath_statement is called when production math_statement is entered.
func (s *BaseChronicleLogstashParserListener) EnterMath_statement(ctx *Math_statementContext) {}

// ExitMath_statement is called when production math_statement is exited.
func (s *BaseChronicleLogstashParserListener) ExitMath_statement(ctx *Math_statementContext) {}

// EnterMath_expression is called when production math_expression is entered.
func (s *BaseChronicleLogstashParserListener) EnterMath_expression(ctx *Math_expressionContext) {}

// ExitMath_expression is called when production math_expression is exited.
func (s *BaseChronicleLogstashParserListener) ExitMath_expression(ctx *Math_expressionContext) {}

// EnterBoolean_op is called when production boolean_op is entered.
func (s *BaseChronicleLogstashParserListener) EnterBoolean_op(ctx *Boolean_opContext) {}

// ExitBoolean_op is called when production boolean_op is exited.
func (s *BaseChronicleLogstashParserListener) ExitBoolean_op(ctx *Boolean_opContext) {}

// EnterBoolean_eval is called when production boolean_eval is entered.
func (s *BaseChronicleLogstashParserListener) EnterBoolean_eval(ctx *Boolean_evalContext) {}

// ExitBoolean_eval is called when production boolean_eval is exited.
func (s *BaseChronicleLogstashParserListener) ExitBoolean_eval(ctx *Boolean_evalContext) {}

// EnterPlugin is called when production plugin is entered.
func (s *BaseChronicleLogstashParserListener) EnterPlugin(ctx *PluginContext) {}

// ExitPlugin is called when production plugin is exited.
func (s *BaseChronicleLogstashParserListener) ExitPlugin(ctx *PluginContext) {}

// EnterKeyvalue is called when production keyvalue is entered.
func (s *BaseChronicleLogstashParserListener) EnterKeyvalue(ctx *KeyvalueContext) {}

// ExitKeyvalue is called when production keyvalue is exited.
func (s *BaseChronicleLogstashParserListener) ExitKeyvalue(ctx *KeyvalueContext) {}

// EnterKv_lvalue is called when production kv_lvalue is entered.
func (s *BaseChronicleLogstashParserListener) EnterKv_lvalue(ctx *Kv_lvalueContext) {}

// ExitKv_lvalue is called when production kv_lvalue is exited.
func (s *BaseChronicleLogstashParserListener) ExitKv_lvalue(ctx *Kv_lvalueContext) {}

// EnterKv_rvalue is called when production kv_rvalue is entered.
func (s *BaseChronicleLogstashParserListener) EnterKv_rvalue(ctx *Kv_rvalueContext) {}

// ExitKv_rvalue is called when production kv_rvalue is exited.
func (s *BaseChronicleLogstashParserListener) ExitKv_rvalue(ctx *Kv_rvalueContext) {}

// EnterHash is called when production hash is entered.
func (s *BaseChronicleLogstashParserListener) EnterHash(ctx *HashContext) {}

// ExitHash is called when production hash is exited.
func (s *BaseChronicleLogstashParserListener) ExitHash(ctx *HashContext) {}

// EnterList is called when production list is entered.
func (s *BaseChronicleLogstashParserListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseChronicleLogstashParserListener) ExitList(ctx *ListContext) {}

// EnterIf_list is called when production if_list is entered.
func (s *BaseChronicleLogstashParserListener) EnterIf_list(ctx *If_listContext) {}

// ExitIf_list is called when production if_list is exited.
func (s *BaseChronicleLogstashParserListener) ExitIf_list(ctx *If_listContext) {}

// EnterList_value is called when production list_value is entered.
func (s *BaseChronicleLogstashParserListener) EnterList_value(ctx *List_valueContext) {}

// ExitList_value is called when production list_value is exited.
func (s *BaseChronicleLogstashParserListener) ExitList_value(ctx *List_valueContext) {}
