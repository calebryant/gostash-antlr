parser grammar ChronicleLogstashParser;

options { tokenVocab=ChronicleLogstashLexer; }

/* PARSER DEFINITON */

filterblock: ID LBRACE (plugin|conditionalblock)* RBRACE EOF;

conditionalblock:
        (IF|ELSEIF) statement LBRACE (plugin|conditionalblock)* RBRACE
    |   ELSE LBRACE (plugin|conditionalblock)* RBRACE
    |   FOR for_var (FORCOMMA for_var)? FORIN for_iterable FOROPENER (plugin|conditionalblock)* RBRACE
    ;

// `in` is also used as a loop variable name in the wild (`for in, file in x`).
for_var: FORID | FORIN;

// The iterable is either a field path, a field path followed by the `map`
// keyword, or a function call such as `xml(message, /Event/Data)` whose
// argument list contains commas.
for_iterable: (FORID | FORCOMMA)+;

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

// Comparisons chain, e.g. `[message] =~ "details.*?etag" != ""`.
binary_expression: expression_val (boolean_eval expression_val)+;

expression_val:
		math_statement
	|	signed_number
	|	NUMBER
	|	if_list
	|	IFSTATEMENTID
	|	STRING
	|	REGEX
	|	BOOLEAN
	|	ID
	|	paren_list
	;

// Negative literals, e.g. `[dpt] >= -2147483648`.
signed_number: MATHOP NUMBER;

// Parenthesized value list, used by `in`/`not in`, e.g. `in ("GET", "POST")`.
// A single-element list also covers the old `(REGEX)` and `(STRING)` forms.
paren_list: LPAREN paren_value (COMMA paren_value)* RPAREN;

paren_value: STRING|REGEX|BOOLEAN|NUMBER|signed_number|ID|IFSTATEMENTID|if_list;

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

kv_rvalue: (NUMBER|list|hash|STRING|BOOLEAN|ID|IFSTATEMENTID);

hash: LBRACE (keyvalue)* RBRACE;

list: LBRACKET (list_value (list_value)*)? RBRACKET;

if_list: LBRACKET ((STRING | BOOLEAN | NUMBER) (list_value)*)? RBRACKET; // lists in if statements cannot start with an ID to avoid ambiguity with if statement IDs

list_value: STRING | ID | BOOLEAN | NUMBER | IFSTATEMENTID | COMMA;