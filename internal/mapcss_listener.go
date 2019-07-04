// Code generated from MapCSS.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // MapCSS

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MapCSSListener is a complete listener for a parse tree produced by MapCSSParser.
type MapCSSListener interface {
	antlr.ParseTreeListener

	// EnterStylesheet is called when entering the stylesheet production.
	EnterStylesheet(c *StylesheetContext)

	// EnterEntry is called when entering the entry production.
	EnterEntry(c *EntryContext)

	// EnterRule_ is called when entering the rule_ production.
	EnterRule_(c *Rule_Context)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterDecl_block is called when entering the decl_block production.
	EnterDecl_block(c *Decl_blockContext)

	// EnterCanvas_rule is called when entering the canvas_rule production.
	EnterCanvas_rule(c *Canvas_ruleContext)

	// EnterCanvas_decl_block is called when entering the canvas_decl_block production.
	EnterCanvas_decl_block(c *Canvas_decl_blockContext)

	// EnterCanvas_decl is called when entering the canvas_decl production.
	EnterCanvas_decl(c *Canvas_declContext)

	// EnterAntialiasing_decl is called when entering the antialiasing_decl production.
	EnterAntialiasing_decl(c *Antialiasing_declContext)

	// EnterFill_opacity_decl is called when entering the fill_opacity_decl production.
	EnterFill_opacity_decl(c *Fill_opacity_declContext)

	// EnterFill_color_decl is called when entering the fill_color_decl production.
	EnterFill_color_decl(c *Fill_color_declContext)

	// EnterColor is called when entering the color production.
	EnterColor(c *ColorContext)

	// EnterRgb_color is called when entering the rgb_color production.
	EnterRgb_color(c *Rgb_colorContext)

	// EnterRgba_color is called when entering the rgba_color production.
	EnterRgba_color(c *Rgba_colorContext)

	// ExitStylesheet is called when exiting the stylesheet production.
	ExitStylesheet(c *StylesheetContext)

	// ExitEntry is called when exiting the entry production.
	ExitEntry(c *EntryContext)

	// ExitRule_ is called when exiting the rule_ production.
	ExitRule_(c *Rule_Context)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitDecl_block is called when exiting the decl_block production.
	ExitDecl_block(c *Decl_blockContext)

	// ExitCanvas_rule is called when exiting the canvas_rule production.
	ExitCanvas_rule(c *Canvas_ruleContext)

	// ExitCanvas_decl_block is called when exiting the canvas_decl_block production.
	ExitCanvas_decl_block(c *Canvas_decl_blockContext)

	// ExitCanvas_decl is called when exiting the canvas_decl production.
	ExitCanvas_decl(c *Canvas_declContext)

	// ExitAntialiasing_decl is called when exiting the antialiasing_decl production.
	ExitAntialiasing_decl(c *Antialiasing_declContext)

	// ExitFill_opacity_decl is called when exiting the fill_opacity_decl production.
	ExitFill_opacity_decl(c *Fill_opacity_declContext)

	// ExitFill_color_decl is called when exiting the fill_color_decl production.
	ExitFill_color_decl(c *Fill_color_declContext)

	// ExitColor is called when exiting the color production.
	ExitColor(c *ColorContext)

	// ExitRgb_color is called when exiting the rgb_color production.
	ExitRgb_color(c *Rgb_colorContext)

	// ExitRgba_color is called when exiting the rgba_color production.
	ExitRgba_color(c *Rgba_colorContext)
}
