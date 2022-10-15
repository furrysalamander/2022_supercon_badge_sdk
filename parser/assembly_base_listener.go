// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Assembly
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

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

// EnterSyntheticInstruction is called when production syntheticInstruction is entered.
func (s *BaseAssemblyListener) EnterSyntheticInstruction(ctx *SyntheticInstructionContext) {}

// ExitSyntheticInstruction is called when production syntheticInstruction is exited.
func (s *BaseAssemblyListener) ExitSyntheticInstruction(ctx *SyntheticInstructionContext) {}

// EnterMacro is called when production macro is entered.
func (s *BaseAssemblyListener) EnterMacro(ctx *MacroContext) {}

// ExitMacro is called when production macro is exited.
func (s *BaseAssemblyListener) ExitMacro(ctx *MacroContext) {}

// EnterRegisterSymbol is called when production registerSymbol is entered.
func (s *BaseAssemblyListener) EnterRegisterSymbol(ctx *RegisterSymbolContext) {}

// ExitRegisterSymbol is called when production registerSymbol is exited.
func (s *BaseAssemblyListener) ExitRegisterSymbol(ctx *RegisterSymbolContext) {}

// EnterRg is called when production rg is entered.
func (s *BaseAssemblyListener) EnterRg(ctx *RgContext) {}

// ExitRg is called when production rg is exited.
func (s *BaseAssemblyListener) ExitRg(ctx *RgContext) {}

// EnterFlag is called when production flag is entered.
func (s *BaseAssemblyListener) EnterFlag(ctx *FlagContext) {}

// ExitFlag is called when production flag is exited.
func (s *BaseAssemblyListener) ExitFlag(ctx *FlagContext) {}

// EnterDataByte is called when production dataByte is entered.
func (s *BaseAssemblyListener) EnterDataByte(ctx *DataByteContext) {}

// ExitDataByte is called when production dataByte is exited.
func (s *BaseAssemblyListener) ExitDataByte(ctx *DataByteContext) {}

// EnterNibble is called when production nibble is entered.
func (s *BaseAssemblyListener) EnterNibble(ctx *NibbleContext) {}

// ExitNibble is called when production nibble is exited.
func (s *BaseAssemblyListener) ExitNibble(ctx *NibbleContext) {}

// EnterQuarter is called when production quarter is entered.
func (s *BaseAssemblyListener) EnterQuarter(ctx *QuarterContext) {}

// ExitQuarter is called when production quarter is exited.
func (s *BaseAssemblyListener) ExitQuarter(ctx *QuarterContext) {}

// EnterRegisterCombo is called when production registerCombo is entered.
func (s *BaseAssemblyListener) EnterRegisterCombo(ctx *RegisterComboContext) {}

// ExitRegisterCombo is called when production registerCombo is exited.
func (s *BaseAssemblyListener) ExitRegisterCombo(ctx *RegisterComboContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseAssemblyListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseAssemblyListener) ExitLiteral(ctx *LiteralContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseAssemblyListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseAssemblyListener) ExitDirective(ctx *DirectiveContext) {}
