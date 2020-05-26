// Code generated from ObParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ObParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 17, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 3, 4, 3, 6, 6, 11, 11, 2, 13, 2, 8,
	3, 2, 2, 2, 4, 11, 3, 2, 2, 2, 6, 14, 3, 2, 2, 2, 8, 9, 5, 4, 3, 2, 9,
	10, 5, 6, 4, 2, 10, 3, 3, 2, 2, 2, 11, 12, 7, 7, 2, 2, 12, 13, 7, 19, 2,
	2, 13, 5, 3, 2, 2, 2, 14, 15, 9, 2, 2, 2, 15, 7, 3, 2, 2, 2, 2,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "'package'", "'.'", "','", "':'", "';'", "'('", "')'",
	"'['", "']'", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "DelimitedComment", "LineComment", "WS", "NL", "PACKAGE", "DOT", "COMMA",
	"COLON", "SEMICOLON", "LPAREN", "RPAREN", "LSQUARE", "RSQUARE", "LCURL",
	"RCURL", "StringLiteral", "Identifier",
}

var ruleNames = []string{
	"sourceFile", "packageClause", "delimiter",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ObParser struct {
	*antlr.BaseParser
}

func NewObParser(input antlr.TokenStream) *ObParser {
	this := new(ObParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ObParser.g4"

	return this
}

// ObParser tokens.
const (
	ObParserEOF              = antlr.TokenEOF
	ObParserDelimitedComment = 1
	ObParserLineComment      = 2
	ObParserWS               = 3
	ObParserNL               = 4
	ObParserPACKAGE          = 5
	ObParserDOT              = 6
	ObParserCOMMA            = 7
	ObParserCOLON            = 8
	ObParserSEMICOLON        = 9
	ObParserLPAREN           = 10
	ObParserRPAREN           = 11
	ObParserLSQUARE          = 12
	ObParserRSQUARE          = 13
	ObParserLCURL            = 14
	ObParserRCURL            = 15
	ObParserStringLiteral    = 16
	ObParserIdentifier       = 17
)

// ObParser rules.
const (
	ObParserRULE_sourceFile    = 0
	ObParserRULE_packageClause = 1
	ObParserRULE_delimiter     = 2
)

// ISourceFileContext is an interface to support dynamic dispatch.
type ISourceFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceFileContext differentiates from other interfaces.
	IsSourceFileContext()
}

type SourceFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceFileContext() *SourceFileContext {
	var p = new(SourceFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObParserRULE_sourceFile
	return p
}

func (*SourceFileContext) IsSourceFileContext() {}

func NewSourceFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceFileContext {
	var p = new(SourceFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObParserRULE_sourceFile

	return p
}

func (s *SourceFileContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceFileContext) PackageClause() IPackageClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackageClauseContext)
}

func (s *SourceFileContext) Delimiter() IDelimiterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelimiterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelimiterContext)
}

func (s *SourceFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObParserListener); ok {
		listenerT.EnterSourceFile(s)
	}
}

func (s *SourceFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObParserListener); ok {
		listenerT.ExitSourceFile(s)
	}
}

func (p *ObParser) SourceFile() (localctx ISourceFileContext) {
	localctx = NewSourceFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ObParserRULE_sourceFile)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		p.PackageClause()
	}
	{
		p.SetState(7)
		p.Delimiter()
	}

	return localctx
}

// IPackageClauseContext is an interface to support dynamic dispatch.
type IPackageClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageClauseContext differentiates from other interfaces.
	IsPackageClauseContext()
}

type PackageClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageClauseContext() *PackageClauseContext {
	var p = new(PackageClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObParserRULE_packageClause
	return p
}

func (*PackageClauseContext) IsPackageClauseContext() {}

func NewPackageClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageClauseContext {
	var p = new(PackageClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObParserRULE_packageClause

	return p
}

func (s *PackageClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageClauseContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(ObParserPACKAGE, 0)
}

func (s *PackageClauseContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ObParserIdentifier, 0)
}

func (s *PackageClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObParserListener); ok {
		listenerT.EnterPackageClause(s)
	}
}

func (s *PackageClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObParserListener); ok {
		listenerT.ExitPackageClause(s)
	}
}

func (p *ObParser) PackageClause() (localctx IPackageClauseContext) {
	localctx = NewPackageClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ObParserRULE_packageClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(9)
		p.Match(ObParserPACKAGE)
	}
	{
		p.SetState(10)
		p.Match(ObParserIdentifier)
	}

	return localctx
}

// IDelimiterContext is an interface to support dynamic dispatch.
type IDelimiterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelimiterContext differentiates from other interfaces.
	IsDelimiterContext()
}

type DelimiterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelimiterContext() *DelimiterContext {
	var p = new(DelimiterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObParserRULE_delimiter
	return p
}

func (*DelimiterContext) IsDelimiterContext() {}

func NewDelimiterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelimiterContext {
	var p = new(DelimiterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObParserRULE_delimiter

	return p
}

func (s *DelimiterContext) GetParser() antlr.Parser { return s.parser }

func (s *DelimiterContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ObParserSEMICOLON, 0)
}

func (s *DelimiterContext) NL() antlr.TerminalNode {
	return s.GetToken(ObParserNL, 0)
}

func (s *DelimiterContext) EOF() antlr.TerminalNode {
	return s.GetToken(ObParserEOF, 0)
}

func (s *DelimiterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelimiterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelimiterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObParserListener); ok {
		listenerT.EnterDelimiter(s)
	}
}

func (s *DelimiterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObParserListener); ok {
		listenerT.ExitDelimiter(s)
	}
}

func (p *ObParser) Delimiter() (localctx IDelimiterContext) {
	localctx = NewDelimiterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ObParserRULE_delimiter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		_la = p.GetTokenStream().LA(1)

		if !(((_la - -1)&-(0x1f+1)) == 0 && ((1<<uint((_la - -1)))&((1<<(ObParserEOF - -1))|(1<<(ObParserNL - -1))|(1<<(ObParserSEMICOLON - -1)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
