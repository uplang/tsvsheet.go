// Code generated from TsvsheetParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tsvsheetgrammar // TsvsheetParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TsvsheetParser struct {
	*antlr.BaseParser
}

var TsvsheetParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tsvsheetparserParserInit() {
	staticData := &TsvsheetParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'>='", "'<='", "'<>'", "'>'", "'<'", "'header'", "'body'", "'final'",
		"'='", "'('", "')'", "'['", "']'", "':'", "','", "'$'", "'*'", "'+'",
		"'-'", "'/'", "'%'", "'!'", "", "", "", "", "'\\t'",
	}
	staticData.SymbolicNames = []string{
		"", "GE", "LE", "NE", "GT", "LT", "HEADER", "BODY", "FINAL", "EQ", "LPAREN",
		"RPAREN", "LBRACK", "RBRACK", "COLON", "COMMA", "DOLLAR", "STAR", "PLUS",
		"DASH", "SLASH", "PERCENT", "BANG", "NUMBER", "COL", "NAME", "STRING",
		"TAB", "NL", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"worksheet", "line", "sectionCommand", "structuralCommand", "cells",
		"cell", "payload", "formula", "literal", "reference", "rangeRef", "endpoint",
		"cellRef", "colRef", "rowWildcard", "rowRef", "numericRef", "signedInt",
		"numRow", "groupedRange", "modifier", "expression", "functionCall",
		"argList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 243, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 5, 0, 52, 8, 0,
		10, 0, 12, 0, 55, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 62, 8, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 73, 8, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 3, 4, 80, 8, 4, 1, 4, 1, 4, 3, 4, 84, 8, 4, 5,
		4, 86, 8, 4, 10, 4, 12, 4, 89, 9, 4, 1, 5, 1, 5, 1, 5, 3, 5, 94, 8, 5,
		1, 5, 3, 5, 97, 8, 5, 1, 5, 3, 5, 100, 8, 5, 1, 6, 1, 6, 1, 6, 3, 6, 105,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9, 114, 8, 9, 1, 10,
		1, 10, 1, 10, 3, 10, 119, 8, 10, 1, 11, 1, 11, 3, 11, 123, 8, 11, 1, 12,
		1, 12, 3, 12, 127, 8, 12, 1, 12, 3, 12, 130, 8, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 3, 13, 137, 8, 13, 1, 14, 1, 14, 3, 14, 141, 8, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 151, 8, 15, 3,
		15, 153, 8, 15, 1, 16, 1, 16, 3, 16, 157, 8, 16, 1, 16, 1, 16, 3, 16, 161,
		8, 16, 3, 16, 163, 8, 16, 1, 16, 1, 16, 1, 17, 3, 17, 168, 8, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 177, 8, 18, 3, 18, 179,
		8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 187, 8, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 195, 8, 19, 3, 19, 197, 8,
		19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 3, 21, 212, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 223, 8, 21, 10, 21, 12, 21, 226,
		9, 21, 1, 22, 1, 22, 1, 22, 3, 22, 231, 8, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 23, 5, 23, 238, 8, 23, 10, 23, 12, 23, 241, 9, 23, 1, 23, 0, 1,
		42, 24, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 0, 6, 2, 0, 23, 23, 25, 26, 1, 0, 18, 19, 2,
		0, 4, 5, 22, 22, 2, 0, 17, 17, 20, 21, 2, 0, 1, 5, 9, 9, 1, 0, 24, 25,
		265, 0, 48, 1, 0, 0, 0, 2, 61, 1, 0, 0, 0, 4, 72, 1, 0, 0, 0, 6, 74, 1,
		0, 0, 0, 8, 79, 1, 0, 0, 0, 10, 99, 1, 0, 0, 0, 12, 104, 1, 0, 0, 0, 14,
		106, 1, 0, 0, 0, 16, 109, 1, 0, 0, 0, 18, 113, 1, 0, 0, 0, 20, 115, 1,
		0, 0, 0, 22, 122, 1, 0, 0, 0, 24, 129, 1, 0, 0, 0, 26, 136, 1, 0, 0, 0,
		28, 138, 1, 0, 0, 0, 30, 152, 1, 0, 0, 0, 32, 154, 1, 0, 0, 0, 34, 167,
		1, 0, 0, 0, 36, 178, 1, 0, 0, 0, 38, 196, 1, 0, 0, 0, 40, 198, 1, 0, 0,
		0, 42, 211, 1, 0, 0, 0, 44, 227, 1, 0, 0, 0, 46, 234, 1, 0, 0, 0, 48, 53,
		3, 2, 1, 0, 49, 50, 5, 28, 0, 0, 50, 52, 3, 2, 1, 0, 51, 49, 1, 0, 0, 0,
		52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 56, 1,
		0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 57, 5, 0, 0, 1, 57, 1, 1, 0, 0, 0, 58,
		62, 3, 4, 2, 0, 59, 62, 3, 6, 3, 0, 60, 62, 3, 8, 4, 0, 61, 58, 1, 0, 0,
		0, 61, 59, 1, 0, 0, 0, 61, 60, 1, 0, 0, 0, 62, 3, 1, 0, 0, 0, 63, 64, 5,
		9, 0, 0, 64, 65, 5, 6, 0, 0, 65, 66, 5, 10, 0, 0, 66, 67, 5, 23, 0, 0,
		67, 73, 5, 11, 0, 0, 68, 69, 5, 9, 0, 0, 69, 73, 5, 7, 0, 0, 70, 71, 5,
		9, 0, 0, 71, 73, 5, 8, 0, 0, 72, 63, 1, 0, 0, 0, 72, 68, 1, 0, 0, 0, 72,
		70, 1, 0, 0, 0, 73, 5, 1, 0, 0, 0, 74, 75, 5, 9, 0, 0, 75, 76, 3, 18, 9,
		0, 76, 77, 3, 40, 20, 0, 77, 7, 1, 0, 0, 0, 78, 80, 3, 10, 5, 0, 79, 78,
		1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 87, 1, 0, 0, 0, 81, 83, 5, 27, 0, 0,
		82, 84, 3, 10, 5, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 86, 1,
		0, 0, 0, 85, 81, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87,
		88, 1, 0, 0, 0, 88, 9, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 100, 3, 14,
		7, 0, 91, 93, 3, 18, 9, 0, 92, 94, 3, 40, 20, 0, 93, 92, 1, 0, 0, 0, 93,
		94, 1, 0, 0, 0, 94, 96, 1, 0, 0, 0, 95, 97, 3, 12, 6, 0, 96, 95, 1, 0,
		0, 0, 96, 97, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 100, 3, 16, 8, 0, 99,
		90, 1, 0, 0, 0, 99, 91, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 11, 1, 0,
		0, 0, 101, 105, 3, 14, 7, 0, 102, 103, 5, 9, 0, 0, 103, 105, 3, 16, 8,
		0, 104, 101, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105, 13, 1, 0, 0, 0, 106,
		107, 5, 9, 0, 0, 107, 108, 3, 42, 21, 0, 108, 15, 1, 0, 0, 0, 109, 110,
		7, 0, 0, 0, 110, 17, 1, 0, 0, 0, 111, 114, 3, 20, 10, 0, 112, 114, 3, 38,
		19, 0, 113, 111, 1, 0, 0, 0, 113, 112, 1, 0, 0, 0, 114, 19, 1, 0, 0, 0,
		115, 118, 3, 22, 11, 0, 116, 117, 5, 14, 0, 0, 117, 119, 3, 22, 11, 0,
		118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 21, 1, 0, 0, 0, 120, 123,
		3, 24, 12, 0, 121, 123, 3, 32, 16, 0, 122, 120, 1, 0, 0, 0, 122, 121, 1,
		0, 0, 0, 123, 23, 1, 0, 0, 0, 124, 126, 3, 26, 13, 0, 125, 127, 3, 30,
		15, 0, 126, 125, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0,
		128, 130, 3, 28, 14, 0, 129, 124, 1, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130,
		25, 1, 0, 0, 0, 131, 132, 5, 16, 0, 0, 132, 137, 5, 24, 0, 0, 133, 137,
		5, 24, 0, 0, 134, 137, 5, 16, 0, 0, 135, 137, 5, 26, 0, 0, 136, 131, 1,
		0, 0, 0, 136, 133, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 135, 1, 0, 0,
		0, 137, 27, 1, 0, 0, 0, 138, 140, 5, 17, 0, 0, 139, 141, 3, 30, 15, 0,
		140, 139, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 29, 1, 0, 0, 0, 142, 153,
		5, 23, 0, 0, 143, 144, 5, 18, 0, 0, 144, 153, 5, 23, 0, 0, 145, 153, 5,
		17, 0, 0, 146, 150, 5, 16, 0, 0, 147, 148, 7, 1, 0, 0, 148, 151, 5, 23,
		0, 0, 149, 151, 5, 23, 0, 0, 150, 147, 1, 0, 0, 0, 150, 149, 1, 0, 0, 0,
		150, 151, 1, 0, 0, 0, 151, 153, 1, 0, 0, 0, 152, 142, 1, 0, 0, 0, 152,
		143, 1, 0, 0, 0, 152, 145, 1, 0, 0, 0, 152, 146, 1, 0, 0, 0, 153, 31, 1,
		0, 0, 0, 154, 156, 5, 12, 0, 0, 155, 157, 3, 34, 17, 0, 156, 155, 1, 0,
		0, 0, 156, 157, 1, 0, 0, 0, 157, 162, 1, 0, 0, 0, 158, 160, 5, 15, 0, 0,
		159, 161, 3, 36, 18, 0, 160, 159, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161,
		163, 1, 0, 0, 0, 162, 158, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164,
		1, 0, 0, 0, 164, 165, 5, 13, 0, 0, 165, 33, 1, 0, 0, 0, 166, 168, 5, 19,
		0, 0, 167, 166, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0,
		169, 170, 5, 23, 0, 0, 170, 35, 1, 0, 0, 0, 171, 179, 3, 34, 17, 0, 172,
		176, 5, 16, 0, 0, 173, 174, 7, 1, 0, 0, 174, 177, 5, 23, 0, 0, 175, 177,
		5, 23, 0, 0, 176, 173, 1, 0, 0, 0, 176, 175, 1, 0, 0, 0, 176, 177, 1, 0,
		0, 0, 177, 179, 1, 0, 0, 0, 178, 171, 1, 0, 0, 0, 178, 172, 1, 0, 0, 0,
		179, 37, 1, 0, 0, 0, 180, 181, 5, 10, 0, 0, 181, 182, 3, 26, 13, 0, 182,
		183, 5, 14, 0, 0, 183, 184, 3, 26, 13, 0, 184, 186, 5, 11, 0, 0, 185, 187,
		3, 30, 15, 0, 186, 185, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 197, 1,
		0, 0, 0, 188, 189, 5, 10, 0, 0, 189, 190, 3, 32, 16, 0, 190, 191, 5, 14,
		0, 0, 191, 192, 3, 32, 16, 0, 192, 194, 5, 11, 0, 0, 193, 195, 3, 30, 15,
		0, 194, 193, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0, 196,
		180, 1, 0, 0, 0, 196, 188, 1, 0, 0, 0, 197, 39, 1, 0, 0, 0, 198, 199, 7,
		2, 0, 0, 199, 41, 1, 0, 0, 0, 200, 201, 6, 21, -1, 0, 201, 202, 5, 10,
		0, 0, 202, 203, 3, 42, 21, 0, 203, 204, 5, 11, 0, 0, 204, 212, 1, 0, 0,
		0, 205, 206, 7, 1, 0, 0, 206, 212, 3, 42, 21, 8, 207, 212, 3, 44, 22, 0,
		208, 212, 3, 18, 9, 0, 209, 212, 5, 23, 0, 0, 210, 212, 5, 26, 0, 0, 211,
		200, 1, 0, 0, 0, 211, 205, 1, 0, 0, 0, 211, 207, 1, 0, 0, 0, 211, 208,
		1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211, 210, 1, 0, 0, 0, 212, 224, 1, 0,
		0, 0, 213, 214, 10, 7, 0, 0, 214, 215, 7, 3, 0, 0, 215, 223, 3, 42, 21,
		8, 216, 217, 10, 6, 0, 0, 217, 218, 7, 1, 0, 0, 218, 223, 3, 42, 21, 7,
		219, 220, 10, 5, 0, 0, 220, 221, 7, 4, 0, 0, 221, 223, 3, 42, 21, 6, 222,
		213, 1, 0, 0, 0, 222, 216, 1, 0, 0, 0, 222, 219, 1, 0, 0, 0, 223, 226,
		1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 43, 1, 0,
		0, 0, 226, 224, 1, 0, 0, 0, 227, 228, 7, 5, 0, 0, 228, 230, 5, 10, 0, 0,
		229, 231, 3, 46, 23, 0, 230, 229, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231,
		232, 1, 0, 0, 0, 232, 233, 5, 11, 0, 0, 233, 45, 1, 0, 0, 0, 234, 239,
		3, 42, 21, 0, 235, 236, 5, 15, 0, 0, 236, 238, 3, 42, 21, 0, 237, 235,
		1, 0, 0, 0, 238, 241, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0,
		0, 0, 240, 47, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 33, 53, 61, 72, 79, 83,
		87, 93, 96, 99, 104, 113, 118, 122, 126, 129, 136, 140, 150, 152, 156,
		160, 162, 167, 176, 178, 186, 194, 196, 211, 222, 224, 230, 239,
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

// TsvsheetParserInit initializes any static state used to implement TsvsheetParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTsvsheetParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TsvsheetParserInit() {
	staticData := &TsvsheetParserParserStaticData
	staticData.once.Do(tsvsheetparserParserInit)
}

// NewTsvsheetParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTsvsheetParser(input antlr.TokenStream) *TsvsheetParser {
	TsvsheetParserInit()
	this := new(TsvsheetParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TsvsheetParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TsvsheetParser.g4"

	return this
}

// TsvsheetParser tokens.
const (
	TsvsheetParserEOF     = antlr.TokenEOF
	TsvsheetParserGE      = 1
	TsvsheetParserLE      = 2
	TsvsheetParserNE      = 3
	TsvsheetParserGT      = 4
	TsvsheetParserLT      = 5
	TsvsheetParserHEADER  = 6
	TsvsheetParserBODY    = 7
	TsvsheetParserFINAL   = 8
	TsvsheetParserEQ      = 9
	TsvsheetParserLPAREN  = 10
	TsvsheetParserRPAREN  = 11
	TsvsheetParserLBRACK  = 12
	TsvsheetParserRBRACK  = 13
	TsvsheetParserCOLON   = 14
	TsvsheetParserCOMMA   = 15
	TsvsheetParserDOLLAR  = 16
	TsvsheetParserSTAR    = 17
	TsvsheetParserPLUS    = 18
	TsvsheetParserDASH    = 19
	TsvsheetParserSLASH   = 20
	TsvsheetParserPERCENT = 21
	TsvsheetParserBANG    = 22
	TsvsheetParserNUMBER  = 23
	TsvsheetParserCOL     = 24
	TsvsheetParserNAME    = 25
	TsvsheetParserSTRING  = 26
	TsvsheetParserTAB     = 27
	TsvsheetParserNL      = 28
	TsvsheetParserCOMMENT = 29
	TsvsheetParserWS      = 30
)

// TsvsheetParser rules.
const (
	TsvsheetParserRULE_worksheet         = 0
	TsvsheetParserRULE_line              = 1
	TsvsheetParserRULE_sectionCommand    = 2
	TsvsheetParserRULE_structuralCommand = 3
	TsvsheetParserRULE_cells             = 4
	TsvsheetParserRULE_cell              = 5
	TsvsheetParserRULE_payload           = 6
	TsvsheetParserRULE_formula           = 7
	TsvsheetParserRULE_literal           = 8
	TsvsheetParserRULE_reference         = 9
	TsvsheetParserRULE_rangeRef          = 10
	TsvsheetParserRULE_endpoint          = 11
	TsvsheetParserRULE_cellRef           = 12
	TsvsheetParserRULE_colRef            = 13
	TsvsheetParserRULE_rowWildcard       = 14
	TsvsheetParserRULE_rowRef            = 15
	TsvsheetParserRULE_numericRef        = 16
	TsvsheetParserRULE_signedInt         = 17
	TsvsheetParserRULE_numRow            = 18
	TsvsheetParserRULE_groupedRange      = 19
	TsvsheetParserRULE_modifier          = 20
	TsvsheetParserRULE_expression        = 21
	TsvsheetParserRULE_functionCall      = 22
	TsvsheetParserRULE_argList           = 23
)

// IWorksheetContext is an interface to support dynamic dispatch.
type IWorksheetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLine() []ILineContext
	Line(i int) ILineContext
	EOF() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsWorksheetContext differentiates from other interfaces.
	IsWorksheetContext()
}

type WorksheetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWorksheetContext() *WorksheetContext {
	var p = new(WorksheetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_worksheet
	return p
}

func InitEmptyWorksheetContext(p *WorksheetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_worksheet
}

func (*WorksheetContext) IsWorksheetContext() {}

func NewWorksheetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WorksheetContext {
	var p = new(WorksheetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_worksheet

	return p
}

func (s *WorksheetContext) GetParser() antlr.Parser { return s.parser }

func (s *WorksheetContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *WorksheetContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
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

	return t.(ILineContext)
}

func (s *WorksheetContext) EOF() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserEOF, 0)
}

func (s *WorksheetContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(TsvsheetParserNL)
}

func (s *WorksheetContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(TsvsheetParserNL, i)
}

func (s *WorksheetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WorksheetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WorksheetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterWorksheet(s)
	}
}

func (s *WorksheetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitWorksheet(s)
	}
}

func (s *WorksheetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitWorksheet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) Worksheet() (localctx IWorksheetContext) {
	localctx = NewWorksheetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TsvsheetParserRULE_worksheet)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Line()
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TsvsheetParserNL {
		{
			p.SetState(49)
			p.Match(TsvsheetParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Line()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
		p.Match(TsvsheetParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SectionCommand() ISectionCommandContext
	StructuralCommand() IStructuralCommandContext
	Cells() ICellsContext

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_line
	return p
}

func InitEmptyLineContext(p *LineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_line
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) SectionCommand() ISectionCommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISectionCommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISectionCommandContext)
}

func (s *LineContext) StructuralCommand() IStructuralCommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructuralCommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructuralCommandContext)
}

func (s *LineContext) Cells() ICellsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICellsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICellsContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitLine(s)
	}
}

func (s *LineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TsvsheetParserRULE_line)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.SectionCommand()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.StructuralCommand()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
			p.Cells()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISectionCommandContext is an interface to support dynamic dispatch.
type ISectionCommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	HEADER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	BODY() antlr.TerminalNode
	FINAL() antlr.TerminalNode

	// IsSectionCommandContext differentiates from other interfaces.
	IsSectionCommandContext()
}

type SectionCommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionCommandContext() *SectionCommandContext {
	var p = new(SectionCommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_sectionCommand
	return p
}

func InitEmptySectionCommandContext(p *SectionCommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_sectionCommand
}

func (*SectionCommandContext) IsSectionCommandContext() {}

func NewSectionCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionCommandContext {
	var p = new(SectionCommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_sectionCommand

	return p
}

func (s *SectionCommandContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionCommandContext) EQ() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserEQ, 0)
}

func (s *SectionCommandContext) HEADER() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserHEADER, 0)
}

func (s *SectionCommandContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserLPAREN, 0)
}

func (s *SectionCommandContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserNUMBER, 0)
}

func (s *SectionCommandContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserRPAREN, 0)
}

func (s *SectionCommandContext) BODY() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserBODY, 0)
}

func (s *SectionCommandContext) FINAL() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserFINAL, 0)
}

func (s *SectionCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionCommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterSectionCommand(s)
	}
}

func (s *SectionCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitSectionCommand(s)
	}
}

func (s *SectionCommandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitSectionCommand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) SectionCommand() (localctx ISectionCommandContext) {
	localctx = NewSectionCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TsvsheetParserRULE_sectionCommand)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Match(TsvsheetParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(TsvsheetParserHEADER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Match(TsvsheetParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(TsvsheetParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(TsvsheetParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.Match(TsvsheetParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(TsvsheetParserBODY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(70)
			p.Match(TsvsheetParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(TsvsheetParserFINAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructuralCommandContext is an interface to support dynamic dispatch.
type IStructuralCommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	Reference() IReferenceContext
	Modifier() IModifierContext

	// IsStructuralCommandContext differentiates from other interfaces.
	IsStructuralCommandContext()
}

type StructuralCommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructuralCommandContext() *StructuralCommandContext {
	var p = new(StructuralCommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_structuralCommand
	return p
}

func InitEmptyStructuralCommandContext(p *StructuralCommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_structuralCommand
}

func (*StructuralCommandContext) IsStructuralCommandContext() {}

func NewStructuralCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructuralCommandContext {
	var p = new(StructuralCommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_structuralCommand

	return p
}

func (s *StructuralCommandContext) GetParser() antlr.Parser { return s.parser }

func (s *StructuralCommandContext) EQ() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserEQ, 0)
}

func (s *StructuralCommandContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *StructuralCommandContext) Modifier() IModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *StructuralCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructuralCommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructuralCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterStructuralCommand(s)
	}
}

func (s *StructuralCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitStructuralCommand(s)
	}
}

func (s *StructuralCommandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitStructuralCommand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) StructuralCommand() (localctx IStructuralCommandContext) {
	localctx = NewStructuralCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TsvsheetParserRULE_structuralCommand)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(TsvsheetParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Reference()
	}
	{
		p.SetState(76)
		p.Modifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICellsContext is an interface to support dynamic dispatch.
type ICellsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCell() []ICellContext
	Cell(i int) ICellContext
	AllTAB() []antlr.TerminalNode
	TAB(i int) antlr.TerminalNode

	// IsCellsContext differentiates from other interfaces.
	IsCellsContext()
}

type CellsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCellsContext() *CellsContext {
	var p = new(CellsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_cells
	return p
}

func InitEmptyCellsContext(p *CellsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_cells
}

func (*CellsContext) IsCellsContext() {}

func NewCellsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CellsContext {
	var p = new(CellsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_cells

	return p
}

func (s *CellsContext) GetParser() antlr.Parser { return s.parser }

func (s *CellsContext) AllCell() []ICellContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICellContext); ok {
			len++
		}
	}

	tst := make([]ICellContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICellContext); ok {
			tst[i] = t.(ICellContext)
			i++
		}
	}

	return tst
}

func (s *CellsContext) Cell(i int) ICellContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICellContext); ok {
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

	return t.(ICellContext)
}

func (s *CellsContext) AllTAB() []antlr.TerminalNode {
	return s.GetTokens(TsvsheetParserTAB)
}

func (s *CellsContext) TAB(i int) antlr.TerminalNode {
	return s.GetToken(TsvsheetParserTAB, i)
}

func (s *CellsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CellsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CellsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterCells(s)
	}
}

func (s *CellsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitCells(s)
	}
}

func (s *CellsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitCells(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) Cells() (localctx ICellsContext) {
	localctx = NewCellsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TsvsheetParserRULE_cells)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126031360) != 0 {
		{
			p.SetState(78)
			p.Cell()
		}

	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TsvsheetParserTAB {
		{
			p.SetState(81)
			p.Match(TsvsheetParserTAB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126031360) != 0 {
			{
				p.SetState(82)
				p.Cell()
			}

		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICellContext is an interface to support dynamic dispatch.
type ICellContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Formula() IFormulaContext
	Reference() IReferenceContext
	Modifier() IModifierContext
	Payload() IPayloadContext
	Literal() ILiteralContext

	// IsCellContext differentiates from other interfaces.
	IsCellContext()
}

type CellContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCellContext() *CellContext {
	var p = new(CellContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_cell
	return p
}

func InitEmptyCellContext(p *CellContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_cell
}

func (*CellContext) IsCellContext() {}

func NewCellContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CellContext {
	var p = new(CellContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_cell

	return p
}

func (s *CellContext) GetParser() antlr.Parser { return s.parser }

func (s *CellContext) Formula() IFormulaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormulaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormulaContext)
}

func (s *CellContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *CellContext) Modifier() IModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *CellContext) Payload() IPayloadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPayloadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPayloadContext)
}

func (s *CellContext) Literal() ILiteralContext {
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

func (s *CellContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CellContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CellContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterCell(s)
	}
}

func (s *CellContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitCell(s)
	}
}

func (s *CellContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitCell(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) Cell() (localctx ICellContext) {
	localctx = NewCellContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TsvsheetParserRULE_cell)
	var _la int

	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Formula()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Reference()
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4194352) != 0 {
			{
				p.SetState(92)
				p.Modifier()
			}

		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TsvsheetParserEQ {
			{
				p.SetState(95)
				p.Payload()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.Literal()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPayloadContext is an interface to support dynamic dispatch.
type IPayloadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Formula() IFormulaContext
	EQ() antlr.TerminalNode
	Literal() ILiteralContext

	// IsPayloadContext differentiates from other interfaces.
	IsPayloadContext()
}

type PayloadContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPayloadContext() *PayloadContext {
	var p = new(PayloadContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_payload
	return p
}

func InitEmptyPayloadContext(p *PayloadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_payload
}

func (*PayloadContext) IsPayloadContext() {}

func NewPayloadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PayloadContext {
	var p = new(PayloadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_payload

	return p
}

func (s *PayloadContext) GetParser() antlr.Parser { return s.parser }

func (s *PayloadContext) Formula() IFormulaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormulaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormulaContext)
}

func (s *PayloadContext) EQ() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserEQ, 0)
}

func (s *PayloadContext) Literal() ILiteralContext {
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

func (s *PayloadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PayloadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PayloadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterPayload(s)
	}
}

func (s *PayloadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitPayload(s)
	}
}

func (s *PayloadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitPayload(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) Payload() (localctx IPayloadContext) {
	localctx = NewPayloadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TsvsheetParserRULE_payload)
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Formula()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Match(TsvsheetParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Literal()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFormulaContext is an interface to support dynamic dispatch.
type IFormulaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	Expression() IExpressionContext

	// IsFormulaContext differentiates from other interfaces.
	IsFormulaContext()
}

type FormulaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormulaContext() *FormulaContext {
	var p = new(FormulaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_formula
	return p
}

func InitEmptyFormulaContext(p *FormulaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_formula
}

func (*FormulaContext) IsFormulaContext() {}

func NewFormulaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormulaContext {
	var p = new(FormulaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_formula

	return p
}

func (s *FormulaContext) GetParser() antlr.Parser { return s.parser }

func (s *FormulaContext) EQ() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserEQ, 0)
}

func (s *FormulaContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FormulaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormulaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormulaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterFormula(s)
	}
}

func (s *FormulaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitFormula(s)
	}
}

func (s *FormulaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitFormula(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) Formula() (localctx IFormulaContext) {
	localctx = NewFormulaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TsvsheetParserRULE_formula)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(TsvsheetParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NAME() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserNAME, 0)
}

func (s *LiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserNUMBER, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserSTRING, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TsvsheetParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&109051904) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RangeRef() IRangeRefContext
	GroupedRange() IGroupedRangeContext

	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_reference
	return p
}

func InitEmptyReferenceContext(p *ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_reference
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) RangeRef() IRangeRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeRefContext)
}

func (s *ReferenceContext) GroupedRange() IGroupedRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupedRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupedRangeContext)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitReference(s)
	}
}

func (s *ReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitReference(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) Reference() (localctx IReferenceContext) {
	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TsvsheetParserRULE_reference)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TsvsheetParserLBRACK, TsvsheetParserDOLLAR, TsvsheetParserSTAR, TsvsheetParserCOL, TsvsheetParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.RangeRef()
		}

	case TsvsheetParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.GroupedRange()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeRefContext is an interface to support dynamic dispatch.
type IRangeRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEndpoint() []IEndpointContext
	Endpoint(i int) IEndpointContext
	COLON() antlr.TerminalNode

	// IsRangeRefContext differentiates from other interfaces.
	IsRangeRefContext()
}

type RangeRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeRefContext() *RangeRefContext {
	var p = new(RangeRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_rangeRef
	return p
}

func InitEmptyRangeRefContext(p *RangeRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_rangeRef
}

func (*RangeRefContext) IsRangeRefContext() {}

func NewRangeRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeRefContext {
	var p = new(RangeRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_rangeRef

	return p
}

func (s *RangeRefContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeRefContext) AllEndpoint() []IEndpointContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEndpointContext); ok {
			len++
		}
	}

	tst := make([]IEndpointContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEndpointContext); ok {
			tst[i] = t.(IEndpointContext)
			i++
		}
	}

	return tst
}

func (s *RangeRefContext) Endpoint(i int) IEndpointContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndpointContext); ok {
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

	return t.(IEndpointContext)
}

func (s *RangeRefContext) COLON() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserCOLON, 0)
}

func (s *RangeRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterRangeRef(s)
	}
}

func (s *RangeRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitRangeRef(s)
	}
}

func (s *RangeRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitRangeRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) RangeRef() (localctx IRangeRefContext) {
	localctx = NewRangeRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TsvsheetParserRULE_rangeRef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Endpoint()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(116)
			p.Match(TsvsheetParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Endpoint()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEndpointContext is an interface to support dynamic dispatch.
type IEndpointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CellRef() ICellRefContext
	NumericRef() INumericRefContext

	// IsEndpointContext differentiates from other interfaces.
	IsEndpointContext()
}

type EndpointContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndpointContext() *EndpointContext {
	var p = new(EndpointContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_endpoint
	return p
}

func InitEmptyEndpointContext(p *EndpointContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_endpoint
}

func (*EndpointContext) IsEndpointContext() {}

func NewEndpointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndpointContext {
	var p = new(EndpointContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_endpoint

	return p
}

func (s *EndpointContext) GetParser() antlr.Parser { return s.parser }

func (s *EndpointContext) CellRef() ICellRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICellRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICellRefContext)
}

func (s *EndpointContext) NumericRef() INumericRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericRefContext)
}

func (s *EndpointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndpointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndpointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterEndpoint(s)
	}
}

func (s *EndpointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitEndpoint(s)
	}
}

func (s *EndpointContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitEndpoint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) Endpoint() (localctx IEndpointContext) {
	localctx = NewEndpointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TsvsheetParserRULE_endpoint)
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TsvsheetParserDOLLAR, TsvsheetParserSTAR, TsvsheetParserCOL, TsvsheetParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.CellRef()
		}

	case TsvsheetParserLBRACK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.NumericRef()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICellRefContext is an interface to support dynamic dispatch.
type ICellRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ColRef() IColRefContext
	RowRef() IRowRefContext
	RowWildcard() IRowWildcardContext

	// IsCellRefContext differentiates from other interfaces.
	IsCellRefContext()
}

type CellRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCellRefContext() *CellRefContext {
	var p = new(CellRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_cellRef
	return p
}

func InitEmptyCellRefContext(p *CellRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_cellRef
}

func (*CellRefContext) IsCellRefContext() {}

func NewCellRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CellRefContext {
	var p = new(CellRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_cellRef

	return p
}

func (s *CellRefContext) GetParser() antlr.Parser { return s.parser }

func (s *CellRefContext) ColRef() IColRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColRefContext)
}

func (s *CellRefContext) RowRef() IRowRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRowRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRowRefContext)
}

func (s *CellRefContext) RowWildcard() IRowWildcardContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRowWildcardContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRowWildcardContext)
}

func (s *CellRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CellRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CellRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterCellRef(s)
	}
}

func (s *CellRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitCellRef(s)
	}
}

func (s *CellRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitCellRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) CellRef() (localctx ICellRefContext) {
	localctx = NewCellRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TsvsheetParserRULE_cellRef)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TsvsheetParserDOLLAR, TsvsheetParserCOL, TsvsheetParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.ColRef()
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(125)
				p.RowRef()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case TsvsheetParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.RowWildcard()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColRefContext is an interface to support dynamic dispatch.
type IColRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOLLAR() antlr.TerminalNode
	COL() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsColRefContext differentiates from other interfaces.
	IsColRefContext()
}

type ColRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColRefContext() *ColRefContext {
	var p = new(ColRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_colRef
	return p
}

func InitEmptyColRefContext(p *ColRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_colRef
}

func (*ColRefContext) IsColRefContext() {}

func NewColRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColRefContext {
	var p = new(ColRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_colRef

	return p
}

func (s *ColRefContext) GetParser() antlr.Parser { return s.parser }

func (s *ColRefContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserDOLLAR, 0)
}

func (s *ColRefContext) COL() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserCOL, 0)
}

func (s *ColRefContext) STRING() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserSTRING, 0)
}

func (s *ColRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterColRef(s)
	}
}

func (s *ColRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitColRef(s)
	}
}

func (s *ColRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitColRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) ColRef() (localctx IColRefContext) {
	localctx = NewColRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TsvsheetParserRULE_colRef)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Match(TsvsheetParserDOLLAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(TsvsheetParserCOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.Match(TsvsheetParserCOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.Match(TsvsheetParserDOLLAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(135)
			p.Match(TsvsheetParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRowWildcardContext is an interface to support dynamic dispatch.
type IRowWildcardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	RowRef() IRowRefContext

	// IsRowWildcardContext differentiates from other interfaces.
	IsRowWildcardContext()
}

type RowWildcardContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowWildcardContext() *RowWildcardContext {
	var p = new(RowWildcardContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_rowWildcard
	return p
}

func InitEmptyRowWildcardContext(p *RowWildcardContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_rowWildcard
}

func (*RowWildcardContext) IsRowWildcardContext() {}

func NewRowWildcardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowWildcardContext {
	var p = new(RowWildcardContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_rowWildcard

	return p
}

func (s *RowWildcardContext) GetParser() antlr.Parser { return s.parser }

func (s *RowWildcardContext) STAR() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserSTAR, 0)
}

func (s *RowWildcardContext) RowRef() IRowRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRowRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRowRefContext)
}

func (s *RowWildcardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowWildcardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowWildcardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterRowWildcard(s)
	}
}

func (s *RowWildcardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitRowWildcard(s)
	}
}

func (s *RowWildcardContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitRowWildcard(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) RowWildcard() (localctx IRowWildcardContext) {
	localctx = NewRowWildcardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TsvsheetParserRULE_rowWildcard)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(TsvsheetParserSTAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(139)
			p.RowRef()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRowRefContext is an interface to support dynamic dispatch.
type IRowRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	STAR() antlr.TerminalNode
	DOLLAR() antlr.TerminalNode
	DASH() antlr.TerminalNode

	// IsRowRefContext differentiates from other interfaces.
	IsRowRefContext()
}

type RowRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowRefContext() *RowRefContext {
	var p = new(RowRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_rowRef
	return p
}

func InitEmptyRowRefContext(p *RowRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_rowRef
}

func (*RowRefContext) IsRowRefContext() {}

func NewRowRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowRefContext {
	var p = new(RowRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_rowRef

	return p
}

func (s *RowRefContext) GetParser() antlr.Parser { return s.parser }

func (s *RowRefContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserNUMBER, 0)
}

func (s *RowRefContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserPLUS, 0)
}

func (s *RowRefContext) STAR() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserSTAR, 0)
}

func (s *RowRefContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserDOLLAR, 0)
}

func (s *RowRefContext) DASH() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserDASH, 0)
}

func (s *RowRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterRowRef(s)
	}
}

func (s *RowRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitRowRef(s)
	}
}

func (s *RowRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitRowRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) RowRef() (localctx IRowRefContext) {
	localctx = NewRowRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TsvsheetParserRULE_rowRef)
	var _la int

	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TsvsheetParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Match(TsvsheetParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TsvsheetParserPLUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.Match(TsvsheetParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.Match(TsvsheetParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TsvsheetParserSTAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(145)
			p.Match(TsvsheetParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TsvsheetParserDOLLAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(146)
			p.Match(TsvsheetParserDOLLAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(147)
				_la = p.GetTokenStream().LA(1)

				if !(_la == TsvsheetParserPLUS || _la == TsvsheetParserDASH) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(148)
				p.Match(TsvsheetParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(149)
				p.Match(TsvsheetParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericRefContext is an interface to support dynamic dispatch.
type INumericRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	SignedInt() ISignedIntContext
	COMMA() antlr.TerminalNode
	NumRow() INumRowContext

	// IsNumericRefContext differentiates from other interfaces.
	IsNumericRefContext()
}

type NumericRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericRefContext() *NumericRefContext {
	var p = new(NumericRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_numericRef
	return p
}

func InitEmptyNumericRefContext(p *NumericRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_numericRef
}

func (*NumericRefContext) IsNumericRefContext() {}

func NewNumericRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericRefContext {
	var p = new(NumericRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_numericRef

	return p
}

func (s *NumericRefContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericRefContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserLBRACK, 0)
}

func (s *NumericRefContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserRBRACK, 0)
}

func (s *NumericRefContext) SignedInt() ISignedIntContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignedIntContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignedIntContext)
}

func (s *NumericRefContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserCOMMA, 0)
}

func (s *NumericRefContext) NumRow() INumRowContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumRowContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumRowContext)
}

func (s *NumericRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterNumericRef(s)
	}
}

func (s *NumericRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitNumericRef(s)
	}
}

func (s *NumericRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitNumericRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) NumericRef() (localctx INumericRefContext) {
	localctx = NewNumericRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TsvsheetParserRULE_numericRef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(TsvsheetParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TsvsheetParserDASH || _la == TsvsheetParserNUMBER {
		{
			p.SetState(155)
			p.SignedInt()
		}

	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TsvsheetParserCOMMA {
		{
			p.SetState(158)
			p.Match(TsvsheetParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8978432) != 0 {
			{
				p.SetState(159)
				p.NumRow()
			}

		}

	}
	{
		p.SetState(164)
		p.Match(TsvsheetParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISignedIntContext is an interface to support dynamic dispatch.
type ISignedIntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	DASH() antlr.TerminalNode

	// IsSignedIntContext differentiates from other interfaces.
	IsSignedIntContext()
}

type SignedIntContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignedIntContext() *SignedIntContext {
	var p = new(SignedIntContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_signedInt
	return p
}

func InitEmptySignedIntContext(p *SignedIntContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_signedInt
}

func (*SignedIntContext) IsSignedIntContext() {}

func NewSignedIntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignedIntContext {
	var p = new(SignedIntContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_signedInt

	return p
}

func (s *SignedIntContext) GetParser() antlr.Parser { return s.parser }

func (s *SignedIntContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserNUMBER, 0)
}

func (s *SignedIntContext) DASH() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserDASH, 0)
}

func (s *SignedIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignedIntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignedIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterSignedInt(s)
	}
}

func (s *SignedIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitSignedInt(s)
	}
}

func (s *SignedIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitSignedInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) SignedInt() (localctx ISignedIntContext) {
	localctx = NewSignedIntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TsvsheetParserRULE_signedInt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TsvsheetParserDASH {
		{
			p.SetState(166)
			p.Match(TsvsheetParserDASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(169)
		p.Match(TsvsheetParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumRowContext is an interface to support dynamic dispatch.
type INumRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SignedInt() ISignedIntContext
	DOLLAR() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	DASH() antlr.TerminalNode

	// IsNumRowContext differentiates from other interfaces.
	IsNumRowContext()
}

type NumRowContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumRowContext() *NumRowContext {
	var p = new(NumRowContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_numRow
	return p
}

func InitEmptyNumRowContext(p *NumRowContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_numRow
}

func (*NumRowContext) IsNumRowContext() {}

func NewNumRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumRowContext {
	var p = new(NumRowContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_numRow

	return p
}

func (s *NumRowContext) GetParser() antlr.Parser { return s.parser }

func (s *NumRowContext) SignedInt() ISignedIntContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignedIntContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignedIntContext)
}

func (s *NumRowContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserDOLLAR, 0)
}

func (s *NumRowContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserNUMBER, 0)
}

func (s *NumRowContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserPLUS, 0)
}

func (s *NumRowContext) DASH() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserDASH, 0)
}

func (s *NumRowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumRowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumRowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterNumRow(s)
	}
}

func (s *NumRowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitNumRow(s)
	}
}

func (s *NumRowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitNumRow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) NumRow() (localctx INumRowContext) {
	localctx = NewNumRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TsvsheetParserRULE_numRow)
	var _la int

	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TsvsheetParserDASH, TsvsheetParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.SignedInt()
		}

	case TsvsheetParserDOLLAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Match(TsvsheetParserDOLLAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		switch p.GetTokenStream().LA(1) {
		case TsvsheetParserPLUS, TsvsheetParserDASH:
			{
				p.SetState(173)
				_la = p.GetTokenStream().LA(1)

				if !(_la == TsvsheetParserPLUS || _la == TsvsheetParserDASH) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(174)
				p.Match(TsvsheetParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case TsvsheetParserNUMBER:
			{
				p.SetState(175)
				p.Match(TsvsheetParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case TsvsheetParserRBRACK:

		default:
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupedRangeContext is an interface to support dynamic dispatch.
type IGroupedRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllColRef() []IColRefContext
	ColRef(i int) IColRefContext
	COLON() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	RowRef() IRowRefContext
	AllNumericRef() []INumericRefContext
	NumericRef(i int) INumericRefContext

	// IsGroupedRangeContext differentiates from other interfaces.
	IsGroupedRangeContext()
}

type GroupedRangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupedRangeContext() *GroupedRangeContext {
	var p = new(GroupedRangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_groupedRange
	return p
}

func InitEmptyGroupedRangeContext(p *GroupedRangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_groupedRange
}

func (*GroupedRangeContext) IsGroupedRangeContext() {}

func NewGroupedRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupedRangeContext {
	var p = new(GroupedRangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_groupedRange

	return p
}

func (s *GroupedRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupedRangeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserLPAREN, 0)
}

func (s *GroupedRangeContext) AllColRef() []IColRefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColRefContext); ok {
			len++
		}
	}

	tst := make([]IColRefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColRefContext); ok {
			tst[i] = t.(IColRefContext)
			i++
		}
	}

	return tst
}

func (s *GroupedRangeContext) ColRef(i int) IColRefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColRefContext); ok {
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

	return t.(IColRefContext)
}

func (s *GroupedRangeContext) COLON() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserCOLON, 0)
}

func (s *GroupedRangeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserRPAREN, 0)
}

func (s *GroupedRangeContext) RowRef() IRowRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRowRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRowRefContext)
}

func (s *GroupedRangeContext) AllNumericRef() []INumericRefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumericRefContext); ok {
			len++
		}
	}

	tst := make([]INumericRefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumericRefContext); ok {
			tst[i] = t.(INumericRefContext)
			i++
		}
	}

	return tst
}

func (s *GroupedRangeContext) NumericRef(i int) INumericRefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericRefContext); ok {
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

	return t.(INumericRefContext)
}

func (s *GroupedRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupedRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupedRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterGroupedRange(s)
	}
}

func (s *GroupedRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitGroupedRange(s)
	}
}

func (s *GroupedRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitGroupedRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) GroupedRange() (localctx IGroupedRangeContext) {
	localctx = NewGroupedRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TsvsheetParserRULE_groupedRange)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Match(TsvsheetParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.ColRef()
		}
		{
			p.SetState(182)
			p.Match(TsvsheetParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.ColRef()
		}
		{
			p.SetState(184)
			p.Match(TsvsheetParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(185)
				p.RowRef()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Match(TsvsheetParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.NumericRef()
		}
		{
			p.SetState(190)
			p.Match(TsvsheetParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.NumericRef()
		}
		{
			p.SetState(192)
			p.Match(TsvsheetParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(193)
				p.RowRef()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModifierContext is an interface to support dynamic dispatch.
type IModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GT() antlr.TerminalNode
	LT() antlr.TerminalNode
	BANG() antlr.TerminalNode

	// IsModifierContext differentiates from other interfaces.
	IsModifierContext()
}

type ModifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifierContext() *ModifierContext {
	var p = new(ModifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_modifier
	return p
}

func InitEmptyModifierContext(p *ModifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_modifier
}

func (*ModifierContext) IsModifierContext() {}

func NewModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifierContext {
	var p = new(ModifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_modifier

	return p
}

func (s *ModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ModifierContext) GT() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserGT, 0)
}

func (s *ModifierContext) LT() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserLT, 0)
}

func (s *ModifierContext) BANG() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserBANG, 0)
}

func (s *ModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterModifier(s)
	}
}

func (s *ModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitModifier(s)
	}
}

func (s *ModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitModifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) Modifier() (localctx IModifierContext) {
	localctx = NewModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TsvsheetParserRULE_modifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4194352) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringExprContext struct {
	ExpressionContext
}

func NewStringExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExprContext {
	var p = new(StringExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserSTRING, 0)
}

func (s *StringExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterStringExpr(s)
	}
}

func (s *StringExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitStringExpr(s)
	}
}

func (s *StringExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitStringExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExprContext struct {
	ExpressionContext
	op antlr.Token
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExprContext) GetOp() antlr.Token { return s.op }

func (s *UnaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserPLUS, 0)
}

func (s *UnaryExprContext) DASH() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserDASH, 0)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExprContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddExprContext) GetOp() antlr.Token { return s.op }

func (s *AddExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *AddExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserPLUS, 0)
}

func (s *AddExprContext) DASH() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserDASH, 0)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RefExprContext struct {
	ExpressionContext
}

func NewRefExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RefExprContext {
	var p = new(RefExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RefExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RefExprContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *RefExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterRefExpr(s)
	}
}

func (s *RefExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitRefExpr(s)
	}
}

func (s *RefExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitRefExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberExprContext struct {
	ExpressionContext
}

func NewNumberExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberExprContext {
	var p = new(NumberExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NumberExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserNUMBER, 0)
}

func (s *NumberExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterNumberExpr(s)
	}
}

func (s *NumberExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitNumberExpr(s)
	}
}

func (s *NumberExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitNumberExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulExprContext struct {
	ExpressionContext
	op antlr.Token
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulExprContext) GetOp() antlr.Token { return s.op }

func (s *MulExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *MulExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserSTAR, 0)
}

func (s *MulExprContext) SLASH() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserSLASH, 0)
}

func (s *MulExprContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserPERCENT, 0)
}

func (s *MulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterMulExpr(s)
	}
}

func (s *MulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitMulExpr(s)
	}
}

func (s *MulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExprContext struct {
	ExpressionContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	ExpressionContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserLPAREN, 0)
}

func (s *ParenExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareExprContext struct {
	ExpressionContext
	op antlr.Token
}

func NewCompareExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExprContext {
	var p = new(CompareExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CompareExprContext) GetOp() antlr.Token { return s.op }

func (s *CompareExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CompareExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *CompareExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserEQ, 0)
}

func (s *CompareExprContext) NE() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserNE, 0)
}

func (s *CompareExprContext) LT() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserLT, 0)
}

func (s *CompareExprContext) LE() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserLE, 0)
}

func (s *CompareExprContext) GT() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserGT, 0)
}

func (s *CompareExprContext) GE() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserGE, 0)
}

func (s *CompareExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterCompareExpr(s)
	}
}

func (s *CompareExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitCompareExpr(s)
	}
}

func (s *CompareExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitCompareExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TsvsheetParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, TsvsheetParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(201)
			p.Match(TsvsheetParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.expression(0)
		}
		{
			p.SetState(203)
			p.Match(TsvsheetParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(205)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == TsvsheetParserPLUS || _la == TsvsheetParserDASH) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(206)
			p.expression(8)
		}

	case 3:
		localctx = NewCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(207)
			p.FunctionCall()
		}

	case 4:
		localctx = NewRefExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(208)
			p.Reference()
		}

	case 5:
		localctx = NewNumberExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(209)
			p.Match(TsvsheetParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStringExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(210)
			p.Match(TsvsheetParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TsvsheetParserRULE_expression)
				p.SetState(213)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(214)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3276800) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(215)
					p.expression(8)
				}

			case 2:
				localctx = NewAddExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TsvsheetParserRULE_expression)
				p.SetState(216)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(217)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TsvsheetParserPLUS || _la == TsvsheetParserDASH) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(218)
					p.expression(7)
				}

			case 3:
				localctx = NewCompareExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TsvsheetParserRULE_expression)
				p.SetState(219)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(220)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompareExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&574) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompareExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(221)
					p.expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	NAME() antlr.TerminalNode
	COL() antlr.TerminalNode
	ArgList() IArgListContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserRPAREN, 0)
}

func (s *FunctionCallContext) NAME() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserNAME, 0)
}

func (s *FunctionCallContext) COL() antlr.TerminalNode {
	return s.GetToken(TsvsheetParserCOL, 0)
}

func (s *FunctionCallContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TsvsheetParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TsvsheetParserCOL || _la == TsvsheetParserNAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(228)
		p.Match(TsvsheetParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126817280) != 0 {
		{
			p.SetState(229)
			p.ArgList()
		}

	}
	{
		p.SetState(232)
		p.Match(TsvsheetParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_argList
	return p
}

func InitEmptyArgListContext(p *ArgListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TsvsheetParserRULE_argList
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TsvsheetParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TsvsheetParserCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TsvsheetParserCOMMA, i)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TsvsheetParserListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TsvsheetParserVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TsvsheetParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TsvsheetParserRULE_argList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.expression(0)
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TsvsheetParserCOMMA {
		{
			p.SetState(235)
			p.Match(TsvsheetParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.expression(0)
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *TsvsheetParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 21:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TsvsheetParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
