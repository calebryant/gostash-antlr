parser grammar ChronicleLogstashParser;

options { tokenVocab=ChronicleLogstashLexer; }

/* PARSER DEFINITON */

filterblock: ID LBRACE (plugin|conditionalblock)* RBRACE EOF;

conditionalblock:
        (IF|ELSEIF) statement LBRACE (plugin|conditionalblock)* RBRACE
    |   ELSE LBRACE (plugin|conditionalblock)* RBRACE
    |   FOR (ID COMMA)? ID IN ID LBRACE (plugin|conditionalblock)* RBRACE
    ;

// Conditional statement rule definitions
statement:
		LPAREN statement RPAREN
	|	statement boolean_op statement
	|	NOT statement
	|	expression
	;

expression: 
		binary_expression
	|	unary_expression
	;

unary_expression:
		if_statement_id
	|	BOOLEAN
	;

binary_expression: expression_val boolean_eval expression_val;

expression_val: if_statement_id|math_statement|number|list|STRING|REGEX|BOOLEAN;

math_statement: 
		LPAREN math_statement RPAREN
	|	math_statement MATHOP math_statement
	|	math_expression
	;

math_expression: (if_statement_id|number) MATHOP (if_statement_id|number);

if_statement_id: (LBRACKET (ID|INTEGER)+ RBRACKET)+;

boolean_op:
		AND
	|	OR
	;

boolean_eval:
		EQUAL
	|	NOTEQUAL
	|	LESSTHAN
	|	GREATERTHAN
	|	LTEQUAL
	|	GTEQUAL
	|	REMATCH
	|	RENOTMATCH
	|	IN
	;

plugin: ID LBRACE keyvalue* RBRACE;

keyvalue: kv_lvalue KVSEPARATOR kv_rvalue COMMA?;

kv_lvalue: (ID|STRING);

kv_rvalue: (number|list|hash|STRING|BOOLEAN|ID);

hash: LBRACE (keyvalue)* RBRACE;

list: LBRACKET listval (COMMA? listval)* RBRACKET;

listval: STRING;

number: INTEGER | FLOAT;