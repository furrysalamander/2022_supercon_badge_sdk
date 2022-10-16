package main

import (
	"assembler/parser"
)

type PassOneListener struct {
	*parser.BaseAssemblyListener

	programCounter int
	symbolTable    map[string]int
}

func (s *PassOneListener) EnterStart(ctx *parser.StartContext) {
	s.symbolTable = make(map[string]int)
}

func (s *PassOneListener) EnterInstruction(ctx *parser.InstructionContext) {
	s.programCounter++
}

func (s *PassOneListener) EnterSyntheticInstruction(ctx *parser.SyntheticInstructionContext) {
	switch ctx.GetChild(0) {
	case ctx.RLC():
		fallthrough
	case ctx.SL():
		fallthrough
	case ctx.LSR():
		fallthrough
	case ctx.NEG():
		s.programCounter += 2

	case ctx.CPL():
		if ctx.RegisterSymbol(1) == nil {
			s.programCounter++
		} else {
			s.programCounter += 2
		}

	case ctx.NOP():
		s.programCounter++
	default:
		panic("Synthetic instruction not implemented!")
	}
}

func (s *PassOneListener) EnterDirective(ctx *parser.DirectiveContext) {
	if label := ctx.LABEL(); label != nil {
		s.symbolTable[label.GetText()] = s.programCounter
	}
}
