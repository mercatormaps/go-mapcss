package mapcss

import (
	"fmt"
	"strconv"

	parser "github.com/mercatormaps/go-mapcss/internal"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Canvas struct {
	Antialiasing Antialiasing
	FillOpacity  float32
}

func (c *Canvas) parse(ctx *parser.Canvas_declaration_blockContext) error {
	c.Antialiasing = AntialiasingFull
	c.FillOpacity = 1

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
		if c.Antialiasing, err = AntialiasingString(tt.GetText()); err != nil {
			return err
		}
	case *parser.Fill_opacityContext:
		v, err := strconv.ParseFloat(tt.GetText(), 32)
		if err != nil {
			return err
		} else if v > 1 {
			return fmt.Errorf("fill-opacity invalid value '%f'", v)
		}
		c.FillOpacity = float32(v)
	default:
		for i := 0; i < tt.GetChildCount(); i++ {
			if err := c.walk(listener, tt.GetChild(i)); err != nil {
				return err
			}
		}
	}

	return nil
}
