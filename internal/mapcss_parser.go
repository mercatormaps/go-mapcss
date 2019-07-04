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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 124,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 7, 2, 34, 10, 2,
	12, 2, 14, 2, 37, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 43, 10, 3, 3, 4,
	3, 4, 3, 4, 6, 4, 48, 10, 4, 13, 4, 14, 4, 49, 3, 4, 3, 4, 3, 5, 3, 5,
	6, 5, 56, 10, 5, 13, 5, 14, 5, 57, 3, 6, 3, 6, 5, 6, 62, 10, 6, 3, 6, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 6, 9, 75, 10,
	9, 13, 9, 14, 9, 76, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 5, 10, 84, 10, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 5, 14, 104, 10, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 2, 2, 17, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 2, 4, 3, 2, 9, 11,
	4, 2, 24, 24, 26, 26, 2, 118, 2, 35, 3, 2, 2, 2, 4, 42, 3, 2, 2, 2, 6,
	44, 3, 2, 2, 2, 8, 53, 3, 2, 2, 2, 10, 59, 3, 2, 2, 2, 12, 66, 3, 2, 2,
	2, 14, 69, 3, 2, 2, 2, 16, 72, 3, 2, 2, 2, 18, 83, 3, 2, 2, 2, 20, 85,
	3, 2, 2, 2, 22, 90, 3, 2, 2, 2, 24, 95, 3, 2, 2, 2, 26, 103, 3, 2, 2, 2,
	28, 105, 3, 2, 2, 2, 30, 113, 3, 2, 2, 2, 32, 34, 5, 4, 3, 2, 33, 32, 3,
	2, 2, 2, 34, 37, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36,
	38, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 38, 39, 7, 2, 2, 3, 39, 3, 3, 2, 2,
	2, 40, 43, 5, 14, 8, 2, 41, 43, 5, 6, 4, 2, 42, 40, 3, 2, 2, 2, 42, 41,
	3, 2, 2, 2, 43, 5, 3, 2, 2, 2, 44, 47, 5, 8, 5, 2, 45, 46, 7, 3, 2, 2,
	46, 48, 5, 8, 5, 2, 47, 45, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 47, 3,
	2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 52, 5, 12, 7, 2, 52,
	7, 3, 2, 2, 2, 53, 55, 7, 29, 2, 2, 54, 56, 5, 10, 6, 2, 55, 54, 3, 2,
	2, 2, 56, 57, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 9,
	3, 2, 2, 2, 59, 61, 7, 4, 2, 2, 60, 62, 7, 5, 2, 2, 61, 60, 3, 2, 2, 2,
	61, 62, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 7, 29, 2, 2, 64, 65, 7,
	6, 2, 2, 65, 11, 3, 2, 2, 2, 66, 67, 7, 18, 2, 2, 67, 68, 7, 19, 2, 2,
	68, 13, 3, 2, 2, 2, 69, 70, 7, 7, 2, 2, 70, 71, 5, 16, 9, 2, 71, 15, 3,
	2, 2, 2, 72, 74, 7, 18, 2, 2, 73, 75, 5, 18, 10, 2, 74, 73, 3, 2, 2, 2,
	75, 76, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 3,
	2, 2, 2, 78, 79, 7, 19, 2, 2, 79, 17, 3, 2, 2, 2, 80, 84, 5, 20, 11, 2,
	81, 84, 5, 22, 12, 2, 82, 84, 5, 24, 13, 2, 83, 80, 3, 2, 2, 2, 83, 81,
	3, 2, 2, 2, 83, 82, 3, 2, 2, 2, 84, 19, 3, 2, 2, 2, 85, 86, 7, 8, 2, 2,
	86, 87, 7, 20, 2, 2, 87, 88, 9, 2, 2, 2, 88, 89, 7, 21, 2, 2, 89, 21, 3,
	2, 2, 2, 90, 91, 7, 12, 2, 2, 91, 92, 7, 20, 2, 2, 92, 93, 9, 3, 2, 2,
	93, 94, 7, 21, 2, 2, 94, 23, 3, 2, 2, 2, 95, 96, 7, 13, 2, 2, 96, 97, 7,
	20, 2, 2, 97, 98, 5, 26, 14, 2, 98, 99, 7, 21, 2, 2, 99, 25, 3, 2, 2, 2,
	100, 104, 7, 28, 2, 2, 101, 104, 5, 28, 15, 2, 102, 104, 5, 30, 16, 2,
	103, 100, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 102, 3, 2, 2, 2, 104,
	27, 3, 2, 2, 2, 105, 106, 7, 14, 2, 2, 106, 107, 7, 24, 2, 2, 107, 108,
	7, 3, 2, 2, 108, 109, 7, 24, 2, 2, 109, 110, 7, 3, 2, 2, 110, 111, 7, 24,
	2, 2, 111, 112, 7, 15, 2, 2, 112, 29, 3, 2, 2, 2, 113, 114, 7, 16, 2, 2,
	114, 115, 7, 24, 2, 2, 115, 116, 7, 3, 2, 2, 116, 117, 7, 24, 2, 2, 117,
	118, 7, 3, 2, 2, 118, 119, 7, 24, 2, 2, 119, 120, 7, 3, 2, 2, 120, 121,
	9, 3, 2, 2, 121, 122, 7, 15, 2, 2, 122, 31, 3, 2, 2, 2, 10, 35, 42, 49,
	57, 61, 76, 83, 103,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'['", "'!'", "']'", "'canvas'", "'antialiasing'", "'full'",
	"'text'", "'none'", "'fill-opacity'", "'fill-color'", "'rgb('", "')'",
	"'rgba('", "", "'{'", "'}'", "':'", "';'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "WS", "LBRACE",
	"RBRACE", "COLON", "SEMICOLON", "DQUOTED_STRING", "SQUOTED_STRING", "POSITIVE_INT",
	"NEGATIVE_INT", "POSITIVE_FLOAT", "NEGATIVE_FLOAT", "HEX", "IDENTIFIER",
}

var ruleNames = []string{
	"stylesheet", "entry", "rule_", "selector", "attribute", "decl_block",
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
	MapCSSParserWS             = 15
	MapCSSParserLBRACE         = 16
	MapCSSParserRBRACE         = 17
	MapCSSParserCOLON          = 18
	MapCSSParserSEMICOLON      = 19
	MapCSSParserDQUOTED_STRING = 20
	MapCSSParserSQUOTED_STRING = 21
	MapCSSParserPOSITIVE_INT   = 22
	MapCSSParserNEGATIVE_INT   = 23
	MapCSSParserPOSITIVE_FLOAT = 24
	MapCSSParserNEGATIVE_FLOAT = 25
	MapCSSParserHEX            = 26
	MapCSSParserIDENTIFIER     = 27
)

// MapCSSParser rules.
const (
	MapCSSParserRULE_stylesheet        = 0
	MapCSSParserRULE_entry             = 1
	MapCSSParserRULE_rule_             = 2
	MapCSSParserRULE_selector          = 3
	MapCSSParserRULE_attribute         = 4
	MapCSSParserRULE_decl_block        = 5
	MapCSSParserRULE_canvas_rule       = 6
	MapCSSParserRULE_canvas_decl_block = 7
	MapCSSParserRULE_canvas_decl       = 8
	MapCSSParserRULE_antialiasing_decl = 9
	MapCSSParserRULE_fill_opacity_decl = 10
	MapCSSParserRULE_fill_color_decl   = 11
	MapCSSParserRULE_color             = 12
	MapCSSParserRULE_rgb_color         = 13
	MapCSSParserRULE_rgba_color        = 14
)

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
	p.EnterRule(localctx, 0, MapCSSParserRULE_stylesheet)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MapCSSParserT__4 || _la == MapCSSParserIDENTIFIER {
		{
			p.SetState(30)
			p.Entry()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
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
	p.EnterRule(localctx, 2, MapCSSParserRULE_entry)

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

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MapCSSParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Canvas_rule()
		}

	case MapCSSParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
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
	p.EnterRule(localctx, 4, MapCSSParserRULE_rule_)
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
		p.SetState(42)
		p.Selector()
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MapCSSParserT__0 {
		{
			p.SetState(43)
			p.Match(MapCSSParserT__0)
		}
		{
			p.SetState(44)
			p.Selector()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(49)
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
	p.EnterRule(localctx, 6, MapCSSParserRULE_selector)
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
		p.SetState(51)

		var _m = p.Match(MapCSSParserIDENTIFIER)

		localctx.(*SelectorContext).typ = _m
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MapCSSParserT__1 {
		{
			p.SetState(52)
			p.Attribute()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 8, MapCSSParserRULE_attribute)
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
		p.SetState(57)
		p.Match(MapCSSParserT__1)
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MapCSSParserT__2 {
		{
			p.SetState(58)

			var _m = p.Match(MapCSSParserT__2)

			localctx.(*AttributeContext).neg = _m
		}

	}
	{
		p.SetState(61)

		var _m = p.Match(MapCSSParserIDENTIFIER)

		localctx.(*AttributeContext).name = _m
	}
	{
		p.SetState(62)
		p.Match(MapCSSParserT__3)
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
	p.EnterRule(localctx, 10, MapCSSParserRULE_decl_block)

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
		p.SetState(64)
		p.Match(MapCSSParserLBRACE)
	}
	{
		p.SetState(65)
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
	p.EnterRule(localctx, 12, MapCSSParserRULE_canvas_rule)

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
		p.SetState(67)
		p.Match(MapCSSParserT__4)
	}
	{
		p.SetState(68)
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
	p.EnterRule(localctx, 14, MapCSSParserRULE_canvas_decl_block)
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
		p.SetState(70)
		p.Match(MapCSSParserLBRACE)
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MapCSSParserT__5)|(1<<MapCSSParserT__9)|(1<<MapCSSParserT__10))) != 0) {
		{
			p.SetState(71)
			p.Canvas_decl()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(76)
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
	p.EnterRule(localctx, 16, MapCSSParserRULE_canvas_decl)

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

	p.SetState(81)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MapCSSParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Antialiasing_decl()
		}

	case MapCSSParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Fill_opacity_decl()
		}

	case MapCSSParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
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
	p.EnterRule(localctx, 18, MapCSSParserRULE_antialiasing_decl)
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
		p.SetState(83)
		p.Match(MapCSSParserT__5)
	}
	{
		p.SetState(84)
		p.Match(MapCSSParserCOLON)
	}
	{
		p.SetState(85)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Antialiasing_declContext).v = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MapCSSParserT__6)|(1<<MapCSSParserT__7)|(1<<MapCSSParserT__8))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Antialiasing_declContext).v = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(86)
		p.Match(MapCSSParserSEMICOLON)
	}

	return localctx
}

// IFill_opacity_declContext is an interface to support dynamic dispatch.
type IFill_opacity_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// IsFill_opacity_declContext differentiates from other interfaces.
	IsFill_opacity_declContext()
}

type Fill_opacity_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      antlr.Token
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

func (s *Fill_opacity_declContext) GetV() antlr.Token { return s.v }

func (s *Fill_opacity_declContext) SetV(v antlr.Token) { s.v = v }

func (s *Fill_opacity_declContext) COLON() antlr.TerminalNode {
	return s.GetToken(MapCSSParserCOLON, 0)
}

func (s *Fill_opacity_declContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MapCSSParserSEMICOLON, 0)
}

func (s *Fill_opacity_declContext) POSITIVE_INT() antlr.TerminalNode {
	return s.GetToken(MapCSSParserPOSITIVE_INT, 0)
}

func (s *Fill_opacity_declContext) POSITIVE_FLOAT() antlr.TerminalNode {
	return s.GetToken(MapCSSParserPOSITIVE_FLOAT, 0)
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
	p.EnterRule(localctx, 20, MapCSSParserRULE_fill_opacity_decl)
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
		p.SetState(88)
		p.Match(MapCSSParserT__9)
	}
	{
		p.SetState(89)
		p.Match(MapCSSParserCOLON)
	}
	{
		p.SetState(90)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Fill_opacity_declContext).v = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == MapCSSParserPOSITIVE_INT || _la == MapCSSParserPOSITIVE_FLOAT) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Fill_opacity_declContext).v = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(91)
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
	p.EnterRule(localctx, 22, MapCSSParserRULE_fill_color_decl)

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
		p.SetState(93)
		p.Match(MapCSSParserT__10)
	}
	{
		p.SetState(94)
		p.Match(MapCSSParserCOLON)
	}
	{
		p.SetState(95)
		p.Color()
	}
	{
		p.SetState(96)
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
	p.EnterRule(localctx, 24, MapCSSParserRULE_color)

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

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MapCSSParserHEX:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Match(MapCSSParserHEX)
		}

	case MapCSSParserT__11:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Rgb_color()
		}

	case MapCSSParserT__13:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
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

	// GetR returns the r token.
	GetR() antlr.Token

	// GetG returns the g token.
	GetG() antlr.Token

	// GetB returns the b token.
	GetB() antlr.Token

	// SetR sets the r token.
	SetR(antlr.Token)

	// SetG sets the g token.
	SetG(antlr.Token)

	// SetB sets the b token.
	SetB(antlr.Token)

	// IsRgb_colorContext differentiates from other interfaces.
	IsRgb_colorContext()
}

type Rgb_colorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	r      antlr.Token
	g      antlr.Token
	b      antlr.Token
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

func (s *Rgb_colorContext) GetR() antlr.Token { return s.r }

func (s *Rgb_colorContext) GetG() antlr.Token { return s.g }

func (s *Rgb_colorContext) GetB() antlr.Token { return s.b }

func (s *Rgb_colorContext) SetR(v antlr.Token) { s.r = v }

func (s *Rgb_colorContext) SetG(v antlr.Token) { s.g = v }

func (s *Rgb_colorContext) SetB(v antlr.Token) { s.b = v }

func (s *Rgb_colorContext) AllPOSITIVE_INT() []antlr.TerminalNode {
	return s.GetTokens(MapCSSParserPOSITIVE_INT)
}

func (s *Rgb_colorContext) POSITIVE_INT(i int) antlr.TerminalNode {
	return s.GetToken(MapCSSParserPOSITIVE_INT, i)
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
	p.EnterRule(localctx, 26, MapCSSParserRULE_rgb_color)

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
		p.SetState(103)
		p.Match(MapCSSParserT__11)
	}
	{
		p.SetState(104)

		var _m = p.Match(MapCSSParserPOSITIVE_INT)

		localctx.(*Rgb_colorContext).r = _m
	}
	{
		p.SetState(105)
		p.Match(MapCSSParserT__0)
	}
	{
		p.SetState(106)

		var _m = p.Match(MapCSSParserPOSITIVE_INT)

		localctx.(*Rgb_colorContext).g = _m
	}
	{
		p.SetState(107)
		p.Match(MapCSSParserT__0)
	}
	{
		p.SetState(108)

		var _m = p.Match(MapCSSParserPOSITIVE_INT)

		localctx.(*Rgb_colorContext).b = _m
	}
	{
		p.SetState(109)
		p.Match(MapCSSParserT__12)
	}

	return localctx
}

// IRgba_colorContext is an interface to support dynamic dispatch.
type IRgba_colorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetR returns the r token.
	GetR() antlr.Token

	// GetG returns the g token.
	GetG() antlr.Token

	// GetB returns the b token.
	GetB() antlr.Token

	// GetA returns the a token.
	GetA() antlr.Token

	// SetR sets the r token.
	SetR(antlr.Token)

	// SetG sets the g token.
	SetG(antlr.Token)

	// SetB sets the b token.
	SetB(antlr.Token)

	// SetA sets the a token.
	SetA(antlr.Token)

	// IsRgba_colorContext differentiates from other interfaces.
	IsRgba_colorContext()
}

type Rgba_colorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	r      antlr.Token
	g      antlr.Token
	b      antlr.Token
	a      antlr.Token
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

func (s *Rgba_colorContext) GetR() antlr.Token { return s.r }

func (s *Rgba_colorContext) GetG() antlr.Token { return s.g }

func (s *Rgba_colorContext) GetB() antlr.Token { return s.b }

func (s *Rgba_colorContext) GetA() antlr.Token { return s.a }

func (s *Rgba_colorContext) SetR(v antlr.Token) { s.r = v }

func (s *Rgba_colorContext) SetG(v antlr.Token) { s.g = v }

func (s *Rgba_colorContext) SetB(v antlr.Token) { s.b = v }

func (s *Rgba_colorContext) SetA(v antlr.Token) { s.a = v }

func (s *Rgba_colorContext) AllPOSITIVE_INT() []antlr.TerminalNode {
	return s.GetTokens(MapCSSParserPOSITIVE_INT)
}

func (s *Rgba_colorContext) POSITIVE_INT(i int) antlr.TerminalNode {
	return s.GetToken(MapCSSParserPOSITIVE_INT, i)
}

func (s *Rgba_colorContext) POSITIVE_FLOAT() antlr.TerminalNode {
	return s.GetToken(MapCSSParserPOSITIVE_FLOAT, 0)
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
	p.EnterRule(localctx, 28, MapCSSParserRULE_rgba_color)
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
		p.SetState(111)
		p.Match(MapCSSParserT__13)
	}
	{
		p.SetState(112)

		var _m = p.Match(MapCSSParserPOSITIVE_INT)

		localctx.(*Rgba_colorContext).r = _m
	}
	{
		p.SetState(113)
		p.Match(MapCSSParserT__0)
	}
	{
		p.SetState(114)

		var _m = p.Match(MapCSSParserPOSITIVE_INT)

		localctx.(*Rgba_colorContext).g = _m
	}
	{
		p.SetState(115)
		p.Match(MapCSSParserT__0)
	}
	{
		p.SetState(116)

		var _m = p.Match(MapCSSParserPOSITIVE_INT)

		localctx.(*Rgba_colorContext).b = _m
	}
	{
		p.SetState(117)
		p.Match(MapCSSParserT__0)
	}
	{
		p.SetState(118)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Rgba_colorContext).a = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == MapCSSParserPOSITIVE_INT || _la == MapCSSParserPOSITIVE_FLOAT) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Rgba_colorContext).a = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(119)
		p.Match(MapCSSParserT__12)
	}

	return localctx
}
