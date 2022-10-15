// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Assembly
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type AssemblyParser struct {
	*antlr.BaseParser
}

var assemblyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func assemblyParserInit() {
	staticData := &assemblyParserStaticData
	staticData.literalNames = []string{
		"", "'z'", "'nz'", "'c'", "'nc'", "'R0'", "'R1'", "'R2'", "'R3'", "'R4'",
		"'R5'", "'R6'", "'R7'", "'R8'", "'R9'", "'OUT'", "'IN'", "'JSR'", "'PCL'",
		"'PCM'", "'PCH'", "'ADD'", "'ADC'", "'SUB'", "'SBB'", "'OR'", "'AND'",
		"'XOR'", "'MOV'", "'JR'", "'CP'", "'INC'", "'DEC'", "'DSZ'", "'EXR'",
		"'BIT'", "'BSET'", "'BCLR'", "'BTG'", "'RRC'", "'RET'", "'SKIP'", "'RLC'",
		"'SL'", "'LSR'", "'CPL'", "'NEG'", "'NOP'", "'LPC'", "'['", "']'", "'string'",
		"'.'", "','", "':'", "", "", "", "", "' '", "'\\n'", "'\\r'", "'\\t'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "R0", "R1", "R2", "R3", "R4", "R5", "R6", "R7",
		"R8", "R9", "OUT", "IN", "JSR", "PCL", "PCM", "PCH", "ADD", "ADC", "SUB",
		"SBB", "OR", "AND", "XOR", "MOV", "JR", "CP", "INC", "DEC", "DSZ", "EXR",
		"BIT", "BSET", "BCLR", "BTG", "RRC", "RET", "SKIPI", "RLC", "SL", "LSR",
		"CPL", "NEG", "NOP", "LPC", "L_BRACKET", "R_BRACKET", "STRING", "DOT",
		"COMMA", "COLON", "CHARACTER", "LINE_COMMENT", "LABEL", "NUMBER", "WHITESPACE",
		"NEWLINE", "CARRIAGE", "TAB", "UNKNOWN",
	}
	staticData.ruleNames = []string{
		"start", "instruction", "syntheticInstruction", "macro", "registerSymbol",
		"rg", "flag", "dataByte", "nibble", "quarter", "registerCombo", "literal",
		"directive",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 63, 229, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 3, 0, 28, 8, 0, 1, 0, 1, 0, 3, 0,
		32, 8, 0, 5, 0, 34, 8, 0, 10, 0, 12, 0, 37, 9, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 172, 8,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 199, 8, 2, 1, 3, 1, 3, 1, 3, 3, 3, 204, 8, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 0,
		0, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 5, 1, 0, 5, 20,
		1, 0, 5, 7, 1, 0, 1, 4, 2, 0, 55, 55, 58, 58, 2, 0, 51, 51, 57, 57, 255,
		0, 35, 1, 0, 0, 0, 2, 171, 1, 0, 0, 0, 4, 198, 1, 0, 0, 0, 6, 200, 1, 0,
		0, 0, 8, 205, 1, 0, 0, 0, 10, 207, 1, 0, 0, 0, 12, 209, 1, 0, 0, 0, 14,
		211, 1, 0, 0, 0, 16, 213, 1, 0, 0, 0, 18, 215, 1, 0, 0, 0, 20, 217, 1,
		0, 0, 0, 22, 223, 1, 0, 0, 0, 24, 225, 1, 0, 0, 0, 26, 28, 3, 24, 12, 0,
		27, 26, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29, 32, 3,
		2, 1, 0, 30, 32, 3, 4, 2, 0, 31, 29, 1, 0, 0, 0, 31, 30, 1, 0, 0, 0, 32,
		34, 1, 0, 0, 0, 33, 27, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0,
		0, 35, 36, 1, 0, 0, 0, 36, 38, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 39,
		5, 0, 0, 1, 39, 1, 1, 0, 0, 0, 40, 41, 5, 21, 0, 0, 41, 42, 3, 8, 4, 0,
		42, 43, 5, 53, 0, 0, 43, 44, 3, 8, 4, 0, 44, 172, 1, 0, 0, 0, 45, 46, 5,
		22, 0, 0, 46, 47, 3, 8, 4, 0, 47, 48, 5, 53, 0, 0, 48, 49, 3, 8, 4, 0,
		49, 172, 1, 0, 0, 0, 50, 51, 5, 23, 0, 0, 51, 52, 3, 8, 4, 0, 52, 53, 5,
		53, 0, 0, 53, 54, 3, 8, 4, 0, 54, 172, 1, 0, 0, 0, 55, 56, 5, 24, 0, 0,
		56, 57, 3, 8, 4, 0, 57, 58, 5, 53, 0, 0, 58, 59, 3, 8, 4, 0, 59, 172, 1,
		0, 0, 0, 60, 61, 5, 25, 0, 0, 61, 62, 3, 8, 4, 0, 62, 63, 5, 53, 0, 0,
		63, 64, 3, 8, 4, 0, 64, 172, 1, 0, 0, 0, 65, 66, 5, 26, 0, 0, 66, 67, 3,
		8, 4, 0, 67, 68, 5, 53, 0, 0, 68, 69, 3, 8, 4, 0, 69, 172, 1, 0, 0, 0,
		70, 71, 5, 27, 0, 0, 71, 72, 3, 8, 4, 0, 72, 73, 5, 53, 0, 0, 73, 74, 3,
		8, 4, 0, 74, 172, 1, 0, 0, 0, 75, 76, 5, 28, 0, 0, 76, 77, 3, 8, 4, 0,
		77, 78, 5, 53, 0, 0, 78, 79, 3, 8, 4, 0, 79, 172, 1, 0, 0, 0, 80, 81, 5,
		28, 0, 0, 81, 82, 3, 8, 4, 0, 82, 83, 5, 53, 0, 0, 83, 84, 3, 16, 8, 0,
		84, 172, 1, 0, 0, 0, 85, 86, 5, 28, 0, 0, 86, 87, 3, 20, 10, 0, 87, 88,
		5, 53, 0, 0, 88, 89, 5, 5, 0, 0, 89, 172, 1, 0, 0, 0, 90, 91, 5, 28, 0,
		0, 91, 92, 5, 5, 0, 0, 92, 93, 5, 53, 0, 0, 93, 172, 3, 20, 10, 0, 94,
		95, 5, 28, 0, 0, 95, 96, 5, 5, 0, 0, 96, 97, 5, 53, 0, 0, 97, 98, 5, 49,
		0, 0, 98, 99, 3, 14, 7, 0, 99, 100, 5, 50, 0, 0, 100, 172, 1, 0, 0, 0,
		101, 102, 5, 28, 0, 0, 102, 103, 5, 49, 0, 0, 103, 104, 3, 14, 7, 0, 104,
		105, 5, 50, 0, 0, 105, 106, 5, 53, 0, 0, 106, 107, 5, 5, 0, 0, 107, 172,
		1, 0, 0, 0, 108, 109, 5, 48, 0, 0, 109, 172, 3, 14, 7, 0, 110, 111, 5,
		29, 0, 0, 111, 172, 3, 14, 7, 0, 112, 113, 5, 30, 0, 0, 113, 114, 5, 5,
		0, 0, 114, 115, 5, 53, 0, 0, 115, 172, 3, 16, 8, 0, 116, 117, 5, 21, 0,
		0, 117, 118, 5, 5, 0, 0, 118, 119, 5, 53, 0, 0, 119, 172, 3, 16, 8, 0,
		120, 121, 5, 31, 0, 0, 121, 172, 3, 8, 4, 0, 122, 123, 5, 32, 0, 0, 123,
		172, 3, 8, 4, 0, 124, 125, 5, 33, 0, 0, 125, 172, 3, 8, 4, 0, 126, 127,
		5, 25, 0, 0, 127, 128, 5, 5, 0, 0, 128, 129, 5, 53, 0, 0, 129, 172, 3,
		16, 8, 0, 130, 131, 5, 26, 0, 0, 131, 132, 5, 5, 0, 0, 132, 133, 5, 53,
		0, 0, 133, 172, 3, 16, 8, 0, 134, 135, 5, 27, 0, 0, 135, 136, 5, 5, 0,
		0, 136, 137, 5, 53, 0, 0, 137, 172, 3, 16, 8, 0, 138, 139, 5, 34, 0, 0,
		139, 172, 3, 16, 8, 0, 140, 141, 5, 35, 0, 0, 141, 142, 3, 10, 5, 0, 142,
		143, 5, 53, 0, 0, 143, 144, 3, 18, 9, 0, 144, 172, 1, 0, 0, 0, 145, 146,
		5, 36, 0, 0, 146, 147, 3, 10, 5, 0, 147, 148, 5, 53, 0, 0, 148, 149, 3,
		18, 9, 0, 149, 172, 1, 0, 0, 0, 150, 151, 5, 37, 0, 0, 151, 152, 3, 10,
		5, 0, 152, 153, 5, 53, 0, 0, 153, 154, 3, 18, 9, 0, 154, 172, 1, 0, 0,
		0, 155, 156, 5, 38, 0, 0, 156, 157, 3, 10, 5, 0, 157, 158, 5, 53, 0, 0,
		158, 159, 3, 18, 9, 0, 159, 172, 1, 0, 0, 0, 160, 161, 5, 39, 0, 0, 161,
		172, 3, 8, 4, 0, 162, 163, 5, 40, 0, 0, 163, 164, 5, 5, 0, 0, 164, 165,
		5, 53, 0, 0, 165, 172, 3, 16, 8, 0, 166, 167, 5, 41, 0, 0, 167, 168, 3,
		12, 6, 0, 168, 169, 5, 53, 0, 0, 169, 170, 3, 18, 9, 0, 170, 172, 1, 0,
		0, 0, 171, 40, 1, 0, 0, 0, 171, 45, 1, 0, 0, 0, 171, 50, 1, 0, 0, 0, 171,
		55, 1, 0, 0, 0, 171, 60, 1, 0, 0, 0, 171, 65, 1, 0, 0, 0, 171, 70, 1, 0,
		0, 0, 171, 75, 1, 0, 0, 0, 171, 80, 1, 0, 0, 0, 171, 85, 1, 0, 0, 0, 171,
		90, 1, 0, 0, 0, 171, 94, 1, 0, 0, 0, 171, 101, 1, 0, 0, 0, 171, 108, 1,
		0, 0, 0, 171, 110, 1, 0, 0, 0, 171, 112, 1, 0, 0, 0, 171, 116, 1, 0, 0,
		0, 171, 120, 1, 0, 0, 0, 171, 122, 1, 0, 0, 0, 171, 124, 1, 0, 0, 0, 171,
		126, 1, 0, 0, 0, 171, 130, 1, 0, 0, 0, 171, 134, 1, 0, 0, 0, 171, 138,
		1, 0, 0, 0, 171, 140, 1, 0, 0, 0, 171, 145, 1, 0, 0, 0, 171, 150, 1, 0,
		0, 0, 171, 155, 1, 0, 0, 0, 171, 160, 1, 0, 0, 0, 171, 162, 1, 0, 0, 0,
		171, 166, 1, 0, 0, 0, 172, 3, 1, 0, 0, 0, 173, 174, 5, 42, 0, 0, 174, 175,
		3, 8, 4, 0, 175, 176, 5, 53, 0, 0, 176, 177, 3, 8, 4, 0, 177, 199, 1, 0,
		0, 0, 178, 179, 5, 43, 0, 0, 179, 180, 3, 8, 4, 0, 180, 181, 5, 53, 0,
		0, 181, 182, 3, 8, 4, 0, 182, 199, 1, 0, 0, 0, 183, 184, 5, 44, 0, 0, 184,
		199, 3, 8, 4, 0, 185, 186, 5, 45, 0, 0, 186, 199, 5, 5, 0, 0, 187, 188,
		5, 45, 0, 0, 188, 189, 3, 8, 4, 0, 189, 190, 5, 53, 0, 0, 190, 191, 3,
		8, 4, 0, 191, 199, 1, 0, 0, 0, 192, 193, 5, 46, 0, 0, 193, 194, 3, 8, 4,
		0, 194, 195, 5, 53, 0, 0, 195, 196, 3, 8, 4, 0, 196, 199, 1, 0, 0, 0, 197,
		199, 5, 47, 0, 0, 198, 173, 1, 0, 0, 0, 198, 178, 1, 0, 0, 0, 198, 183,
		1, 0, 0, 0, 198, 185, 1, 0, 0, 0, 198, 187, 1, 0, 0, 0, 198, 192, 1, 0,
		0, 0, 198, 197, 1, 0, 0, 0, 199, 5, 1, 0, 0, 0, 200, 203, 5, 52, 0, 0,
		201, 204, 5, 51, 0, 0, 202, 204, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203,
		202, 1, 0, 0, 0, 204, 7, 1, 0, 0, 0, 205, 206, 7, 0, 0, 0, 206, 9, 1, 0,
		0, 0, 207, 208, 7, 1, 0, 0, 208, 11, 1, 0, 0, 0, 209, 210, 7, 2, 0, 0,
		210, 13, 1, 0, 0, 0, 211, 212, 3, 22, 11, 0, 212, 15, 1, 0, 0, 0, 213,
		214, 3, 22, 11, 0, 214, 17, 1, 0, 0, 0, 215, 216, 3, 22, 11, 0, 216, 19,
		1, 0, 0, 0, 217, 218, 5, 49, 0, 0, 218, 219, 3, 8, 4, 0, 219, 220, 5, 54,
		0, 0, 220, 221, 3, 8, 4, 0, 221, 222, 5, 50, 0, 0, 222, 21, 1, 0, 0, 0,
		223, 224, 7, 3, 0, 0, 224, 23, 1, 0, 0, 0, 225, 226, 5, 52, 0, 0, 226,
		227, 7, 4, 0, 0, 227, 25, 1, 0, 0, 0, 6, 27, 31, 35, 171, 198, 203,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// AssemblyParserInit initializes any static state used to implement AssemblyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAssemblyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AssemblyParserInit() {
	staticData := &assemblyParserStaticData
	staticData.once.Do(assemblyParserInit)
}

// NewAssemblyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAssemblyParser(input antlr.TokenStream) *AssemblyParser {
	AssemblyParserInit()
	this := new(AssemblyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &assemblyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// AssemblyParser tokens.
const (
	AssemblyParserEOF          = antlr.TokenEOF
	AssemblyParserT__0         = 1
	AssemblyParserT__1         = 2
	AssemblyParserT__2         = 3
	AssemblyParserT__3         = 4
	AssemblyParserR0           = 5
	AssemblyParserR1           = 6
	AssemblyParserR2           = 7
	AssemblyParserR3           = 8
	AssemblyParserR4           = 9
	AssemblyParserR5           = 10
	AssemblyParserR6           = 11
	AssemblyParserR7           = 12
	AssemblyParserR8           = 13
	AssemblyParserR9           = 14
	AssemblyParserOUT          = 15
	AssemblyParserIN           = 16
	AssemblyParserJSR          = 17
	AssemblyParserPCL          = 18
	AssemblyParserPCM          = 19
	AssemblyParserPCH          = 20
	AssemblyParserADD          = 21
	AssemblyParserADC          = 22
	AssemblyParserSUB          = 23
	AssemblyParserSBB          = 24
	AssemblyParserOR           = 25
	AssemblyParserAND          = 26
	AssemblyParserXOR          = 27
	AssemblyParserMOV          = 28
	AssemblyParserJR           = 29
	AssemblyParserCP           = 30
	AssemblyParserINC          = 31
	AssemblyParserDEC          = 32
	AssemblyParserDSZ          = 33
	AssemblyParserEXR          = 34
	AssemblyParserBIT          = 35
	AssemblyParserBSET         = 36
	AssemblyParserBCLR         = 37
	AssemblyParserBTG          = 38
	AssemblyParserRRC          = 39
	AssemblyParserRET          = 40
	AssemblyParserSKIPI        = 41
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
	AssemblyParserLINE_COMMENT = 56
	AssemblyParserLABEL        = 57
	AssemblyParserNUMBER       = 58
	AssemblyParserWHITESPACE   = 59
	AssemblyParserNEWLINE      = 60
	AssemblyParserCARRIAGE     = 61
	AssemblyParserTAB          = 62
	AssemblyParserUNKNOWN      = 63
)

// AssemblyParser rules.
const (
	AssemblyParserRULE_start                = 0
	AssemblyParserRULE_instruction          = 1
	AssemblyParserRULE_syntheticInstruction = 2
	AssemblyParserRULE_macro                = 3
	AssemblyParserRULE_registerSymbol       = 4
	AssemblyParserRULE_rg                   = 5
	AssemblyParserRULE_flag                 = 6
	AssemblyParserRULE_dataByte             = 7
	AssemblyParserRULE_nibble               = 8
	AssemblyParserRULE_quarter              = 9
	AssemblyParserRULE_registerCombo        = 10
	AssemblyParserRULE_literal              = 11
	AssemblyParserRULE_directive            = 12
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

func (s *StartContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *StartContext) AllSyntheticInstruction() []ISyntheticInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISyntheticInstructionContext); ok {
			len++
		}
	}

	tst := make([]ISyntheticInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISyntheticInstructionContext); ok {
			tst[i] = t.(ISyntheticInstructionContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) SyntheticInstruction(i int) ISyntheticInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISyntheticInstructionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISyntheticInstructionContext)
}

func (s *StartContext) AllDirective() []IDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDirectiveContext); ok {
			len++
		}
	}

	tst := make([]IDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDirectiveContext); ok {
			tst[i] = t.(IDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) Directive(i int) IDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectiveContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
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
	this := p
	_ = this

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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5066549578694656) != 0 {
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AssemblyParserDOT {
			{
				p.SetState(26)
				p.Directive()
			}

		}
		p.SetState(31)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AssemblyParserADD, AssemblyParserADC, AssemblyParserSUB, AssemblyParserSBB, AssemblyParserOR, AssemblyParserAND, AssemblyParserXOR, AssemblyParserMOV, AssemblyParserJR, AssemblyParserCP, AssemblyParserINC, AssemblyParserDEC, AssemblyParserDSZ, AssemblyParserEXR, AssemblyParserBIT, AssemblyParserBSET, AssemblyParserBCLR, AssemblyParserBTG, AssemblyParserRRC, AssemblyParserRET, AssemblyParserSKIPI, AssemblyParserLPC:
			{
				p.SetState(29)
				p.Instruction()
			}

		case AssemblyParserRLC, AssemblyParserSL, AssemblyParserLSR, AssemblyParserCPL, AssemblyParserNEG, AssemblyParserNOP:
			{
				p.SetState(30)
				p.SyntheticInstruction()
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

func (s *InstructionContext) AllRegisterSymbol() []IRegisterSymbolContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegisterSymbolContext); ok {
			len++
		}
	}

	tst := make([]IRegisterSymbolContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegisterSymbolContext); ok {
			tst[i] = t.(IRegisterSymbolContext)
			i++
		}
	}

	return tst
}

func (s *InstructionContext) RegisterSymbol(i int) IRegisterSymbolContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterSymbolContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegisterSymbolContext)
}

func (s *InstructionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(AssemblyParserCOMMA, 0)
}

func (s *InstructionContext) ADC() antlr.TerminalNode {
	return s.GetToken(AssemblyParserADC, 0)
}

func (s *InstructionContext) SUB() antlr.TerminalNode {
	return s.GetToken(AssemblyParserSUB, 0)
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INibbleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INibbleContext)
}

func (s *InstructionContext) RegisterCombo() IRegisterComboContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterComboContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegisterComboContext)
}

func (s *InstructionContext) R0() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR0, 0)
}

func (s *InstructionContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(AssemblyParserL_BRACKET, 0)
}

func (s *InstructionContext) DataByte() IDataByteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataByteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataByteContext)
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRgContext)
}

func (s *InstructionContext) Quarter() IQuarterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuarterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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

func (s *InstructionContext) SKIPI() antlr.TerminalNode {
	return s.GetToken(AssemblyParserSKIPI, 0)
}

func (s *InstructionContext) Flag() IFlagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFlagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Match(AssemblyParserADD)
		}
		{
			p.SetState(41)
			p.RegisterSymbol()
		}
		{
			p.SetState(42)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(43)
			p.RegisterSymbol()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(AssemblyParserADC)
		}
		{
			p.SetState(46)
			p.RegisterSymbol()
		}
		{
			p.SetState(47)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(48)
			p.RegisterSymbol()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Match(AssemblyParserSUB)
		}
		{
			p.SetState(51)
			p.RegisterSymbol()
		}
		{
			p.SetState(52)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(53)
			p.RegisterSymbol()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(55)
			p.Match(AssemblyParserSBB)
		}
		{
			p.SetState(56)
			p.RegisterSymbol()
		}
		{
			p.SetState(57)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(58)
			p.RegisterSymbol()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(60)
			p.Match(AssemblyParserOR)
		}
		{
			p.SetState(61)
			p.RegisterSymbol()
		}
		{
			p.SetState(62)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(63)
			p.RegisterSymbol()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(65)
			p.Match(AssemblyParserAND)
		}
		{
			p.SetState(66)
			p.RegisterSymbol()
		}
		{
			p.SetState(67)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(68)
			p.RegisterSymbol()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(70)
			p.Match(AssemblyParserXOR)
		}
		{
			p.SetState(71)
			p.RegisterSymbol()
		}
		{
			p.SetState(72)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(73)
			p.RegisterSymbol()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(75)
			p.Match(AssemblyParserMOV)
		}
		{
			p.SetState(76)
			p.RegisterSymbol()
		}
		{
			p.SetState(77)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(78)
			p.RegisterSymbol()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(80)
			p.Match(AssemblyParserMOV)
		}
		{
			p.SetState(81)
			p.RegisterSymbol()
		}
		{
			p.SetState(82)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(83)
			p.Nibble()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(85)
			p.Match(AssemblyParserMOV)
		}
		{
			p.SetState(86)
			p.RegisterCombo()
		}
		{
			p.SetState(87)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(88)
			p.Match(AssemblyParserR0)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(90)
			p.Match(AssemblyParserMOV)
		}
		{
			p.SetState(91)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(92)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(93)
			p.RegisterCombo()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(94)
			p.Match(AssemblyParserMOV)
		}
		{
			p.SetState(95)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(96)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(97)
			p.Match(AssemblyParserL_BRACKET)
		}
		{
			p.SetState(98)
			p.DataByte()
		}
		{
			p.SetState(99)
			p.Match(AssemblyParserR_BRACKET)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(101)
			p.Match(AssemblyParserMOV)
		}
		{
			p.SetState(102)
			p.Match(AssemblyParserL_BRACKET)
		}
		{
			p.SetState(103)
			p.DataByte()
		}
		{
			p.SetState(104)
			p.Match(AssemblyParserR_BRACKET)
		}
		{
			p.SetState(105)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(106)
			p.Match(AssemblyParserR0)
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(108)
			p.Match(AssemblyParserLPC)
		}
		{
			p.SetState(109)
			p.DataByte()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(110)
			p.Match(AssemblyParserJR)
		}
		{
			p.SetState(111)
			p.DataByte()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(112)
			p.Match(AssemblyParserCP)
		}
		{
			p.SetState(113)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(114)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(115)
			p.Nibble()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(116)
			p.Match(AssemblyParserADD)
		}
		{
			p.SetState(117)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(118)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(119)
			p.Nibble()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(120)
			p.Match(AssemblyParserINC)
		}
		{
			p.SetState(121)
			p.RegisterSymbol()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(122)
			p.Match(AssemblyParserDEC)
		}
		{
			p.SetState(123)
			p.RegisterSymbol()
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(124)
			p.Match(AssemblyParserDSZ)
		}
		{
			p.SetState(125)
			p.RegisterSymbol()
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(126)
			p.Match(AssemblyParserOR)
		}
		{
			p.SetState(127)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(128)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(129)
			p.Nibble()
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(130)
			p.Match(AssemblyParserAND)
		}
		{
			p.SetState(131)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(132)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(133)
			p.Nibble()
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(134)
			p.Match(AssemblyParserXOR)
		}
		{
			p.SetState(135)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(136)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(137)
			p.Nibble()
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(138)
			p.Match(AssemblyParserEXR)
		}
		{
			p.SetState(139)
			p.Nibble()
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(140)
			p.Match(AssemblyParserBIT)
		}
		{
			p.SetState(141)
			p.Rg()
		}
		{
			p.SetState(142)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(143)
			p.Quarter()
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(145)
			p.Match(AssemblyParserBSET)
		}
		{
			p.SetState(146)
			p.Rg()
		}
		{
			p.SetState(147)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(148)
			p.Quarter()
		}

	case 27:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(150)
			p.Match(AssemblyParserBCLR)
		}
		{
			p.SetState(151)
			p.Rg()
		}
		{
			p.SetState(152)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(153)
			p.Quarter()
		}

	case 28:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(155)
			p.Match(AssemblyParserBTG)
		}
		{
			p.SetState(156)
			p.Rg()
		}
		{
			p.SetState(157)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(158)
			p.Quarter()
		}

	case 29:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(160)
			p.Match(AssemblyParserRRC)
		}
		{
			p.SetState(161)
			p.RegisterSymbol()
		}

	case 30:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(162)
			p.Match(AssemblyParserRET)
		}
		{
			p.SetState(163)
			p.Match(AssemblyParserR0)
		}
		{
			p.SetState(164)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(165)
			p.Nibble()
		}

	case 31:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(166)
			p.Match(AssemblyParserSKIPI)
		}
		{
			p.SetState(167)
			p.Flag()
		}
		{
			p.SetState(168)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(169)
			p.Quarter()
		}

	}

	return localctx
}

// ISyntheticInstructionContext is an interface to support dynamic dispatch.
type ISyntheticInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSyntheticInstructionContext differentiates from other interfaces.
	IsSyntheticInstructionContext()
}

type SyntheticInstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySyntheticInstructionContext() *SyntheticInstructionContext {
	var p = new(SyntheticInstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_syntheticInstruction
	return p
}

func (*SyntheticInstructionContext) IsSyntheticInstructionContext() {}

func NewSyntheticInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntheticInstructionContext {
	var p = new(SyntheticInstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_syntheticInstruction

	return p
}

func (s *SyntheticInstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntheticInstructionContext) RLC() antlr.TerminalNode {
	return s.GetToken(AssemblyParserRLC, 0)
}

func (s *SyntheticInstructionContext) AllRegisterSymbol() []IRegisterSymbolContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegisterSymbolContext); ok {
			len++
		}
	}

	tst := make([]IRegisterSymbolContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegisterSymbolContext); ok {
			tst[i] = t.(IRegisterSymbolContext)
			i++
		}
	}

	return tst
}

func (s *SyntheticInstructionContext) RegisterSymbol(i int) IRegisterSymbolContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterSymbolContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegisterSymbolContext)
}

func (s *SyntheticInstructionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(AssemblyParserCOMMA, 0)
}

func (s *SyntheticInstructionContext) SL() antlr.TerminalNode {
	return s.GetToken(AssemblyParserSL, 0)
}

func (s *SyntheticInstructionContext) LSR() antlr.TerminalNode {
	return s.GetToken(AssemblyParserLSR, 0)
}

func (s *SyntheticInstructionContext) CPL() antlr.TerminalNode {
	return s.GetToken(AssemblyParserCPL, 0)
}

func (s *SyntheticInstructionContext) R0() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR0, 0)
}

func (s *SyntheticInstructionContext) NEG() antlr.TerminalNode {
	return s.GetToken(AssemblyParserNEG, 0)
}

func (s *SyntheticInstructionContext) NOP() antlr.TerminalNode {
	return s.GetToken(AssemblyParserNOP, 0)
}

func (s *SyntheticInstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntheticInstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntheticInstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterSyntheticInstruction(s)
	}
}

func (s *SyntheticInstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitSyntheticInstruction(s)
	}
}

func (p *AssemblyParser) SyntheticInstruction() (localctx ISyntheticInstructionContext) {
	this := p
	_ = this

	localctx = NewSyntheticInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AssemblyParserRULE_syntheticInstruction)

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

	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.Match(AssemblyParserRLC)
		}
		{
			p.SetState(174)
			p.RegisterSymbol()
		}
		{
			p.SetState(175)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(176)
			p.RegisterSymbol()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
			p.Match(AssemblyParserSL)
		}
		{
			p.SetState(179)
			p.RegisterSymbol()
		}
		{
			p.SetState(180)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(181)
			p.RegisterSymbol()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(183)
			p.Match(AssemblyParserLSR)
		}
		{
			p.SetState(184)
			p.RegisterSymbol()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(185)
			p.Match(AssemblyParserCPL)
		}
		{
			p.SetState(186)
			p.Match(AssemblyParserR0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(187)
			p.Match(AssemblyParserCPL)
		}
		{
			p.SetState(188)
			p.RegisterSymbol()
		}
		{
			p.SetState(189)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(190)
			p.RegisterSymbol()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(192)
			p.Match(AssemblyParserNEG)
		}
		{
			p.SetState(193)
			p.RegisterSymbol()
		}
		{
			p.SetState(194)
			p.Match(AssemblyParserCOMMA)
		}
		{
			p.SetState(195)
			p.RegisterSymbol()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(197)
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
	this := p
	_ = this

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
		p.SetState(200)
		p.Match(AssemblyParserDOT)
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AssemblyParserSTRING:
		{
			p.SetState(201)
			p.Match(AssemblyParserSTRING)
		}

	case AssemblyParserEOF:

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRegisterSymbolContext is an interface to support dynamic dispatch.
type IRegisterSymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegisterSymbolContext differentiates from other interfaces.
	IsRegisterSymbolContext()
}

type RegisterSymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegisterSymbolContext() *RegisterSymbolContext {
	var p = new(RegisterSymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_registerSymbol
	return p
}

func (*RegisterSymbolContext) IsRegisterSymbolContext() {}

func NewRegisterSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegisterSymbolContext {
	var p = new(RegisterSymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_registerSymbol

	return p
}

func (s *RegisterSymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *RegisterSymbolContext) R0() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR0, 0)
}

func (s *RegisterSymbolContext) R1() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR1, 0)
}

func (s *RegisterSymbolContext) R2() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR2, 0)
}

func (s *RegisterSymbolContext) R3() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR3, 0)
}

func (s *RegisterSymbolContext) R4() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR4, 0)
}

func (s *RegisterSymbolContext) R5() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR5, 0)
}

func (s *RegisterSymbolContext) R6() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR6, 0)
}

func (s *RegisterSymbolContext) R7() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR7, 0)
}

func (s *RegisterSymbolContext) R8() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR8, 0)
}

func (s *RegisterSymbolContext) R9() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR9, 0)
}

func (s *RegisterSymbolContext) OUT() antlr.TerminalNode {
	return s.GetToken(AssemblyParserOUT, 0)
}

func (s *RegisterSymbolContext) IN() antlr.TerminalNode {
	return s.GetToken(AssemblyParserIN, 0)
}

func (s *RegisterSymbolContext) JSR() antlr.TerminalNode {
	return s.GetToken(AssemblyParserJSR, 0)
}

func (s *RegisterSymbolContext) PCL() antlr.TerminalNode {
	return s.GetToken(AssemblyParserPCL, 0)
}

func (s *RegisterSymbolContext) PCM() antlr.TerminalNode {
	return s.GetToken(AssemblyParserPCM, 0)
}

func (s *RegisterSymbolContext) PCH() antlr.TerminalNode {
	return s.GetToken(AssemblyParserPCH, 0)
}

func (s *RegisterSymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegisterSymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegisterSymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterRegisterSymbol(s)
	}
}

func (s *RegisterSymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitRegisterSymbol(s)
	}
}

func (p *AssemblyParser) RegisterSymbol() (localctx IRegisterSymbolContext) {
	this := p
	_ = this

	localctx = NewRegisterSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AssemblyParserRULE_registerSymbol)
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
		p.SetState(205)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097120) != 0) {
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
	this := p
	_ = this

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
		p.SetState(207)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&224) != 0) {
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
	this := p
	_ = this

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
		p.SetState(209)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDataByteContext is an interface to support dynamic dispatch.
type IDataByteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataByteContext differentiates from other interfaces.
	IsDataByteContext()
}

type DataByteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataByteContext() *DataByteContext {
	var p = new(DataByteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_dataByte
	return p
}

func (*DataByteContext) IsDataByteContext() {}

func NewDataByteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataByteContext {
	var p = new(DataByteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_dataByte

	return p
}

func (s *DataByteContext) GetParser() antlr.Parser { return s.parser }

func (s *DataByteContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *DataByteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataByteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataByteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterDataByte(s)
	}
}

func (s *DataByteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitDataByte(s)
	}
}

func (p *AssemblyParser) DataByte() (localctx IDataByteContext) {
	this := p
	_ = this

	localctx = NewDataByteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AssemblyParserRULE_dataByte)

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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
		p.SetState(213)
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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
		p.SetState(215)
		p.Literal()
	}

	return localctx
}

// IRegisterComboContext is an interface to support dynamic dispatch.
type IRegisterComboContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegisterComboContext differentiates from other interfaces.
	IsRegisterComboContext()
}

type RegisterComboContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegisterComboContext() *RegisterComboContext {
	var p = new(RegisterComboContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AssemblyParserRULE_registerCombo
	return p
}

func (*RegisterComboContext) IsRegisterComboContext() {}

func NewRegisterComboContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegisterComboContext {
	var p = new(RegisterComboContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblyParserRULE_registerCombo

	return p
}

func (s *RegisterComboContext) GetParser() antlr.Parser { return s.parser }

func (s *RegisterComboContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(AssemblyParserL_BRACKET, 0)
}

func (s *RegisterComboContext) AllRegisterSymbol() []IRegisterSymbolContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegisterSymbolContext); ok {
			len++
		}
	}

	tst := make([]IRegisterSymbolContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegisterSymbolContext); ok {
			tst[i] = t.(IRegisterSymbolContext)
			i++
		}
	}

	return tst
}

func (s *RegisterComboContext) RegisterSymbol(i int) IRegisterSymbolContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterSymbolContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegisterSymbolContext)
}

func (s *RegisterComboContext) COLON() antlr.TerminalNode {
	return s.GetToken(AssemblyParserCOLON, 0)
}

func (s *RegisterComboContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(AssemblyParserR_BRACKET, 0)
}

func (s *RegisterComboContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegisterComboContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegisterComboContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.EnterRegisterCombo(s)
	}
}

func (s *RegisterComboContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblyListener); ok {
		listenerT.ExitRegisterCombo(s)
	}
}

func (p *AssemblyParser) RegisterCombo() (localctx IRegisterComboContext) {
	this := p
	_ = this

	localctx = NewRegisterComboContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AssemblyParserRULE_registerCombo)

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
		p.SetState(217)
		p.Match(AssemblyParserL_BRACKET)
	}
	{
		p.SetState(218)
		p.RegisterSymbol()
	}
	{
		p.SetState(219)
		p.Match(AssemblyParserCOLON)
	}
	{
		p.SetState(220)
		p.RegisterSymbol()
	}
	{
		p.SetState(221)
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
	this := p
	_ = this

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
		p.SetState(223)
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

func (s *DirectiveContext) LABEL() antlr.TerminalNode {
	return s.GetToken(AssemblyParserLABEL, 0)
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
	this := p
	_ = this

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
		p.SetState(225)
		p.Match(AssemblyParserDOT)
	}
	{
		p.SetState(226)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AssemblyParserSTRING || _la == AssemblyParserLABEL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
