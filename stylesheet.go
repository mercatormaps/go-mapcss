package mapcss

import (
	"fmt"
	"io"
	"io/ioutil"

	parser "github.com/mercatormaps/go-mapcss/internal"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pkg/errors"
)

type Stylesheet struct {
	Canvas Canvas
}

func Parse(r io.Reader) (*Stylesheet, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	lex := parser.NewMapCSSLexer(&antlr.FileStream{
		InputStream: antlr.NewInputStream(string(buf)),
	})

	stream := antlr.NewCommonTokenStream(lex, 0)
	p := parser.NewMapCSSParser(stream)
	p.BuildParseTrees = true

	tree := p.Stylesheet()
	listener := &parser.BaseMapCSSListener{}

	s := &Stylesheet{}
	return s, s.walk(listener, tree)
}

func (s *Stylesheet) walk(listener *parser.BaseMapCSSListener, t antlr.Tree) error {
	ctx, err := walk(listener, t)
	if err != nil {
		return err
	} else if ctx == nil {
		return nil
	}

	switch tt := ctx.(type) {
	case *parser.Canvas_declaration_blockContext:
		if err := s.Canvas.parse(tt); err != nil {
			return errors.Wrap(err, "canvas parse error")
		}
	default:
		for i := 0; i < tt.GetChildCount(); i++ {
			if err := s.walk(listener, tt.GetChild(i)); err != nil {
				return err
			}
		}
	}

	return nil
}

func walk(listener *parser.BaseMapCSSListener, t antlr.Tree) (antlr.ParserRuleContext, error) {
	switch tt := t.(type) {
	case antlr.ErrorNode:
		listener.VisitErrorNode(tt)
		return nil, fmt.Errorf(tt.GetText())
	case antlr.TerminalNode:
		listener.VisitTerminal(tt)
		return nil, nil
	default:
		return t.(antlr.RuleNode).GetRuleContext().(antlr.ParserRuleContext), nil
	}
}
