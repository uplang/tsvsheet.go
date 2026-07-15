package sheet_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uplang/tsvsheet.go/internal/sheet"
)

func TestArray_SequenceSpills(t *testing.T) {
	t.Parallel()

	g := compute(t, "=sequence(2, 3)\n")
	require.Len(t, g, 2)
	assert.Equal(t, []string{"1", "2", "3"}, g[0]) // spills right
	assert.Equal(t, []string{"4", "5", "6"}, g[1]) // and down (grid grew)
}

func TestArray_SequenceSingleArg(t *testing.T) {
	t.Parallel()

	g := compute(t, "=sequence(3)\n") // one column
	assert.Equal(t, "1", cellAt(t, g, 0, 0))
	assert.Equal(t, "2", cellAt(t, g, 1, 0))
	assert.Equal(t, "3", cellAt(t, g, 2, 0))
}

func TestArray_SpillBlocked(t *testing.T) {
	t.Parallel()

	// A literal below the anchor blocks the spill.
	g := compute(t, "=sequence(3, 1)\nX\n")
	assert.Equal(t, string(sheet.ErrSpill), cellAt(t, g, 0, 0))
	assert.Equal(t, "X", cellAt(t, g, 1, 0)) // the blocker is untouched
}

func TestArray_TransposeUniqueSortFilterFlatten(t *testing.T) {
	t.Parallel()

	// Transpose a row into a column.
	tr := compute(t, "1\t2\t3\n=transpose(A1:C1)\n")
	assert.Equal(t, "1", cellAt(t, tr, 1, 0))
	assert.Equal(t, "3", cellAt(t, tr, 3, 0))

	// Unique keeps first occurrences.
	uq := compute(t, "a\na\nb\n=unique(A1:A3)\n")
	assert.Equal(t, "a", cellAt(t, uq, 3, 0))
	assert.Equal(t, "b", cellAt(t, uq, 4, 0))

	// Sort a numeric column ascending, and a text column.
	sn := compute(t, "3\n1\n2\n=sort(A1:A3)\n")
	assert.Equal(t, "1", cellAt(t, sn, 3, 0))
	assert.Equal(t, "3", cellAt(t, sn, 5, 0))
	st := compute(t, "c\na\nb\n=sort(A1:A3)\n")
	assert.Equal(t, "a", cellAt(t, st, 3, 0)) // lexicographic

	// Filter keeps rows whose condition is truthy.
	fl := compute(t, "10\t1\n20\t0\n30\t1\n=filter(A1:A3, B1:B3)\n")
	assert.Equal(t, "10", cellAt(t, fl, 3, 0))
	assert.Equal(t, "30", cellAt(t, fl, 4, 0))

	// Flatten stacks all cells into a column.
	ft := compute(t, "1\t2\n3\t4\n=flatten(A1:B2)\n")
	assert.Equal(t, "1", cellAt(t, ft, 2, 0))
	assert.Equal(t, "4", cellAt(t, ft, 5, 0))
}

func TestArray_ScalarContext(t *testing.T) {
	t.Parallel()

	// An array in a scalar context reduces to its top-left value.
	assert.Equal(t, "1", formula1(t, "sequence(2, 2) + 0"))
}

func TestArray_Errors(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"sequence()":        string(sheet.ErrValue), // arity
		"sequence(0)":       string(sheet.ErrValue), // rows < 1
		"sequence(1, 0)":    string(sheet.ErrValue), // cols < 1
		"transpose(A1, A2)": string(sheet.ErrValue), // arity
		"unique(A1, A2)":    string(sheet.ErrValue),
		"sort(A1, A2)":      string(sheet.ErrValue),
		"filter(A1)":        string(sheet.ErrValue),
		"flatten(A1, A2)":   string(sheet.ErrValue),
	}
	for expr, want := range cases {
		t.Run(expr, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, want, formula1(t, expr))
		})
	}

	// FILTER with no truthy condition is #N/A.
	assert.Equal(t, string(sheet.ErrNA),
		cellAt(t, compute(t, "10\t0\n20\t0\n=filter(A1:A2, B1:B2)\n"), 2, 0))
	// A short condition range leaves later rows unmatched.
	assert.Equal(t, "10",
		cellAt(t, compute(t, "10\t1\n20\t1\n=filter(A1:A2, B1:B1)\n"), 2, 0))

	// Non-numeric dimension arguments propagate #VALUE! (A1 holds text).
	for _, expr := range []string{"=sequence(A1)", "=sequence(1, A1)"} {
		t.Run(expr, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, "#VALUE!", cellAt(t, compute(t, "hi\t"+expr+"\n"), 0, 1))
		})
	}
}
