package mapcss

import (
	"fmt"
	"strconv"

	parser "github.com/mercatormaps/go-mapcss/internal"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Canvas struct {
	Antialiasing *Antialiasing
	FillOpacity  *float32
	FillColor    *Color
}

func (c *Canvas) parse(ctx *parser.Canvas_ruleContext) error {
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

func antialiasing(str string) (*Antialiasing, error) {
	val := AntialiasingFull
	if str == "" {
		return &val, nil
	}

	var err error
	val, err = AntialiasingString(str)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func zeroToOneValue(str string) (*float32, error) {
	var val float32 = 1
	if str == "" {
		return &val, nil
	}

	v, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return nil, err
	} else if v < 0 || v > 1 {
		return nil, fmt.Errorf("fill-opacity invalid value '%f'", val)
	}

	val = float32(v)
	return &val, nil
}
