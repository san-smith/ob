package main

import (
	"fmt"
	"os"

	"./parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseObParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewObLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewObParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.SourceFile()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
