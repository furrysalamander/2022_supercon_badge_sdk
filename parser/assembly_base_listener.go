// Code generated from Assembly.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Assembly

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAssemblyListener is a complete listener for a parse tree produced by AssemblyParser.
type BaseAssemblyListener struct{}

var _ AssemblyListener = &BaseAssemblyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAssemblyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAssemblyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAssemblyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAssemblyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseAssemblyListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseAssemblyListener) ExitStart(ctx *StartContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseAssemblyListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseAssemblyListener) ExitInstruction(ctx *InstructionContext) {}

// EnterSynthetic_instruction is called when production synthetic_instruction is entered.
func (s *BaseAssemblyListener) EnterSynthetic_instruction(ctx *Synthetic_instructionContext) {}

// ExitSynthetic_instruction is called when production synthetic_instruction is exited.
func (s *BaseAssemblyListener) ExitSynthetic_instruction(ctx *Synthetic_instructionContext) {}

// EnterMacro is called when production macro is entered.
func (s *BaseAssemblyListener) EnterMacro(ctx *MacroContext) {}

// ExitMacro is called when production macro is exited.
func (s *BaseAssemblyListener) ExitMacro(ctx *MacroContext) {}

// EnterRegister_symbol is called when production register_symbol is entered.
func (s *BaseAssemblyListener) EnterRegister_symbol(ctx *Register_symbolContext) {}

// ExitRegister_symbol is called when production register_symbol is exited.
func (s *BaseAssemblyListener) ExitRegister_symbol(ctx *Register_symbolContext) {}

// EnterRg is called when production rg is entered.
func (s *BaseAssemblyListener) EnterRg(ctx *RgContext) {}

// ExitRg is called when production rg is exited.
func (s *BaseAssemblyListener) ExitRg(ctx *RgContext) {}

// EnterFlag is called when production flag is entered.
func (s *BaseAssemblyListener) EnterFlag(ctx *FlagContext) {}

// ExitFlag is called when production flag is exited.
func (s *BaseAssemblyListener) ExitFlag(ctx *FlagContext) {}

// EnterData_byte is called when production data_byte is entered.
func (s *BaseAssemblyListener) EnterData_byte(ctx *Data_byteContext) {}

// ExitData_byte is called when production data_byte is exited.
func (s *BaseAssemblyListener) ExitData_byte(ctx *Data_byteContext) {}

// EnterNibble is called when production nibble is entered.
func (s *BaseAssemblyListener) EnterNibble(ctx *NibbleContext) {}

// ExitNibble is called when production nibble is exited.
func (s *BaseAssemblyListener) ExitNibble(ctx *NibbleContext) {}

// EnterQuarter is called when production quarter is entered.
func (s *BaseAssemblyListener) EnterQuarter(ctx *QuarterContext) {}

// ExitQuarter is called when production quarter is exited.
func (s *BaseAssemblyListener) ExitQuarter(ctx *QuarterContext) {}

// EnterRegister_combo is called when production register_combo is entered.
func (s *BaseAssemblyListener) EnterRegister_combo(ctx *Register_comboContext) {}

// ExitRegister_combo is called when production register_combo is exited.
func (s *BaseAssemblyListener) ExitRegister_combo(ctx *Register_comboContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseAssemblyListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseAssemblyListener) ExitLiteral(ctx *LiteralContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseAssemblyListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseAssemblyListener) ExitDirective(ctx *DirectiveContext) {}

// EnterMn_SKIP is called when production mn_SKIP is entered.
func (s *BaseAssemblyListener) EnterMn_SKIP(ctx *Mn_SKIPContext) {}

// ExitMn_SKIP is called when production mn_SKIP is exited.
func (s *BaseAssemblyListener) ExitMn_SKIP(ctx *Mn_SKIPContext) {}
