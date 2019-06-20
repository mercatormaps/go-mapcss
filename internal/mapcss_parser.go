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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 43, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 7, 2, 16, 10, 2, 12, 2, 14, 2, 19, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 6, 5, 30, 10, 5, 13, 5, 14, 5, 31, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8,
	10, 12, 2, 2, 2, 38, 2, 17, 3, 2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 24, 3, 2,
	2, 2, 8, 26, 3, 2, 2, 2, 10, 35, 3, 2, 2, 2, 12, 40, 3, 2, 2, 2, 14, 16,
	5, 4, 3, 2, 15, 14, 3, 2, 2, 2, 16, 19, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2,
	17, 18, 3, 2, 2, 2, 18, 20, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 20, 21, 7,
	2, 2, 3, 21, 3, 3, 2, 2, 2, 22, 23, 5, 6, 4, 2, 23, 5, 3, 2, 2, 2, 24,
	25, 5, 8, 5, 2, 25, 7, 3, 2, 2, 2, 26, 27, 7, 3, 2, 2, 27, 29, 7, 5, 2,
	2, 28, 30, 5, 10, 6, 2, 29, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 29,
	3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 34, 7, 6, 2, 2,
	34, 9, 3, 2, 2, 2, 35, 36, 7, 11, 2, 2, 36, 37, 7, 7, 2, 2, 37, 38, 5,
	12, 7, 2, 38, 39, 7, 8, 2, 2, 39, 11, 3, 2, 2, 2, 40, 41, 7, 12, 2, 2,
	41, 13, 3, 2, 2, 2, 4, 17, 31,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'canvas'", "", "'{'", "'}'", "':'", "';'", "", "", "'antialiasing'",
}
var symbolicNames = []string{
	"", "", "WS", "LBRACE", "RBRACE", "COLON", "SEMICOLON", "DQUOTED_STRING",
	"SQUOTED_STRING", "PROP_ANTIALIASING", "PROP_ANTIALIASING_VALUES",
}

var ruleNames = []string{
	"stylesheet", "entry", "rule_", "canvas_declaration_block", "canvas_declaration",
	"antialiasing",
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
	MapCSSParserPROP_ANTIALIASING        = 9
	MapCSSParserPROP_ANTIALIASING_VALUES = 10
)

// MapCSSParser rules.
const (
	MapCSSParserRULE_stylesheet               = 0
	MapCSSParserRULE_entry                    = 1
	MapCSSParserRULE_rule_                    = 2
	MapCSSParserRULE_canvas_declaration_block = 3
	MapCSSParserRULE_canvas_declaration       = 4
	MapCSSParserRULE_antialiasing             = 5
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MapCSSParserT__0 {
		{
			p.SetState(12)
			p.Entry()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(18)
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
		p.SetState(20)
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
		p.SetState(22)
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
		p.SetState(24)
		p.Match(MapCSSParserT__0)
	}
	{
		p.SetState(25)
		p.Match(MapCSSParserLBRACE)
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MapCSSParserPROP_ANTIALIASING {
		{
			p.SetState(26)
			p.Canvas_declaration()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(MapCSSParserPROP_ANTIALIASING)
	}
	{
		p.SetState(34)
		p.Match(MapCSSParserCOLON)
	}
	{
		p.SetState(35)
		p.Antialiasing()
	}
	{
		p.SetState(36)
		p.Match(MapCSSParserSEMICOLON)
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
		p.SetState(38)
		p.Match(MapCSSParserPROP_ANTIALIASING_VALUES)
	}

	return localctx
}
