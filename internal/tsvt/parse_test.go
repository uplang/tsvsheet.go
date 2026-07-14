package tsvt_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uplang/tsvsheet.go/internal/constants"
	"github.com/uplang/tsvsheet.go/internal/tsvt"
)

// parse is a test helper that parses source and fails the test on error.
func parse(t *testing.T, src string) tsvt.Template {
	t.Helper()
	tmpl, err := tsvt.Parse(tsvt.Source(src))
	require.NoError(t, err)
	return tmpl
}

// firstRow returns the single Row of a one-line template.
func firstRow(t *testing.T, src string) tsvt.Row {
	t.Helper()
	tmpl := parse(t, src)
	require.Len(t, tmpl.Lines, 1)
	row, ok := tmpl.Lines[0].(tsvt.Row)
	require.True(t, ok, "expected a Row, got %T", tmpl.Lines[0])
	return row
}

// firstCell returns the single cell of a one-cell, one-line template.
func firstCell(t *testing.T, src string) tsvt.Cell {
	t.Helper()
	row := firstRow(t, src)
	require.Len(t, row.Cells, 1)
	return row.Cells[0]
}

// placedRef parses `<ref>` as a bare placement cell and returns its Reference.
func placedRef(t *testing.T, src string) tsvt.Reference {
	t.Helper()
	cell, ok := firstCell(t, src).(tsvt.PlacementCell)
	require.True(t, ok, "expected a PlacementCell, got %T", firstCell(t, src))
	return cell.Ref
}

// formulaExpr parses `=<expr>` as a positional formula and returns its Expr.
func formulaExpr(t *testing.T, expr string) tsvt.Expr {
	t.Helper()
	cell, ok := firstCell(t, "="+expr).(tsvt.FormulaCell)
	require.True(t, ok, "expected a FormulaCell, got %T", firstCell(t, "="+expr))
	return cell.Expr
}

func TestParse_Testdata(t *testing.T) {
	t.Parallel()

	files, err := filepath.Glob("testdata/*.tsvt")
	require.NoError(t, err)
	require.NotEmpty(t, files)

	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			t.Parallel()
			src, err := os.ReadFile(file) //nolint:gosec // fixed testdata path
			require.NoError(t, err)

			tmpl, err := tsvt.Parse(src)
			require.NoError(t, err)
			assert.NotEmpty(t, tmpl.Lines)
		})
	}
}

func TestParse_SyntaxError(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"unbalanced paren":  "=sum(A",
		"trailing operator": "=3 +",
		"bad numeric ref":   "[3,",
		"empty grouped":     "()",
		"lexer error":       "=\x00",
		"double section eq": "==body",
	}
	for name, src := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := tsvt.Parse(tsvt.Source(src))
			require.Error(t, err)
			assert.ErrorIs(t, err, constants.ErrSyntax)
		})
	}
}

func TestParse_FirstErrorOnly(t *testing.T) {
	t.Parallel()

	// Two malformed lines; only the first syntax error is reported.
	_, err := tsvt.Parse(tsvt.Source("=sum(\n=min("))
	require.Error(t, err)
	assert.ErrorIs(t, err, constants.ErrSyntax)
}

func TestParse_FractionalHeaderCount(t *testing.T) {
	t.Parallel()

	_, err := tsvt.Parse(tsvt.Source("=header(1.5)"))
	require.Error(t, err)
	assert.ErrorIs(t, err, constants.ErrSyntax)
}

func TestParse_FractionalRow(t *testing.T) {
	t.Parallel()

	_, err := tsvt.Parse(tsvt.Source("C1.5"))
	require.Error(t, err)
	assert.ErrorIs(t, err, constants.ErrSyntax)
}

func TestParse_Empty(t *testing.T) {
	t.Parallel()

	tmpl := parse(t, "")
	require.Len(t, tmpl.Lines, 1)
	row, ok := tmpl.Lines[0].(tsvt.Row)
	require.True(t, ok)
	assert.Empty(t, row.Cells)
}
