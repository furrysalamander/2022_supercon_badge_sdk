package main

import (
	"assembler/parser"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream("test.asm")
	lexer := parser.NewAssemblyLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAssemblyParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	var passOne PassOneListener
	var passTwo PassTwoListener

	parseTree := p.Start()

	// Build the symbol table
	antlr.ParseTreeWalkerDefault.Walk(&passOne, parseTree)
	fmt.Println(passOne.symbolTable)
	// Assemble the program
	antlr.ParseTreeWalkerDefault.Walk(&passTwo, parseTree)

	hexString := passTwo.progHex.String()

	if len(hexString)%2 != 0 {
		fmt.Printf("Resulting program length (%d) is odd, padding by one nibble.\n", len(hexString))
		hexString += "0"
	}

	fmt.Println(hexString)

	programBytes, err := hex.DecodeString(hexString)

	if err != nil {
		fmt.Println(err)
		return
	}

	os.WriteFile("out.bin", programBytes, 0644)
}
