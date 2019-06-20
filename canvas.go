package mapcss

import (
	parser "github.com/mercatormaps/go-mapcss/internal"
	"github.com/mercatormaps/go-mapcss/property"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Canvas struct {
	property.Antialiasing
}

func (c *Canvas) Parse(ctx *parser.Canvas_declaration_blockContext) error {
	c.Antialiasing = property.AntialiasingFull
	return c.walk(&parser.BaseMapCSSListener{}, ctx)
}

func (c *Canvas) walk(listener *parser.BaseMapCSSListener, t antlr.Tree) error {
	ctx, err := walk(listener, t)
	if err != nil {
		return err
	} else if ctx == nil {
		return nil
	}

	switch tt := ctx.(type) {
	case *parser.AntialiasingContext:
		if c.Antialiasing, err = property.AntialiasingString(tt.GetText()); err != nil {
			return err
		}
	default:
		for i := 0; i < tt.GetChildCount(); i++ {
			if err := c.walk(listener, tt.GetChild(i)); err != nil {
				return err
			}
		}
	}

	return nil
}
