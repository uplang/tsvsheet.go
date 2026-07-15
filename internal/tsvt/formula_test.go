package tsvt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uplang/tsvsheet.go/internal/constants"
	"github.com/uplang/tsvsheet.go/internal/tsvt"
)

func TestParseFormula_Expression(t *testing.T) {
	t.Parallel()

	// A formula is the expression sublanguage: references, operators, and calls
	// compose into the typed Expr AST.
	expr, err := tsvt.ParseFormula("B2 + C2")
	require.NoError(t, err)
	binary, ok := expr.(tsvt.Binary)
	require.True(t, ok, "expected Binary, got %T", expr)
	assert.Equal(t, tsvt.OpAdd, binary.Op)
}

func TestParseFormula_Reference(t *testing.T) {
	t.Parallel()

	// A bare reference is a valid formula: an operand with no operator.
	expr, err := tsvt.ParseFormula("D2:D4")
	require.NoError(t, err)
	_, ok := expr.(tsvt.RefOperand)
	assert.True(t, ok, "expected RefOperand, got %T", expr)
}

func TestParseFormula_SyntaxError(t *testing.T) {
	t.Parallel()

	_, err := tsvt.ParseFormula("sum(")
	require.Error(t, err)
	assert.ErrorIs(t, err, constants.ErrSyntax)
}

func TestParseFormula_TrailingInput(t *testing.T) {
	t.Parallel()

	// A complete expression followed by more tokens is rejected — the whole
	// cell must be one formula, not a formula plus leftovers.
	_, err := tsvt.ParseFormula("1 2")
	require.Error(t, err)
	assert.ErrorIs(t, err, constants.ErrSyntax)
}
