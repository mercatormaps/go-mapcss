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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 25, 96, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 3, 2, 7, 2, 28, 10, 2, 12, 2, 14, 2, 31, 11, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 6, 5, 42, 10, 5, 13, 5, 14, 5, 43, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 63, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 76, 10, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 2, 2, 14, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 2, 4, 3, 2, 21, 23, 4, 2, 15, 15, 17, 17,
	2, 89, 2, 29, 3, 2, 2, 2, 4, 34, 3, 2, 2, 2, 6, 36, 3, 2, 2, 2, 8, 38,
	3, 2, 2, 2, 10, 62, 3, 2, 2, 2, 12, 64, 3, 2, 2, 2, 14, 66, 3, 2, 2, 2,
	16, 68, 3, 2, 2, 2, 18, 70, 3, 2, 2, 2, 20, 75, 3, 2, 2, 2, 22, 77, 3,
	2, 2, 2, 24, 85, 3, 2, 2, 2, 26, 28, 5, 4, 3, 2, 27, 26, 3, 2, 2, 2, 28,
	31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 32, 3, 2, 2,
	2, 31, 29, 3, 2, 2, 2, 32, 33, 7, 2, 2, 3, 33, 3, 3, 2, 2, 2, 34, 35, 5,
	6, 4, 2, 35, 5, 3, 2, 2, 2, 36, 37, 5, 8, 5, 2, 37, 7, 3, 2, 2, 2, 38,
	39, 7, 3, 2, 2, 39, 41, 7, 9, 2, 2, 40, 42, 5, 10, 6, 2, 41, 40, 3, 2,
	2, 2, 42, 43, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 45,
	3, 2, 2, 2, 45, 46, 7, 10, 2, 2, 46, 9, 3, 2, 2, 2, 47, 48, 7, 20, 2, 2,
	48, 49, 7, 11, 2, 2, 49, 50, 5, 12, 7, 2, 50, 51, 7, 12, 2, 2, 51, 63,
	3, 2, 2, 2, 52, 53, 7, 24, 2, 2, 53, 54, 7, 11, 2, 2, 54, 55, 5, 14, 8,
	2, 55, 56, 7, 12, 2, 2, 56, 63, 3, 2, 2, 2, 57, 58, 7, 25, 2, 2, 58, 59,
	7, 11, 2, 2, 59, 60, 5, 16, 9, 2, 60, 61, 7, 12, 2, 2, 61, 63, 3, 2, 2,
	2, 62, 47, 3, 2, 2, 2, 62, 52, 3, 2, 2, 2, 62, 57, 3, 2, 2, 2, 63, 11,
	3, 2, 2, 2, 64, 65, 9, 2, 2, 2, 65, 13, 3, 2, 2, 2, 66, 67, 9, 3, 2, 2,
	67, 15, 3, 2, 2, 2, 68, 69, 5, 20, 11, 2, 69, 17, 3, 2, 2, 2, 70, 71, 7,
	15, 2, 2, 71, 19, 3, 2, 2, 2, 72, 76, 7, 19, 2, 2, 73, 76, 5, 22, 12, 2,
	74, 76, 5, 24, 13, 2, 75, 72, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 74, 3,
	2, 2, 2, 76, 21, 3, 2, 2, 2, 77, 78, 7, 4, 2, 2, 78, 79, 7, 15, 2, 2, 79,
	80, 7, 5, 2, 2, 80, 81, 7, 15, 2, 2, 81, 82, 7, 5, 2, 2, 82, 83, 7, 15,
	2, 2, 83, 84, 7, 6, 2, 2, 84, 23, 3, 2, 2, 2, 85, 86, 7, 7, 2, 2, 86, 87,
	7, 15, 2, 2, 87, 88, 7, 5, 2, 2, 88, 89, 7, 15, 2, 2, 89, 90, 7, 5, 2,
	2, 90, 91, 7, 15, 2, 2, 91, 92, 7, 5, 2, 2, 92, 93, 9, 3, 2, 2, 93, 94,
	7, 6, 2, 2, 94, 25, 3, 2, 2, 2, 6, 29, 43, 62, 75,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'canvas'", "'rgb('", "','", "')'", "'rgba('", "", "'{'", "'}'", "':'",
	"';'", "", "", "", "", "", "", "", "'antialiasing'", "'full'", "'text'",
	"'none'", "'fill-opacity'", "'fill-color'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "WS", "LBRACE", "RBRACE", "COLON", "SEMICOLON",
	"DQUOTED_STRING", "SQUOTED_STRING", "POSITIVE_INT", "NEGATIVE_INT", "POSITIVE_FLOAT",
	"NEGATIVE_FLOAT", "HEX", "PROP_ANTIALIASING", "PROP_ANTIALIASING_FULL",
	"PROP_ANTIALIASING_TEXT", "PROP_ANTIALIASING_NONE", "PROP_FILL_OPACITY",
	"PROP_FILL_COLOR",
}

var ruleNames = []string{
	"stylesheet", "entry", "rule_", "canvas_declaration_block", "canvas_declaration",
	"antialiasing", "fill_opacity", "fill_color", "alpha", "color", "rgb_color",
	"rgba_color",
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
	MapCSSParserEOF                    = antlr.TokenEOF
	MapCSSParserT__0                   = 1
	MapCSSParserT__1                   = 2
	MapCSSParserT__2                   = 3
	MapCSSParserT__3                   = 4
	MapCSSParserT__4                   = 5
	MapCSSParserWS                     = 6
	MapCSSParserLBRACE                 = 7
	MapCSSParserRBRACE                 = 8
	MapCSSParserCOLON                  = 9
	MapCSSParserSEMICOLON              = 10
	MapCSSParserDQUOTED_STRING         = 11
	MapCSSParserSQUOTED_STRING         = 12
	MapCSSParserPOSITIVE_INT           = 13
	MapCSSParserNEGATIVE_INT           = 14
	MapCSSParserPOSITIVE_FLOAT         = 15
	MapCSSParserNEGATIVE_FLOAT         = 16
	MapCSSParserHEX                    = 17
	MapCSSParserPROP_ANTIALIASING      = 18
	MapCSSParserPROP_ANTIALIASING_FULL = 19
	MapCSSParserPROP_ANTIALIASING_TEXT = 20
	MapCSSParserPROP_ANTIALIASING_NONE = 21
	MapCSSParserPROP_FILL_OPACITY      = 22
	MapCSSParserPROP_FILL_COLOR        = 23
)

// MapCSSParser rules.
const (
	MapCSSParserRULE_stylesheet               = 0
	MapCSSParserRULE_entry                    = 1
	MapCSSParserRULE_rule_                    = 2
	MapCSSParserRULE_canvas_declaration_block = 3
	MapCSSParserRULE_canvas_declaration       = 4
	MapCSSParserRULE_antialiasing             = 5
	MapCSSParserRULE_fill_opacity             = 6
	MapCSSParserRULE_fill_color               = 7
	MapCSSParserRULE_alpha                    = 8
	MapCSSParserRULE_color                    = 9
	MapCSSParserRULE_rgb_color                = 10
	MapCSSParserRULE_rgba_color               = 11
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MapCSSParserT__0 {
		{
			p.SetState(24)
			p.Entry()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(30)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Rule_()
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

func (s *Rule_Context) Canvas_declaration_block() ICanvas_declaration_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICanvas_declaration_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICanvas_declaration_blockContext)
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
		p.SetState(34)
		p.Canvas_declaration_block()
	}

	return localctx
}

// ICanvas_declaration_blockContext is an interface to support dynamic dispatch.
type ICanvas_declaration_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCanvas_declaration_blockContext differentiates from other interfaces.
	IsCanvas_declaration_blockContext()
}

type Canvas_declaration_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCanvas_declaration_blockContext() *Canvas_declaration_blockContext {
	var p = new(Canvas_declaration_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_canvas_declaration_block
	return p
}

func (*Canvas_declaration_blockContext) IsCanvas_declaration_blockContext() {}

func NewCanvas_declaration_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Canvas_declaration_blockContext {
	var p = new(Canvas_declaration_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_canvas_declaration_block

	return p
}

func (s *Canvas_declaration_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Canvas_declaration_blockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MapCSSParserLBRACE, 0)
}

func (s *Canvas_declaration_blockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MapCSSParserRBRACE, 0)
}

func (s *Canvas_declaration_blockContext) AllCanvas_declaration() []ICanvas_declarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICanvas_declarationContext)(nil)).Elem())
	var tst = make([]ICanvas_declarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICanvas_declarationContext)
		}
	}

	return tst
}

func (s *Canvas_declaration_blockContext) Canvas_declaration(i int) ICanvas_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICanvas_declarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICanvas_declarationContext)
}

func (s *Canvas_declaration_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Canvas_declaration_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Canvas_declaration_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterCanvas_declaration_block(s)
	}
}

func (s *Canvas_declaration_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitCanvas_declaration_block(s)
	}
}

func (p *MapCSSParser) Canvas_declaration_block() (localctx ICanvas_declaration_blockContext) {
	localctx = NewCanvas_declaration_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MapCSSParserRULE_canvas_declaration_block)
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
		p.SetState(36)
		p.Match(MapCSSParserT__0)
	}
	{
		p.SetState(37)
		p.Match(MapCSSParserLBRACE)
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MapCSSParserPROP_ANTIALIASING)|(1<<MapCSSParserPROP_FILL_OPACITY)|(1<<MapCSSParserPROP_FILL_COLOR))) != 0) {
		{
			p.SetState(38)
			p.Canvas_declaration()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(43)
		p.Match(MapCSSParserRBRACE)
	}

	return localctx
}

// ICanvas_declarationContext is an interface to support dynamic dispatch.
type ICanvas_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCanvas_declarationContext differentiates from other interfaces.
	IsCanvas_declarationContext()
}

type Canvas_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCanvas_declarationContext() *Canvas_declarationContext {
	var p = new(Canvas_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_canvas_declaration
	return p
}

func (*Canvas_declarationContext) IsCanvas_declarationContext() {}

func NewCanvas_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Canvas_declarationContext {
	var p = new(Canvas_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_canvas_declaration

	return p
}

func (s *Canvas_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Canvas_declarationContext) PROP_ANTIALIASING() antlr.TerminalNode {
	return s.GetToken(MapCSSParserPROP_ANTIALIASING, 0)
}

func (s *Canvas_declarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(MapCSSParserCOLON, 0)
}

func (s *Canvas_declarationContext) Antialiasing() IAntialiasingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAntialiasingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAntialiasingContext)
}

func (s *Canvas_declarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MapCSSParserSEMICOLON, 0)
}

func (s *Canvas_declarationContext) PROP_FILL_OPACITY() antlr.TerminalNode {
	return s.GetToken(MapCSSParserPROP_FILL_OPACITY, 0)
}

func (s *Canvas_declarationContext) Fill_opacity() IFill_opacityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFill_opacityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFill_opacityContext)
}

func (s *Canvas_declarationContext) PROP_FILL_COLOR() antlr.TerminalNode {
	return s.GetToken(MapCSSParserPROP_FILL_COLOR, 0)
}

func (s *Canvas_declarationContext) Fill_color() IFill_colorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFill_colorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFill_colorContext)
}

func (s *Canvas_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Canvas_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Canvas_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterCanvas_declaration(s)
	}
}

func (s *Canvas_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitCanvas_declaration(s)
	}
}

func (p *MapCSSParser) Canvas_declaration() (localctx ICanvas_declarationContext) {
	localctx = NewCanvas_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MapCSSParserRULE_canvas_declaration)

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

	p.SetState(60)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MapCSSParserPROP_ANTIALIASING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(MapCSSParserPROP_ANTIALIASING)
		}
		{
			p.SetState(46)
			p.Match(MapCSSParserCOLON)
		}
		{
			p.SetState(47)
			p.Antialiasing()
		}
		{
			p.SetState(48)
			p.Match(MapCSSParserSEMICOLON)
		}

	case MapCSSParserPROP_FILL_OPACITY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
			p.Match(MapCSSParserPROP_FILL_OPACITY)
		}
		{
			p.SetState(51)
			p.Match(MapCSSParserCOLON)
		}
		{
			p.SetState(52)
			p.Fill_opacity()
		}
		{
			p.SetState(53)
			p.Match(MapCSSParserSEMICOLON)
		}

	case MapCSSParserPROP_FILL_COLOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Match(MapCSSParserPROP_FILL_COLOR)
		}
		{
			p.SetState(56)
			p.Match(MapCSSParserCOLON)
		}
		{
			p.SetState(57)
			p.Fill_color()
		}
		{
			p.SetState(58)
			p.Match(MapCSSParserSEMICOLON)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAntialiasingContext is an interface to support dynamic dispatch.
type IAntialiasingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAntialiasingContext differentiates from other interfaces.
	IsAntialiasingContext()
}

type AntialiasingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAntialiasingContext() *AntialiasingContext {
	var p = new(AntialiasingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_antialiasing
	return p
}

func (*AntialiasingContext) IsAntialiasingContext() {}

func NewAntialiasingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AntialiasingContext {
	var p = new(AntialiasingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_antialiasing

	return p
}

func (s *AntialiasingContext) GetParser() antlr.Parser { return s.parser }

func (s *AntialiasingContext) PROP_ANTIALIASING_FULL() antlr.TerminalNode {
	return s.GetToken(MapCSSParserPROP_ANTIALIASING_FULL, 0)
}

func (s *AntialiasingContext) PROP_ANTIALIASING_TEXT() antlr.TerminalNode {
	return s.GetToken(MapCSSParserPROP_ANTIALIASING_TEXT, 0)
}

func (s *AntialiasingContext) PROP_ANTIALIASING_NONE() antlr.TerminalNode {
	return s.GetToken(MapCSSParserPROP_ANTIALIASING_NONE, 0)
}

func (s *AntialiasingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AntialiasingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AntialiasingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterAntialiasing(s)
	}
}

func (s *AntialiasingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitAntialiasing(s)
	}
}

func (p *MapCSSParser) Antialiasing() (localctx IAntialiasingContext) {
	localctx = NewAntialiasingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MapCSSParserRULE_antialiasing)
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
		p.SetState(62)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MapCSSParserPROP_ANTIALIASING_FULL)|(1<<MapCSSParserPROP_ANTIALIASING_TEXT)|(1<<MapCSSParserPROP_ANTIALIASING_NONE))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFill_opacityContext is an interface to support dynamic dispatch.
type IFill_opacityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFill_opacityContext differentiates from other interfaces.
	IsFill_opacityContext()
}

type Fill_opacityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFill_opacityContext() *Fill_opacityContext {
	var p = new(Fill_opacityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_fill_opacity
	return p
}

func (*Fill_opacityContext) IsFill_opacityContext() {}

func NewFill_opacityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fill_opacityContext {
	var p = new(Fill_opacityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_fill_opacity

	return p
}

func (s *Fill_opacityContext) GetParser() antlr.Parser { return s.parser }

func (s *Fill_opacityContext) POSITIVE_INT() antlr.TerminalNode {
	return s.GetToken(MapCSSParserPOSITIVE_INT, 0)
}

func (s *Fill_opacityContext) POSITIVE_FLOAT() antlr.TerminalNode {
	return s.GetToken(MapCSSParserPOSITIVE_FLOAT, 0)
}

func (s *Fill_opacityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fill_opacityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fill_opacityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterFill_opacity(s)
	}
}

func (s *Fill_opacityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitFill_opacity(s)
	}
}

func (p *MapCSSParser) Fill_opacity() (localctx IFill_opacityContext) {
	localctx = NewFill_opacityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MapCSSParserRULE_fill_opacity)
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
		p.SetState(64)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MapCSSParserPOSITIVE_INT || _la == MapCSSParserPOSITIVE_FLOAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFill_colorContext is an interface to support dynamic dispatch.
type IFill_colorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFill_colorContext differentiates from other interfaces.
	IsFill_colorContext()
}

type Fill_colorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFill_colorContext() *Fill_colorContext {
	var p = new(Fill_colorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_fill_color
	return p
}

func (*Fill_colorContext) IsFill_colorContext() {}

func NewFill_colorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fill_colorContext {
	var p = new(Fill_colorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_fill_color

	return p
}

func (s *Fill_colorContext) GetParser() antlr.Parser { return s.parser }

func (s *Fill_colorContext) Color() IColorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColorContext)
}

func (s *Fill_colorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fill_colorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fill_colorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterFill_color(s)
	}
}

func (s *Fill_colorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitFill_color(s)
	}
}

func (p *MapCSSParser) Fill_color() (localctx IFill_colorContext) {
	localctx = NewFill_colorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MapCSSParserRULE_fill_color)

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
		p.SetState(66)
		p.Color()
	}

	return localctx
}

// IAlphaContext is an interface to support dynamic dispatch.
type IAlphaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlphaContext differentiates from other interfaces.
	IsAlphaContext()
}

type AlphaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlphaContext() *AlphaContext {
	var p = new(AlphaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MapCSSParserRULE_alpha
	return p
}

func (*AlphaContext) IsAlphaContext() {}

func NewAlphaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlphaContext {
	var p = new(AlphaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MapCSSParserRULE_alpha

	return p
}

func (s *AlphaContext) GetParser() antlr.Parser { return s.parser }

func (s *AlphaContext) POSITIVE_INT() antlr.TerminalNode {
	return s.GetToken(MapCSSParserPOSITIVE_INT, 0)
}

func (s *AlphaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlphaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlphaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.EnterAlpha(s)
	}
}

func (s *AlphaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MapCSSListener); ok {
		listenerT.ExitAlpha(s)
	}
}

func (p *MapCSSParser) Alpha() (localctx IAlphaContext) {
	localctx = NewAlphaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MapCSSParserRULE_alpha)

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
		p.SetState(68)
		p.Match(MapCSSParserPOSITIVE_INT)
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
	p.EnterRule(localctx, 18, MapCSSParserRULE_color)

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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MapCSSParserHEX:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Match(MapCSSParserHEX)
		}

	case MapCSSParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Rgb_color()
		}

	case MapCSSParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
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
	p.EnterRule(localctx, 20, MapCSSParserRULE_rgb_color)

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
		p.SetState(75)
		p.Match(MapCSSParserT__1)
	}
	{
		p.SetState(76)

		var _m = p.Match(MapCSSParserPOSITIVE_INT)

		localctx.(*Rgb_colorContext).r = _m
	}
	{
		p.SetState(77)
		p.Match(MapCSSParserT__2)
	}
	{
		p.SetState(78)

		var _m = p.Match(MapCSSParserPOSITIVE_INT)

		localctx.(*Rgb_colorContext).g = _m
	}
	{
		p.SetState(79)
		p.Match(MapCSSParserT__2)
	}
	{
		p.SetState(80)

		var _m = p.Match(MapCSSParserPOSITIVE_INT)

		localctx.(*Rgb_colorContext).b = _m
	}
	{
		p.SetState(81)
		p.Match(MapCSSParserT__3)
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
	p.EnterRule(localctx, 22, MapCSSParserRULE_rgba_color)
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
		p.Match(MapCSSParserT__4)
	}
	{
		p.SetState(84)

		var _m = p.Match(MapCSSParserPOSITIVE_INT)

		localctx.(*Rgba_colorContext).r = _m
	}
	{
		p.SetState(85)
		p.Match(MapCSSParserT__2)
	}
	{
		p.SetState(86)

		var _m = p.Match(MapCSSParserPOSITIVE_INT)

		localctx.(*Rgba_colorContext).g = _m
	}
	{
		p.SetState(87)
		p.Match(MapCSSParserT__2)
	}
	{
		p.SetState(88)

		var _m = p.Match(MapCSSParserPOSITIVE_INT)

		localctx.(*Rgba_colorContext).b = _m
	}
	{
		p.SetState(89)
		p.Match(MapCSSParserT__2)
	}
	{
		p.SetState(90)

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
		p.SetState(91)
		p.Match(MapCSSParserT__3)
	}

	return localctx
}
