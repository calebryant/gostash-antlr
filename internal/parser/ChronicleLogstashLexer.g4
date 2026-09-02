lexer grammar ChronicleLogstashLexer;

/* LEXER DEFINITION */

WS: [ \t\n\r\f\u00A0\uFEFF]+ -> skip; // whitespace (configs in the wild contain NBSP and BOM)
COMMENT: ('#' ~[\n\r]*) -> skip;

// Keyword token definitions
IF: 'if';
ELSEIF: 'else if';
ELSE: 'else';
FOR: 'for' -> pushMode(FORMODE);
IN: 'in' | 'IN';
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
// Strings may span multiple lines: grok patterns and long regex alternations are
// routinely wrapped across lines in these configs.
STRING: STRING_DOUBLE | STRING_SINGLE;
fragment STRING_DOUBLE: '"' ( ESC | ~["\\] )* '"' ;
fragment STRING_SINGLE: '\'' ( ESC | ~['\\] )* '\'' ;
fragment ESC: '\\' . ; // any escaped character, including '\ ' and an escaped newline
REGEX: ( '/' ( ESC | ~[/\\\r\n] )*? '/' ) | ( '\\/' ( ESC | ~[/\\\r\n] )*? '\\/' );
// Field references. Inner padding is allowed, e.g. `[ title ]`.
IFSTATEMENTID: (LBRACKET HSPACE* (LETTERS|DIGIT|[_\-@.])+ HSPACE* RBRACKET)+;
fragment HSPACE: [ \t];
ID: LETTERS (LETTERS | DIGIT | [_])*;
NUMBER: (DIGIT* [.] DIGIT+) | DIGIT+;
fragment DIGIT: [0-9];
fragment LETTERS: [a-zA-Z];

// Mode for for statements
mode FORMODE;
FORWS: WS -> skip ; // whitespace
FORCOMMA: COMMA ;
FORIN: IN ;
// A for-header word. `%{...}` interpolations are consumed whole so that braces
// inside an xpath (e.g. /violation[%{ridx}]/sig_data) do not end the header.
FORID: ( '%{' ~[}\r\n]* '}' | ~[ ,{}\t\n\r] )+ ;
FOROPENER: '{' -> popMode ;