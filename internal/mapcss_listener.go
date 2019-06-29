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

	// EnterCanvas_declaration_block is called when entering the canvas_declaration_block production.
	EnterCanvas_declaration_block(c *Canvas_declaration_blockContext)

	// EnterCanvas_declaration is called when entering the canvas_declaration production.
	EnterCanvas_declaration(c *Canvas_declarationContext)

	// EnterAntialiasing is called when entering the antialiasing production.
	EnterAntialiasing(c *AntialiasingContext)

	// EnterFill_opacity is called when entering the fill_opacity production.
	EnterFill_opacity(c *Fill_opacityContext)

	// EnterFill_color is called when entering the fill_color production.
	EnterFill_color(c *Fill_colorContext)

	// EnterAlpha is called when entering the alpha production.
	EnterAlpha(c *AlphaContext)

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

	// ExitCanvas_declaration_block is called when exiting the canvas_declaration_block production.
	ExitCanvas_declaration_block(c *Canvas_declaration_blockContext)

	// ExitCanvas_declaration is called when exiting the canvas_declaration production.
	ExitCanvas_declaration(c *Canvas_declarationContext)

	// ExitAntialiasing is called when exiting the antialiasing production.
	ExitAntialiasing(c *AntialiasingContext)

	// ExitFill_opacity is called when exiting the fill_opacity production.
	ExitFill_opacity(c *Fill_opacityContext)

	// ExitFill_color is called when exiting the fill_color production.
	ExitFill_color(c *Fill_colorContext)

	// ExitAlpha is called when exiting the alpha production.
	ExitAlpha(c *AlphaContext)

	// ExitColor is called when exiting the color production.
	ExitColor(c *ColorContext)

	// ExitRgb_color is called when exiting the rgb_color production.
	ExitRgb_color(c *Rgb_colorContext)

	// ExitRgba_color is called when exiting the rgba_color production.
	ExitRgba_color(c *Rgba_colorContext)
}
