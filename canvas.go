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
	FillColor    *Color
}

func (c *Canvas) parse(ctx *parser.Canvas_ruleContext) error {
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
	case *parser.Antialiasing_declContext:
		if c.Antialiasing, err = antialiasing(tt.GetV().GetText()); err != nil {
			return err
		}
	case *parser.Fill_opacity_declContext:
		if c.FillOpacity, err = zeroToOneValue(tt.GetV().GetText()); err != nil {
			return err
		}
	case *parser.ColorContext:
		color, err := parseColor(tt)
		if err != nil {
			return err
		}

		switch tt.GetParent().(type) {
		case *parser.Fill_color_declContext:
			c.FillColor = color
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

func antialiasing(str string) (Antialiasing, error) {
	if str == "" {
		return AntialiasingFull, nil
	}
	return AntialiasingString(str)
}

func zeroToOneValue(str string) (float32, error) {
	if str == "" {
		return 1, nil
	}

	val, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	} else if val < 0 || val > 1 {
		return 0, fmt.Errorf("fill-opacity invalid value '%f'", val)
	}
	return float32(val), nil
}
