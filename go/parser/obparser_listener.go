// Code generated from ObParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ObParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ObParserListener is a complete listener for a parse tree produced by ObParser.
type ObParserListener interface {
	antlr.ParseTreeListener

	// EnterSourceFile is called when entering the sourceFile production.
	EnterSourceFile(c *SourceFileContext)

	// EnterPackageClause is called when entering the packageClause production.
	EnterPackageClause(c *PackageClauseContext)

	// EnterDelimiter is called when entering the delimiter production.
	EnterDelimiter(c *DelimiterContext)

	// ExitSourceFile is called when exiting the sourceFile production.
	ExitSourceFile(c *SourceFileContext)

	// ExitPackageClause is called when exiting the packageClause production.
	ExitPackageClause(c *PackageClauseContext)

	// ExitDelimiter is called when exiting the delimiter production.
	ExitDelimiter(c *DelimiterContext)
}
