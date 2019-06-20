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

POSITIVE_INT:   [0-9]+;
NEGATIVE_INT:   '-' POSITIVE_INT;
POSITIVE_FLOAT: [0-9]+ | [0-9]* '.' [0-9]+;
NEGATIVE_FLOAT: '-' POSITIVE_FLOAT;

// Properties

PROP_ANTIALIASING:        'antialiasing';
PROP_ANTIALIASING_VALUES: ('full' | 'text' | 'none');

PROP_FILL_OPACITY: 'fill-opacity';

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
    : PROP_ANTIALIASING COLON antialiasing? SEMICOLON
    | PROP_FILL_OPACITY COLON fill_opacity? SEMICOLON
    ;

// Properties

antialiasing
    : PROP_ANTIALIASING_VALUES
    ;

fill_opacity
    : POSITIVE_INT
    | POSITIVE_FLOAT
    ;
