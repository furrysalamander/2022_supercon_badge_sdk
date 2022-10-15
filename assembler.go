package main

import (
	"os"

	"assembler/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewAssemblyLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAssemblyParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	p.BuildParseTrees = true
	_ = p.Assembly()

}
