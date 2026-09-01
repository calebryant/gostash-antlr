parser grammar ChronicleLogstashParser;

options { tokenVocab=ChronicleLogstashLexer; }

/* PARSER DEFINITON */

filterblock: ID LBRACE (plugin|conditionalblock)* RBRACE EOF;

conditionalblock:
        (IF|ELSEIF) statement LBRACE (plugin|conditionalblock)* RBRACE
    |   ELSE LBRACE (plugin|conditionalblock)* RBRACE
    |   FOR (FORID FORCOMMA)? FORID FORIN FORID FOROPENER (plugin|conditionalblock)* RBRACE
    ;

// Conditional statement rule definitions
statement:
		LPAREN statement RPAREN
	|	LBRACKET statement RBRACKET
	|	statement boolean_op statement
	|	BOOLNOT statement
	|	NOT statement
	|	expression
	;

expression:
		binary_expression
	|	unary_expression
	// |	is_in_expression
	;

unary_expression:
		IFSTATEMENTID
	|	BOOLEAN
	;

// is_in_expression:
// 		IFSTATEMENTID NOT? IN list
// 	|	IFSTATEMENTID NOT? IN LPAREN list RPAREN
// 	;

binary_expression: expression_val boolean_eval expression_val;

expression_val: math_statement|NUMBER|if_list|IFSTATEMENTID|STRING|REGEX|BOOLEAN|ID|LPAREN REGEX RPAREN|LPAREN STRING RPAREN;

math_statement:
		LPAREN math_statement RPAREN
	|	math_statement MATHOP math_statement
	|	math_expression
	;

math_expression: (IFSTATEMENTID|NUMBER) MATHOP (IFSTATEMENTID|NUMBER);

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
	|	MATCH
	|	NOTMATCH
	|	IN
	|	NOT IN
	;

plugin: ID LBRACE keyvalue* RBRACE;

keyvalue: kv_lvalue KVSEPARATOR kv_rvalue COMMA?;

kv_lvalue: (ID|STRING);

kv_rvalue: (NUMBER|list|hash|STRING|BOOLEAN|ID);

hash: LBRACE (keyvalue)* RBRACE;

list: LBRACKET (list_value (list_value)*)? RBRACKET;

if_list: LBRACKET ((STRING | BOOLEAN | NUMBER) (list_value)*)? RBRACKET; // lists in if statements cannot start with an ID to avoid ambiguity with if statement IDs

list_value: STRING | ID | BOOLEAN | NUMBER | COMMA;