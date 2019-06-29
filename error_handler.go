package mapcss

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ErrorHandler struct {
	*antlr.DefaultErrorListener

	Reporter ErrorReporter
}

type ErrorReporter func(string)

func (e *ErrorHandler) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, ex antlr.RecognitionException) {
	if e.Reporter != nil {
		e.Reporter(fmt.Sprintf("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg))
	}
}
