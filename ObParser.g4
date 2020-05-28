
parser grammar ObParser;

options {
    tokenVocab=ObLexer;
}

sourceFile
    : packageClause delimiter
        (importDecl delimiter)* 
        ((varDecl | constDecl) delimiter)*
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
    : VAR varSpec
    ;

constDecl
    : CONST varSpec
    ;

varSpec
    : Identifier (
        ( COMMA Identifier)* COLON type
        | (COLON type)? ASSIGNMENT expression
        )
    ;

type
    : typeName
    ;

typeName
    : Identifier
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
    | StringLiteral
    | NULL
    ;

integer
    : DECIMAL_LIT
    | OCTAL_LIT
    | HEX_LIT
    | IMAGINARY_LIT
    | RUNE_LIT
    ;