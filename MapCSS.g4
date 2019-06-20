grammar MapCSS;

WS: (' ' | '\t' | '\n' | '\r' | '\f') -> skip;

LBRACE:    '{';
RBRACE:    '}';
COLON:     ':';
SEMICOLON: ';';

fragment EBACKSLASH: '\\\\';
fragment UNICODE:    '\u0080'..'\uFFFE';
fragment EDQUOTE:    '\\"';
fragment ESQUOTE:    '\\\'';
DQUOTED_STRING:      '"' (' ' | '!' | '#'..'[' | ']'..'~' | '°' | UNICODE | EDQUOTE | EBACKSLASH)* '"';
SQUOTED_STRING:      '\'' (' '..'&' | '('..'[' | ']'..'~' | '°' | UNICODE | ESQUOTE | EBACKSLASH)* '\'';

PROP_ANTIALIASING:        'antialiasing';
PROP_ANTIALIASING_VALUES: ('full' | 'text' | 'none');

// Structure

stylesheet
    : entry* EOF
    ;

entry
    : rule_
    ;

rule_
    : canvas_declaration_block
    ;

canvas_declaration_block
    : 'canvas' LBRACE (canvas_declaration)+ RBRACE
    ;

canvas_declaration
    : PROP_ANTIALIASING COLON antialiasing SEMICOLON
    ;

// Properties

antialiasing
    : PROP_ANTIALIASING_VALUES
    ;
