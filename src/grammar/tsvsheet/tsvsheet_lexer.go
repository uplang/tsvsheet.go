// Code generated from TsvsheetLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tsvsheetgrammar
import (
	"fmt"
  	"sync"
	"unicode"
	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type TsvsheetLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var TsvsheetLexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  ChannelNames           []string
  ModeNames              []string
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func tsvsheetlexerLexerInit() {
  staticData := &TsvsheetLexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE",
  }
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
    "GE", "LE", "NE", "GT", "LT", "HEADER", "BODY", "FINAL", "EQ", "LPAREN", 
    "RPAREN", "LBRACK", "RBRACK", "COLON", "COMMA", "DOLLAR", "STAR", "PLUS", 
    "DASH", "SLASH", "PERCENT", "BANG", "NUMBER", "COL", "NAME", "STRING", 
    "TAB", "NL", "COMMENT", "WS",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 30, 177, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 
	20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 
	2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1, 
	0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 
	5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 
	7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 
	1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 
	16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 
	1, 22, 4, 22, 122, 8, 22, 11, 22, 12, 22, 123, 1, 22, 1, 22, 4, 22, 128, 
	8, 22, 11, 22, 12, 22, 129, 3, 22, 132, 8, 22, 1, 23, 4, 23, 135, 8, 23, 
	11, 23, 12, 23, 136, 1, 24, 4, 24, 140, 8, 24, 11, 24, 12, 24, 141, 1, 
	25, 1, 25, 5, 25, 146, 8, 25, 10, 25, 12, 25, 149, 9, 25, 1, 25, 1, 25, 
	1, 26, 1, 26, 1, 27, 3, 27, 156, 8, 27, 1, 27, 1, 27, 3, 27, 160, 8, 27, 
	1, 28, 1, 28, 5, 28, 164, 8, 28, 10, 28, 12, 28, 167, 9, 28, 1, 28, 1, 
	28, 1, 29, 4, 29, 172, 8, 29, 11, 29, 12, 29, 173, 1, 29, 1, 29, 0, 0, 
	30, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 
	11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 
	20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 
	29, 59, 30, 1, 0, 5, 1, 0, 48, 57, 1, 0, 65, 90, 2, 0, 65, 90, 97, 122, 
	3, 0, 10, 10, 13, 13, 34, 34, 2, 0, 10, 10, 13, 13, 186, 0, 1, 1, 0, 0, 
	0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 
	0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 
	0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 
	0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 
	1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 
	41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 
	0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 
	0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 1, 61, 1, 0, 0, 0, 3, 64, 1, 0, 
	0, 0, 5, 67, 1, 0, 0, 0, 7, 70, 1, 0, 0, 0, 9, 72, 1, 0, 0, 0, 11, 74, 
	1, 0, 0, 0, 13, 81, 1, 0, 0, 0, 15, 86, 1, 0, 0, 0, 17, 92, 1, 0, 0, 0, 
	19, 94, 1, 0, 0, 0, 21, 96, 1, 0, 0, 0, 23, 98, 1, 0, 0, 0, 25, 100, 1, 
	0, 0, 0, 27, 102, 1, 0, 0, 0, 29, 104, 1, 0, 0, 0, 31, 106, 1, 0, 0, 0, 
	33, 108, 1, 0, 0, 0, 35, 110, 1, 0, 0, 0, 37, 112, 1, 0, 0, 0, 39, 114, 
	1, 0, 0, 0, 41, 116, 1, 0, 0, 0, 43, 118, 1, 0, 0, 0, 45, 121, 1, 0, 0, 
	0, 47, 134, 1, 0, 0, 0, 49, 139, 1, 0, 0, 0, 51, 143, 1, 0, 0, 0, 53, 152, 
	1, 0, 0, 0, 55, 159, 1, 0, 0, 0, 57, 161, 1, 0, 0, 0, 59, 171, 1, 0, 0, 
	0, 61, 62, 5, 62, 0, 0, 62, 63, 5, 61, 0, 0, 63, 2, 1, 0, 0, 0, 64, 65, 
	5, 60, 0, 0, 65, 66, 5, 61, 0, 0, 66, 4, 1, 0, 0, 0, 67, 68, 5, 60, 0, 
	0, 68, 69, 5, 62, 0, 0, 69, 6, 1, 0, 0, 0, 70, 71, 5, 62, 0, 0, 71, 8, 
	1, 0, 0, 0, 72, 73, 5, 60, 0, 0, 73, 10, 1, 0, 0, 0, 74, 75, 5, 104, 0, 
	0, 75, 76, 5, 101, 0, 0, 76, 77, 5, 97, 0, 0, 77, 78, 5, 100, 0, 0, 78, 
	79, 5, 101, 0, 0, 79, 80, 5, 114, 0, 0, 80, 12, 1, 0, 0, 0, 81, 82, 5, 
	98, 0, 0, 82, 83, 5, 111, 0, 0, 83, 84, 5, 100, 0, 0, 84, 85, 5, 121, 0, 
	0, 85, 14, 1, 0, 0, 0, 86, 87, 5, 102, 0, 0, 87, 88, 5, 105, 0, 0, 88, 
	89, 5, 110, 0, 0, 89, 90, 5, 97, 0, 0, 90, 91, 5, 108, 0, 0, 91, 16, 1, 
	0, 0, 0, 92, 93, 5, 61, 0, 0, 93, 18, 1, 0, 0, 0, 94, 95, 5, 40, 0, 0, 
	95, 20, 1, 0, 0, 0, 96, 97, 5, 41, 0, 0, 97, 22, 1, 0, 0, 0, 98, 99, 5, 
	91, 0, 0, 99, 24, 1, 0, 0, 0, 100, 101, 5, 93, 0, 0, 101, 26, 1, 0, 0, 
	0, 102, 103, 5, 58, 0, 0, 103, 28, 1, 0, 0, 0, 104, 105, 5, 44, 0, 0, 105, 
	30, 1, 0, 0, 0, 106, 107, 5, 36, 0, 0, 107, 32, 1, 0, 0, 0, 108, 109, 5, 
	42, 0, 0, 109, 34, 1, 0, 0, 0, 110, 111, 5, 43, 0, 0, 111, 36, 1, 0, 0, 
	0, 112, 113, 5, 45, 0, 0, 113, 38, 1, 0, 0, 0, 114, 115, 5, 47, 0, 0, 115, 
	40, 1, 0, 0, 0, 116, 117, 5, 37, 0, 0, 117, 42, 1, 0, 0, 0, 118, 119, 5, 
	33, 0, 0, 119, 44, 1, 0, 0, 0, 120, 122, 7, 0, 0, 0, 121, 120, 1, 0, 0, 
	0, 122, 123, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 
	131, 1, 0, 0, 0, 125, 127, 5, 46, 0, 0, 126, 128, 7, 0, 0, 0, 127, 126, 
	1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 
	0, 0, 130, 132, 1, 0, 0, 0, 131, 125, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 
	132, 46, 1, 0, 0, 0, 133, 135, 7, 1, 0, 0, 134, 133, 1, 0, 0, 0, 135, 136, 
	1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 48, 1, 0, 
	0, 0, 138, 140, 7, 2, 0, 0, 139, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 
	141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 50, 1, 0, 0, 0, 143, 147, 
	5, 34, 0, 0, 144, 146, 8, 3, 0, 0, 145, 144, 1, 0, 0, 0, 146, 149, 1, 0, 
	0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 
	149, 147, 1, 0, 0, 0, 150, 151, 5, 34, 0, 0, 151, 52, 1, 0, 0, 0, 152, 
	153, 5, 9, 0, 0, 153, 54, 1, 0, 0, 0, 154, 156, 5, 13, 0, 0, 155, 154, 
	1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 160, 5, 10, 
	0, 0, 158, 160, 5, 13, 0, 0, 159, 155, 1, 0, 0, 0, 159, 158, 1, 0, 0, 0, 
	160, 56, 1, 0, 0, 0, 161, 165, 5, 35, 0, 0, 162, 164, 8, 4, 0, 0, 163, 
	162, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 
	1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 168, 169, 6, 28, 
	0, 0, 169, 58, 1, 0, 0, 0, 170, 172, 5, 32, 0, 0, 171, 170, 1, 0, 0, 0, 
	172, 173, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 
	175, 1, 0, 0, 0, 175, 176, 6, 29, 0, 0, 176, 60, 1, 0, 0, 0, 11, 0, 123, 
	129, 131, 136, 141, 147, 155, 159, 165, 173, 1, 6, 0, 0,
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

// TsvsheetLexerInit initializes any static state used to implement TsvsheetLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTsvsheetLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TsvsheetLexerInit() {
  staticData := &TsvsheetLexerLexerStaticData
  staticData.once.Do(tsvsheetlexerLexerInit)
}

// NewTsvsheetLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTsvsheetLexer(input antlr.CharStream) *TsvsheetLexer {
  TsvsheetLexerInit()
	l := new(TsvsheetLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &TsvsheetLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "TsvsheetLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TsvsheetLexer tokens.
const (
	TsvsheetLexerGE = 1
	TsvsheetLexerLE = 2
	TsvsheetLexerNE = 3
	TsvsheetLexerGT = 4
	TsvsheetLexerLT = 5
	TsvsheetLexerHEADER = 6
	TsvsheetLexerBODY = 7
	TsvsheetLexerFINAL = 8
	TsvsheetLexerEQ = 9
	TsvsheetLexerLPAREN = 10
	TsvsheetLexerRPAREN = 11
	TsvsheetLexerLBRACK = 12
	TsvsheetLexerRBRACK = 13
	TsvsheetLexerCOLON = 14
	TsvsheetLexerCOMMA = 15
	TsvsheetLexerDOLLAR = 16
	TsvsheetLexerSTAR = 17
	TsvsheetLexerPLUS = 18
	TsvsheetLexerDASH = 19
	TsvsheetLexerSLASH = 20
	TsvsheetLexerPERCENT = 21
	TsvsheetLexerBANG = 22
	TsvsheetLexerNUMBER = 23
	TsvsheetLexerCOL = 24
	TsvsheetLexerNAME = 25
	TsvsheetLexerSTRING = 26
	TsvsheetLexerTAB = 27
	TsvsheetLexerNL = 28
	TsvsheetLexerCOMMENT = 29
	TsvsheetLexerWS = 30
)

