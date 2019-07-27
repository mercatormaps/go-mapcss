// Code generated from MapCSS.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // MapCSS

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 193,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 5, 2, 48, 10, 2, 3, 2, 6, 2, 51, 10, 2, 13, 2, 14, 2, 52, 3, 3, 6,
	3, 56, 10, 3, 13, 3, 14, 3, 57, 3, 4, 5, 4, 61, 10, 4, 3, 4, 6, 4, 64,
	10, 4, 13, 4, 14, 4, 65, 3, 4, 7, 4, 69, 10, 4, 12, 4, 14, 4, 72, 11, 4,
	3, 4, 3, 4, 6, 4, 76, 10, 4, 13, 4, 14, 4, 77, 5, 4, 80, 10, 4, 3, 5, 7,
	5, 83, 10, 5, 12, 5, 14, 5, 86, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 92,
	10, 6, 3, 7, 3, 7, 3, 7, 6, 7, 97, 10, 7, 13, 7, 14, 7, 98, 3, 7, 3, 7,
	3, 8, 3, 8, 5, 8, 105, 10, 8, 3, 8, 6, 8, 108, 10, 8, 13, 8, 14, 8, 109,
	3, 9, 3, 9, 3, 9, 5, 9, 115, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 5, 13, 131,
	10, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15,
	3, 16, 3, 16, 6, 16, 144, 10, 16, 13, 16, 14, 16, 145, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 5, 17, 153, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 21, 5, 21, 173, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 2, 2, 24, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 2, 3, 3, 2, 11, 13, 2,
	191, 2, 47, 3, 2, 2, 2, 4, 55, 3, 2, 2, 2, 6, 79, 3, 2, 2, 2, 8, 84, 3,
	2, 2, 2, 10, 91, 3, 2, 2, 2, 12, 93, 3, 2, 2, 2, 14, 102, 3, 2, 2, 2, 16,
	114, 3, 2, 2, 2, 18, 116, 3, 2, 2, 2, 20, 121, 3, 2, 2, 2, 22, 125, 3,
	2, 2, 2, 24, 128, 3, 2, 2, 2, 26, 135, 3, 2, 2, 2, 28, 138, 3, 2, 2, 2,
	30, 141, 3, 2, 2, 2, 32, 152, 3, 2, 2, 2, 34, 154, 3, 2, 2, 2, 36, 159,
	3, 2, 2, 2, 38, 164, 3, 2, 2, 2, 40, 172, 3, 2, 2, 2, 42, 174, 3, 2, 2,
	2, 44, 182, 3, 2, 2, 2, 46, 48, 7, 24, 2, 2, 47, 46, 3, 2, 2, 2, 47, 48,
	3, 2, 2, 2, 48, 50, 3, 2, 2, 2, 49, 51, 7, 27, 2, 2, 50, 49, 3, 2, 2, 2,
	51, 52, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 3, 3, 2,
	2, 2, 54, 56, 7, 27, 2, 2, 55, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57,
	55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 5, 3, 2, 2, 2, 59, 61, 7, 24, 2,
	2, 60, 59, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 63, 3, 2, 2, 2, 62, 64,
	7, 27, 2, 2, 63, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2,
	65, 66, 3, 2, 2, 2, 66, 80, 3, 2, 2, 2, 67, 69, 7, 27, 2, 2, 68, 67, 3,
	2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71,
	73, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 73, 75, 7, 3, 2, 2, 74, 76, 7, 27,
	2, 2, 75, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78,
	3, 2, 2, 2, 78, 80, 3, 2, 2, 2, 79, 60, 3, 2, 2, 2, 79, 70, 3, 2, 2, 2,
	80, 7, 3, 2, 2, 2, 81, 83, 5, 10, 6, 2, 82, 81, 3, 2, 2, 2, 83, 86, 3,
	2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 87, 3, 2, 2, 2, 86,
	84, 3, 2, 2, 2, 87, 88, 7, 2, 2, 3, 88, 9, 3, 2, 2, 2, 89, 92, 5, 28, 15,
	2, 90, 92, 5, 12, 7, 2, 91, 89, 3, 2, 2, 2, 91, 90, 3, 2, 2, 2, 92, 11,
	3, 2, 2, 2, 93, 96, 5, 14, 8, 2, 94, 95, 7, 4, 2, 2, 95, 97, 5, 14, 8,
	2, 96, 94, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 99,
	3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 101, 5, 26, 14, 2, 101, 13, 3, 2,
	2, 2, 102, 104, 7, 29, 2, 2, 103, 105, 5, 16, 9, 2, 104, 103, 3, 2, 2,
	2, 104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2, 2, 106, 108, 5, 24, 13, 2,
	107, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109,
	110, 3, 2, 2, 2, 110, 15, 3, 2, 2, 2, 111, 115, 5, 18, 10, 2, 112, 115,
	5, 20, 11, 2, 113, 115, 5, 22, 12, 2, 114, 111, 3, 2, 2, 2, 114, 112, 3,
	2, 2, 2, 114, 113, 3, 2, 2, 2, 115, 17, 3, 2, 2, 2, 116, 117, 7, 5, 2,
	2, 117, 118, 5, 4, 3, 2, 118, 119, 7, 24, 2, 2, 119, 120, 5, 4, 3, 2, 120,
	19, 3, 2, 2, 2, 121, 122, 7, 5, 2, 2, 122, 123, 5, 4, 3, 2, 123, 124, 7,
	24, 2, 2, 124, 21, 3, 2, 2, 2, 125, 126, 7, 5, 2, 2, 126, 127, 5, 4, 3,
	2, 127, 23, 3, 2, 2, 2, 128, 130, 7, 6, 2, 2, 129, 131, 7, 7, 2, 2, 130,
	129, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133,
	7, 29, 2, 2, 133, 134, 7, 8, 2, 2, 134, 25, 3, 2, 2, 2, 135, 136, 7, 20,
	2, 2, 136, 137, 7, 21, 2, 2, 137, 27, 3, 2, 2, 2, 138, 139, 7, 9, 2, 2,
	139, 140, 5, 30, 16, 2, 140, 29, 3, 2, 2, 2, 141, 143, 7, 20, 2, 2, 142,
	144, 5, 32, 17, 2, 143, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 143,
	3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 148, 7, 21,
	2, 2, 148, 31, 3, 2, 2, 2, 149, 153, 5, 34, 18, 2, 150, 153, 5, 36, 19,
	2, 151, 153, 5, 38, 20, 2, 152, 149, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2,
	152, 151, 3, 2, 2, 2, 153, 33, 3, 2, 2, 2, 154, 155, 7, 10, 2, 2, 155,
	156, 7, 22, 2, 2, 156, 157, 9, 2, 2, 2, 157, 158, 7, 23, 2, 2, 158, 35,
	3, 2, 2, 2, 159, 160, 7, 14, 2, 2, 160, 161, 7, 22, 2, 2, 161, 162, 5,
	6, 4, 2, 162, 163, 7, 23, 2, 2, 163, 37, 3, 2, 2, 2, 164, 165, 7, 15, 2,
	2, 165, 166, 7, 22, 2, 2, 166, 167, 5, 40, 21, 2, 167, 168, 7, 23, 2, 2,
	168, 39, 3, 2, 2, 2, 169, 173, 7, 28, 2, 2, 170, 173, 5, 42, 22, 2, 171,
	173, 5, 44, 23, 2, 172, 169, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 172, 171,
	3, 2, 2, 2, 173, 41, 3, 2, 2, 2, 174, 175, 7, 16, 2, 2, 175, 176, 5, 4,
	3, 2, 176, 177, 7, 4, 2, 2, 177, 178, 5, 4, 3, 2, 178, 179, 7, 4, 2, 2,
	179, 180, 5, 4, 3, 2, 180, 181, 7, 17, 2, 2, 181, 43, 3, 2, 2, 2, 182,
	183, 7, 18, 2, 2, 183, 184, 5, 4, 3, 2, 184, 185, 7, 4, 2, 2, 185, 186,
	5, 4, 3, 2, 186, 187, 7, 4, 2, 2, 187, 188, 5, 4, 3, 2, 188, 189, 7, 4,
	2, 2, 189, 190, 5, 6, 4, 2, 190, 191, 7, 17, 2, 2, 191, 45, 3, 2, 2, 2,
	20, 47, 52, 57, 60, 65, 70, 77, 79, 84, 91, 98, 104, 109, 114, 130, 145,
	152, 172,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'.'", "','", "'|z'", "'['", "'!'", "']'", "'canvas'", "'antialiasing'",
	"'full'", "'text'", "'none'", "'fill-opacity'", "'fill-color'", "'rgb('",
	"')'", "'rgba('", "", "'{'", "'}'", "':'", "';'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "WS",
	"LBRACE", "RBRACE", "COLON", "SEMICOLON", "HYPHEN", "DQUOTED_STRING", "SQUOTED_STRING",
	"DIGIT", "HEX", "IDENTIFIER",
}

var ruleNames = []string{
	"int_", "uint_", "float_", "stylesheet", "entry", "rule_", "selector",
	"zoom", "zoom_range", "min_zoom", "exact_zoom", "attribute", "decl_block",
	"canvas_rule", "canvas_decl_block", "canvas_decl", "antialiasing_decl",
	"fill_opacity_decl", "fill_color_decl", "color", "rgb_color", "rgba_color",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MapCSSParser struct {
	*antlr.BaseParser
}

func NewMapCSSParser(input antlr.TokenStream) *MapCSSParser {
	this := new(MapCSSParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "MapCSS.g4"

	return this
}

// MapCSSParser tokens.
const (
	MapCSSParserEOF            = antlr.TokenEOF
	MapCSSParserT__0           = 1
	MapCSSParserT__1           = 2
	MapCSSParserT__2           = 3
	MapCSSParserT__3           = 4
	MapCSSParserT__4           = 5
	MapCSSParserT__5           = 6
	MapCSSParserT__6           = 7
	MapCSSParserT__7           = 8
	MapCSSParserT__8           = 9
	MapCSSParserT__9           = 10
	MapCSSParserT__10          = 11
	MapCSSParserT__11          = 12
	MapCSSParserT__12          = 13
	MapCSSParserT__13          = 14
	MapCSSParserT__14          = 15
	MapCSSParserT__15          = 16
	MapCSSParserWS             = 17
	MapCSSParserLBRACE         = 18
	MapCSSParserRBRACE         = 19
	MapCSSParserCOLON          = 20
	MapCSSParserSEMICOLON      = 21
	MapCSSParserHYPHEN         = 22
	MapCSSParserDQUOTED_STRING = 23
	MapCSSParserSQUOTED_STRING = 24
	MapCSSParserDIGIT          = 25
	MapCSSParserHEX            = 26
	MapCSSParserIDENTIFIER     = 27
)

// MapCSSParser rules.
const (
	MapCSSParserRULE_int_              = 0
	MapCSSParserRULE_uint_             = 1
	MapCSSParserRULE_float_            = 2
	MapCSSParserRULE_stylesheet        = 3
	MapCSSParserRULE_entry             = 4
	MapCSSParserRULE_rule_             = 5
	MapCSSParserRULE_selector          = 6
	MapCSSParserRULE_zoom              = 7
	MapCSSParserRULE_zoom_range        = 8
	MapCSSParserRULE_min_zoom          = 9
	MapCSSParserRULE_exact_zoom        = 10
	MapCSSParserRULE_attribute         = 11
	MapCSSParserRULE_decl_block        = 12
	MapCSSParserRULE_canvas_rule       = 13
	MapCSSParserRULE_canvas_decl_block = 14
	MapCSSParserRULE_canvas_decl       = 15
	MapCSSParserRULE_antialiasing_decl = 16
	MapCSSParserRULE_fill_opacity_decl = 17
	MapCSSParserRULE_fill_color_decl   = 18
	MapCSSParserRULE_color             = 19
	MapCSSParserRULE_rgb_color         = 20
	MapCSSParserRULE_rgba_color        = 21
)

// IInt_Context is an interface to support dynamic dispatch.
type IInt_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInt_Context differentiates from other interfaces.
	IsInt_Context()
}

type Int_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInt_Context() *Int_Context {
	var p = new(Int_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_int_
	return p
}

func (*Int_Context) IsInt_Context() {}

func NewInt_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Int_Context {
	var p = new(Int_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_int_

	return p
}

func (s *Int_Context) GetParser() antlr.Parser { return s.parser }

func (s *Int_Context) HYPHEN() antlr.TerminalNode {
	return s.GetToken(MapCSSParserHYPHEN, 0)
}

func (s *Int_Context) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(MapCSSParserDIGIT)
}

func (s *Int_Context) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(MapCSSParserDIGIT, i)
}

func (s *Int_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Int_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Int_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterInt_(s)
	}
}

func (s *Int_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitInt_(s)
	}
}

func (p *MapCSSParser) Int_() (localctx IInt_Context) {
	localctx = NewInt_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MapCSSParserRULE_int_)
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MapCSSParserHYPHEN {
		{
			p.SetState(44)
			p.Match(MapCSSParserHYPHEN)
		}

	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MapCSSParserDIGIT {
		{
			p.SetState(47)
			p.Match(MapCSSParserDIGIT)
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUint_Context is an interface to support dynamic dispatch.
type IUint_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUint_Context differentiates from other interfaces.
	IsUint_Context()
}

type Uint_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUint_Context() *Uint_Context {
	var p = new(Uint_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_uint_
	return p
}

func (*Uint_Context) IsUint_Context() {}

func NewUint_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Uint_Context {
	var p = new(Uint_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_uint_

	return p
}

func (s *Uint_Context) GetParser() antlr.Parser { return s.parser }

func (s *Uint_Context) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(MapCSSParserDIGIT)
}

func (s *Uint_Context) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(MapCSSParserDIGIT, i)
}

func (s *Uint_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Uint_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Uint_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterUint_(s)
	}
}

func (s *Uint_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitUint_(s)
	}
}

func (p *MapCSSParser) Uint_() (localctx IUint_Context) {
	localctx = NewUint_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MapCSSParserRULE_uint_)
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
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MapCSSParserDIGIT {
		{
			p.SetState(52)
			p.Match(MapCSSParserDIGIT)
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFloat_Context is an interface to support dynamic dispatch.
type IFloat_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloat_Context differentiates from other interfaces.
	IsFloat_Context()
}

type Float_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloat_Context() *Float_Context {
	var p = new(Float_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_float_
	return p
}

func (*Float_Context) IsFloat_Context() {}

func NewFloat_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Float_Context {
	var p = new(Float_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_float_

	return p
}

func (s *Float_Context) GetParser() antlr.Parser { return s.parser }

func (s *Float_Context) HYPHEN() antlr.TerminalNode {
	return s.GetToken(MapCSSParserHYPHEN, 0)
}

func (s *Float_Context) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(MapCSSParserDIGIT)
}

func (s *Float_Context) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(MapCSSParserDIGIT, i)
}

func (s *Float_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Float_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Float_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterFloat_(s)
	}
}

func (s *Float_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitFloat_(s)
	}
}

func (p *MapCSSParser) Float_() (localctx IFloat_Context) {
	localctx = NewFloat_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MapCSSParserRULE_float_)
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

	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MapCSSParserHYPHEN {
			{
				p.SetState(57)
				p.Match(MapCSSParserHYPHEN)
			}

		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == MapCSSParserDIGIT {
			{
				p.SetState(60)
				p.Match(MapCSSParserDIGIT)
			}

			p.SetState(63)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MapCSSParserDIGIT {
			{
				p.SetState(65)
				p.Match(MapCSSParserDIGIT)
			}

			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(71)
			p.Match(MapCSSParserT__0)
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == MapCSSParserDIGIT {
			{
				p.SetState(72)
				p.Match(MapCSSParserDIGIT)
			}

			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IStylesheetContext is an interface to support dynamic dispatch.
type IStylesheetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStylesheetContext differentiates from other interfaces.
	IsStylesheetContext()
}

type StylesheetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStylesheetContext() *StylesheetContext {
	var p = new(StylesheetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_stylesheet
	return p
}

func (*StylesheetContext) IsStylesheetContext() {}

func NewStylesheetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StylesheetContext {
	var p = new(StylesheetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_stylesheet

	return p
}

func (s *StylesheetContext) GetParser() antlr.Parser { return s.parser }

func (s *StylesheetContext) EOF() antlr.TerminalNode {
	return s.GetToken(MapCSSParserEOF, 0)
}

func (s *StylesheetContext) AllEntry() []IEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEntryContext)(nil)).Elem())
	var tst = make([]IEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEntryContext)
		}
	}

	return tst
}

func (s *StylesheetContext) Entry(i int) IEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEntryContext)
}

func (s *StylesheetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StylesheetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StylesheetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterStylesheet(s)
	}
}

func (s *StylesheetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitStylesheet(s)
	}
}

func (p *MapCSSParser) Stylesheet() (localctx IStylesheetContext) {
	localctx = NewStylesheetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MapCSSParserRULE_stylesheet)
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
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MapCSSParserT__6 || _la == MapCSSParserIDENTIFIER {
		{
			p.SetState(79)
			p.Entry()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(85)
		p.Match(MapCSSParserEOF)
	}

	return localctx
}

// IEntryContext is an interface to support dynamic dispatch.
type IEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntryContext differentiates from other interfaces.
	IsEntryContext()
}

type EntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntryContext() *EntryContext {
	var p = new(EntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_entry
	return p
}

func (*EntryContext) IsEntryContext() {}

func NewEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryContext {
	var p = new(EntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_entry

	return p
}

func (s *EntryContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryContext) Canvas_rule() ICanvas_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICanvas_ruleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICanvas_ruleContext)
}

func (s *EntryContext) Rule_() IRule_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRule_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRule_Context)
}

func (s *EntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterEntry(s)
	}
}

func (s *EntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitEntry(s)
	}
}

func (p *MapCSSParser) Entry() (localctx IEntryContext) {
	localctx = NewEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MapCSSParserRULE_entry)

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

	p.SetState(89)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MapCSSParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Canvas_rule()
		}

	case MapCSSParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			p.Rule_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRule_Context is an interface to support dynamic dispatch.
type IRule_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRule_Context differentiates from other interfaces.
	IsRule_Context()
}

type Rule_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRule_Context() *Rule_Context {
	var p = new(Rule_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_rule_
	return p
}

func (*Rule_Context) IsRule_Context() {}

func NewRule_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_Context {
	var p = new(Rule_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_rule_

	return p
}

func (s *Rule_Context) GetParser() antlr.Parser { return s.parser }

func (s *Rule_Context) AllSelector() []ISelectorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectorContext)(nil)).Elem())
	var tst = make([]ISelectorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectorContext)
		}
	}

	return tst
}

func (s *Rule_Context) Selector(i int) ISelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *Rule_Context) Decl_block() IDecl_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecl_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecl_blockContext)
}

func (s *Rule_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterRule_(s)
	}
}

func (s *Rule_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitRule_(s)
	}
}

func (p *MapCSSParser) Rule_() (localctx IRule_Context) {
	localctx = NewRule_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MapCSSParserRULE_rule_)
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
		p.SetState(91)
		p.Selector()
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MapCSSParserT__1 {
		{
			p.SetState(92)
			p.Match(MapCSSParserT__1)
		}
		{
			p.SetState(93)
			p.Selector()
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(98)
		p.Decl_block()
	}

	return localctx
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTyp returns the typ token.
	GetTyp() antlr.Token

	// SetTyp sets the typ token.
	SetTyp(antlr.Token)

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	typ    antlr.Token
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_selector
	return p
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) GetTyp() antlr.Token { return s.typ }

func (s *SelectorContext) SetTyp(v antlr.Token) { s.typ = v }

func (s *SelectorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MapCSSParserIDENTIFIER, 0)
}

func (s *SelectorContext) Zoom() IZoomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IZoomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IZoomContext)
}

func (s *SelectorContext) AllAttribute() []IAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeContext)(nil)).Elem())
	var tst = make([]IAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeContext)
		}
	}

	return tst
}

func (s *SelectorContext) Attribute(i int) IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (p *MapCSSParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MapCSSParserRULE_selector)
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
		p.SetState(100)

		var _m = p.Match(MapCSSParserIDENTIFIER)

		localctx.(*SelectorContext).typ = _m
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MapCSSParserT__2 {
		{
			p.SetState(101)
			p.Zoom()
		}

	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MapCSSParserT__3 {
		{
			p.SetState(104)
			p.Attribute()
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IZoomContext is an interface to support dynamic dispatch.
type IZoomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsZoomContext differentiates from other interfaces.
	IsZoomContext()
}

type ZoomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZoomContext() *ZoomContext {
	var p = new(ZoomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_zoom
	return p
}

func (*ZoomContext) IsZoomContext() {}

func NewZoomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZoomContext {
	var p = new(ZoomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_zoom

	return p
}

func (s *ZoomContext) GetParser() antlr.Parser { return s.parser }

func (s *ZoomContext) Zoom_range() IZoom_rangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IZoom_rangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IZoom_rangeContext)
}

func (s *ZoomContext) Min_zoom() IMin_zoomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMin_zoomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMin_zoomContext)
}

func (s *ZoomContext) Exact_zoom() IExact_zoomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExact_zoomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExact_zoomContext)
}

func (s *ZoomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZoomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZoomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterZoom(s)
	}
}

func (s *ZoomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitZoom(s)
	}
}

func (p *MapCSSParser) Zoom() (localctx IZoomContext) {
	localctx = NewZoomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MapCSSParserRULE_zoom)

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

	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Zoom_range()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Min_zoom()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			p.Exact_zoom()
		}

	}

	return localctx
}

// IZoom_rangeContext is an interface to support dynamic dispatch.
type IZoom_rangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMin returns the min rule contexts.
	GetMin() IUint_Context

	// GetMax returns the max rule contexts.
	GetMax() IUint_Context

	// SetMin sets the min rule contexts.
	SetMin(IUint_Context)

	// SetMax sets the max rule contexts.
	SetMax(IUint_Context)

	// IsZoom_rangeContext differentiates from other interfaces.
	IsZoom_rangeContext()
}

type Zoom_rangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	min    IUint_Context
	max    IUint_Context
}

func NewEmptyZoom_rangeContext() *Zoom_rangeContext {
	var p = new(Zoom_rangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_zoom_range
	return p
}

func (*Zoom_rangeContext) IsZoom_rangeContext() {}

func NewZoom_rangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Zoom_rangeContext {
	var p = new(Zoom_rangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_zoom_range

	return p
}

func (s *Zoom_rangeContext) GetParser() antlr.Parser { return s.parser }

func (s *Zoom_rangeContext) GetMin() IUint_Context { return s.min }

func (s *Zoom_rangeContext) GetMax() IUint_Context { return s.max }

func (s *Zoom_rangeContext) SetMin(v IUint_Context) { s.min = v }

func (s *Zoom_rangeContext) SetMax(v IUint_Context) { s.max = v }

func (s *Zoom_rangeContext) HYPHEN() antlr.TerminalNode {
	return s.GetToken(MapCSSParserHYPHEN, 0)
}

func (s *Zoom_rangeContext) AllUint_() []IUint_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUint_Context)(nil)).Elem())
	var tst = make([]IUint_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUint_Context)
		}
	}

	return tst
}

func (s *Zoom_rangeContext) Uint_(i int) IUint_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUint_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUint_Context)
}

func (s *Zoom_rangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Zoom_rangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Zoom_rangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterZoom_range(s)
	}
}

func (s *Zoom_rangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitZoom_range(s)
	}
}

func (p *MapCSSParser) Zoom_range() (localctx IZoom_rangeContext) {
	localctx = NewZoom_rangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MapCSSParserRULE_zoom_range)

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
		p.SetState(114)
		p.Match(MapCSSParserT__2)
	}
	{
		p.SetState(115)

		var _x = p.Uint_()

		localctx.(*Zoom_rangeContext).min = _x
	}
	{
		p.SetState(116)
		p.Match(MapCSSParserHYPHEN)
	}
	{
		p.SetState(117)

		var _x = p.Uint_()

		localctx.(*Zoom_rangeContext).max = _x
	}

	return localctx
}

// IMin_zoomContext is an interface to support dynamic dispatch.
type IMin_zoomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMin returns the min rule contexts.
	GetMin() IUint_Context

	// SetMin sets the min rule contexts.
	SetMin(IUint_Context)

	// IsMin_zoomContext differentiates from other interfaces.
	IsMin_zoomContext()
}

type Min_zoomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	min    IUint_Context
}

func NewEmptyMin_zoomContext() *Min_zoomContext {
	var p = new(Min_zoomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_min_zoom
	return p
}

func (*Min_zoomContext) IsMin_zoomContext() {}

func NewMin_zoomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Min_zoomContext {
	var p = new(Min_zoomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_min_zoom

	return p
}

func (s *Min_zoomContext) GetParser() antlr.Parser { return s.parser }

func (s *Min_zoomContext) GetMin() IUint_Context { return s.min }

func (s *Min_zoomContext) SetMin(v IUint_Context) { s.min = v }

func (s *Min_zoomContext) HYPHEN() antlr.TerminalNode {
	return s.GetToken(MapCSSParserHYPHEN, 0)
}

func (s *Min_zoomContext) Uint_() IUint_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUint_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUint_Context)
}

func (s *Min_zoomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Min_zoomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Min_zoomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterMin_zoom(s)
	}
}

func (s *Min_zoomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitMin_zoom(s)
	}
}

func (p *MapCSSParser) Min_zoom() (localctx IMin_zoomContext) {
	localctx = NewMin_zoomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MapCSSParserRULE_min_zoom)

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
		p.SetState(119)
		p.Match(MapCSSParserT__2)
	}
	{
		p.SetState(120)

		var _x = p.Uint_()

		localctx.(*Min_zoomContext).min = _x
	}
	{
		p.SetState(121)
		p.Match(MapCSSParserHYPHEN)
	}

	return localctx
}

// IExact_zoomContext is an interface to support dynamic dispatch.
type IExact_zoomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMin returns the min rule contexts.
	GetMin() IUint_Context

	// SetMin sets the min rule contexts.
	SetMin(IUint_Context)

	// IsExact_zoomContext differentiates from other interfaces.
	IsExact_zoomContext()
}

type Exact_zoomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	min    IUint_Context
}

func NewEmptyExact_zoomContext() *Exact_zoomContext {
	var p = new(Exact_zoomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_exact_zoom
	return p
}

func (*Exact_zoomContext) IsExact_zoomContext() {}

func NewExact_zoomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exact_zoomContext {
	var p = new(Exact_zoomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_exact_zoom

	return p
}

func (s *Exact_zoomContext) GetParser() antlr.Parser { return s.parser }

func (s *Exact_zoomContext) GetMin() IUint_Context { return s.min }

func (s *Exact_zoomContext) SetMin(v IUint_Context) { s.min = v }

func (s *Exact_zoomContext) Uint_() IUint_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUint_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUint_Context)
}

func (s *Exact_zoomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exact_zoomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exact_zoomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterExact_zoom(s)
	}
}

func (s *Exact_zoomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitExact_zoom(s)
	}
}

func (p *MapCSSParser) Exact_zoom() (localctx IExact_zoomContext) {
	localctx = NewExact_zoomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MapCSSParserRULE_exact_zoom)

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
		p.SetState(123)
		p.Match(MapCSSParserT__2)
	}
	{
		p.SetState(124)

		var _x = p.Uint_()

		localctx.(*Exact_zoomContext).min = _x
	}

	return localctx
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNeg returns the neg token.
	GetNeg() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// SetNeg sets the neg token.
	SetNeg(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	neg    antlr.Token
	name   antlr.Token
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_attribute
	return p
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) GetNeg() antlr.Token { return s.neg }

func (s *AttributeContext) GetName() antlr.Token { return s.name }

func (s *AttributeContext) SetNeg(v antlr.Token) { s.neg = v }

func (s *AttributeContext) SetName(v antlr.Token) { s.name = v }

func (s *AttributeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MapCSSParserIDENTIFIER, 0)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (p *MapCSSParser) Attribute() (localctx IAttributeContext) {
	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MapCSSParserRULE_attribute)
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
		p.SetState(126)
		p.Match(MapCSSParserT__3)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MapCSSParserT__4 {
		{
			p.SetState(127)

			var _m = p.Match(MapCSSParserT__4)

			localctx.(*AttributeContext).neg = _m
		}

	}
	{
		p.SetState(130)

		var _m = p.Match(MapCSSParserIDENTIFIER)

		localctx.(*AttributeContext).name = _m
	}
	{
		p.SetState(131)
		p.Match(MapCSSParserT__5)
	}

	return localctx
}

// IDecl_blockContext is an interface to support dynamic dispatch.
type IDecl_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecl_blockContext differentiates from other interfaces.
	IsDecl_blockContext()
}

type Decl_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecl_blockContext() *Decl_blockContext {
	var p = new(Decl_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_decl_block
	return p
}

func (*Decl_blockContext) IsDecl_blockContext() {}

func NewDecl_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decl_blockContext {
	var p = new(Decl_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_decl_block

	return p
}

func (s *Decl_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Decl_blockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MapCSSParserLBRACE, 0)
}

func (s *Decl_blockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MapCSSParserRBRACE, 0)
}

func (s *Decl_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decl_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Decl_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterDecl_block(s)
	}
}

func (s *Decl_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitDecl_block(s)
	}
}

func (p *MapCSSParser) Decl_block() (localctx IDecl_blockContext) {
	localctx = NewDecl_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MapCSSParserRULE_decl_block)

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
		p.SetState(133)
		p.Match(MapCSSParserLBRACE)
	}
	{
		p.SetState(134)
		p.Match(MapCSSParserRBRACE)
	}

	return localctx
}

// ICanvas_ruleContext is an interface to support dynamic dispatch.
type ICanvas_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCanvas_ruleContext differentiates from other interfaces.
	IsCanvas_ruleContext()
}

type Canvas_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCanvas_ruleContext() *Canvas_ruleContext {
	var p = new(Canvas_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_canvas_rule
	return p
}

func (*Canvas_ruleContext) IsCanvas_ruleContext() {}

func NewCanvas_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Canvas_ruleContext {
	var p = new(Canvas_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_canvas_rule

	return p
}

func (s *Canvas_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Canvas_ruleContext) Canvas_decl_block() ICanvas_decl_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICanvas_decl_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICanvas_decl_blockContext)
}

func (s *Canvas_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Canvas_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Canvas_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterCanvas_rule(s)
	}
}

func (s *Canvas_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitCanvas_rule(s)
	}
}

func (p *MapCSSParser) Canvas_rule() (localctx ICanvas_ruleContext) {
	localctx = NewCanvas_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MapCSSParserRULE_canvas_rule)

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
		p.SetState(136)
		p.Match(MapCSSParserT__6)
	}
	{
		p.SetState(137)
		p.Canvas_decl_block()
	}

	return localctx
}

// ICanvas_decl_blockContext is an interface to support dynamic dispatch.
type ICanvas_decl_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCanvas_decl_blockContext differentiates from other interfaces.
	IsCanvas_decl_blockContext()
}

type Canvas_decl_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCanvas_decl_blockContext() *Canvas_decl_blockContext {
	var p = new(Canvas_decl_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_canvas_decl_block
	return p
}

func (*Canvas_decl_blockContext) IsCanvas_decl_blockContext() {}

func NewCanvas_decl_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Canvas_decl_blockContext {
	var p = new(Canvas_decl_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_canvas_decl_block

	return p
}

func (s *Canvas_decl_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Canvas_decl_blockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MapCSSParserLBRACE, 0)
}

func (s *Canvas_decl_blockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MapCSSParserRBRACE, 0)
}

func (s *Canvas_decl_blockContext) AllCanvas_decl() []ICanvas_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICanvas_declContext)(nil)).Elem())
	var tst = make([]ICanvas_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICanvas_declContext)
		}
	}

	return tst
}

func (s *Canvas_decl_blockContext) Canvas_decl(i int) ICanvas_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICanvas_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICanvas_declContext)
}

func (s *Canvas_decl_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Canvas_decl_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Canvas_decl_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterCanvas_decl_block(s)
	}
}

func (s *Canvas_decl_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitCanvas_decl_block(s)
	}
}

func (p *MapCSSParser) Canvas_decl_block() (localctx ICanvas_decl_blockContext) {
	localctx = NewCanvas_decl_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MapCSSParserRULE_canvas_decl_block)
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
		p.SetState(139)
		p.Match(MapCSSParserLBRACE)
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MapCSSParserT__7)|(1<<MapCSSParserT__11)|(1<<MapCSSParserT__12))) != 0) {
		{
			p.SetState(140)
			p.Canvas_decl()
		}

		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(145)
		p.Match(MapCSSParserRBRACE)
	}

	return localctx
}

// ICanvas_declContext is an interface to support dynamic dispatch.
type ICanvas_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCanvas_declContext differentiates from other interfaces.
	IsCanvas_declContext()
}

type Canvas_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCanvas_declContext() *Canvas_declContext {
	var p = new(Canvas_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_canvas_decl
	return p
}

func (*Canvas_declContext) IsCanvas_declContext() {}

func NewCanvas_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Canvas_declContext {
	var p = new(Canvas_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_canvas_decl

	return p
}

func (s *Canvas_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Canvas_declContext) Antialiasing_decl() IAntialiasing_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAntialiasing_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAntialiasing_declContext)
}

func (s *Canvas_declContext) Fill_opacity_decl() IFill_opacity_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFill_opacity_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFill_opacity_declContext)
}

func (s *Canvas_declContext) Fill_color_decl() IFill_color_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFill_color_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFill_color_declContext)
}

func (s *Canvas_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Canvas_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Canvas_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterCanvas_decl(s)
	}
}

func (s *Canvas_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitCanvas_decl(s)
	}
}

func (p *MapCSSParser) Canvas_decl() (localctx ICanvas_declContext) {
	localctx = NewCanvas_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MapCSSParserRULE_canvas_decl)

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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MapCSSParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.Antialiasing_decl()
		}

	case MapCSSParserT__11:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Fill_opacity_decl()
		}

	case MapCSSParserT__12:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(149)
			p.Fill_color_decl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAntialiasing_declContext is an interface to support dynamic dispatch.
type IAntialiasing_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// IsAntialiasing_declContext differentiates from other interfaces.
	IsAntialiasing_declContext()
}

type Antialiasing_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      antlr.Token
}

func NewEmptyAntialiasing_declContext() *Antialiasing_declContext {
	var p = new(Antialiasing_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_antialiasing_decl
	return p
}

func (*Antialiasing_declContext) IsAntialiasing_declContext() {}

func NewAntialiasing_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Antialiasing_declContext {
	var p = new(Antialiasing_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_antialiasing_decl

	return p
}

func (s *Antialiasing_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Antialiasing_declContext) GetV() antlr.Token { return s.v }

func (s *Antialiasing_declContext) SetV(v antlr.Token) { s.v = v }

func (s *Antialiasing_declContext) COLON() antlr.TerminalNode {
	return s.GetToken(MapCSSParserCOLON, 0)
}

func (s *Antialiasing_declContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MapCSSParserSEMICOLON, 0)
}

func (s *Antialiasing_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Antialiasing_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Antialiasing_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterAntialiasing_decl(s)
	}
}

func (s *Antialiasing_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitAntialiasing_decl(s)
	}
}

func (p *MapCSSParser) Antialiasing_decl() (localctx IAntialiasing_declContext) {
	localctx = NewAntialiasing_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MapCSSParserRULE_antialiasing_decl)
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
		p.SetState(152)
		p.Match(MapCSSParserT__7)
	}
	{
		p.SetState(153)
		p.Match(MapCSSParserCOLON)
	}
	{
		p.SetState(154)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Antialiasing_declContext).v = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MapCSSParserT__8)|(1<<MapCSSParserT__9)|(1<<MapCSSParserT__10))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Antialiasing_declContext).v = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(155)
		p.Match(MapCSSParserSEMICOLON)
	}

	return localctx
}

// IFill_opacity_declContext is an interface to support dynamic dispatch.
type IFill_opacity_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v rule contexts.
	GetV() IFloat_Context

	// SetV sets the v rule contexts.
	SetV(IFloat_Context)

	// IsFill_opacity_declContext differentiates from other interfaces.
	IsFill_opacity_declContext()
}

type Fill_opacity_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      IFloat_Context
}

func NewEmptyFill_opacity_declContext() *Fill_opacity_declContext {
	var p = new(Fill_opacity_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_fill_opacity_decl
	return p
}

func (*Fill_opacity_declContext) IsFill_opacity_declContext() {}

func NewFill_opacity_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fill_opacity_declContext {
	var p = new(Fill_opacity_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_fill_opacity_decl

	return p
}

func (s *Fill_opacity_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Fill_opacity_declContext) GetV() IFloat_Context { return s.v }

func (s *Fill_opacity_declContext) SetV(v IFloat_Context) { s.v = v }

func (s *Fill_opacity_declContext) COLON() antlr.TerminalNode {
	return s.GetToken(MapCSSParserCOLON, 0)
}

func (s *Fill_opacity_declContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MapCSSParserSEMICOLON, 0)
}

func (s *Fill_opacity_declContext) Float_() IFloat_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloat_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloat_Context)
}

func (s *Fill_opacity_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fill_opacity_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fill_opacity_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterFill_opacity_decl(s)
	}
}

func (s *Fill_opacity_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitFill_opacity_decl(s)
	}
}

func (p *MapCSSParser) Fill_opacity_decl() (localctx IFill_opacity_declContext) {
	localctx = NewFill_opacity_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MapCSSParserRULE_fill_opacity_decl)

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
		p.SetState(157)
		p.Match(MapCSSParserT__11)
	}
	{
		p.SetState(158)
		p.Match(MapCSSParserCOLON)
	}
	{
		p.SetState(159)

		var _x = p.Float_()

		localctx.(*Fill_opacity_declContext).v = _x
	}
	{
		p.SetState(160)
		p.Match(MapCSSParserSEMICOLON)
	}

	return localctx
}

// IFill_color_declContext is an interface to support dynamic dispatch.
type IFill_color_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFill_color_declContext differentiates from other interfaces.
	IsFill_color_declContext()
}

type Fill_color_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFill_color_declContext() *Fill_color_declContext {
	var p = new(Fill_color_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_fill_color_decl
	return p
}

func (*Fill_color_declContext) IsFill_color_declContext() {}

func NewFill_color_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fill_color_declContext {
	var p = new(Fill_color_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_fill_color_decl

	return p
}

func (s *Fill_color_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Fill_color_declContext) COLON() antlr.TerminalNode {
	return s.GetToken(MapCSSParserCOLON, 0)
}

func (s *Fill_color_declContext) Color() IColorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColorContext)
}

func (s *Fill_color_declContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MapCSSParserSEMICOLON, 0)
}

func (s *Fill_color_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fill_color_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fill_color_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterFill_color_decl(s)
	}
}

func (s *Fill_color_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitFill_color_decl(s)
	}
}

func (p *MapCSSParser) Fill_color_decl() (localctx IFill_color_declContext) {
	localctx = NewFill_color_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MapCSSParserRULE_fill_color_decl)

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
		p.SetState(162)
		p.Match(MapCSSParserT__12)
	}
	{
		p.SetState(163)
		p.Match(MapCSSParserCOLON)
	}
	{
		p.SetState(164)
		p.Color()
	}
	{
		p.SetState(165)
		p.Match(MapCSSParserSEMICOLON)
	}

	return localctx
}

// IColorContext is an interface to support dynamic dispatch.
type IColorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColorContext differentiates from other interfaces.
	IsColorContext()
}

type ColorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColorContext() *ColorContext {
	var p = new(ColorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_color
	return p
}

func (*ColorContext) IsColorContext() {}

func NewColorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColorContext {
	var p = new(ColorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_color

	return p
}

func (s *ColorContext) GetParser() antlr.Parser { return s.parser }

func (s *ColorContext) HEX() antlr.TerminalNode {
	return s.GetToken(MapCSSParserHEX, 0)
}

func (s *ColorContext) Rgb_color() IRgb_colorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRgb_colorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRgb_colorContext)
}

func (s *ColorContext) Rgba_color() IRgba_colorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRgba_colorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRgba_colorContext)
}

func (s *ColorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterColor(s)
	}
}

func (s *ColorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitColor(s)
	}
}

func (p *MapCSSParser) Color() (localctx IColorContext) {
	localctx = NewColorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MapCSSParserRULE_color)

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

	p.SetState(170)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MapCSSParserHEX:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Match(MapCSSParserHEX)
		}

	case MapCSSParserT__13:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(168)
			p.Rgb_color()
		}

	case MapCSSParserT__15:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(169)
			p.Rgba_color()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRgb_colorContext is an interface to support dynamic dispatch.
type IRgb_colorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetR returns the r rule contexts.
	GetR() IUint_Context

	// GetG returns the g rule contexts.
	GetG() IUint_Context

	// GetB returns the b rule contexts.
	GetB() IUint_Context

	// SetR sets the r rule contexts.
	SetR(IUint_Context)

	// SetG sets the g rule contexts.
	SetG(IUint_Context)

	// SetB sets the b rule contexts.
	SetB(IUint_Context)

	// IsRgb_colorContext differentiates from other interfaces.
	IsRgb_colorContext()
}

type Rgb_colorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	r      IUint_Context
	g      IUint_Context
	b      IUint_Context
}

func NewEmptyRgb_colorContext() *Rgb_colorContext {
	var p = new(Rgb_colorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_rgb_color
	return p
}

func (*Rgb_colorContext) IsRgb_colorContext() {}

func NewRgb_colorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rgb_colorContext {
	var p = new(Rgb_colorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_rgb_color

	return p
}

func (s *Rgb_colorContext) GetParser() antlr.Parser { return s.parser }

func (s *Rgb_colorContext) GetR() IUint_Context { return s.r }

func (s *Rgb_colorContext) GetG() IUint_Context { return s.g }

func (s *Rgb_colorContext) GetB() IUint_Context { return s.b }

func (s *Rgb_colorContext) SetR(v IUint_Context) { s.r = v }

func (s *Rgb_colorContext) SetG(v IUint_Context) { s.g = v }

func (s *Rgb_colorContext) SetB(v IUint_Context) { s.b = v }

func (s *Rgb_colorContext) AllUint_() []IUint_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUint_Context)(nil)).Elem())
	var tst = make([]IUint_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUint_Context)
		}
	}

	return tst
}

func (s *Rgb_colorContext) Uint_(i int) IUint_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUint_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUint_Context)
}

func (s *Rgb_colorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rgb_colorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rgb_colorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterRgb_color(s)
	}
}

func (s *Rgb_colorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitRgb_color(s)
	}
}

func (p *MapCSSParser) Rgb_color() (localctx IRgb_colorContext) {
	localctx = NewRgb_colorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MapCSSParserRULE_rgb_color)

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
		p.SetState(172)
		p.Match(MapCSSParserT__13)
	}
	{
		p.SetState(173)

		var _x = p.Uint_()

		localctx.(*Rgb_colorContext).r = _x
	}
	{
		p.SetState(174)
		p.Match(MapCSSParserT__1)
	}
	{
		p.SetState(175)

		var _x = p.Uint_()

		localctx.(*Rgb_colorContext).g = _x
	}
	{
		p.SetState(176)
		p.Match(MapCSSParserT__1)
	}
	{
		p.SetState(177)

		var _x = p.Uint_()

		localctx.(*Rgb_colorContext).b = _x
	}
	{
		p.SetState(178)
		p.Match(MapCSSParserT__14)
	}

	return localctx
}

// IRgba_colorContext is an interface to support dynamic dispatch.
type IRgba_colorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetR returns the r rule contexts.
	GetR() IUint_Context

	// GetG returns the g rule contexts.
	GetG() IUint_Context

	// GetB returns the b rule contexts.
	GetB() IUint_Context

	// GetA returns the a rule contexts.
	GetA() IFloat_Context

	// SetR sets the r rule contexts.
	SetR(IUint_Context)

	// SetG sets the g rule contexts.
	SetG(IUint_Context)

	// SetB sets the b rule contexts.
	SetB(IUint_Context)

	// SetA sets the a rule contexts.
	SetA(IFloat_Context)

	// IsRgba_colorContext differentiates from other interfaces.
	IsRgba_colorContext()
}

type Rgba_colorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	r      IUint_Context
	g      IUint_Context
	b      IUint_Context
	a      IFloat_Context
}

func NewEmptyRgba_colorContext() *Rgba_colorContext {
	var p = new(Rgba_colorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_rgba_color
	return p
}

func (*Rgba_colorContext) IsRgba_colorContext() {}

func NewRgba_colorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rgba_colorContext {
	var p = new(Rgba_colorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_rgba_color

	return p
}

func (s *Rgba_colorContext) GetParser() antlr.Parser { return s.parser }

func (s *Rgba_colorContext) GetR() IUint_Context { return s.r }

func (s *Rgba_colorContext) GetG() IUint_Context { return s.g }

func (s *Rgba_colorContext) GetB() IUint_Context { return s.b }

func (s *Rgba_colorContext) GetA() IFloat_Context { return s.a }

func (s *Rgba_colorContext) SetR(v IUint_Context) { s.r = v }

func (s *Rgba_colorContext) SetG(v IUint_Context) { s.g = v }

func (s *Rgba_colorContext) SetB(v IUint_Context) { s.b = v }

func (s *Rgba_colorContext) SetA(v IFloat_Context) { s.a = v }

func (s *Rgba_colorContext) AllUint_() []IUint_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUint_Context)(nil)).Elem())
	var tst = make([]IUint_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUint_Context)
		}
	}

	return tst
}

func (s *Rgba_colorContext) Uint_(i int) IUint_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUint_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUint_Context)
}

func (s *Rgba_colorContext) Float_() IFloat_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloat_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloat_Context)
}

func (s *Rgba_colorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rgba_colorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rgba_colorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterRgba_color(s)
	}
}

func (s *Rgba_colorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitRgba_color(s)
	}
}

func (p *MapCSSParser) Rgba_color() (localctx IRgba_colorContext) {
	localctx = NewRgba_colorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MapCSSParserRULE_rgba_color)

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
		p.SetState(180)
		p.Match(MapCSSParserT__15)
	}
	{
		p.SetState(181)

		var _x = p.Uint_()

		localctx.(*Rgba_colorContext).r = _x
	}
	{
		p.SetState(182)
		p.Match(MapCSSParserT__1)
	}
	{
		p.SetState(183)

		var _x = p.Uint_()

		localctx.(*Rgba_colorContext).g = _x
	}
	{
		p.SetState(184)
		p.Match(MapCSSParserT__1)
	}
	{
		p.SetState(185)

		var _x = p.Uint_()

		localctx.(*Rgba_colorContext).b = _x
	}
	{
		p.SetState(186)
		p.Match(MapCSSParserT__1)
	}
	{
		p.SetState(187)

		var _x = p.Float_()

		localctx.(*Rgba_colorContext).a = _x
	}
	{
		p.SetState(188)
		p.Match(MapCSSParserT__14)
	}

	return localctx
}
