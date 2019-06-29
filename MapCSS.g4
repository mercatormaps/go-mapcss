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

// Colors

fragment HEX_CHAR:     [0-9a-fA-F];
fragment HEX_3_DIGITS: '#' HEX_CHAR HEX_CHAR HEX_CHAR;
fragment HEX_4_DIGITS: '#' HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR;
fragment HEX_6_DIGITS: '#' HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR;
fragment HEX_8_DIGITS: '#' HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR;
HEX:                   (HEX_3_DIGITS | HEX_4_DIGITS | HEX_6_DIGITS | HEX_8_DIGITS);

//RGBA: 'rgba(' POSITIVE_INT ',' POSITIVE_INT ',' POSITIVE_INT ',' ALPHA ')';

// Properties

PROP_ANTIALIASING:      'antialiasing';
PROP_ANTIALIASING_FULL: 'full';
PROP_ANTIALIASING_TEXT: 'text';
PROP_ANTIALIASING_NONE: 'none';

PROP_FILL_OPACITY: 'fill-opacity';

PROP_FILL_COLOR: 'fill-color';

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
    | PROP_FILL_OPACITY COLON fill_opacity SEMICOLON
    | PROP_FILL_COLOR COLON fill_color SEMICOLON
    ;

// Properties

antialiasing
    : PROP_ANTIALIASING_FULL
    | PROP_ANTIALIASING_TEXT
    | PROP_ANTIALIASING_NONE
    ;

fill_opacity
    : POSITIVE_INT
    | POSITIVE_FLOAT
    ;

fill_color
    : color
    ;

alpha
    : POSITIVE_INT
    ;

color
    : HEX
    | rgb_color
    | rgba_color
    ;

rgb_color
    : 'rgb(' r=POSITIVE_INT ',' g=POSITIVE_INT ',' b=POSITIVE_INT ')'
    ;

rgba_color
    : 'rgba(' r=POSITIVE_INT ',' g=POSITIVE_INT ',' b=POSITIVE_INT ',' a=(POSITIVE_INT | POSITIVE_FLOAT) ')'
    ;
