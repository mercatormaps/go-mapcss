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

// EnterInt_ is called when production int_ is entered.
func (s *BaseMapCSSListener) EnterInt_(ctx *Int_Context) {}

// ExitInt_ is called when production int_ is exited.
func (s *BaseMapCSSListener) ExitInt_(ctx *Int_Context) {}

// EnterUint_ is called when production uint_ is entered.
func (s *BaseMapCSSListener) EnterUint_(ctx *Uint_Context) {}

// ExitUint_ is called when production uint_ is exited.
func (s *BaseMapCSSListener) ExitUint_(ctx *Uint_Context) {}

// EnterFloat_ is called when production float_ is entered.
func (s *BaseMapCSSListener) EnterFloat_(ctx *Float_Context) {}

// ExitFloat_ is called when production float_ is exited.
func (s *BaseMapCSSListener) ExitFloat_(ctx *Float_Context) {}

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

// EnterSelector is called when production selector is entered.
func (s *BaseMapCSSListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseMapCSSListener) ExitSelector(ctx *SelectorContext) {}

// EnterZoom is called when production zoom is entered.
func (s *BaseMapCSSListener) EnterZoom(ctx *ZoomContext) {}

// ExitZoom is called when production zoom is exited.
func (s *BaseMapCSSListener) ExitZoom(ctx *ZoomContext) {}

// EnterZoom_range is called when production zoom_range is entered.
func (s *BaseMapCSSListener) EnterZoom_range(ctx *Zoom_rangeContext) {}

// ExitZoom_range is called when production zoom_range is exited.
func (s *BaseMapCSSListener) ExitZoom_range(ctx *Zoom_rangeContext) {}

// EnterMin_zoom is called when production min_zoom is entered.
func (s *BaseMapCSSListener) EnterMin_zoom(ctx *Min_zoomContext) {}

// ExitMin_zoom is called when production min_zoom is exited.
func (s *BaseMapCSSListener) ExitMin_zoom(ctx *Min_zoomContext) {}

// EnterExact_zoom is called when production exact_zoom is entered.
func (s *BaseMapCSSListener) EnterExact_zoom(ctx *Exact_zoomContext) {}

// ExitExact_zoom is called when production exact_zoom is exited.
func (s *BaseMapCSSListener) ExitExact_zoom(ctx *Exact_zoomContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseMapCSSListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseMapCSSListener) ExitAttribute(ctx *AttributeContext) {}

// EnterDecl_block is called when production decl_block is entered.
func (s *BaseMapCSSListener) EnterDecl_block(ctx *Decl_blockContext) {}

// ExitDecl_block is called when production decl_block is exited.
func (s *BaseMapCSSListener) ExitDecl_block(ctx *Decl_blockContext) {}

// EnterCanvas_rule is called when production canvas_rule is entered.
func (s *BaseMapCSSListener) EnterCanvas_rule(ctx *Canvas_ruleContext) {}

// ExitCanvas_rule is called when production canvas_rule is exited.
func (s *BaseMapCSSListener) ExitCanvas_rule(ctx *Canvas_ruleContext) {}

// EnterCanvas_decl_block is called when production canvas_decl_block is entered.
func (s *BaseMapCSSListener) EnterCanvas_decl_block(ctx *Canvas_decl_blockContext) {}

// ExitCanvas_decl_block is called when production canvas_decl_block is exited.
func (s *BaseMapCSSListener) ExitCanvas_decl_block(ctx *Canvas_decl_blockContext) {}

// EnterCanvas_decl is called when production canvas_decl is entered.
func (s *BaseMapCSSListener) EnterCanvas_decl(ctx *Canvas_declContext) {}

// ExitCanvas_decl is called when production canvas_decl is exited.
func (s *BaseMapCSSListener) ExitCanvas_decl(ctx *Canvas_declContext) {}

// EnterAntialiasing_decl is called when production antialiasing_decl is entered.
func (s *BaseMapCSSListener) EnterAntialiasing_decl(ctx *Antialiasing_declContext) {}

// ExitAntialiasing_decl is called when production antialiasing_decl is exited.
func (s *BaseMapCSSListener) ExitAntialiasing_decl(ctx *Antialiasing_declContext) {}

// EnterFill_opacity_decl is called when production fill_opacity_decl is entered.
func (s *BaseMapCSSListener) EnterFill_opacity_decl(ctx *Fill_opacity_declContext) {}

// ExitFill_opacity_decl is called when production fill_opacity_decl is exited.
func (s *BaseMapCSSListener) ExitFill_opacity_decl(ctx *Fill_opacity_declContext) {}

// EnterFill_color_decl is called when production fill_color_decl is entered.
func (s *BaseMapCSSListener) EnterFill_color_decl(ctx *Fill_color_declContext) {}

// ExitFill_color_decl is called when production fill_color_decl is exited.
func (s *BaseMapCSSListener) ExitFill_color_decl(ctx *Fill_color_declContext) {}

// EnterColor is called when production color is entered.
func (s *BaseMapCSSListener) EnterColor(ctx *ColorContext) {}

// ExitColor is called when production color is exited.
func (s *BaseMapCSSListener) ExitColor(ctx *ColorContext) {}

// EnterRgb_color is called when production rgb_color is entered.
func (s *BaseMapCSSListener) EnterRgb_color(ctx *Rgb_colorContext) {}

// ExitRgb_color is called when production rgb_color is exited.
func (s *BaseMapCSSListener) ExitRgb_color(ctx *Rgb_colorContext) {}

// EnterRgba_color is called when production rgba_color is entered.
func (s *BaseMapCSSListener) EnterRgba_color(ctx *Rgba_colorContext) {}

// ExitRgba_color is called when production rgba_color is exited.
func (s *BaseMapCSSListener) ExitRgba_color(ctx *Rgba_colorContext) {}
