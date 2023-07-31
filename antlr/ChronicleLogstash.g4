grammar ChronicleLogstash;

/* PARSER DEFINITON */

filterblock: FILTER LBRACE (plugin|conditionalblock)* RBRACE EOF;

conditionalblock:
        (IF|ELSEIF) statement LBRACE (plugin|conditionalblock)* RBRACE
    |   ELSE LBRACE (plugin|conditionalblock)* RBRACE
    |   FOR (CONDITIONALID COMMA)? CONDITIONALID IN CONDITIONALID LBRACE (plugin|conditionalblock)* RBRACE
    ;

// Conditional statement rule definitions
statement:
		BOOLNOT? (binary_expression | unary_expression) (boolean_op statement)*
	|	BOOLNOT? LPAREN (binary_expression | unary_expression) (boolean_op statement)* RPAREN
	;

unary_expression:
		BOOLNOT? expression_val
	|	BOOLNOT? LPAREN (BOOLNOT? expression_val) RPAREN
	;

binary_expression:
		BOOLNOT? expression_val binary_eval BOOLNOT? expression_val
	|	BOOLNOT? LPAREN expression_val binary_eval expression_val RPAREN
	;

expression_val: CONDITIONALID|math_val|NUMBERS|STRING|REGEX|BOOLEAN;
math_val: (CONDITIONALID|NUMBERS) MATHOP (CONDITIONALID|NUMBERS);

boolean_op:
		AND
	|	OR
	;

binary_eval:
		EQUAL
	|	NOTEQUAL
	|	LESSTHAN
	|	GREATERTHAN
	|	LTEQUAL
	|	GTEQUAL
	|	REMATCH
	|	RENOTMATCH
	;

// filter plugin options
// plugins:
//         grok
//  	|	json
//  	|	xml
//  	|	kv
//  	|	csv
//  	|	mutate
//  	|	base64
//  	|	date
//  	|	drop
//  	|	statedump
// 	;

plugin: ID LBRACE keyvalue* RBRACE;

// // grok plugin rule definition
// grok: GROK LBRACE grokconfig RBRACE;
// grokconfig: ((grokmatch | overwrite | on_error) COMMA?)*;

// // json plugin rule definition
// json: JSON LBRACE jsonconfig RBRACE;
// jsonconfig: ((source | target | on_error | arrayfunction) COMMA?)*;

// // xml plugin rule definition
// xml: XML LBRACE xmlconfig RBRACE;
// xmlconfig: ((source | target | on_error | xpath) COMMA?)*;

// // kv plugin rule definition
// kv: KV LBRACE kvconfig RBRACE;
// kvconfig: ((source | target | fieldsplit | fieldsplitpattern | valuesplit | valuesplitpattern | whitespace | trim_value | on_error) COMMA?)*;

// // csv plugin rule definition
// csv: CSV LBRACE csvconfig RBRACE;
// csvconfig: ((source | target | separator | on_error) COMMA?)*;

// // base64 plugin rule definition
// base64: BASE64 LBRACE base64config RBRACE;
// base64config: ((source | target | encoding | on_error) COMMA?)*;

// // date plugin rule definition
// date: DATE LBRACE dateconfig RBRACE;
// dateconfig: ((datematch | target | timezone | rebase | on_error) COMMA?)*;

// // drop plugin rule definition
// drop: DROP LBRACE dropconfig? RBRACE;
// dropconfig: (tag COMMA?);

// // statedump plugin rule definition
// statedump: STATEDUMP LBRACE statedumpconfig? RBRACE;
// statedumpconfig: (label COMMA?);

// // mutate pligin rule definition
// mutate: MUTATE LBRACE mutateconfig RBRACE;
// mutateconfig: ((mutate_function|on_error) COMMA?)*;

// mutate function list
// mutate_function:
// 		convert
// 	|	merge
// 	|	rename
// 	|	replace
// 	|	copy
// 	|	gsub
// 	|	lowercase
// 	|	uppercase
// 	|	remove_field
// 	|	split
// 	;

// Function rule definitions
// grokmatch: MATCH KVSEPARATOR hash;
// datematch: MATCH KVSEPARATOR list;
// convert: CONVERT KVSEPARATOR LBRACE (keyvalue)* RBRACE;
// merge: MERGE KVSEPARATOR LBRACE (keyvalue)* RBRACE;
// rename: RENAME KVSEPARATOR LBRACE (keyvalue)* RBRACE;
// replace: REPLACE KVSEPARATOR LBRACE (keyvalue)* RBRACE;
// copy: COPY KVSEPARATOR LBRACE (keyvalue)* RBRACE;
// gsub: GSUB KVSEPARATOR list;
// lowercase: LOWERCASE KVSEPARATOR list;
// uppercase: UPPERCASE KVSEPARATOR list;
// remove_field: REMOVEFIELD KVSEPARATOR LBRACE list RBRACE;
// split: SPLIT KVSEPARATOR LBRACE (source | separator | target)* RBRACE;

// Function option rule definitions
// overwrite: OVERWRITE KVSEPARATOR list COMMA?;
// on_error: ONERROR KVSEPARATOR STRING;
// source: SOURCE KVSEPARATOR STRING;
// target: TARGET KVSEPARATOR STRING;
// arrayfunction: ARRAYFUNCTION KVSEPARATOR STRING;
// xpath: XPATH KVSEPARATOR hash;
// fieldsplit: FIELDSPLIT KVSEPARATOR STRING;
// valuesplit: VALUESPLIT KVSEPARATOR STRING;
// fieldsplitpattern: FIELDSPLITPATTERN KVSEPARATOR STRING;
// valuesplitpattern: VALUESPLITPATTERN KVSEPARATOR STRING;
// whitespace: WHITESPACE KVSEPARATOR STRING;
// trim_value: TRIMVALUE KVSEPARATOR STRING;
// separator: SEPARATOR KVSEPARATOR STRING;
// encoding: ENCODING KVSEPARATOR STRING;
// timezone: TIMEZONE KVSEPARATOR STRING;
// rebase: REBASE KVSEPARATOR BOOLEAN;
// tag: TAG KVSEPARATOR STRING;
// label: LABEL KVSEPARATOR STRING;

keyvalue: (ID|STRING) KVSEPARATOR (STRING|NUMBERS|BOOLEAN|ID|list|hash) COMMA?;

list: LBRACKET (listval | COMMA)* RBRACKET;

listval: STRING|NUMBERS|BOOLEAN|ID;

hash: LBRACE (keyvalue)* RBRACE;

/* LEXER DEFINITION */

WS: [ \t\n\r]+ -> skip; // whitespace
COMMENT: ('#' .*? [\n\r]) -> skip;

// Keyword token definitions
FILTER: 'filter';
// GROK: 'grok';
// JSON: 'json';
// XML: 'xml';
// KV: 'kv';
// CSV: 'csv';
// MUTATE: 'mutate';
// CONVERT: 'convert';
// GSUB: 'gsub';
// LOWERCASE: 'lowercase';
// MERGE: 'merge';
// RENAME: 'rename';
// REPLACE: 'replace';
// UPPERCASE: 'uppsercase';
// REMOVEFIELD: 'remove_field';
// COPY: 'copy';
// SPLIT: 'split';
// BASE64: 'base64';
// DATE: 'date';
// DROP: 'drop';
// STATEDUMP: 'statedump';
IF: 'if';
ELSEIF: 'else if';
ELSE: 'else';
FOR: 'for';
IN: 'in';
AND: 'and' | '&&';
OR: 'or' | '||';
BOOLEAN: TRUE | FALSE;
fragment TRUE: 'true';
fragment FALSE: 'false';
// ONERROR: 'on_error';
// SOURCE: 'source';
// TARGET: 'target';
// FIELDSPLIT: 'field_split';
// FIELDSPLITPATTERN: 'field_split_pattern';
// VALUESPLIT: 'value_split';
// VALUESPLITPATTERN: 'value_split_pattern';
// WHITESPACE: 'whitespace';
// TRIMVALUE: 'trim_value';
// XPATH: 'xpath';
// ARRAYFUNCTION: 'array_function';
// MATCH: 'match';
// OVERWRITE: 'overwrite';
// SEPARATOR: 'separator';
// ENCODING: 'encoding';
// TIMEZONE: 'timezone';
// REBASE: 'rebase';
// TAG: 'tag';
// LABEL: 'label';

// Literal token definitions
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';
LPAREN: '(';
RPAREN: ')';
KVSEPARATOR: ARROW | HEADLESSARROW | COLON;
fragment ARROW: '=>';
fragment HEADLESSARROW: '=';
fragment COLON: ':';
COMMA: ',';
PERCENT: '%';
BOOLNOT: '!';
EQUAL: '==';
NOTEQUAL: '!=';
LESSTHAN: '<';
GREATERTHAN: '>';
LTEQUAL: '<=';
GTEQUAL: '>=';
REMATCH: '=~';
RENOTMATCH: '!~';
MATHOP: [+\-*/];

CONDITIONALID: (LBRACKET (LETTERS+ | INTEGER | [_\-@]+) RBRACKET)+;
ID: LETTERS (LETTERS | INTEGER | [_])+;
NUMBERS: INTEGER ([.] INTEGER)?;
fragment INTEGER: [0-9]+;
fragment LETTERS: [a-zA-Z];

STRING: STRING_DOUBLE | STRING_SINGLE;
fragment STRING_DOUBLE: '"' ( ESC | ~["\r\n] )* '"';
fragment STRING_SINGLE: '\'' ( ESC | ~['\r\n] )* '\'';
fragment ESC: '\\' ["'\\];
REGEX: ( '/' ( ESC | ~[/\r\n] )*? '/' ) | ( '\\/' ( ESC | ~[/\r\n] )*? '\\/' );