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
	parser *parser.MapCSSParser
}

func NewStylesheet(r io.Reader) (*Stylesheet, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	lex := parser.NewMapCSSLexer(&antlr.FileStream{
		InputStream: antlr.NewInputStream(string(buf)),
	})

	stream := antlr.NewCommonTokenStream(lex, 0)
	return &Stylesheet{
		parser: parser.NewMapCSSParser(stream),
	}, nil
}

func (p *Stylesheet) Parse() (err error) {
	p.parser.BuildParseTrees = true
	tree := p.parser.Stylesheet()
	listener := &parser.BaseMapCSSListener{}
	return p.walk(listener, tree)
}

func (p *Stylesheet) walk(listener *parser.BaseMapCSSListener, t antlr.Tree) error {
	ctx, err := walk(listener, t)
	if err != nil {
		return err
	} else if ctx == nil {
		return nil
	}

	switch tt := ctx.(type) {
	case *parser.Canvas_declaration_blockContext:
		canvas := &Canvas{}
		if err := canvas.Parse(tt); err != nil {
			return errors.Wrap(err, "canvas parse error")
		}
	default:
		for i := 0; i < tt.GetChildCount(); i++ {
			if err := p.walk(listener, tt.GetChild(i)); err != nil {
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
