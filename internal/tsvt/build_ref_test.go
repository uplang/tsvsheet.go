package tsvt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uplang/tsvsheet.go/internal/constants"
	"github.com/uplang/tsvsheet.go/internal/tsvt"
)

// refCol extracts the column of a single-endpoint cell reference.
func refCol(t *testing.T, ref tsvt.Reference) tsvt.Col {
	t.Helper()
	return refEndpoint(t, ref).Col
}

// refEndpoint extracts the From cell endpoint of a single-endpoint range.
func refEndpoint(t *testing.T, ref tsvt.Reference) tsvt.CellEndpoint {
	t.Helper()
	rangeRef, ok := ref.(tsvt.RangeRef)
	require.True(t, ok, "expected RangeRef, got %T", ref)
	endpoint, ok := rangeRef.From.(tsvt.CellEndpoint)
	require.True(t, ok, "expected CellEndpoint, got %T", rangeRef.From)
	return endpoint
}

func TestRef_Columns(t *testing.T) {
	t.Parallel()

	cases := map[string]tsvt.Col{
		"A":     tsvt.ColLetters{Name: "A"},
		"AA":    tsvt.ColLetters{Name: "AA"},
		"$B":    tsvt.ColLetters{Name: "B", IsAbs: true},
		"$":     tsvt.ColLast{},
		`"Sum"`: tsvt.ColNamed{Name: "Sum"},
		"[3]":   tsvt.ColIndex{Index: 3},
		"[-1]":  tsvt.ColIndex{Index: -1},
	}
	for src, want := range cases {
		t.Run(src, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, want, refCol(t, placedRef(t, src)))
		})
	}
}

func TestRef_Rows(t *testing.T) {
	t.Parallel()

	cases := map[string]tsvt.RowRef{
		"C0":    tsvt.RowBefore{N: 0},
		"C1":    tsvt.RowBefore{N: 1},
		"C+1":   tsvt.RowAfter{N: 1},
		"E*":    tsvt.RowAll{},
		"C$":    tsvt.RowLast{},
		"$F$-1": tsvt.RowLast{Offset: -1},
		"C$4":   tsvt.RowAbs{N: 4},
	}
	for src, want := range cases {
		t.Run(src, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, want, refEndpoint(t, placedRef(t, src)).Row)
		})
	}
}

func TestRef_LastRowPlusOffset(t *testing.T) {
	t.Parallel()

	// *$+1: append-a-row selector (§5.2).
	ref := placedRef(t, "*$+1")
	rangeRef, ok := ref.(tsvt.RangeRef)
	require.True(t, ok)
	selector, ok := rangeRef.From.(tsvt.RowSelector)
	require.True(t, ok)
	assert.Equal(t, tsvt.RowLast{Offset: 1}, selector.Row)
}

func TestRef_RowWildcardBare(t *testing.T) {
	t.Parallel()

	ref := placedRef(t, "*")
	rangeRef, ok := ref.(tsvt.RangeRef)
	require.True(t, ok)
	selector, ok := rangeRef.From.(tsvt.RowSelector)
	require.True(t, ok)
	assert.Nil(t, selector.Row)
}

func TestRef_NumericCell(t *testing.T) {
	t.Parallel()

	// [3,1]: column 3, one row before current.
	endpoint := refEndpoint(t, placedRef(t, "[3,1]"))
	assert.Equal(t, tsvt.ColIndex{Index: 3}, endpoint.Col)
	assert.Equal(t, tsvt.RowBefore{N: 1}, endpoint.Row)
}

func TestRef_NumericFromEnd(t *testing.T) {
	t.Parallel()

	// [-3,-5]: 3 columns from the right, 5 rows from the bottom.
	endpoint := refEndpoint(t, placedRef(t, "[-3,-5]"))
	assert.Equal(t, tsvt.ColIndex{Index: -3}, endpoint.Col)
	assert.Equal(t, tsvt.RowFromEnd{N: 5}, endpoint.Row)
}

func TestRef_NumericLastRow(t *testing.T) {
	t.Parallel()

	// [3,$]: last cell of column 3.
	endpoint := refEndpoint(t, placedRef(t, "[3,$]"))
	assert.Equal(t, tsvt.RowLast{}, endpoint.Row)
}

func TestRef_NumericElidedColumnAppend(t *testing.T) {
	t.Parallel()

	// [,$+1]: append a row, column elided.
	endpoint := refEndpoint(t, placedRef(t, "[,$+1]"))
	assert.Equal(t, tsvt.ColElided{}, endpoint.Col)
	assert.Equal(t, tsvt.RowLast{Offset: 1}, endpoint.Row)
}

func TestRef_Matrix(t *testing.T) {
	t.Parallel()

	// C1:E3 — a two-endpoint matrix.
	rangeRef, ok := placedRef(t, "C1:E3").(tsvt.RangeRef)
	require.True(t, ok)
	require.NotNil(t, rangeRef.To)

	from, ok := rangeRef.From.(tsvt.CellEndpoint)
	require.True(t, ok)
	assert.Equal(t, tsvt.ColLetters{Name: "C"}, from.Col)
	assert.Equal(t, tsvt.RowBefore{N: 1}, from.Row)

	to, ok := rangeRef.To.(tsvt.CellEndpoint)
	require.True(t, ok)
	assert.Equal(t, tsvt.ColLetters{Name: "E"}, to.Col)
	assert.Equal(t, tsvt.RowBefore{N: 3}, to.Row)
}

func TestRef_AbsoluteMatrix(t *testing.T) {
	t.Parallel()

	// $B$1:$F$-1 (§5.3).
	rangeRef, ok := placedRef(t, "$B$1:$F$-1").(tsvt.RangeRef)
	require.True(t, ok)
	from, ok := rangeRef.From.(tsvt.CellEndpoint)
	require.True(t, ok)
	assert.Equal(t, tsvt.ColLetters{Name: "B", IsAbs: true}, from.Col)
	assert.Equal(t, tsvt.RowAbs{N: 1}, from.Row)
}

func TestRef_NumericMatrix(t *testing.T) {
	t.Parallel()

	// [3,1]:[5,3].
	rangeRef, ok := placedRef(t, "[3,1]:[5,3]").(tsvt.RangeRef)
	require.True(t, ok)
	require.NotNil(t, rangeRef.To)
}

func TestRef_GroupedRangeLetters(t *testing.T) {
	t.Parallel()

	// (C:E)1 — columns C, D, E, one row before current (§5.3).
	grouped, ok := placedRef(t, "(C:E)1").(tsvt.GroupedRange)
	require.True(t, ok)
	assert.Equal(t, tsvt.ColLetters{Name: "C"}, grouped.FromCol)
	assert.Equal(t, tsvt.ColLetters{Name: "E"}, grouped.ToCol)
	assert.Equal(t, tsvt.RowBefore{N: 1}, grouped.Row)
}

func TestRef_GroupedRangeNoRow(t *testing.T) {
	t.Parallel()

	grouped, ok := placedRef(t, "(C:E)").(tsvt.GroupedRange)
	require.True(t, ok)
	assert.Nil(t, grouped.Row)
}

func TestRef_GroupedRangeNumeric(t *testing.T) {
	t.Parallel()

	// ([3]:[5])1.
	grouped, ok := placedRef(t, "([3]:[5])1").(tsvt.GroupedRange)
	require.True(t, ok)
	assert.Equal(t, tsvt.ColIndex{Index: 3}, grouped.FromCol)
	assert.Equal(t, tsvt.ColIndex{Index: 5}, grouped.ToCol)
}

func TestRef_GroupedRangeNumericRejectsRow(t *testing.T) {
	t.Parallel()

	// A row inside a grouped-range column is rejected.
	_, err := tsvt.Parse(tsvt.Source("([3,1]:[5])"))
	require.Error(t, err)
	assert.ErrorIs(t, err, constants.ErrSyntax)
}
