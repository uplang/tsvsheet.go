package tsvt

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/uplang/tsvsheet.go/internal/constants"
	grammar "github.com/uplang/tsvsheet.go/src/grammar/tsvsheet"
)

// Parse turns template source into its typed AST, or constants.ErrSyntax
// carrying line, column, and message detail. It never prints and never
// returns a partial tree.
func Parse(src Source) (Template, error) {
	collector := &errorCollector{}

	lexer := grammar.NewTsvsheetLexer(antlr.NewInputStream(string(src)))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(collector)

	parser := grammar.NewTsvsheetParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	parser.RemoveErrorListeners()
	parser.AddErrorListener(collector)

	worksheet := parser.Worksheet()
	if collector.err != nil {
		return Template{}, collector.err
	}
	return buildTemplate(worksheet)
}

// errorCollector records the first syntax error as a sentinel. It has pointer
// receivers because antlr's ErrorListener contract mutates listener state and
// the antlr API requires the interface implementation it is handed.
type errorCollector struct {
	antlr.DefaultErrorListener
	err error
}

// SyntaxError implements antlr.ErrorListener, converting the report into
// constants.ErrSyntax; only the first error is kept.
func (c *errorCollector) SyntaxError(_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException) {
	if c.err == nil {
		c.err = constants.ErrSyntax.With(nil, "line", line, "column", column, "message", msg)
	}
}
