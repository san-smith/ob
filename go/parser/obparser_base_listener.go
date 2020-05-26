// Code generated from ObParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ObParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseObParserListener is a complete listener for a parse tree produced by ObParser.
type BaseObParserListener struct{}

var _ ObParserListener = &BaseObParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseObParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseObParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseObParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseObParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceFile is called when production sourceFile is entered.
func (s *BaseObParserListener) EnterSourceFile(ctx *SourceFileContext) {}

// ExitSourceFile is called when production sourceFile is exited.
func (s *BaseObParserListener) ExitSourceFile(ctx *SourceFileContext) {}

// EnterPackageClause is called when production packageClause is entered.
func (s *BaseObParserListener) EnterPackageClause(ctx *PackageClauseContext) {}

// ExitPackageClause is called when production packageClause is exited.
func (s *BaseObParserListener) ExitPackageClause(ctx *PackageClauseContext) {}

// EnterDelimiter is called when production delimiter is entered.
func (s *BaseObParserListener) EnterDelimiter(ctx *DelimiterContext) {}

// ExitDelimiter is called when production delimiter is exited.
func (s *BaseObParserListener) ExitDelimiter(ctx *DelimiterContext) {}
