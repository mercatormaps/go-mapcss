package mapcss

import (
	"fmt"
	"strconv"

	parser "github.com/mercatormaps/go-mapcss/internal"
	"github.com/pkg/errors"
)

type Color struct {
	R, G, B uint8
	A       float32
}

func parseColor(ctx *parser.ColorContext) (*Color, error) {
	if hex := ctx.HEX(); hex != nil {
		return parseHexColor(hex.GetText())
	}

	var err error
	parseUint8 := func(s string, dest *uint8) {
		if err == nil {
			var v uint64
			v, err = strconv.ParseUint(s, 10, 8)
			if v > 255 {
				err = fmt.Errorf("invalid value: %d > 255", v)
			} else {
				*dest = uint8(v)
			}
		}
	}

	if rgb := ctx.Rgb_color(); rgb != nil {
		c := &Color{A: 1}
		parseUint8(rgb.GetR().GetText(), &c.R)
		parseUint8(rgb.GetG().GetText(), &c.G)
		parseUint8(rgb.GetB().GetText(), &c.B)
		return c, err
	}

	if rgba := ctx.Rgba_color(); rgba != nil {
		c := &Color{}
		if v, err := strconv.ParseFloat(rgba.GetA().GetText(), 32); err != nil {
			return nil, err
		} else if v < 0 {
			return nil, fmt.Errorf("invalid value: %f < 0", v)
		} else if v > 1 {
			return nil, fmt.Errorf("invalid value: %f > 0", v)
		} else {
			c.A = float32(v)
		}

		parseUint8(rgba.GetR().GetText(), &c.R)
		parseUint8(rgba.GetG().GetText(), &c.G)
		parseUint8(rgba.GetB().GetText(), &c.B)
		return c, err
	}

	return nil, fmt.Errorf("unknown color format")
}

func parseHexColor(hex string) (*Color, error) {
	var err error
	rgb := func(hex string, dest *uint8) {
		if err == nil {
			var v uint8
			v, err = hexToUint8(hex)
			*dest = v
		}
	}

	alpha := func(hex string, dest *float32) {
		if err == nil {
			var v uint8
			v, err = hexToUint8(hex)
			*dest = float32(v) / 255
		}
	}

	c := &Color{A: 1}
	switch len(hex) {
	case 5:
		alpha(hex[4:5]+hex[4:5], &c.A)
		fallthrough
	case 4:
		rgb(hex[1:2]+hex[1:2], &c.R)
		rgb(hex[2:3]+hex[2:3], &c.G)
		rgb(hex[3:4]+hex[3:4], &c.B)
	case 9:
		alpha(hex[7:9], &c.A)
		fallthrough
	case 7:
		rgb(hex[1:3], &c.R)
		rgb(hex[3:5], &c.G)
		rgb(hex[5:7], &c.B)
	default:
		return nil, fmt.Errorf("unknown hex color '%s'", hex)
	}

	return c, err
}

func hexToUint8(hex string) (uint8, error) {
	v, err := strconv.ParseUint(hex, 16, 8)
	if err != nil {
		return 0, errors.Wrapf(err, "invalid hex '%s'", hex)
	} else if v > 255 {
		return 0, errors.Wrapf(err, "invalid value: %d > 255", v)
	}
	return uint8(v), nil
}
