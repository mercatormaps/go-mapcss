package mapcss

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// ErrorHandler is invoked when error is encounted by the parser.
type ErrorHandler struct {
	*antlr.DefaultErrorListener
	Report ErrorReporter
}

// ErrorReporter is passed an error message.
type ErrorReporter func(string)

// SyntaxError is invoked when a syntax error is encounted.
func (e *ErrorHandler) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, ex antlr.RecognitionException) {
	if e.Report != nil {
		e.Report(fmt.Sprintf("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg))
	}
}
