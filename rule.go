package mapcss

import (
	"fmt"
	"strconv"

	parser "github.com/mercatormaps/go-mapcss/internal"
)

type Rule struct {
	Selectors []Selector
}

func (r *Rule) parse(ctx *parser.Rule_Context) error {
	for _, ctx := range ctx.AllSelector() {
		sel := Selector{}
		if err := sel.parse(ctx.(*parser.SelectorContext)); err != nil {
			return err
		}
		r.Selectors = append(r.Selectors, sel)
	}
	return nil
}

type Selector struct {
	Type       string
	Zoom       *Zoom
	Attributes []Attribute
}

func (s *Selector) parse(ctx *parser.SelectorContext) error {
	s.Type = ctx.GetTyp().GetText()

	if zoom := ctx.Zoom(); zoom != nil {
		z := Zoom{}
		if err := z.parse(zoom.(*parser.ZoomContext)); err != nil {
			return err
		}
		s.Zoom = &z
	}

	for _, ctx := range ctx.AllAttribute() {
		attr := Attribute{}
		if err := attr.parse(ctx.(*parser.AttributeContext)); err != nil {
			return err
		}
		s.Attributes = append(s.Attributes, attr)
	}
	return nil
}

type Zoom struct {
	Type ZoomType
	Min  uint
	Max  uint
}

type ZoomType int

const (
	ZoomRange ZoomType = iota
	MinZoom
	ExactZoom
)

func (z *Zoom) parse(ctx *parser.ZoomContext) error {
	var err error
	parse := func(s string, dest *uint) {
		if err == nil {
			var v uint64
			v, err = strconv.ParseUint(s, 10, 32)
			*dest = uint(v)
		}
	}

	if zoom := ctx.Zoom_range(); zoom != nil {
		z.Type = ZoomRange
		parse(zoom.GetMin().GetText(), &z.Min)
		parse(zoom.GetMax().GetText(), &z.Max)
	} else if zoom := ctx.Min_zoom(); zoom != nil {
		z.Type = MinZoom
		parse(zoom.GetMin().GetText(), &z.Min)
	} else if zoom := ctx.Exact_zoom(); zoom != nil {
		z.Type = ExactZoom
		parse(zoom.GetMin().GetText(), &z.Min)
		z.Max = z.Min
	} else {
		return fmt.Errorf("unknown zoom type")
	}
	return err
}

type Attribute struct {
	Name string
	Not  bool
}

func (a *Attribute) parse(ctx *parser.AttributeContext) error {
	a.Name = ctx.GetName().GetText()
	a.Not = ctx.GetNeg() != nil
	return nil
}
