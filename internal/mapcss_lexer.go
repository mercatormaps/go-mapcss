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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 165,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 77, 10, 12, 12, 12, 14,
	12, 80, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13,
	89, 10, 13, 12, 13, 14, 13, 92, 11, 13, 3, 13, 3, 13, 3, 14, 6, 14, 97,
	10, 14, 13, 14, 14, 14, 98, 3, 15, 3, 15, 3, 15, 3, 16, 6, 16, 105, 10,
	16, 13, 16, 14, 16, 106, 3, 16, 7, 16, 110, 10, 16, 12, 16, 14, 16, 113,
	11, 16, 3, 16, 3, 16, 6, 16, 117, 10, 16, 13, 16, 14, 16, 118, 5, 16, 121,
	10, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 151,
	10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 2, 2, 21, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 2, 17, 2, 19, 2, 21, 2, 23, 9, 25, 10, 27, 11, 29, 12, 31, 13, 33,
	14, 35, 15, 37, 16, 39, 17, 3, 2, 6, 5, 2, 11, 12, 14, 15, 34, 34, 6, 2,
	34, 35, 37, 93, 95, 128, 178, 178, 6, 2, 34, 40, 42, 93, 95, 128, 178,
	178, 3, 2, 50, 59, 2, 175, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 23,
	3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2,
	31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2,
	2, 39, 3, 2, 2, 2, 3, 41, 3, 2, 2, 2, 5, 48, 3, 2, 2, 2, 7, 52, 3, 2, 2,
	2, 9, 54, 3, 2, 2, 2, 11, 56, 3, 2, 2, 2, 13, 58, 3, 2, 2, 2, 15, 60, 3,
	2, 2, 2, 17, 63, 3, 2, 2, 2, 19, 65, 3, 2, 2, 2, 21, 68, 3, 2, 2, 2, 23,
	71, 3, 2, 2, 2, 25, 83, 3, 2, 2, 2, 27, 96, 3, 2, 2, 2, 29, 100, 3, 2,
	2, 2, 31, 120, 3, 2, 2, 2, 33, 122, 3, 2, 2, 2, 35, 125, 3, 2, 2, 2, 37,
	150, 3, 2, 2, 2, 39, 152, 3, 2, 2, 2, 41, 42, 7, 101, 2, 2, 42, 43, 7,
	99, 2, 2, 43, 44, 7, 112, 2, 2, 44, 45, 7, 120, 2, 2, 45, 46, 7, 99, 2,
	2, 46, 47, 7, 117, 2, 2, 47, 4, 3, 2, 2, 2, 48, 49, 9, 2, 2, 2, 49, 50,
	3, 2, 2, 2, 50, 51, 8, 3, 2, 2, 51, 6, 3, 2, 2, 2, 52, 53, 7, 125, 2, 2,
	53, 8, 3, 2, 2, 2, 54, 55, 7, 127, 2, 2, 55, 10, 3, 2, 2, 2, 56, 57, 7,
	60, 2, 2, 57, 12, 3, 2, 2, 2, 58, 59, 7, 61, 2, 2, 59, 14, 3, 2, 2, 2,
	60, 61, 7, 94, 2, 2, 61, 62, 7, 94, 2, 2, 62, 16, 3, 2, 2, 2, 63, 64, 4,
	130, 0, 2, 64, 18, 3, 2, 2, 2, 65, 66, 7, 94, 2, 2, 66, 67, 7, 36, 2, 2,
	67, 20, 3, 2, 2, 2, 68, 69, 7, 94, 2, 2, 69, 70, 7, 41, 2, 2, 70, 22, 3,
	2, 2, 2, 71, 78, 7, 36, 2, 2, 72, 77, 9, 3, 2, 2, 73, 77, 5, 17, 9, 2,
	74, 77, 5, 19, 10, 2, 75, 77, 5, 15, 8, 2, 76, 72, 3, 2, 2, 2, 76, 73,
	3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 75, 3, 2, 2, 2, 77, 80, 3, 2, 2, 2,
	78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 81, 3, 2, 2, 2, 80, 78, 3,
	2, 2, 2, 81, 82, 7, 36, 2, 2, 82, 24, 3, 2, 2, 2, 83, 90, 7, 41, 2, 2,
	84, 89, 9, 4, 2, 2, 85, 89, 5, 17, 9, 2, 86, 89, 5, 21, 11, 2, 87, 89,
	5, 15, 8, 2, 88, 84, 3, 2, 2, 2, 88, 85, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2,
	88, 87, 3, 2, 2, 2, 89, 92, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3,
	2, 2, 2, 91, 93, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 94, 7, 41, 2, 2, 94,
	26, 3, 2, 2, 2, 95, 97, 9, 5, 2, 2, 96, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2,
	2, 98, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 28, 3, 2, 2, 2, 100, 101,
	7, 47, 2, 2, 101, 102, 5, 27, 14, 2, 102, 30, 3, 2, 2, 2, 103, 105, 9,
	5, 2, 2, 104, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 104, 3, 2, 2,
	2, 106, 107, 3, 2, 2, 2, 107, 121, 3, 2, 2, 2, 108, 110, 9, 5, 2, 2, 109,
	108, 3, 2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112,
	3, 2, 2, 2, 112, 114, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114, 116, 7, 48,
	2, 2, 115, 117, 9, 5, 2, 2, 116, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2,
	118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 121, 3, 2, 2, 2, 120,
	104, 3, 2, 2, 2, 120, 111, 3, 2, 2, 2, 121, 32, 3, 2, 2, 2, 122, 123, 7,
	47, 2, 2, 123, 124, 5, 31, 16, 2, 124, 34, 3, 2, 2, 2, 125, 126, 7, 99,
	2, 2, 126, 127, 7, 112, 2, 2, 127, 128, 7, 118, 2, 2, 128, 129, 7, 107,
	2, 2, 129, 130, 7, 99, 2, 2, 130, 131, 7, 110, 2, 2, 131, 132, 7, 107,
	2, 2, 132, 133, 7, 99, 2, 2, 133, 134, 7, 117, 2, 2, 134, 135, 7, 107,
	2, 2, 135, 136, 7, 112, 2, 2, 136, 137, 7, 105, 2, 2, 137, 36, 3, 2, 2,
	2, 138, 139, 7, 104, 2, 2, 139, 140, 7, 119, 2, 2, 140, 141, 7, 110, 2,
	2, 141, 151, 7, 110, 2, 2, 142, 143, 7, 118, 2, 2, 143, 144, 7, 103, 2,
	2, 144, 145, 7, 122, 2, 2, 145, 151, 7, 118, 2, 2, 146, 147, 7, 112, 2,
	2, 147, 148, 7, 113, 2, 2, 148, 149, 7, 112, 2, 2, 149, 151, 7, 103, 2,
	2, 150, 138, 3, 2, 2, 2, 150, 142, 3, 2, 2, 2, 150, 146, 3, 2, 2, 2, 151,
	38, 3, 2, 2, 2, 152, 153, 7, 104, 2, 2, 153, 154, 7, 107, 2, 2, 154, 155,
	7, 110, 2, 2, 155, 156, 7, 110, 2, 2, 156, 157, 7, 47, 2, 2, 157, 158,
	7, 113, 2, 2, 158, 159, 7, 114, 2, 2, 159, 160, 7, 99, 2, 2, 160, 161,
	7, 101, 2, 2, 161, 162, 7, 107, 2, 2, 162, 163, 7, 118, 2, 2, 163, 164,
	7, 123, 2, 2, 164, 40, 3, 2, 2, 2, 13, 2, 76, 78, 88, 90, 98, 106, 111,
	118, 120, 150, 3, 8, 2, 2,
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
	"", "'canvas'", "", "'{'", "'}'", "':'", "';'", "", "", "", "", "", "",
	"'antialiasing'", "", "'fill-opacity'",
}

var lexerSymbolicNames = []string{
	"", "", "WS", "LBRACE", "RBRACE", "COLON", "SEMICOLON", "DQUOTED_STRING",
	"SQUOTED_STRING", "POSITIVE_INT", "NEGATIVE_INT", "POSITIVE_FLOAT", "NEGATIVE_FLOAT",
	"PROP_ANTIALIASING", "PROP_ANTIALIASING_VALUES", "PROP_FILL_OPACITY",
}

var lexerRuleNames = []string{
	"T__0", "WS", "LBRACE", "RBRACE", "COLON", "SEMICOLON", "EBACKSLASH", "UNICODE",
	"EDQUOTE", "ESQUOTE", "DQUOTED_STRING", "SQUOTED_STRING", "POSITIVE_INT",
	"NEGATIVE_INT", "POSITIVE_FLOAT", "NEGATIVE_FLOAT", "PROP_ANTIALIASING",
	"PROP_ANTIALIASING_VALUES", "PROP_FILL_OPACITY",
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
	MapCSSLexerPOSITIVE_INT             = 9
	MapCSSLexerNEGATIVE_INT             = 10
	MapCSSLexerPOSITIVE_FLOAT           = 11
	MapCSSLexerNEGATIVE_FLOAT           = 12
	MapCSSLexerPROP_ANTIALIASING        = 13
	MapCSSLexerPROP_ANTIALIASING_VALUES = 14
	MapCSSLexerPROP_FILL_OPACITY        = 15
)
