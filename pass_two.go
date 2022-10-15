package main

import (
	"fmt"
	"strconv"
	"strings"

	"assembler/parser"
)

type PassTwoListener struct {
	*parser.BaseAssemblyListener

	// progHex is the hex string representing the program.
	progHex strings.Builder
}

func (s *PassTwoListener) WriteRegisterCode(register parser.IRegisterSymbolContext) {
	registerMap := map[string]string{
		"R0":  "0",
		"R1":  "1",
		"R2":  "2",
		"R3":  "3",
		"R4":  "4",
		"R5":  "5",
		"R6":  "6",
		"R7":  "7",
		"R8":  "8",
		"R9":  "9",
		"OUT": "A",
		"IN":  "B",
		"JSR": "C",
		"PCL": "D",
		"PCM": "E",
		"PCH": "F",
	}
	s.WriteNibbles(registerMap[register.GetText()])
}

// TODO Support for negatives and characters.
func parseHexDataFromString(data string, maxLength int) string {
	parsedInt, err := strconv.ParseInt(data, 0, 8)
	if err != nil {
		panic("") // TODO Error Message
	}
	hexString := strconv.FormatInt(parsedInt, 16)
	if len(hexString) > maxLength {
		panic("") // TODO Error Message
	}
	return hexString
}

func (s *PassTwoListener) WriteNibbles(hex string) {
	s.progHex.WriteString(hex)
}

func (s *PassTwoListener) WriteImmediateNibble(data string) {
	s.WriteNibbles(parseHexDataFromString(data, 1))
}

func (s *PassTwoListener) WriteImmediateByte(data string) {
	s.WriteNibbles(parseHexDataFromString(data, 2))
}

func (s *PassTwoListener) WriteDoubleRegInstruction(ctx *parser.InstructionContext, opcode string) {
	s.WriteNibbles(opcode)
	s.WriteRegisterCode(ctx.RegisterSymbol(0))
	s.WriteRegisterCode(ctx.RegisterSymbol(1))
}

func (s *PassTwoListener) WriteRegZeroNibbleInstruction(ctx *parser.InstructionContext, opcode string) {
	s.WriteNibbles(opcode)
	s.WriteImmediateNibble(ctx.Nibble().GetText())
}

func (s *PassTwoListener) WriteAluInstruction(ctx *parser.InstructionContext, baseOpcode string) {
	if ctx.RegisterSymbol(1) != nil {
		s.WriteDoubleRegInstruction(ctx, baseOpcode)
	} else {
		s.WriteRegZeroNibbleInstruction(ctx, fmt.Sprintf("0%s", baseOpcode))
	}
}

func (s *PassTwoListener) WriteDataByteInstruction(ctx *parser.InstructionContext, opcode string) {
	s.WriteNibbles(opcode)
	s.WriteImmediateByte(ctx.DataByte().GetText())
}

func (s *PassTwoListener) WriteSingleRegInstruction(ctx *parser.InstructionContext, opcode string) {
	s.WriteNibbles(opcode)
	s.WriteRegisterCode(ctx.RegisterSymbol(0))
}

func (s *PassTwoListener) EnterInstruction(ctx *parser.InstructionContext) {
	instructionMnemonic := ctx.GetChild(0)

	switch instructionMnemonic {
	case ctx.ADD():
		fmt.Println("ADD")
		s.WriteAluInstruction(ctx, "1")
	case ctx.ADC():
		fmt.Println("ADC")
		s.WriteAluInstruction(ctx, "2")
	case ctx.SUB():
		fmt.Println("SUB")
		s.WriteAluInstruction(ctx, "3")
	case ctx.SBB():
		fmt.Println("SBB")
		s.WriteAluInstruction(ctx, "4")
	case ctx.OR():
		fmt.Println("OR")
		s.WriteAluInstruction(ctx, "5")
	case ctx.AND():
		fmt.Println("AND")
		s.WriteAluInstruction(ctx, "6")
	case ctx.XOR():
		fmt.Println("XOR")
		s.WriteAluInstruction(ctx, "7")
	case ctx.MOV():
		fmt.Println("MOV")
		s.WriteDoubleRegInstruction(ctx, "8")
		// s.WriteNibbles("9") // TODO
		// s.WriteNibbles("A") // TODO
		// s.WriteNibbles("B") // TODO
		// s.WriteNibbles("C") // TODO
		// s.WriteNibbles("D") // TODO
		// s.WriteNibbles("E") // TODO
	case ctx.LPC():
		fmt.Println("LPC")
		s.WriteDataByteInstruction(ctx, "E")
	case ctx.JR():
		fmt.Println("JR")
		s.WriteDataByteInstruction(ctx, "F")
	case ctx.CP():
		fmt.Println("CP")
		s.WriteRegZeroNibbleInstruction(ctx, "00")
	case ctx.INC():
		fmt.Println("INC")
		s.WriteSingleRegInstruction(ctx, "02")
	case ctx.DEC():
		fmt.Println("DEC")
		s.WriteSingleRegInstruction(ctx, "03")
	case ctx.DSZ():
		fmt.Println("DSZ")
		s.WriteSingleRegInstruction(ctx, "04")
	case ctx.EXR():
		fmt.Println("EXR")
		s.WriteRegZeroNibbleInstruction(ctx, "08")
	case ctx.BIT():
		fmt.Println("BIT")
		s.WriteNibbles("09") // TODO
	case ctx.BSET():
		fmt.Println("BSET")
		s.WriteNibbles("0A") // TODO
	case ctx.BCLR():
		fmt.Println("BCLR")
		s.WriteNibbles("0B") // TODO
	case ctx.BTG():
		fmt.Println("BTG")
		s.WriteNibbles("0C") // TODO
	case ctx.RRC():
		fmt.Println("RRC")
		s.WriteSingleRegInstruction(ctx, "0D")
	case ctx.RET():
		fmt.Println("RET")
		s.WriteRegZeroNibbleInstruction(ctx, "0E")
	case ctx.SKIPI():
		fmt.Println("SKIP")
		s.WriteNibbles("0F") // TODO
	default:
		fmt.Println("Instruction Not Yet Implemented")
	}
}

// TODO
func (s *PassTwoListener) EnterDirective(ctx *parser.DirectiveContext) {

}

// TODO
func (s *PassTwoListener) EnterSyntheticInstruction(ctx *parser.SyntheticInstructionContext) {
	instructionMnemonic := ctx.GetChild(0)
	switch instructionMnemonic {
	case ctx.CPL():
		fmt.Println("CPL")
	}
}