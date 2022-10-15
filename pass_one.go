package main

import (
	"assembler/parser"
)

type PassOneListener struct {
	*parser.BaseAssemblyListener
}

func validateDataSize(s string, bits int) bool {
	return true
}

func (s *PassOneListener) EnterDataByte(ctx *parser.DataByteContext) {
	if !validateDataSize(ctx.GetText(), 8) {
		panic("") // TODO: Add better erorr messages.
	}
}

func (s *PassOneListener) EnterNibble(ctx *parser.NibbleContext) {
	if !validateDataSize(ctx.GetText(), 4) {
		panic("") // TODO: Add better erorr messages.
	}
}

func (s *PassOneListener) EnterQuarter(ctx *parser.QuarterContext) {
	if !validateDataSize(ctx.GetText(), 2) {
		panic("") // TODO: Add better erorr messages.
	}
}

func (s *PassOneListener) EnterInstruction(ctx *parser.InstructionContext) {

}

func (s *PassOneListener) EnterDirective(ctx *parser.DirectiveContext) {

}
