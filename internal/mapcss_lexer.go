// Code generated from MapCSS.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 112,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 67, 10, 12, 12, 12,
	14, 12, 70, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7,
	13, 79, 10, 13, 12, 13, 14, 13, 82, 11, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 5, 15, 111, 10, 15, 2, 2, 16, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 2, 17, 2, 19, 2, 21, 2, 23, 9, 25, 10, 27, 11, 29, 12, 3,
	2, 5, 5, 2, 11, 12, 14, 15, 34, 34, 6, 2, 34, 35, 37, 93, 95, 128, 178,
	178, 6, 2, 34, 40, 42, 93, 95, 128, 178, 178, 2, 117, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2,
	2, 2, 29, 3, 2, 2, 2, 3, 31, 3, 2, 2, 2, 5, 38, 3, 2, 2, 2, 7, 42, 3, 2,
	2, 2, 9, 44, 3, 2, 2, 2, 11, 46, 3, 2, 2, 2, 13, 48, 3, 2, 2, 2, 15, 50,
	3, 2, 2, 2, 17, 53, 3, 2, 2, 2, 19, 55, 3, 2, 2, 2, 21, 58, 3, 2, 2, 2,
	23, 61, 3, 2, 2, 2, 25, 73, 3, 2, 2, 2, 27, 85, 3, 2, 2, 2, 29, 110, 3,
	2, 2, 2, 31, 32, 7, 101, 2, 2, 32, 33, 7, 99, 2, 2, 33, 34, 7, 112, 2,
	2, 34, 35, 7, 120, 2, 2, 35, 36, 7, 99, 2, 2, 36, 37, 7, 117, 2, 2, 37,
	4, 3, 2, 2, 2, 38, 39, 9, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 41, 8, 3, 2,
	2, 41, 6, 3, 2, 2, 2, 42, 43, 7, 125, 2, 2, 43, 8, 3, 2, 2, 2, 44, 45,
	7, 127, 2, 2, 45, 10, 3, 2, 2, 2, 46, 47, 7, 60, 2, 2, 47, 12, 3, 2, 2,
	2, 48, 49, 7, 61, 2, 2, 49, 14, 3, 2, 2, 2, 50, 51, 7, 94, 2, 2, 51, 52,
	7, 94, 2, 2, 52, 16, 3, 2, 2, 2, 53, 54, 4, 130, 0, 2, 54, 18, 3, 2, 2,
	2, 55, 56, 7, 94, 2, 2, 56, 57, 7, 36, 2, 2, 57, 20, 3, 2, 2, 2, 58, 59,
	7, 94, 2, 2, 59, 60, 7, 41, 2, 2, 60, 22, 3, 2, 2, 2, 61, 68, 7, 36, 2,
	2, 62, 67, 9, 3, 2, 2, 63, 67, 5, 17, 9, 2, 64, 67, 5, 19, 10, 2, 65, 67,
	5, 15, 8, 2, 66, 62, 3, 2, 2, 2, 66, 63, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2,
	66, 65, 3, 2, 2, 2, 67, 70, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3,
	2, 2, 2, 69, 71, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 72, 7, 36, 2, 2, 72,
	24, 3, 2, 2, 2, 73, 80, 7, 41, 2, 2, 74, 79, 9, 4, 2, 2, 75, 79, 5, 17,
	9, 2, 76, 79, 5, 21, 11, 2, 77, 79, 5, 15, 8, 2, 78, 74, 3, 2, 2, 2, 78,
	75, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 77, 3, 2, 2, 2, 79, 82, 3, 2, 2,
	2, 80, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 83, 3, 2, 2, 2, 82, 80,
	3, 2, 2, 2, 83, 84, 7, 41, 2, 2, 84, 26, 3, 2, 2, 2, 85, 86, 7, 99, 2,
	2, 86, 87, 7, 112, 2, 2, 87, 88, 7, 118, 2, 2, 88, 89, 7, 107, 2, 2, 89,
	90, 7, 99, 2, 2, 90, 91, 7, 110, 2, 2, 91, 92, 7, 107, 2, 2, 92, 93, 7,
	99, 2, 2, 93, 94, 7, 117, 2, 2, 94, 95, 7, 107, 2, 2, 95, 96, 7, 112, 2,
	2, 96, 97, 7, 105, 2, 2, 97, 28, 3, 2, 2, 2, 98, 99, 7, 104, 2, 2, 99,
	100, 7, 119, 2, 2, 100, 101, 7, 110, 2, 2, 101, 111, 7, 110, 2, 2, 102,
	103, 7, 118, 2, 2, 103, 104, 7, 103, 2, 2, 104, 105, 7, 122, 2, 2, 105,
	111, 7, 118, 2, 2, 106, 107, 7, 112, 2, 2, 107, 108, 7, 113, 2, 2, 108,
	109, 7, 112, 2, 2, 109, 111, 7, 103, 2, 2, 110, 98, 3, 2, 2, 2, 110, 102,
	3, 2, 2, 2, 110, 106, 3, 2, 2, 2, 111, 30, 3, 2, 2, 2, 8, 2, 66, 68, 78,
	80, 110, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'canvas'", "", "'{'", "'}'", "':'", "';'", "", "", "'antialiasing'",
}

var lexerSymbolicNames = []string{
	"", "", "WS", "LBRACE", "RBRACE", "COLON", "SEMICOLON", "DQUOTED_STRING",
	"SQUOTED_STRING", "PROP_ANTIALIASING", "PROP_ANTIALIASING_VALUES",
}

var lexerRuleNames = []string{
	"T__0", "WS", "LBRACE", "RBRACE", "COLON", "SEMICOLON", "EBACKSLASH", "UNICODE",
	"EDQUOTE", "ESQUOTE", "DQUOTED_STRING", "SQUOTED_STRING", "PROP_ANTIALIASING",
	"PROP_ANTIALIASING_VALUES",
}

type MapCSSLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewMapCSSLexer(input antlr.CharStream) *MapCSSLexer {

	l := new(MapCSSLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "MapCSS.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MapCSSLexer tokens.
const (
	MapCSSLexerT__0                     = 1
	MapCSSLexerWS                       = 2
	MapCSSLexerLBRACE                   = 3
	MapCSSLexerRBRACE                   = 4
	MapCSSLexerCOLON                    = 5
	MapCSSLexerSEMICOLON                = 6
	MapCSSLexerDQUOTED_STRING           = 7
	MapCSSLexerSQUOTED_STRING           = 8
	MapCSSLexerPROP_ANTIALIASING        = 9
	MapCSSLexerPROP_ANTIALIASING_VALUES = 10
)