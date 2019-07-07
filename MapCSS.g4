grammar MapCSS;

fragment COMMENT: ('/*' .*? '*/') | ('//' .*? '\n');
WS: (' ' | '\t' | '\n' | '\r' | '\f' | COMMENT) -> skip;

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
POSITIVE_FLOAT: [0-9]+ | [0-9]* '.' [0-9]+;

// Colors

fragment HEX_CHAR:     [0-9a-fA-F];
fragment HEX_3_DIGITS: '#' HEX_CHAR HEX_CHAR HEX_CHAR;
fragment HEX_4_DIGITS: '#' HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR;
fragment HEX_6_DIGITS: '#' HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR;
fragment HEX_8_DIGITS: '#' HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR;
HEX:                   (HEX_3_DIGITS | HEX_4_DIGITS | HEX_6_DIGITS | HEX_8_DIGITS);

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
    : '|z' min=POSITIVE_INT '-' max=POSITIVE_INT
    ;

min_zoom
    : '|z' min=POSITIVE_INT '-'
    ;

exact_zoom
    : '|z' min=POSITIVE_INT
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
    : 'fill-opacity' COLON v=(POSITIVE_INT | POSITIVE_FLOAT) SEMICOLON
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
    : 'rgb(' r=POSITIVE_INT ',' g=POSITIVE_INT ',' b=POSITIVE_INT ')'
    ;

rgba_color
    : 'rgba(' r=POSITIVE_INT ',' g=POSITIVE_INT ',' b=POSITIVE_INT ',' a=(POSITIVE_INT | POSITIVE_FLOAT) ')'
    ;
