package sheet_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uplang/tsvsheet.go/internal/constants"
	"github.com/uplang/tsvsheet.go/internal/sheet"
)

// parse is a test helper that parses a sheet, failing on error.
func parse(t *testing.T, src string) sheet.Sheet {
	t.Helper()
	s, err := sheet.Parse([]byte(src))
	require.NoError(t, err)
	return s
}

func TestCheck_Clean(t *testing.T) {
	t.Parallel()

	assert.Empty(t, sheet.Check(parse(t, "1\t2\t=A1 + B1\n")))
}

func TestCheck_UnknownFunction(t *testing.T) {
	t.Parallel()

	diags := sheet.Check(parse(t, "1\t=bogus(A1)\n"))
	require.Len(t, diags, 1)
	assert.Equal(t, "B1", diags[0].Cell)
	assert.Contains(t, diags[0].Message, "bogus")
	assert.False(t, diags[0].IsFatal)
}

func TestCheck_NonA1Reference(t *testing.T) {
	t.Parallel()

	diags := sheet.Check(parse(t, "1\t=A0 + 1\n")) // A0 is not a valid A1 cell
	require.Len(t, diags, 1)
	assert.Equal(t, "B1", diags[0].Cell)
	assert.Contains(t, diags[0].Message, "A1 reference")
}

func TestExplain_Formula(t *testing.T) {
	t.Parallel()

	// C1 = A1 + B1 over 2 and 3.
	trace, err := sheet.Explain(parse(t, "2\t3\t=A1 + B1\n"), sheet.Address{Row: 0, Col: 2})
	require.NoError(t, err)
	assert.Equal(t, "C1", trace.Cell)
	assert.Equal(t, "5", trace.Value)
	assert.Equal(t, "A1 + B1", trace.Formula)
	assert.Equal(t, []sheet.TraceInput{{Ref: "A1", Value: "2"}, {Ref: "B1", Value: "3"}}, trace.Inputs)
}

func TestExplain_Literal(t *testing.T) {
	t.Parallel()

	trace, err := sheet.Explain(parse(t, "hello\t=A1\n"), sheet.Address{Row: 0, Col: 0})
	require.NoError(t, err)
	assert.Equal(t, "hello", trace.Value)
	assert.Empty(t, trace.Formula)
	assert.Empty(t, trace.Inputs)
}

func TestExplain_OutOfGrid(t *testing.T) {
	t.Parallel()

	_, err := sheet.Explain(parse(t, "1\t2\n"), sheet.Address{Row: 9, Col: 9})
	require.Error(t, err)
	assert.ErrorIs(t, err, constants.ErrNotFound)
}

// TestExplain_RendersReferenceForms drives the reference renderer through a
// formula whose trace inputs cover the letter, absolute, range, named, numeric,
// last, and wildcard render branches.
func TestExplain_RendersReferenceForms(t *testing.T) {
	t.Parallel()

	// The formula sits in a cell of its own row so the traced refs don't include
	// it. Data: A1=1 B1=2; the formula is in A2.
	const formula = `=sum(A1, $B$1, A$1, "X", [3], [,$], [0,-1], A1:B1, $, A+1, A$, A$+1, A$-1, A*, *, (A:B)1, ([0]:[2])0)`
	s := parse(t, "1\t2\n"+formula+"\n")
	trace, err := sheet.Explain(s, sheet.Address{Row: 1, Col: 0})
	require.NoError(t, err)

	refs := make([]string, len(trace.Inputs))
	for i, in := range trace.Inputs {
		refs[i] = in.Ref
	}
	assert.Equal(t, []string{
		"A1", "$B$1", "A$1", `"X"`, "[3]", "[,$]", "[0,-1]", "A1:B1", "$",
		"A+1", "A$", "A$+1", "A$-1", "A*", "*", "(A:B)1", "([0]:[2])0",
	}, refs)
}

func TestExplain_FormulaShapes(t *testing.T) {
	t.Parallel()

	// A number, a unary, and a call formula exercise RenderExpr's branches.
	number, err := sheet.Explain(parse(t, "=42\n"), sheet.Address{Row: 0, Col: 0})
	require.NoError(t, err)
	assert.Equal(t, "42", number.Formula)

	unary, err := sheet.Explain(parse(t, "5\t=-A1\n"), sheet.Address{Row: 0, Col: 1})
	require.NoError(t, err)
	assert.Equal(t, "-A1", unary.Formula)

	call, err := sheet.Explain(parse(t, "5\t=abs(A1)\n"), sheet.Address{Row: 0, Col: 1})
	require.NoError(t, err)
	assert.Equal(t, "abs(A1)", call.Formula)
}

func TestCheck_KnownFunctionAndBadRange(t *testing.T) {
	t.Parallel()

	// `if` is a known function (no diagnostic); a range with a non-A1 endpoint
	// is flagged.
	assert.Empty(t, sheet.Check(parse(t, "1\t2\t=if(A1 > B1, A1, B1)\n")))

	diags := sheet.Check(parse(t, "1\t=sum(A1:$)\n")) // To endpoint `$` is not A1
	require.Len(t, diags, 1)
	assert.Contains(t, diags[0].Message, "A1 reference")
}

func TestCheck_NumberFormulaHasNoRefs(t *testing.T) {
	t.Parallel()

	// A formula with no references or calls yields no diagnostics (walk no-ops).
	assert.Empty(t, sheet.Check(parse(t, "=1 + 2\n")))
}

func TestCheck_UnaryWithCall(t *testing.T) {
	t.Parallel()

	// A unary wrapping an unknown call exercises the walkers' unary branch.
	diags := sheet.Check(parse(t, "1\t=-bogus(A1)\n"))
	require.Len(t, diags, 1)
	assert.Contains(t, diags[0].Message, "bogus")
}

func TestCheck_GroupedRangeNotA1(t *testing.T) {
	t.Parallel()

	// A grouped range is not an A1 reference (not even a RangeRef).
	diags := sheet.Check(parse(t, "1\t=sum((A:B)1)\n"))
	require.Len(t, diags, 1)
	assert.Contains(t, diags[0].Message, "A1 reference")
}
