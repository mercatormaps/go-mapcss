grammar MapCSS;

fragment COMMENT: ('/*' .*? '*/') | ('//' .*? '\n');
WS: (' ' | '\t' | '\n' | '\r' | '\f' | COMMENT) -> skip;

LBRACE:    '{';
RBRACE:    '}';
COLON:     ':';
SEMICOLON: ';';
HYPHEN:    '-';

fragment EBACKSLASH: '\\\\';
fragment UNICODE:    '\u0080'..'\uFFFE';
fragment EDQUOTE:    '\\"';
fragment ESQUOTE:    '\\\'';
DQUOTED_STRING:      '"' (' ' | '!' | '#'..'[' | ']'..'~' | '°' | UNICODE | EDQUOTE | EBACKSLASH)* '"';
SQUOTED_STRING:      '\'' (' '..'&' | '('..'[' | ']'..'~' | '°' | UNICODE | ESQUOTE | EBACKSLASH)* '\'';

DIGIT: [0-9];

// Colors

fragment HEX_CHAR:     [0-9a-fA-F];
fragment HEX_3_DIGITS: '#' HEX_CHAR HEX_CHAR HEX_CHAR;
fragment HEX_4_DIGITS: '#' HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR;
fragment HEX_6_DIGITS: '#' HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR;
fragment HEX_8_DIGITS: '#' HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR;
HEX:                   (HEX_3_DIGITS | HEX_4_DIGITS | HEX_6_DIGITS | HEX_8_DIGITS);

// Numbers

int_
    : '-'? DIGIT+
    ;

uint_
    : DIGIT+
    ;

float_
    : '-'? DIGIT+ | DIGIT* '.' DIGIT+;

// Properties

IDENTIFIER: [A-Za-z]+ [A-Za-z0-9\-]*;

// Structure

stylesheet
    : entry* EOF
    ;

entry
    : canvas_rule
    | rule_
    ;

// Selectors

rule_
    : selector (',' selector)+ decl_block
    ;

selector
    : typ=IDENTIFIER zoom? (attribute)+
    ;

zoom
    : zoom_range
    | min_zoom
    | exact_zoom
    ;

zoom_range
    : '|z' min=uint_ HYPHEN max=uint_
    ;

min_zoom
    : '|z' min=uint_ '-'
    ;

exact_zoom
    : '|z' min=uint_
    ;

attribute
    : '[' neg='!'? name=IDENTIFIER ']'
    ;

decl_block
    : LBRACE RBRACE
    ;

// Canvas rule

canvas_rule
    : 'canvas' canvas_decl_block
    ;

canvas_decl_block
    : LBRACE (canvas_decl)+ RBRACE
    ;

canvas_decl
    : antialiasing_decl
    | fill_opacity_decl
    | fill_color_decl
    ;

// Properties

antialiasing_decl
    : 'antialiasing' COLON v=('full' | 'text' | 'none') SEMICOLON
    ;

fill_opacity_decl
    : 'fill-opacity' COLON v=float_ SEMICOLON
    ;

fill_color_decl
    : 'fill-color' COLON color SEMICOLON
    ;

color
    : HEX
    | rgb_color
    | rgba_color
    ;

rgb_color
    : 'rgb(' r=uint_ ',' g=uint_ ',' b=uint_ ')'
    ;

rgba_color
    : 'rgba(' r=uint_ ',' g=uint_ ',' b=uint_ ',' a=float_ ')'
    ;
