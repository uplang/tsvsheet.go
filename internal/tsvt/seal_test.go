package tsvt

import "testing"

// TestSeal exercises the sealed-interface marker methods. They have empty
// bodies and exist only to bound each interface's variant set at compile time,
// so nothing else ever calls them; invoking one representative of each keeps
// the markers covered.
func TestSeal(t *testing.T) {
	t.Parallel()

	HeaderMarker{}.isLine()
	EmptyCell{}.isCell()
	FormulaPayload{}.isPayload()
	RangeRef{}.isReference()
	CellEndpoint{}.isEndpoint()
	ColLast{}.isCol()
	RowAll{}.isRow()
	Number{}.isExpr()
}
