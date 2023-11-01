lexer grammar ChronicleLogstashLexer;

/* LEXER DEFINITION */

WS: [ \t\n\r]+ -> skip; // whitespace
COMMENT: ('#' ~[\n\r]*) -> skip;

// Keyword token definitions
IF: 'if';
ELSEIF: 'else if';
ELSE: 'else';
FOR: 'for' -> pushMode(FORMODE);
IN: 'in';
NOT: 'not';
fragment TRUE: 'true';
fragment FALSE: 'false';

// Literal token definitions
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';
LPAREN: '(';
RPAREN: ')';
KVSEPARATOR: ( '=>' | '=' | ':' );
COMMA: ',';
BOOLNOT: '!';
MATHOP: [+\-*/];
EQUAL: '==';
NOTEQUAL: '!=';
LESSTHAN: '<';
GREATERTHAN: '>';
LTEQUAL: '<=';
GTEQUAL: '>=';
MATCH: '=~';
NOTMATCH: '!~';
AND: 'and' | '&&';
OR: 'or' | '||';

// Value token definitions
BOOLEAN: TRUE | FALSE;
STRING: STRING_DOUBLE | STRING_SINGLE;
fragment STRING_DOUBLE: '"' ( ESC | ~["\\\r\n] )* '"' ;
fragment STRING_SINGLE: '\'' ( ESC | ~['\\\r\n] )* '\'' ;
fragment ESC: '\\' ~[ \n\r] ;
REGEX: ( '/' ( ESC | ~[/\r\n] )*? '/' ) | ( '\\/' ( ESC | ~[/\r\n] )*? '\\/' );
IFSTATEMENTID: (LBRACKET (LETTERS|DIGIT+|[_\-@])+ RBRACKET)+;
ID: LETTERS (LETTERS | DIGIT | [_])*;
NUMBER: (DIGIT* [.] DIGIT+) | DIGIT+;
fragment DIGIT: [0-9];
fragment LETTERS: [a-zA-Z];

// Mode for for statements
mode FORMODE;
FORWS: WS -> skip ; // whitespace
FORCOMMA: COMMA ;
FORIN: IN ;
FORID: ~[ ,{}\t\n\r]+ ;
FOROPENER: '{' -> popMode ;