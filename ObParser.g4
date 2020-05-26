
parser grammar ObParser;

options {
    tokenVocab=ObLexer;
}

sourceFile
    : packageClause delimiter
    ;

packageClause
    : PACKAGE Identifier
    ;

delimiter
    : SEMICOLON
    | NL
    | EOF
    ;