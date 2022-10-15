// Code generated from Assembly.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Assembly

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AssemblyListener is a complete listener for a parse tree produced by AssemblyParser.
type AssemblyListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterSynthetic_instruction is called when entering the synthetic_instruction production.
	EnterSynthetic_instruction(c *Synthetic_instructionContext)

	// EnterMacro is called when entering the macro production.
	EnterMacro(c *MacroContext)

	// EnterRegister_symbol is called when entering the register_symbol production.
	EnterRegister_symbol(c *Register_symbolContext)

	// EnterRg is called when entering the rg production.
	EnterRg(c *RgContext)

	// EnterFlag is called when entering the flag production.
	EnterFlag(c *FlagContext)

	// EnterData_byte is called when entering the data_byte production.
	EnterData_byte(c *Data_byteContext)

	// EnterNibble is called when entering the nibble production.
	EnterNibble(c *NibbleContext)

	// EnterQuarter is called when entering the quarter production.
	EnterQuarter(c *QuarterContext)

	// EnterRegister_combo is called when entering the register_combo production.
	EnterRegister_combo(c *Register_comboContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// EnterMn_SKIP is called when entering the mn_SKIP production.
	EnterMn_SKIP(c *Mn_SKIPContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitSynthetic_instruction is called when exiting the synthetic_instruction production.
	ExitSynthetic_instruction(c *Synthetic_instructionContext)

	// ExitMacro is called when exiting the macro production.
	ExitMacro(c *MacroContext)

	// ExitRegister_symbol is called when exiting the register_symbol production.
	ExitRegister_symbol(c *Register_symbolContext)

	// ExitRg is called when exiting the rg production.
	ExitRg(c *RgContext)

	// ExitFlag is called when exiting the flag production.
	ExitFlag(c *FlagContext)

	// ExitData_byte is called when exiting the data_byte production.
	ExitData_byte(c *Data_byteContext)

	// ExitNibble is called when exiting the nibble production.
	ExitNibble(c *NibbleContext)

	// ExitQuarter is called when exiting the quarter production.
	ExitQuarter(c *QuarterContext)

	// ExitRegister_combo is called when exiting the register_combo production.
	ExitRegister_combo(c *Register_comboContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)

	// ExitMn_SKIP is called when exiting the mn_SKIP production.
	ExitMn_SKIP(c *Mn_SKIPContext)
}
