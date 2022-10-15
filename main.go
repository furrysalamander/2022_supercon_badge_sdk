package main

import (
	"assembler/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream("test.asm")
	lexer := parser.NewAssemblyLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAssemblyParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	var listener PassTwoListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())
	fmt.Println(listener.progHex.String())
}
