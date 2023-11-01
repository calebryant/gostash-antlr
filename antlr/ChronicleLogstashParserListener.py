# Generated from /Users/caleb.bryant/Library/CloudStorage/OneDrive-Cyderes/TelEng/github/chronicle-parser-language-server/antlr/ChronicleLogstashParser.g4 by ANTLR 4.13.0
from antlr4 import *
if "." in __name__:
    from .ChronicleLogstashParser import ChronicleLogstashParser
else:
    from ChronicleLogstashParser import ChronicleLogstashParser

# This class defines a complete listener for a parse tree produced by ChronicleLogstashParser.
class ChronicleLogstashParserListener(ParseTreeListener):

    # Enter a parse tree produced by ChronicleLogstashParser#filterblock.
    def enterFilterblock(self, ctx:ChronicleLogstashParser.FilterblockContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#filterblock.
    def exitFilterblock(self, ctx:ChronicleLogstashParser.FilterblockContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#conditionalblock.
    def enterConditionalblock(self, ctx:ChronicleLogstashParser.ConditionalblockContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#conditionalblock.
    def exitConditionalblock(self, ctx:ChronicleLogstashParser.ConditionalblockContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#statement.
    def enterStatement(self, ctx:ChronicleLogstashParser.StatementContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#statement.
    def exitStatement(self, ctx:ChronicleLogstashParser.StatementContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#expression.
    def enterExpression(self, ctx:ChronicleLogstashParser.ExpressionContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#expression.
    def exitExpression(self, ctx:ChronicleLogstashParser.ExpressionContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#unary_expression.
    def enterUnary_expression(self, ctx:ChronicleLogstashParser.Unary_expressionContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#unary_expression.
    def exitUnary_expression(self, ctx:ChronicleLogstashParser.Unary_expressionContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#binary_expression.
    def enterBinary_expression(self, ctx:ChronicleLogstashParser.Binary_expressionContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#binary_expression.
    def exitBinary_expression(self, ctx:ChronicleLogstashParser.Binary_expressionContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#expression_val.
    def enterExpression_val(self, ctx:ChronicleLogstashParser.Expression_valContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#expression_val.
    def exitExpression_val(self, ctx:ChronicleLogstashParser.Expression_valContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#math_statement.
    def enterMath_statement(self, ctx:ChronicleLogstashParser.Math_statementContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#math_statement.
    def exitMath_statement(self, ctx:ChronicleLogstashParser.Math_statementContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#math_expression.
    def enterMath_expression(self, ctx:ChronicleLogstashParser.Math_expressionContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#math_expression.
    def exitMath_expression(self, ctx:ChronicleLogstashParser.Math_expressionContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#boolean_op.
    def enterBoolean_op(self, ctx:ChronicleLogstashParser.Boolean_opContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#boolean_op.
    def exitBoolean_op(self, ctx:ChronicleLogstashParser.Boolean_opContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#boolean_eval.
    def enterBoolean_eval(self, ctx:ChronicleLogstashParser.Boolean_evalContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#boolean_eval.
    def exitBoolean_eval(self, ctx:ChronicleLogstashParser.Boolean_evalContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#plugin.
    def enterPlugin(self, ctx:ChronicleLogstashParser.PluginContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#plugin.
    def exitPlugin(self, ctx:ChronicleLogstashParser.PluginContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#keyvalue.
    def enterKeyvalue(self, ctx:ChronicleLogstashParser.KeyvalueContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#keyvalue.
    def exitKeyvalue(self, ctx:ChronicleLogstashParser.KeyvalueContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#kv_lvalue.
    def enterKv_lvalue(self, ctx:ChronicleLogstashParser.Kv_lvalueContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#kv_lvalue.
    def exitKv_lvalue(self, ctx:ChronicleLogstashParser.Kv_lvalueContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#kv_rvalue.
    def enterKv_rvalue(self, ctx:ChronicleLogstashParser.Kv_rvalueContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#kv_rvalue.
    def exitKv_rvalue(self, ctx:ChronicleLogstashParser.Kv_rvalueContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#hash.
    def enterHash(self, ctx:ChronicleLogstashParser.HashContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#hash.
    def exitHash(self, ctx:ChronicleLogstashParser.HashContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#list.
    def enterList(self, ctx:ChronicleLogstashParser.ListContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#list.
    def exitList(self, ctx:ChronicleLogstashParser.ListContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#if_list.
    def enterIf_list(self, ctx:ChronicleLogstashParser.If_listContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#if_list.
    def exitIf_list(self, ctx:ChronicleLogstashParser.If_listContext):
        pass


    # Enter a parse tree produced by ChronicleLogstashParser#list_value.
    def enterList_value(self, ctx:ChronicleLogstashParser.List_valueContext):
        pass

    # Exit a parse tree produced by ChronicleLogstashParser#list_value.
    def exitList_value(self, ctx:ChronicleLogstashParser.List_valueContext):
        pass



del ChronicleLogstashParser