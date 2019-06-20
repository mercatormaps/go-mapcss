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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 56, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 7, 2, 18, 10, 2, 12, 2, 14, 2, 21, 11, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 6, 5, 32, 10, 5, 13, 5, 14, 5, 33,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 5, 6, 41, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	5, 6, 47, 10, 6, 3, 6, 5, 6, 50, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 2,
	2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 3, 4, 2, 11, 11, 13, 13, 2, 53, 2, 19,
	3, 2, 2, 2, 4, 24, 3, 2, 2, 2, 6, 26, 3, 2, 2, 2, 8, 28, 3, 2, 2, 2, 10,
	49, 3, 2, 2, 2, 12, 51, 3, 2, 2, 2, 14, 53, 3, 2, 2, 2, 16, 18, 5, 4, 3,
	2, 17, 16, 3, 2, 2, 2, 18, 21, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20,
	3, 2, 2, 2, 20, 22, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 22, 23, 7, 2, 2, 3,
	23, 3, 3, 2, 2, 2, 24, 25, 5, 6, 4, 2, 25, 5, 3, 2, 2, 2, 26, 27, 5, 8,
	5, 2, 27, 7, 3, 2, 2, 2, 28, 29, 7, 3, 2, 2, 29, 31, 7, 5, 2, 2, 30, 32,
	5, 10, 6, 2, 31, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2,
	33, 34, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 36, 7, 6, 2, 2, 36, 9, 3, 2,
	2, 2, 37, 38, 7, 15, 2, 2, 38, 40, 7, 7, 2, 2, 39, 41, 5, 12, 7, 2, 40,
	39, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 50, 7, 8, 2,
	2, 43, 44, 7, 17, 2, 2, 44, 46, 7, 7, 2, 2, 45, 47, 5, 14, 8, 2, 46, 45,
	3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 50, 7, 8, 2, 2,
	49, 37, 3, 2, 2, 2, 49, 43, 3, 2, 2, 2, 50, 11, 3, 2, 2, 2, 51, 52, 7,
	16, 2, 2, 52, 13, 3, 2, 2, 2, 53, 54, 9, 2, 2, 2, 54, 15, 3, 2, 2, 2, 7,
	19, 33, 40, 46, 49,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'canvas'", "", "'{'", "'}'", "':'", "';'", "", "", "", "", "", "",
	"'antialiasing'", "", "'fill-opacity'",
}
var symbolicNames = []string{
	"", "", "WS", "LBRACE", "RBRACE", "COLON", "SEMICOLON", "DQUOTED_STRING",
	"SQUOTED_STRING", "POSITIVE_INT", "NEGATIVE_INT", "POSITIVE_FLOAT", "NEGATIVE_FLOAT",
	"PROP_ANTIALIASING", "PROP_ANTIALIASING_VALUES", "PROP_FILL_OPACITY",
}

var ruleNames = []string{
	"stylesheet", "entry", "rule_", "canvas_declaration_block", "canvas_declaration",
	"antialiasing", "fill_opacity",
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
	MapCSSParserEOF                      = antlr.TokenEOF
	MapCSSParserT__0                     = 1
	MapCSSParserWS                       = 2
	MapCSSParserLBRACE                   = 3
	MapCSSParserRBRACE                   = 4
	MapCSSParserCOLON                    = 5
	MapCSSParserSEMICOLON                = 6
	MapCSSParserDQUOTED_STRING           = 7
	MapCSSParserSQUOTED_STRING           = 8
	MapCSSParserPOSITIVE_INT             = 9
	MapCSSParserNEGATIVE_INT             = 10
	MapCSSParserPOSITIVE_FLOAT           = 11
	MapCSSParserNEGATIVE_FLOAT           = 12
	MapCSSParserPROP_ANTIALIASING        = 13
	MapCSSParserPROP_ANTIALIASING_VALUES = 14
	MapCSSParserPROP_FILL_OPACITY        = 15
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MapCSSParserT__0 {
		{
			p.SetState(14)
			p.Entry()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(20)
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
		p.SetState(22)
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
		p.SetState(24)
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
		p.SetState(26)
		p.Match(MapCSSParserT__0)
	}
	{
		p.SetState(27)
		p.Match(MapCSSParserLBRACE)
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MapCSSParserPROP_ANTIALIASING || _la == MapCSSParserPROP_FILL_OPACITY {
		{
			p.SetState(28)
			p.Canvas_declaration()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(33)
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

func (s *Canvas_declarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MapCSSParserSEMICOLON, 0)
}

func (s *Canvas_declarationContext) Antialiasing() IAntialiasingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAntialiasingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAntialiasingContext)
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

	p.SetState(47)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MapCSSParserPROP_ANTIALIASING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.Match(MapCSSParserPROP_ANTIALIASING)
		}
		{
			p.SetState(36)
			p.Match(MapCSSParserCOLON)
		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MapCSSParserPROP_ANTIALIASING_VALUES {
			{
				p.SetState(37)
				p.Antialiasing()
			}

		}
		{
			p.SetState(40)
			p.Match(MapCSSParserSEMICOLON)
		}

	case MapCSSParserPROP_FILL_OPACITY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.Match(MapCSSParserPROP_FILL_OPACITY)
		}
		{
			p.SetState(42)
			p.Match(MapCSSParserCOLON)
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MapCSSParserPOSITIVE_INT || _la == MapCSSParserPOSITIVE_FLOAT {
			{
				p.SetState(43)
				p.Fill_opacity()
			}

		}
		{
			p.SetState(46)
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

func (s *AntialiasingContext) PROP_ANTIALIASING_VALUES() antlr.TerminalNode {
	return s.GetToken(MapCSSParserPROP_ANTIALIASING_VALUES, 0)
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
		p.SetState(49)
		p.Match(MapCSSParserPROP_ANTIALIASING_VALUES)
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
		p.SetState(51)
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
