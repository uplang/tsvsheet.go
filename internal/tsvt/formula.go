package tsvt

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/uplang/tsvsheet.go/internal/constants"
	grammar "github.com/uplang/tsvsheet.go/src/grammar/tsvsheet"
)

// FormulaText is the source of a single formula expression — the part of a
// spreadsheet cell after its leading `=`.
type FormulaText string

// ParseFormula parses one formula expression into the typed Expr AST, or
// constants.ErrSyntax. It reuses the ANTLR-generated expression sublanguage
// (§11) via the grammar's `expression` entry rule, so the A1 spreadsheet model
// compiles each `=formula` cell without a hand-written parser.
func ParseFormula(src FormulaText) (Expr, error) {
	collector := &errorCollector{sink: &errorSink{}}

	lexer := grammar.NewTsvsheetLexer(antlr.NewInputStream(string(src)))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(collector)

	parser := grammar.NewTsvsheetParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	parser.RemoveErrorListeners()
	parser.AddErrorListener(collector)

	expr := parser.Expression()
	if collector.sink.err != nil {
		return nil, collector.sink.err
	}
	if parser.GetCurrentToken().GetTokenType() != antlr.TokenEOF {
		return nil, constants.ErrSyntax.With(nil, "message", "unexpected input after formula")
	}
	return buildExpr(expr)
}
