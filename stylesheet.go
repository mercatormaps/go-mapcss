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
	Rules  []Rule
}

func Parse(r io.Reader, opts ...Option) (*Stylesheet, error) {
	conf := config{}
	for _, opt := range opts {
		opt(&conf)
	}

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	eh := &ErrorHandler{Reporter: conf.reporter}
	antlr.NewDefaultErrorListener()

	lexer := parser.NewMapCSSLexer(&antlr.FileStream{
		InputStream: antlr.NewInputStream(string(buf)),
	})
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(eh)

	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewMapCSSParser(stream)
	p.BuildParseTrees = true
	p.RemoveErrorListeners()
	p.AddErrorListener(eh)

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
	case *parser.Canvas_ruleContext:
		if err := s.Canvas.parse(tt); err != nil {
			return errors.Wrap(err, "canvas parse error")
		}
	case *parser.Rule_Context:
		r := Rule{}
		if err := r.parse(tt); err != nil {
			return errors.Wrap(err, "rule parse error")
		}
		s.Rules = append(s.Rules, r)
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

type Option func(*config)

func WithErrorReporter(r ErrorReporter) Option {
	return func(c *config) {
		c.reporter = r
	}
}

type config struct {
	reporter ErrorReporter
}
