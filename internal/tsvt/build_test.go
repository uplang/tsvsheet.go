package tsvt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uplang/tsvsheet.go/internal/tsvt"
)

func TestBuild_SectionMarkers(t *testing.T) {
	t.Parallel()

	tmpl := parse(t, "=header(2)\n=body\n=final")
	require.Len(t, tmpl.Lines, 3)

	header, ok := tmpl.Lines[0].(tsvt.HeaderMarker)
	require.True(t, ok)
	assert.Equal(t, 2, header.Count)
	assert.Equal(t, tsvt.LineNumber(1), header.At)

	_, ok = tmpl.Lines[1].(tsvt.BodyMarker)
	assert.True(t, ok)

	final, ok := tmpl.Lines[2].(tsvt.FinalMarker)
	require.True(t, ok)
	assert.Equal(t, tsvt.LineNumber(3), final.At)
}

func TestBuild_StructuralCommand(t *testing.T) {
	t.Parallel()

	tmpl := parse(t, "=A<")
	require.Len(t, tmpl.Lines, 1)
	structural, ok := tmpl.Lines[0].(tsvt.Structural)
	require.True(t, ok)
	assert.Equal(t, tsvt.ModMove, structural.Mod)
	assert.Equal(t, tsvt.ColLetters{Name: "A"}, refCol(t, structural.Ref))
}

func TestBuild_StructuralModifiers(t *testing.T) {
	t.Parallel()

	cases := map[string]tsvt.Modifier{
		"=A>": tsvt.ModShift,
		"=A<": tsvt.ModMove,
		"=A!": tsvt.ModDelete,
	}
	for src, want := range cases {
		t.Run(src, func(t *testing.T) {
			t.Parallel()
			structural, ok := parse(t, src).Lines[0].(tsvt.Structural)
			require.True(t, ok)
			assert.Equal(t, want, structural.Mod)
		})
	}
}

func TestBuild_EmptyCells(t *testing.T) {
	t.Parallel()

	// Leading, interior, and trailing empty cells preserve column positions.
	row := firstRow(t, "\tA\t\tB\t")
	require.Len(t, row.Cells, 5)
	assert.IsType(t, tsvt.EmptyCell{}, row.Cells[0])
	assert.IsType(t, tsvt.PlacementCell{}, row.Cells[1])
	assert.IsType(t, tsvt.EmptyCell{}, row.Cells[2])
	assert.IsType(t, tsvt.PlacementCell{}, row.Cells[3])
	assert.IsType(t, tsvt.EmptyCell{}, row.Cells[4])
}

func TestBuild_FormulaCell(t *testing.T) {
	t.Parallel()

	cell, ok := firstCell(t, "=C + D").(tsvt.FormulaCell)
	require.True(t, ok)
	assert.IsType(t, tsvt.Binary{}, cell.Expr)
}

func TestBuild_LiteralCells(t *testing.T) {
	t.Parallel()

	cases := map[string]tsvt.Literal{
		"Total": {Kind: tsvt.LiteralName, Text: "Total"},
		"42":    {Kind: tsvt.LiteralNumber, Text: "42"},
	}
	for src, want := range cases {
		t.Run(src, func(t *testing.T) {
			t.Parallel()
			cell, ok := firstCell(t, src).(tsvt.LiteralCell)
			require.True(t, ok)
			assert.Equal(t, want, cell.Value)
		})
	}
}

func TestBuild_PlacementFormulaPayload(t *testing.T) {
	t.Parallel()

	cell, ok := firstCell(t, "E=C + D").(tsvt.PlacementCell)
	require.True(t, ok)
	assert.Equal(t, tsvt.ModNone, cell.Mod)
	payload, ok := cell.Payload.(tsvt.FormulaPayload)
	require.True(t, ok)
	assert.IsType(t, tsvt.Binary{}, payload.Expr)
}

func TestBuild_PlacementLiteralPayload(t *testing.T) {
	t.Parallel()

	// A$+1=Total: the leading = is the separator; Total is a literal (§11.1).
	cell, ok := firstCell(t, "A$+1=Total").(tsvt.PlacementCell)
	require.True(t, ok)
	payload, ok := cell.Payload.(tsvt.LiteralPayload)
	require.True(t, ok)
	assert.Equal(t, tsvt.Literal{Kind: tsvt.LiteralName, Text: "Total"}, payload.Value)
}

func TestBuild_PlacementBarewordPayloadIsLiteral(t *testing.T) {
	t.Parallel()

	// A=Total: Total is a bareword, not a valid expression, so the payload is a
	// literal (§11.1), not a formula.
	cell, ok := firstCell(t, `A=Total`).(tsvt.PlacementCell)
	require.True(t, ok)
	assert.IsType(t, tsvt.LiteralPayload{}, cell.Payload)
}

func TestBuild_PlacementModifierNoPayload(t *testing.T) {
	t.Parallel()

	cell, ok := firstCell(t, "C!").(tsvt.PlacementCell)
	require.True(t, ok)
	assert.Equal(t, tsvt.ModDelete, cell.Mod)
	assert.Nil(t, cell.Payload)
}

func TestBuild_QuotedPayloadIsFormulaReference(t *testing.T) {
	t.Parallel()

	// A="hello": "hello" is a valid expression (a named-column reference), so the
	// payload is a formula, not a literal (ADR 0003 rule 16).
	cell, ok := firstCell(t, `A="hello"`).(tsvt.PlacementCell)
	require.True(t, ok)
	payload, ok := cell.Payload.(tsvt.FormulaPayload)
	require.True(t, ok)
	ref, ok := payload.Expr.(tsvt.RefOperand)
	require.True(t, ok)
	assert.Equal(t, tsvt.ColNamed{Name: "hello"}, refCol(t, ref.Ref))
}
