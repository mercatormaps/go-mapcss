package mapcss

import parser "github.com/mercatormaps/go-mapcss/internal"

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
	Attributes []Attribute
}

func (s *Selector) parse(ctx *parser.SelectorContext) error {
	s.Type = ctx.GetTyp().GetText()

	for _, ctx := range ctx.AllAttribute() {
		attr := Attribute{}
		if err := attr.parse(ctx.(*parser.AttributeContext)); err != nil {
			return err
		}
		s.Attributes = append(s.Attributes, attr)
	}
	return nil
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
