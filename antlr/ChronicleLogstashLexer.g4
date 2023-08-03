lexer grammar ChronicleLogstashLexer;

/* LEXER DEFINITION */

WS: [ \t\n\r]+ -> skip; // whitespace
COMMENT: ('#' ~[\n\r]*) -> skip;

// Keyword token definitions
IF: 'if';
ELSEIF: 'else if';
ELSE: 'else';
FOR: 'for';
IN: 'in';
fragment TRUE: 'true';
fragment FALSE: 'false';

// Literal token definitions
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';
LPAREN: '(';
RPAREN: ')';
KVSEPARATOR: '=>' | '=' | ':';
COMMA: ',';
NOT: '!';
MATHOP: [+\-*/];
EQUAL: '==';
NOTEQUAL: '!=';
LESSTHAN: '<';
GREATERTHAN: '>';
LTEQUAL: '<=';
GTEQUAL: '>=';
REMATCH: '=~';
RENOTMATCH: '!~';
AND: 'and' | '&&';
OR: 'or' | '||';

// Value token definitions
BOOLEAN: TRUE | FALSE;
STRING: STRING_DOUBLE | STRING_SINGLE;
fragment STRING_DOUBLE: '"' ( ESC | ~["\r\n] )* '"';
fragment STRING_SINGLE: '\'' ( ESC | ~['\r\n] )* '\'';
fragment ESC: '\\' ["'\\];
REGEX: ( '/' ( ESC | ~[/\r\n] )*? '/' ) | ( '\\/' ( ESC | ~[/\r\n] )*? '\\/' );
ID: LETTERS (LETTERS | DIGIT | [_])*;
FLOAT: (DIGIT* [.] DIGIT+);
INTEGER: DIGIT+;
fragment DIGIT: [0-9];
fragment LETTERS: [a-zA-Z];