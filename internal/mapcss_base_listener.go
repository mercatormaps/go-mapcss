// Code generated from MapCSS.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // MapCSS

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMapCSSListener is a complete listener for a parse tree produced by MapCSSParser.
type BaseMapCSSListener struct{}

var _ MapCSSListener = &BaseMapCSSListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMapCSSListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMapCSSListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMapCSSListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMapCSSListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStylesheet is called when production stylesheet is entered.
func (s *BaseMapCSSListener) EnterStylesheet(ctx *StylesheetContext) {}

// ExitStylesheet is called when production stylesheet is exited.
func (s *BaseMapCSSListener) ExitStylesheet(ctx *StylesheetContext) {}

// EnterEntry is called when production entry is entered.
func (s *BaseMapCSSListener) EnterEntry(ctx *EntryContext) {}

// ExitEntry is called when production entry is exited.
func (s *BaseMapCSSListener) ExitEntry(ctx *EntryContext) {}

// EnterRule_ is called when production rule_ is entered.
func (s *BaseMapCSSListener) EnterRule_(ctx *Rule_Context) {}

// ExitRule_ is called when production rule_ is exited.
func (s *BaseMapCSSListener) ExitRule_(ctx *Rule_Context) {}

// EnterCanvas_declaration_block is called when production canvas_declaration_block is entered.
func (s *BaseMapCSSListener) EnterCanvas_declaration_block(ctx *Canvas_declaration_blockContext) {}

// ExitCanvas_declaration_block is called when production canvas_declaration_block is exited.
func (s *BaseMapCSSListener) ExitCanvas_declaration_block(ctx *Canvas_declaration_blockContext) {}

// EnterCanvas_declaration is called when production canvas_declaration is entered.
func (s *BaseMapCSSListener) EnterCanvas_declaration(ctx *Canvas_declarationContext) {}

// ExitCanvas_declaration is called when production canvas_declaration is exited.
func (s *BaseMapCSSListener) ExitCanvas_declaration(ctx *Canvas_declarationContext) {}

// EnterAntialiasing is called when production antialiasing is entered.
func (s *BaseMapCSSListener) EnterAntialiasing(ctx *AntialiasingContext) {}

// ExitAntialiasing is called when production antialiasing is exited.
func (s *BaseMapCSSListener) ExitAntialiasing(ctx *AntialiasingContext) {}
