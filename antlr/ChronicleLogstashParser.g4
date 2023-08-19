parser grammar ChronicleLogstashParser;

options { tokenVocab=ChronicleLogstashLexer; }

/* PARSER DEFINITON */

filterblock: ID LBRACE (plugin|conditionalblock)* RBRACE EOF;

conditionalblock:
        (IF|ELSEIF) statement LBRACE (plugin|conditionalblock)* RBRACE
    |   ELSE LBRACE (plugin|conditionalblock)* RBRACE
    |   FOR (FORID FORCOMMA)? FORID FORIN FORID FOROPENER (plugin|conditionalblock)* RBRACE
    // |   FOR ((ID|UNDERSCORE) COMMA)? ID IN (ID (DOT ID)*) LBRACE (plugin|conditionalblock)* RBRACE
    ;

// Conditional statement rule definitions
statement:
		LPAREN statement RPAREN
	|	LBRACKET statement RBRACKET
	|	statement boolean_op statement
	|	BOOLNOT statement
	|	expression
	;

expression:
		binary_expression
	|	unary_expression
	|	is_in_expression
	;

unary_expression:
		IFSTATEMENTID
	|	BOOLEAN
	;

is_in_expression:
		IFSTATEMENTID NOT? IN list
	|	IFSTATEMENTID NOT? IN LPAREN list RPAREN
	;

binary_expression: expression_val boolean_eval expression_val;

expression_val: math_statement|number|if_list|IFSTATEMENTID|STRING|REGEX|BOOLEAN|ID;

math_statement:
		LPAREN math_statement RPAREN
	|	math_statement MATHOP math_statement
	|	math_expression
	;

math_expression: (IFSTATEMENTID|number) MATHOP (IFSTATEMENTID|number);

// if_statement_id: (LBRACKET (ID|INTEGER|UNDERSCORE)+ RBRACKET)+;

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
	;

plugin: ID LBRACE keyvalue* RBRACE;

keyvalue: kv_lvalue KVSEPARATOR kv_rvalue COMMA?;

kv_lvalue: (ID|STRING);

kv_rvalue: (number|list|hash|STRING|BOOLEAN|ID);

hash: LBRACE (keyvalue)* RBRACE;

list: LBRACKET (list_value (list_value)*)? RBRACKET;

if_list: LBRACKET ((STRING | BOOLEAN | number) (list_value)*)? RBRACKET; // lists in if statements cannot start with an ID to avoid ambiguity with if statement IDs

list_value: STRING | ID | BOOLEAN | number | COMMA;

number: INTEGER | FLOAT;