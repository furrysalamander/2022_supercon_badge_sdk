// Code generated from Assembly.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Assembly

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 66, 221,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 5, 2, 34, 10, 2, 7,
	2, 36, 10, 2, 12, 2, 14, 2, 39, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 162, 10, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 189, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 194, 10, 5, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 2, 2, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 2,
	7, 3, 2, 8, 23, 3, 2, 8, 10, 3, 2, 3, 6, 4, 2, 57, 57, 61, 61, 4, 2, 53,
	53, 60, 60, 2, 243, 2, 37, 3, 2, 2, 2, 4, 161, 3, 2, 2, 2, 6, 188, 3, 2,
	2, 2, 8, 190, 3, 2, 2, 2, 10, 195, 3, 2, 2, 2, 12, 197, 3, 2, 2, 2, 14,
	199, 3, 2, 2, 2, 16, 201, 3, 2, 2, 2, 18, 203, 3, 2, 2, 2, 20, 205, 3,
	2, 2, 2, 22, 207, 3, 2, 2, 2, 24, 213, 3, 2, 2, 2, 26, 215, 3, 2, 2, 2,
	28, 218, 3, 2, 2, 2, 30, 33, 5, 26, 14, 2, 31, 34, 5, 4, 3, 2, 32, 34,
	5, 6, 4, 2, 33, 31, 3, 2, 2, 2, 33, 32, 3, 2, 2, 2, 34, 36, 3, 2, 2, 2,
	35, 30, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3,
	2, 2, 2, 38, 40, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 41, 7, 2, 2, 3, 41,
	3, 3, 2, 2, 2, 42, 43, 7, 24, 2, 2, 43, 44, 5, 10, 6, 2, 44, 45, 7, 55,
	2, 2, 45, 46, 5, 10, 6, 2, 46, 162, 3, 2, 2, 2, 47, 48, 7, 25, 2, 2, 48,
	49, 5, 10, 6, 2, 49, 50, 7, 55, 2, 2, 50, 51, 5, 10, 6, 2, 51, 162, 3,
	2, 2, 2, 52, 53, 7, 27, 2, 2, 53, 54, 5, 10, 6, 2, 54, 55, 7, 55, 2, 2,
	55, 56, 5, 10, 6, 2, 56, 162, 3, 2, 2, 2, 57, 58, 7, 28, 2, 2, 58, 59,
	5, 10, 6, 2, 59, 60, 7, 55, 2, 2, 60, 61, 5, 10, 6, 2, 61, 162, 3, 2, 2,
	2, 62, 63, 7, 29, 2, 2, 63, 64, 5, 10, 6, 2, 64, 65, 7, 55, 2, 2, 65, 66,
	5, 10, 6, 2, 66, 162, 3, 2, 2, 2, 67, 68, 7, 30, 2, 2, 68, 69, 5, 10, 6,
	2, 69, 70, 7, 55, 2, 2, 70, 71, 5, 10, 6, 2, 71, 162, 3, 2, 2, 2, 72, 73,
	7, 31, 2, 2, 73, 74, 5, 10, 6, 2, 74, 75, 7, 55, 2, 2, 75, 76, 5, 10, 6,
	2, 76, 162, 3, 2, 2, 2, 77, 78, 7, 31, 2, 2, 78, 79, 5, 10, 6, 2, 79, 80,
	7, 55, 2, 2, 80, 81, 5, 18, 10, 2, 81, 162, 3, 2, 2, 2, 82, 83, 7, 31,
	2, 2, 83, 84, 5, 22, 12, 2, 84, 85, 7, 55, 2, 2, 85, 86, 7, 8, 2, 2, 86,
	162, 3, 2, 2, 2, 87, 88, 7, 31, 2, 2, 88, 89, 7, 8, 2, 2, 89, 90, 7, 55,
	2, 2, 90, 162, 5, 22, 12, 2, 91, 92, 7, 31, 2, 2, 92, 93, 7, 8, 2, 2, 93,
	94, 7, 55, 2, 2, 94, 95, 7, 51, 2, 2, 95, 96, 5, 16, 9, 2, 96, 97, 7, 52,
	2, 2, 97, 162, 3, 2, 2, 2, 98, 99, 7, 50, 2, 2, 99, 162, 5, 16, 9, 2, 100,
	101, 7, 32, 2, 2, 101, 162, 5, 16, 9, 2, 102, 103, 7, 33, 2, 2, 103, 104,
	7, 8, 2, 2, 104, 105, 7, 55, 2, 2, 105, 162, 5, 18, 10, 2, 106, 107, 7,
	24, 2, 2, 107, 108, 7, 8, 2, 2, 108, 109, 7, 55, 2, 2, 109, 162, 5, 18,
	10, 2, 110, 111, 7, 34, 2, 2, 111, 162, 5, 10, 6, 2, 112, 113, 7, 35, 2,
	2, 113, 162, 5, 10, 6, 2, 114, 115, 7, 36, 2, 2, 115, 162, 5, 10, 6, 2,
	116, 117, 7, 28, 2, 2, 117, 118, 7, 8, 2, 2, 118, 119, 7, 55, 2, 2, 119,
	162, 5, 18, 10, 2, 120, 121, 7, 29, 2, 2, 121, 122, 7, 8, 2, 2, 122, 123,
	7, 55, 2, 2, 123, 162, 5, 18, 10, 2, 124, 125, 7, 30, 2, 2, 125, 126, 7,
	8, 2, 2, 126, 127, 7, 55, 2, 2, 127, 162, 5, 18, 10, 2, 128, 129, 7, 37,
	2, 2, 129, 162, 5, 18, 10, 2, 130, 131, 7, 38, 2, 2, 131, 132, 5, 12, 7,
	2, 132, 133, 7, 55, 2, 2, 133, 134, 5, 20, 11, 2, 134, 162, 3, 2, 2, 2,
	135, 136, 7, 39, 2, 2, 136, 137, 5, 12, 7, 2, 137, 138, 7, 55, 2, 2, 138,
	139, 5, 20, 11, 2, 139, 162, 3, 2, 2, 2, 140, 141, 7, 40, 2, 2, 141, 142,
	5, 12, 7, 2, 142, 143, 7, 55, 2, 2, 143, 144, 5, 20, 11, 2, 144, 162, 3,
	2, 2, 2, 145, 146, 7, 41, 2, 2, 146, 147, 5, 12, 7, 2, 147, 148, 7, 55,
	2, 2, 148, 149, 5, 20, 11, 2, 149, 162, 3, 2, 2, 2, 150, 151, 7, 42, 2,
	2, 151, 162, 5, 10, 6, 2, 152, 153, 7, 43, 2, 2, 153, 154, 7, 8, 2, 2,
	154, 155, 7, 55, 2, 2, 155, 162, 5, 18, 10, 2, 156, 157, 5, 28, 15, 2,
	157, 158, 5, 14, 8, 2, 158, 159, 7, 55, 2, 2, 159, 160, 5, 20, 11, 2, 160,
	162, 3, 2, 2, 2, 161, 42, 3, 2, 2, 2, 161, 47, 3, 2, 2, 2, 161, 52, 3,
	2, 2, 2, 161, 57, 3, 2, 2, 2, 161, 62, 3, 2, 2, 2, 161, 67, 3, 2, 2, 2,
	161, 72, 3, 2, 2, 2, 161, 77, 3, 2, 2, 2, 161, 82, 3, 2, 2, 2, 161, 87,
	3, 2, 2, 2, 161, 91, 3, 2, 2, 2, 161, 98, 3, 2, 2, 2, 161, 100, 3, 2, 2,
	2, 161, 102, 3, 2, 2, 2, 161, 106, 3, 2, 2, 2, 161, 110, 3, 2, 2, 2, 161,
	112, 3, 2, 2, 2, 161, 114, 3, 2, 2, 2, 161, 116, 3, 2, 2, 2, 161, 120,
	3, 2, 2, 2, 161, 124, 3, 2, 2, 2, 161, 128, 3, 2, 2, 2, 161, 130, 3, 2,
	2, 2, 161, 135, 3, 2, 2, 2, 161, 140, 3, 2, 2, 2, 161, 145, 3, 2, 2, 2,
	161, 150, 3, 2, 2, 2, 161, 152, 3, 2, 2, 2, 161, 156, 3, 2, 2, 2, 162,
	5, 3, 2, 2, 2, 163, 164, 7, 44, 2, 2, 164, 165, 5, 10, 6, 2, 165, 166,
	7, 55, 2, 2, 166, 167, 5, 10, 6, 2, 167, 189, 3, 2, 2, 2, 168, 169, 7,
	45, 2, 2, 169, 170, 5, 10, 6, 2, 170, 171, 7, 55, 2, 2, 171, 172, 5, 10,
	6, 2, 172, 189, 3, 2, 2, 2, 173, 174, 7, 46, 2, 2, 174, 189, 5, 10, 6,
	2, 175, 176, 7, 47, 2, 2, 176, 189, 7, 8, 2, 2, 177, 178, 7, 47, 2, 2,
	178, 179, 5, 10, 6, 2, 179, 180, 7, 55, 2, 2, 180, 181, 5, 10, 6, 2, 181,
	189, 3, 2, 2, 2, 182, 183, 7, 48, 2, 2, 183, 184, 5, 10, 6, 2, 184, 185,
	7, 55, 2, 2, 185, 186, 5, 10, 6, 2, 186, 189, 3, 2, 2, 2, 187, 189, 7,
	49, 2, 2, 188, 163, 3, 2, 2, 2, 188, 168, 3, 2, 2, 2, 188, 173, 3, 2, 2,
	2, 188, 175, 3, 2, 2, 2, 188, 177, 3, 2, 2, 2, 188, 182, 3, 2, 2, 2, 188,
	187, 3, 2, 2, 2, 189, 7, 3, 2, 2, 2, 190, 193, 7, 54, 2, 2, 191, 194, 7,
	53, 2, 2, 192, 194, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 193, 192, 3, 2, 2,
	2, 194, 9, 3, 2, 2, 2, 195, 196, 9, 2, 2, 2, 196, 11, 3, 2, 2, 2, 197,
	198, 9, 3, 2, 2, 198, 13, 3, 2, 2, 2, 199, 200, 9, 4, 2, 2, 200, 15, 3,
	2, 2, 2, 201, 202, 5, 24, 13, 2, 202, 17, 3, 2, 2, 2, 203, 204, 5, 24,
	13, 2, 204, 19, 3, 2, 2, 2, 205, 206, 5, 24, 13, 2, 206, 21, 3, 2, 2, 2,
	207, 208, 7, 51, 2, 2, 208, 209, 5, 10, 6, 2, 209, 210, 7, 56, 2, 2, 210,
	211, 5, 10, 6, 2, 211, 212, 7, 52, 2, 2, 212, 23, 3, 2, 2, 2, 213, 214,
	9, 5, 2, 2, 214, 25, 3, 2, 2, 2, 215, 216, 7, 54, 2, 2, 216, 217, 9, 6,
	2, 2, 217, 27, 3, 2, 2, 2, 218, 219, 7, 7, 2, 2, 219, 29, 3, 2, 2, 2, 7,
	33, 37, 161, 188, 193,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'z'", "'nz'", "'c'", "'nc'", "'SKIP'", "'R0'", "'R1'", "'R2'", "'R3'",
	"'R4'", "'R5'", "'R6'", "'R7'", "'R8'", "'R9'", "'OUT'", "'IN'", "'JSR'",
	"'PCL'", "'PCM'", "'PCH'", "'ADD'", "'ADC'", "'SUB'", "'SBB'", "'OR'",
	"'AND'", "'XOR'", "'MOV'", "'JR'", "'CP'", "'INC'", "'DEC'", "'DSZ'", "'EXR'",
	"'BIT'", "'BSET'", "'BCLR'", "'BTG'", "'RRC'", "'RET'", "'RLC'", "'SL'",
	"'LSR'", "'CPL'", "'NEG'", "'NOP'", "'LPC'", "'['", "']'", "'string'",
	"'.'", "','", "':'", "", "", "", "", "", "' '", "'\n'", "'\r'", "'\t'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "R0", "R1", "R2", "R3", "R4", "R5", "R6", "R7",
	"R8", "R9", "OUT", "IN", "JSR", "PCL", "PCM", "PCH", "ADD", "ADC", "SUB",
	"SBB", "OR", "AND", "XOR", "MOV", "JR", "CP", "INC", "DEC", "DSZ", "EXR",
	"BIT", "BSET", "BCLR", "BTG", "RRC", "RET", "RLC", "SL", "LSR", "CPL",
	"NEG", "NOP", "LPC", "L_BRACKET", "R_BRACKET", "STRING", "DOT", "COMMA",
	"COLON", "CHARACTER", "COMMENT", "LINE_COMMENT", "IDENTIFIER", "NUMBER",
	"WHITESPACE", "NEWLINE", "CARRIAGE", "TAB", "UNKNOWN",
}

var ruleNames = []string{
	"start", "instruction", "synthetic_instruction", "macro", "register_symbol",
	"rg", "flag", "data_byte", "nibble", "quarter", "register_combo", "literal",
	"directive", "mn_SKIP",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AssemblyParser struct {
	*antlr.BaseParser
}

func NewAssemblyParser(input antlr.TokenStream) *AssemblyParser {
	this := new(AssemblyParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Assembly.g4"

	return this
}

// AssemblyParser tokens.
const (
	AssemblyParserEOF          = antlr.TokenEOF
	AssemblyParserT__0         = 1
	AssemblyParserT__1         = 2
	AssemblyParserT__2         = 3
	AssemblyParserT__3         = 4
	AssemblyParserT__4         = 5
	AssemblyParserR0           = 6
	AssemblyParserR1           = 7
	AssemblyParserR2           = 8
	AssemblyParserR3           = 9
	AssemblyParserR4           = 10
	AssemblyParserR5           = 11
	AssemblyParserR6           = 12
	AssemblyParserR7           = 13
	AssemblyParserR8           = 14
	AssemblyParserR9           = 15
	AssemblyParserOUT          = 16
	AssemblyParserIN           = 17
	AssemblyParserJSR          = 18
	AssemblyParserPCL          = 19
	AssemblyParserPCM          = 20
	AssemblyParserPCH          = 21
	AssemblyParserADD          = 22
	AssemblyParserADC          = 23
	AssemblyParserSUB          = 24
	AssemblyParserSBB          = 25
	AssemblyParserOR           = 26
	AssemblyParserAND          = 27
	AssemblyParserXOR          = 28
	AssemblyParserMOV          = 29
	AssemblyParserJR           = 30
	AssemblyParserCP           = 31
	AssemblyParserINC          = 32
	AssemblyParserDEC          = 33
	AssemblyParserDSZ          = 34
	AssemblyParserEXR          = 35
	AssemblyParserBIT          = 36
	AssemblyParserBSET         = 37
	AssemblyParserBCLR         = 38
	AssemblyParserBTG          = 39
	AssemblyParserRRC          = 40
	AssemblyParserRET          = 41
	AssemblyParserRLC          = 42
	AssemblyParserSL           = 43
	AssemblyParserLSR          = 44
	AssemblyParserCPL          = 45
	AssemblyParserNEG          = 46
	AssemblyParserNOP          = 47
	AssemblyParserLPC          = 48
	AssemblyParserL_BRACKET    = 49
	AssemblyParserR_BRACKET    = 50
	AssemblyParserSTRING       = 51
	AssemblyParserDOT          = 52
	AssemblyParserCOMMA        = 53
	AssemblyParserCOLON        = 54
	AssemblyParserCHARACTER    = 55
	AssemblyParserCOMMENT      = 56
	AssemblyParserLINE_COMMENT = 57
	AssemblyParserIDENTIFIER   = 58
	AssemblyParserNUMBER       = 59
	AssemblyParserWHITESPACE   = 60
	AssemblyParserNEWLINE      = 61
	AssemblyParserCARRIAGE     = 62
	AssemblyParserTAB          = 63
	AssemblyParserUNKNOWN      = 64
)

// AssemblyParser rules.
const (
	AssemblyParserRULE_start                 = 0
	AssemblyParserRULE_instruction           = 1
	AssemblyParserRULE_synthetic_instruction = 2
	AssemblyParserRULE_macro                 = 3
	AssemblyParserRULE_register_symbol       = 4
	AssemblyParserRULE_rg                    = 5
	AssemblyParserRULE_flag                  = 6
	AssemblyParserRULE_data_byte             = 7
	AssemblyParserRULE_nibble                = 8
	AssemblyParserRULE_quarter               = 9
	AssemblyParserRULE_register_combo        = 10
	AssemblyParserRULE_literal               = 11
	AssemblyParserRULE_directive             = 12
	AssemblyParserRULE_mn_SKIP               = 13
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(AssemblyParserEOF, 0)
}

func (s *StartContext) AllDirective() []IDirectiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDirectiveContext)(nil)).Elem())
	var tst = make([]IDirectiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDirectiveContext)
		}
	}

	return tst
}

func (s *StartContext) Directive(i int) IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *StartContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *StartContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *StartContext) AllSynthetic_instruction() []ISynthetic_instructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISynthetic_instructionContext)(nil)).Elem())
	var tst = make([]ISynthetic_instructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISynthetic_instructionContext)
		}
	}

	return tst
}

func (s *StartContext) Synthetic_instruction(i int) ISynthetic_instructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISynthetic_instructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISynthetic_instructionContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *AssemblyParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AssemblyParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AssemblyParserDOT {
		{
			p.SetState(28)
			p.Directive()
		}
		p.SetState(31)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AssemblyParserT__4, AssemblyParserADD, AssemblyParserADC, AssemblyParserSBB, AssemblyParserOR, AssemblyParserAND, AssemblyParserXOR, AssemblyParserMOV, AssemblyParserJR, AssemblyParserCP, AssemblyParserINC, AssemblyParserDEC, AssemblyParserDSZ, AssemblyParserEXR, AssemblyParserBIT, AssemblyParserBSET, AssemblyParserBCLR, AssemblyParserBTG, AssemblyParserRRC, AssemblyParserRET, AssemblyParserLPC:
			{
				p.SetState(29)
				p.Instruction()
			}

		case AssemblyParserRLC, AssemblyParserSL, AssemblyParserLSR, AssemblyParserCPL, AssemblyParserNEG, AssemblyParserNOP:
			{
				p.SetState(30)
				p.Synthetic_instruction()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Match(AssemblyParserEOF)
	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) ADD() antlr.TerminalNode {
	return s.GetToken(AssemblyParserADD, 0)
}

func (s *InstructionContext) AllRegister_symbol() []IRegister_symbolContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegister_symbolContext)(nil)).Elem())
	var tst = make([]IRegister_symbolContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegister_symbolContext)
		}
	}

	return tst
}

func (s *InstructionContext) Register_symbol(i int) IRegister_symbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegister_symbolContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegister_symbolContext)
}

func (s *InstructionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(AssemblyParserCOMMA, 0)
}

func (s *InstructionContext) ADC() antlr.TerminalNode {
	return s.GetToken(AssemblyParserADC, 0)
}

func (s *InstructionContext) SBB() antlr.TerminalNode {
	return s.GetToken(AssemblyParserSBB, 0)
}

func (s *InstructionContext) OR() antlr.TerminalNode {
	return s.GetToken(AssemblyParserOR, 0)
}

func (s *InstructionContext) AND() antlr.TerminalNode {
	return s.GetToken(AssemblyParserAND, 0)
}

func (s *InstructionContext) XOR() antlr.TerminalNode {
	return s.GetToken(AssemblyParserXOR, 0)
}

func (s *InstructionContext) MOV() antlr.TerminalNode {
	return s.GetToken(AssemblyParserMOV, 0)
}

func (s *InstructionContext) Nibble() INibbleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INibbleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INibbleContext)
}

func (s *InstructionContext) Register_combo() IRegister_comboContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegister_comboContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegister_comboContext)
}

func (s *InstructionContext) R0() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR0, 0)
}

func (s *InstructionContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(AssemblyParserL_BRACKET, 0)
}

func (s *InstructionContext) Data_byte() IData_byteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_byteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_byteContext)
}

func (s *InstructionContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR_BRACKET, 0)
}

func (s *InstructionContext) LPC() antlr.TerminalNode {
	return s.GetToken(AssemblyParserLPC, 0)
}

func (s *InstructionContext) JR() antlr.TerminalNode {
	return s.GetToken(AssemblyParserJR, 0)
}

func (s *InstructionContext) CP() antlr.TerminalNode {
	return s.GetToken(AssemblyParserCP, 0)
}

func (s *InstructionContext) INC() antlr.TerminalNode {
	return s.GetToken(AssemblyParserINC, 0)
}

func (s *InstructionContext) DEC() antlr.TerminalNode {
	return s.GetToken(AssemblyParserDEC, 0)
}

func (s *InstructionContext) DSZ() antlr.TerminalNode {
	return s.GetToken(AssemblyParserDSZ, 0)
}

func (s *InstructionContext) EXR() antlr.TerminalNode {
	return s.GetToken(AssemblyParserEXR, 0)
}

func (s *InstructionContext) BIT() antlr.TerminalNode {
	return s.GetToken(AssemblyParserBIT, 0)
}

func (s *InstructionContext) Rg() IRgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRgContext)
}

func (s *InstructionContext) Quarter() IQuarterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuarterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuarterContext)
}

func (s *InstructionContext) BSET() antlr.TerminalNode {
	return s.GetToken(AssemblyParserBSET, 0)
}

func (s *InstructionContext) BCLR() antlr.TerminalNode {
	return s.GetToken(AssemblyParserBCLR, 0)
}

func (s *InstructionContext) BTG() antlr.TerminalNode {
	return s.GetToken(AssemblyParserBTG, 0)
}

func (s *InstructionContext) RRC() antlr.TerminalNode {
	return s.GetToken(AssemblyParserRRC, 0)
}

func (s *InstructionContext) RET() antlr.TerminalNode {
	return s.GetToken(AssemblyParserRET, 0)
}

func (s *InstructionContext) Mn_SKIP() IMn_SKIPContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMn_SKIPContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMn_SKIPContext)
}

func (s *InstructionContext) Flag() IFlagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFlagContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *AssemblyParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AssemblyParserRULE_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Match(AssemblyParserADD)
		}
		{
			p.SetState(41)
			p.Register_symbol()
		}
		{
			p.SetState(42)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(43)
			p.Register_symbol()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(AssemblyParserADC)
		}
		{
			p.SetState(46)
			p.Register_symbol()
		}
		{
			p.SetState(47)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(48)
			p.Register_symbol()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Match(AssemblyParserSBB)
		}
		{
			p.SetState(51)
			p.Register_symbol()
		}
		{
			p.SetState(52)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(53)
			p.Register_symbol()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(55)
			p.Match(AssemblyParserOR)
		}
		{
			p.SetState(56)
			p.Register_symbol()
		}
		{
			p.SetState(57)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(58)
			p.Register_symbol()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(60)
			p.Match(AssemblyParserAND)
		}
		{
			p.SetState(61)
			p.Register_symbol()
		}
		{
			p.SetState(62)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(63)
			p.Register_symbol()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(65)
			p.Match(AssemblyParserXOR)
		}
		{
			p.SetState(66)
			p.Register_symbol()
		}
		{
			p.SetState(67)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(68)
			p.Register_symbol()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(70)
			p.Match(AssemblyParserMOV)
		}
		{
			p.SetState(71)
			p.Register_symbol()
		}
		{
			p.SetState(72)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(73)
			p.Register_symbol()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(75)
			p.Match(AssemblyParserMOV)
		}
		{
			p.SetState(76)
			p.Register_symbol()
		}
		{
			p.SetState(77)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(78)
			p.Nibble()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(80)
			p.Match(AssemblyParserMOV)
		}
		{
			p.SetState(81)
			p.Register_combo()
		}
		{
			p.SetState(82)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(83)
			p.Match(AssemblyParserR0)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(85)
			p.Match(AssemblyParserMOV)
		}
		{
			p.SetState(86)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(87)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(88)
			p.Register_combo()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(89)
			p.Match(AssemblyParserMOV)
		}
		{
			p.SetState(90)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(91)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(92)
			p.Match(AssemblyParserL_BRACKET)
		}
		{
			p.SetState(93)
			p.Data_byte()
		}
		{
			p.SetState(94)
			p.Match(AssemblyParserR_BRACKET)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(96)
			p.Match(AssemblyParserLPC)
		}
		{
			p.SetState(97)
			p.Data_byte()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(98)
			p.Match(AssemblyParserJR)
		}
		{
			p.SetState(99)
			p.Data_byte()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(100)
			p.Match(AssemblyParserCP)
		}
		{
			p.SetState(101)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(102)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(103)
			p.Nibble()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(104)
			p.Match(AssemblyParserADD)
		}
		{
			p.SetState(105)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(106)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(107)
			p.Nibble()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(108)
			p.Match(AssemblyParserINC)
		}
		{
			p.SetState(109)
			p.Register_symbol()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(110)
			p.Match(AssemblyParserDEC)
		}
		{
			p.SetState(111)
			p.Register_symbol()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(112)
			p.Match(AssemblyParserDSZ)
		}
		{
			p.SetState(113)
			p.Register_symbol()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(114)
			p.Match(AssemblyParserOR)
		}
		{
			p.SetState(115)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(116)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(117)
			p.Nibble()
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(118)
			p.Match(AssemblyParserAND)
		}
		{
			p.SetState(119)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(120)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(121)
			p.Nibble()
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(122)
			p.Match(AssemblyParserXOR)
		}
		{
			p.SetState(123)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(124)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(125)
			p.Nibble()
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(126)
			p.Match(AssemblyParserEXR)
		}
		{
			p.SetState(127)
			p.Nibble()
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(128)
			p.Match(AssemblyParserBIT)
		}
		{
			p.SetState(129)
			p.Rg()
		}
		{
			p.SetState(130)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(131)
			p.Quarter()
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(133)
			p.Match(AssemblyParserBSET)
		}
		{
			p.SetState(134)
			p.Rg()
		}
		{
			p.SetState(135)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(136)
			p.Quarter()
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(138)
			p.Match(AssemblyParserBCLR)
		}
		{
			p.SetState(139)
			p.Rg()
		}
		{
			p.SetState(140)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(141)
			p.Quarter()
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(143)
			p.Match(AssemblyParserBTG)
		}
		{
			p.SetState(144)
			p.Rg()
		}
		{
			p.SetState(145)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(146)
			p.Quarter()
		}

	case 27:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(148)
			p.Match(AssemblyParserRRC)
		}
		{
			p.SetState(149)
			p.Register_symbol()
		}

	case 28:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(150)
			p.Match(AssemblyParserRET)
		}
		{
			p.SetState(151)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(152)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(153)
			p.Nibble()
		}

	case 29:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(154)
			p.Mn_SKIP()
		}
		{
			p.SetState(155)
			p.Flag()
		}
		{
			p.SetState(156)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(157)
			p.Quarter()
		}

	}

	return localctx
}

// ISynthetic_instructionContext is an interface to support dynamic dispatch.
type ISynthetic_instructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSynthetic_instructionContext differentiates from other interfaces.
	IsSynthetic_instructionContext()
}

type Synthetic_instructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySynthetic_instructionContext() *Synthetic_instructionContext {
	var p = new(Synthetic_instructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_synthetic_instruction
	return p
}

func (*Synthetic_instructionContext) IsSynthetic_instructionContext() {}

func NewSynthetic_instructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Synthetic_instructionContext {
	var p = new(Synthetic_instructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_synthetic_instruction

	return p
}

func (s *Synthetic_instructionContext) GetParser() antlr.Parser { return s.parser }

func (s *Synthetic_instructionContext) RLC() antlr.TerminalNode {
	return s.GetToken(AssemblyParserRLC, 0)
}

func (s *Synthetic_instructionContext) AllRegister_symbol() []IRegister_symbolContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegister_symbolContext)(nil)).Elem())
	var tst = make([]IRegister_symbolContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegister_symbolContext)
		}
	}

	return tst
}

func (s *Synthetic_instructionContext) Register_symbol(i int) IRegister_symbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegister_symbolContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegister_symbolContext)
}

func (s *Synthetic_instructionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(AssemblyParserCOMMA, 0)
}

func (s *Synthetic_instructionContext) SL() antlr.TerminalNode {
	return s.GetToken(AssemblyParserSL, 0)
}

func (s *Synthetic_instructionContext) LSR() antlr.TerminalNode {
	return s.GetToken(AssemblyParserLSR, 0)
}

func (s *Synthetic_instructionContext) CPL() antlr.TerminalNode {
	return s.GetToken(AssemblyParserCPL, 0)
}

func (s *Synthetic_instructionContext) R0() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR0, 0)
}

func (s *Synthetic_instructionContext) NEG() antlr.TerminalNode {
	return s.GetToken(AssemblyParserNEG, 0)
}

func (s *Synthetic_instructionContext) NOP() antlr.TerminalNode {
	return s.GetToken(AssemblyParserNOP, 0)
}

func (s *Synthetic_instructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Synthetic_instructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Synthetic_instructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterSynthetic_instruction(s)
	}
}

func (s *Synthetic_instructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitSynthetic_instruction(s)
	}
}

func (p *AssemblyParser) Synthetic_instruction() (localctx ISynthetic_instructionContext) {
	localctx = NewSynthetic_instructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AssemblyParserRULE_synthetic_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Match(AssemblyParserRLC)
		}
		{
			p.SetState(162)
			p.Register_symbol()
		}
		{
			p.SetState(163)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(164)
			p.Register_symbol()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Match(AssemblyParserSL)
		}
		{
			p.SetState(167)
			p.Register_symbol()
		}
		{
			p.SetState(168)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(169)
			p.Register_symbol()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(171)
			p.Match(AssemblyParserLSR)
		}
		{
			p.SetState(172)
			p.Register_symbol()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(173)
			p.Match(AssemblyParserCPL)
		}
		{
			p.SetState(174)
			p.Match(AssemblyParserR0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(175)
			p.Match(AssemblyParserCPL)
		}
		{
			p.SetState(176)
			p.Register_symbol()
		}
		{
			p.SetState(177)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(178)
			p.Register_symbol()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(180)
			p.Match(AssemblyParserNEG)
		}
		{
			p.SetState(181)
			p.Register_symbol()
		}
		{
			p.SetState(182)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(183)
			p.Register_symbol()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(185)
			p.Match(AssemblyParserNOP)
		}

	}

	return localctx
}

// IMacroContext is an interface to support dynamic dispatch.
type IMacroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMacroContext differentiates from other interfaces.
	IsMacroContext()
}

type MacroContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacroContext() *MacroContext {
	var p = new(MacroContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_macro
	return p
}

func (*MacroContext) IsMacroContext() {}

func NewMacroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MacroContext {
	var p = new(MacroContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_macro

	return p
}

func (s *MacroContext) GetParser() antlr.Parser { return s.parser }

func (s *MacroContext) DOT() antlr.TerminalNode {
	return s.GetToken(AssemblyParserDOT, 0)
}

func (s *MacroContext) STRING() antlr.TerminalNode {
	return s.GetToken(AssemblyParserSTRING, 0)
}

func (s *MacroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MacroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterMacro(s)
	}
}

func (s *MacroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitMacro(s)
	}
}

func (p *AssemblyParser) Macro() (localctx IMacroContext) {
	localctx = NewMacroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AssemblyParserRULE_macro)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(AssemblyParserDOT)
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AssemblyParserSTRING:
		{
			p.SetState(189)
			p.Match(AssemblyParserSTRING)
		}

	case AssemblyParserEOF:

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRegister_symbolContext is an interface to support dynamic dispatch.
type IRegister_symbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegister_symbolContext differentiates from other interfaces.
	IsRegister_symbolContext()
}

type Register_symbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegister_symbolContext() *Register_symbolContext {
	var p = new(Register_symbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_register_symbol
	return p
}

func (*Register_symbolContext) IsRegister_symbolContext() {}

func NewRegister_symbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Register_symbolContext {
	var p = new(Register_symbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_register_symbol

	return p
}

func (s *Register_symbolContext) GetParser() antlr.Parser { return s.parser }

func (s *Register_symbolContext) R0() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR0, 0)
}

func (s *Register_symbolContext) R1() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR1, 0)
}

func (s *Register_symbolContext) R2() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR2, 0)
}

func (s *Register_symbolContext) R3() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR3, 0)
}

func (s *Register_symbolContext) R4() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR4, 0)
}

func (s *Register_symbolContext) R5() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR5, 0)
}

func (s *Register_symbolContext) R6() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR6, 0)
}

func (s *Register_symbolContext) R7() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR7, 0)
}

func (s *Register_symbolContext) R8() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR8, 0)
}

func (s *Register_symbolContext) R9() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR9, 0)
}

func (s *Register_symbolContext) OUT() antlr.TerminalNode {
	return s.GetToken(AssemblyParserOUT, 0)
}

func (s *Register_symbolContext) IN() antlr.TerminalNode {
	return s.GetToken(AssemblyParserIN, 0)
}

func (s *Register_symbolContext) JSR() antlr.TerminalNode {
	return s.GetToken(AssemblyParserJSR, 0)
}

func (s *Register_symbolContext) PCL() antlr.TerminalNode {
	return s.GetToken(AssemblyParserPCL, 0)
}

func (s *Register_symbolContext) PCM() antlr.TerminalNode {
	return s.GetToken(AssemblyParserPCM, 0)
}

func (s *Register_symbolContext) PCH() antlr.TerminalNode {
	return s.GetToken(AssemblyParserPCH, 0)
}

func (s *Register_symbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Register_symbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Register_symbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterRegister_symbol(s)
	}
}

func (s *Register_symbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitRegister_symbol(s)
	}
}

func (p *AssemblyParser) Register_symbol() (localctx IRegister_symbolContext) {
	localctx = NewRegister_symbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AssemblyParserRULE_register_symbol)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AssemblyParserR0)|(1<<AssemblyParserR1)|(1<<AssemblyParserR2)|(1<<AssemblyParserR3)|(1<<AssemblyParserR4)|(1<<AssemblyParserR5)|(1<<AssemblyParserR6)|(1<<AssemblyParserR7)|(1<<AssemblyParserR8)|(1<<AssemblyParserR9)|(1<<AssemblyParserOUT)|(1<<AssemblyParserIN)|(1<<AssemblyParserJSR)|(1<<AssemblyParserPCL)|(1<<AssemblyParserPCM)|(1<<AssemblyParserPCH))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRgContext is an interface to support dynamic dispatch.
type IRgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRgContext differentiates from other interfaces.
	IsRgContext()
}

type RgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRgContext() *RgContext {
	var p = new(RgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_rg
	return p
}

func (*RgContext) IsRgContext() {}

func NewRgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RgContext {
	var p = new(RgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_rg

	return p
}

func (s *RgContext) GetParser() antlr.Parser { return s.parser }

func (s *RgContext) R0() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR0, 0)
}

func (s *RgContext) R1() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR1, 0)
}

func (s *RgContext) R2() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR2, 0)
}

func (s *RgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterRg(s)
	}
}

func (s *RgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitRg(s)
	}
}

func (p *AssemblyParser) Rg() (localctx IRgContext) {
	localctx = NewRgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AssemblyParserRULE_rg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AssemblyParserR0)|(1<<AssemblyParserR1)|(1<<AssemblyParserR2))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFlagContext is an interface to support dynamic dispatch.
type IFlagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlagContext differentiates from other interfaces.
	IsFlagContext()
}

type FlagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlagContext() *FlagContext {
	var p = new(FlagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_flag
	return p
}

func (*FlagContext) IsFlagContext() {}

func NewFlagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlagContext {
	var p = new(FlagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_flag

	return p
}

func (s *FlagContext) GetParser() antlr.Parser { return s.parser }
func (s *FlagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterFlag(s)
	}
}

func (s *FlagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitFlag(s)
	}
}

func (p *AssemblyParser) Flag() (localctx IFlagContext) {
	localctx = NewFlagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AssemblyParserRULE_flag)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AssemblyParserT__0)|(1<<AssemblyParserT__1)|(1<<AssemblyParserT__2)|(1<<AssemblyParserT__3))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IData_byteContext is an interface to support dynamic dispatch.
type IData_byteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsData_byteContext differentiates from other interfaces.
	IsData_byteContext()
}

type Data_byteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_byteContext() *Data_byteContext {
	var p = new(Data_byteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_data_byte
	return p
}

func (*Data_byteContext) IsData_byteContext() {}

func NewData_byteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_byteContext {
	var p = new(Data_byteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_data_byte

	return p
}

func (s *Data_byteContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_byteContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Data_byteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_byteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_byteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterData_byte(s)
	}
}

func (s *Data_byteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitData_byte(s)
	}
}

func (p *AssemblyParser) Data_byte() (localctx IData_byteContext) {
	localctx = NewData_byteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AssemblyParserRULE_data_byte)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Literal()
	}

	return localctx
}

// INibbleContext is an interface to support dynamic dispatch.
type INibbleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNibbleContext differentiates from other interfaces.
	IsNibbleContext()
}

type NibbleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNibbleContext() *NibbleContext {
	var p = new(NibbleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_nibble
	return p
}

func (*NibbleContext) IsNibbleContext() {}

func NewNibbleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NibbleContext {
	var p = new(NibbleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_nibble

	return p
}

func (s *NibbleContext) GetParser() antlr.Parser { return s.parser }

func (s *NibbleContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *NibbleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NibbleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NibbleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterNibble(s)
	}
}

func (s *NibbleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitNibble(s)
	}
}

func (p *AssemblyParser) Nibble() (localctx INibbleContext) {
	localctx = NewNibbleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AssemblyParserRULE_nibble)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Literal()
	}

	return localctx
}

// IQuarterContext is an interface to support dynamic dispatch.
type IQuarterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuarterContext differentiates from other interfaces.
	IsQuarterContext()
}

type QuarterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuarterContext() *QuarterContext {
	var p = new(QuarterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_quarter
	return p
}

func (*QuarterContext) IsQuarterContext() {}

func NewQuarterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuarterContext {
	var p = new(QuarterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_quarter

	return p
}

func (s *QuarterContext) GetParser() antlr.Parser { return s.parser }

func (s *QuarterContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *QuarterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuarterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuarterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterQuarter(s)
	}
}

func (s *QuarterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitQuarter(s)
	}
}

func (p *AssemblyParser) Quarter() (localctx IQuarterContext) {
	localctx = NewQuarterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AssemblyParserRULE_quarter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Literal()
	}

	return localctx
}

// IRegister_comboContext is an interface to support dynamic dispatch.
type IRegister_comboContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegister_comboContext differentiates from other interfaces.
	IsRegister_comboContext()
}

type Register_comboContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegister_comboContext() *Register_comboContext {
	var p = new(Register_comboContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_register_combo
	return p
}

func (*Register_comboContext) IsRegister_comboContext() {}

func NewRegister_comboContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Register_comboContext {
	var p = new(Register_comboContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_register_combo

	return p
}

func (s *Register_comboContext) GetParser() antlr.Parser { return s.parser }

func (s *Register_comboContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(AssemblyParserL_BRACKET, 0)
}

func (s *Register_comboContext) AllRegister_symbol() []IRegister_symbolContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegister_symbolContext)(nil)).Elem())
	var tst = make([]IRegister_symbolContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegister_symbolContext)
		}
	}

	return tst
}

func (s *Register_comboContext) Register_symbol(i int) IRegister_symbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegister_symbolContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegister_symbolContext)
}

func (s *Register_comboContext) COLON() antlr.TerminalNode {
	return s.GetToken(AssemblyParserCOLON, 0)
}

func (s *Register_comboContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR_BRACKET, 0)
}

func (s *Register_comboContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Register_comboContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Register_comboContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterRegister_combo(s)
	}
}

func (s *Register_comboContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitRegister_combo(s)
	}
}

func (p *AssemblyParser) Register_combo() (localctx IRegister_comboContext) {
	localctx = NewRegister_comboContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AssemblyParserRULE_register_combo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(AssemblyParserL_BRACKET)
	}
	{
		p.SetState(206)
		p.Register_symbol()
	}
	{
		p.SetState(207)
		p.Match(AssemblyParserCOLON)
	}
	{
		p.SetState(208)
		p.Register_symbol()
	}
	{
		p.SetState(209)
		p.Match(AssemblyParserR_BRACKET)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AssemblyParserNUMBER, 0)
}

func (s *LiteralContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(AssemblyParserCHARACTER, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *AssemblyParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AssemblyParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AssemblyParserCHARACTER || _la == AssemblyParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) DOT() antlr.TerminalNode {
	return s.GetToken(AssemblyParserDOT, 0)
}

func (s *DirectiveContext) STRING() antlr.TerminalNode {
	return s.GetToken(AssemblyParserSTRING, 0)
}

func (s *DirectiveContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AssemblyParserIDENTIFIER, 0)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (p *AssemblyParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AssemblyParserRULE_directive)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(AssemblyParserDOT)
	}
	{
		p.SetState(214)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AssemblyParserSTRING || _la == AssemblyParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMn_SKIPContext is an interface to support dynamic dispatch.
type IMn_SKIPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMn_SKIPContext differentiates from other interfaces.
	IsMn_SKIPContext()
}

type Mn_SKIPContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMn_SKIPContext() *Mn_SKIPContext {
	var p = new(Mn_SKIPContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_mn_SKIP
	return p
}

func (*Mn_SKIPContext) IsMn_SKIPContext() {}

func NewMn_SKIPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mn_SKIPContext {
	var p = new(Mn_SKIPContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_mn_SKIP

	return p
}

func (s *Mn_SKIPContext) GetParser() antlr.Parser { return s.parser }
func (s *Mn_SKIPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mn_SKIPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mn_SKIPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterMn_SKIP(s)
	}
}

func (s *Mn_SKIPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitMn_SKIP(s)
	}
}

func (p *AssemblyParser) Mn_SKIP() (localctx IMn_SKIPContext) {
	localctx = NewMn_SKIPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AssemblyParserRULE_mn_SKIP)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(AssemblyParserT__4)
	}

	return localctx
}
