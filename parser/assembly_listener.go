// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Assembly
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// AssemblyListener is a complete listener for a parse tree produced by AssemblyParser.
type AssemblyListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterSyntheticInstruction is called when entering the syntheticInstruction production.
	EnterSyntheticInstruction(c *SyntheticInstructionContext)

	// EnterMacro is called when entering the macro production.
	EnterMacro(c *MacroContext)

	// EnterRegisterSymbol is called when entering the registerSymbol production.
	EnterRegisterSymbol(c *RegisterSymbolContext)

	// EnterRg is called when entering the rg production.
	EnterRg(c *RgContext)

	// EnterFlag is called when entering the flag production.
	EnterFlag(c *FlagContext)

	// EnterDataByte is called when entering the dataByte production.
	EnterDataByte(c *DataByteContext)

	// EnterNibble is called when entering the nibble production.
	EnterNibble(c *NibbleContext)

	// EnterQuarter is called when entering the quarter production.
	EnterQuarter(c *QuarterContext)

	// EnterRegisterCombo is called when entering the registerCombo production.
	EnterRegisterCombo(c *RegisterComboContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitSyntheticInstruction is called when exiting the syntheticInstruction production.
	ExitSyntheticInstruction(c *SyntheticInstructionContext)

	// ExitMacro is called when exiting the macro production.
	ExitMacro(c *MacroContext)

	// ExitRegisterSymbol is called when exiting the registerSymbol production.
	ExitRegisterSymbol(c *RegisterSymbolContext)

	// ExitRg is called when exiting the rg production.
	ExitRg(c *RgContext)

	// ExitFlag is called when exiting the flag production.
	ExitFlag(c *FlagContext)

	// ExitDataByte is called when exiting the dataByte production.
	ExitDataByte(c *DataByteContext)

	// ExitNibble is called when exiting the nibble production.
	ExitNibble(c *NibbleContext)

	// ExitQuarter is called when exiting the quarter production.
	ExitQuarter(c *QuarterContext)

	// ExitRegisterCombo is called when exiting the registerCombo production.
	ExitRegisterCombo(c *RegisterComboContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)
}
