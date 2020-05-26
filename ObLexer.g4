/**
 * Ob lexical grammar in ANTLR4 notation
 */

lexer grammar ObLexer;

import UnicodeClasses;

//===============================================================//
// Section: hidden tokens

DelimitedComment
    : '/*' ( DelimitedComment | . )*? '*/'
      -> channel(HIDDEN)
    ;

LineComment
    : '//' ~[\r\n]*
      -> channel(HIDDEN)
    ;

WS
    : [\u0020\u0009\u000C]
      -> channel(HIDDEN)
    ;

NL: '\n' | '\r' '\n'?;


//===============================================================//
// Section: keywords

PACKAGE : 'package';


//===============================================================//
// Section: punctuations

DOT: '.' ;
COMMA: ',' ;
COLON: ':' ;
SEMICOLON: ';' ;


//===============================================================//
// Section: parentheses

LPAREN: '(';
RPAREN: ')' ;
LSQUARE: '[';
RSQUARE: ']' ;
LCURL: '{' ;
RCURL: '}' ;


//===============================================================//
// Section: strings

StringLiteral
    : '"' (~["\\] | EscapedValue)*  '"'
    ;

fragment LETTER
    : UNICODE_LETTER
    | '_'
    ;

fragment EscapedValue
    : '\\' ('u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
           | 'U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
           | [abfnrtv\\'"]
           | OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT
           | 'x' HEX_DIGIT HEX_DIGIT)
    ;


//===============================================================//
// Section: numbers

fragment DECIMALS
    : [0-9]+
    ;

fragment OCTAL_DIGIT
    : [0-7]
    ;

fragment HEX_DIGIT
    : [0-9a-fA-F]
    ;

fragment EXPONENT
    : [eE] [+-]? DECIMALS
    ;


//===============================================================//
// other

Identifier
    : LETTER (LETTER | UNICODE_DIGIT)*
    ;
