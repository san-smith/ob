
parser grammar ObParser;

options {
    tokenVocab=ObLexer;
}

sourceFile
    : packageClause delimiter importDecl*
    ;

packageClause
    : PACKAGE Identifier
    ;

importDecl
    : 
        ( IMPORT StringLiteral 
        | IMPORT StringLiteral AS Identifier
    ) delimiter
    ;

delimiter
    : SEMICOLON
    | NL
    | EOF
    ;