
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
    : IMPORT StringLiteral (AS Identifier)? 
    ;

delimiter
    : SEMICOLON
    | NL
    | EOF
    ;