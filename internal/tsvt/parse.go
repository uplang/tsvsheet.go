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
	collector := &errorCollector{sink: &errorSink{}}

	lexer := grammar.NewTsvsheetLexer(antlr.NewInputStream(string(src)))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(collector)

	parser := grammar.NewTsvsheetParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	parser.RemoveErrorListeners()
	parser.AddErrorListener(collector)

	worksheet := parser.Worksheet()
	if collector.sink.err != nil {
		return Template{}, collector.sink.err
	}
	return buildTemplate(worksheet)
}

// errorSink holds the first collected syntax error so an errorCollector can
// record it from a value-receiver method (the sink is shared by pointer).
type errorSink struct {
	err error
}

// errorCollector records the first syntax error as a sentinel. The mutable
// error lives behind the sink pointer, so the antlr ErrorListener callback is a
// value-receiver method whose write still persists.
type errorCollector struct {
	antlr.DefaultErrorListener
	sink *errorSink
}

// SyntaxError implements antlr.ErrorListener, converting the report into
// constants.ErrSyntax; only the first error is kept.
func (c errorCollector) SyntaxError(
	_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException,
) {
	if c.sink.err == nil {
		c.sink.err = constants.ErrSyntax.With(nil, "line", line, "column", column, "message", msg)
	}
}
