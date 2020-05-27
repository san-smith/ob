
parser grammar ObParser;

options {
    tokenVocab=ObLexer;
}

sourceFile
    : packageClause delimiter
        (importDecl delimiter)* 
        ((varDecl) delimiter)*
    ;

packageClause
    : PACKAGE Identifier
    ;

importDecl
    : IMPORT StringLiteral (AS Identifier)? 
    ;

delimiter
    : SEMICOLON
    | NL
    | EOF
    ;

varDecl
    : VAR Identifier (
        ( COMMA Identifier)* COLON Identifier
        | ASSIGNMENT expression
        )
    ;

expression
    : primaryExpr
    | unaryExpr
    ;

primaryExpr
    : operand
    ;

unaryExpr
    : primaryExpr
    | (ADD | SUB | NOT ) expression
    ;

operand
    : literal
    | LPAREN expression RPAREN
    ;

literal
    : basicLit
    ;

basicLit
    : integer
    | FLOAT_LIT
    | IMAGINARY_LIT
    | RUNE_LIT
    ;

integer
    : DECIMAL_LIT
    | OCTAL_LIT
    | HEX_LIT
    | IMAGINARY_LIT
    | RUNE_LIT
    ;