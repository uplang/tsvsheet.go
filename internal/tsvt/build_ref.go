package tsvt

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/uplang/tsvsheet.go/internal/constants"
	grammar "github.com/uplang/tsvsheet.go/src/grammar/tsvsheet"
)

// buildReference dispatches on the two reference shapes: range and grouped
// range.
func buildReference(ctx grammar.IReferenceContext) (Reference, error) {
	if rangeRef := ctx.RangeRef(); rangeRef != nil {
		return buildRangeRef(rangeRef)
	}
	return buildGroupedRange(ctx.GroupedRange())
}

// buildRangeRef converts a single endpoint or an endpoint:endpoint range.
func buildRangeRef(ctx grammar.IRangeRefContext) (Reference, error) {
	endpoints := ctx.AllEndpoint()
	from, err := buildEndpoint(endpoints[0])
	if err != nil {
		return nil, err
	}
	if len(endpoints) == 1 {
		return RangeRef{From: from}, nil
	}
	to, err := buildEndpoint(endpoints[1])
	if err != nil {
		return nil, err
	}
	return RangeRef{From: from, To: to}, nil
}

// buildEndpoint dispatches on the two endpoint shapes: letter/named cell and
// numeric.
func buildEndpoint(ctx grammar.IEndpointContext) (Endpoint, error) {
	if cellRef := ctx.CellRef(); cellRef != nil {
		return buildCellRef(cellRef)
	}
	return buildNumericEndpoint(ctx.NumericRef())
}

// buildCellRef converts a column-with-optional-row or a row wildcard.
func buildCellRef(ctx grammar.ICellRefContext) (Endpoint, error) {
	if wildcard := ctx.RowWildcard(); wildcard != nil {
		return buildRowWildcard(wildcard)
	}
	row, err := optionalRow(ctx.RowRef())
	if err != nil {
		return nil, err
	}
	return CellEndpoint{Col: buildColRef(ctx.ColRef()), Row: row}, nil
}

// buildColRef converts the four column shapes: `$B`, `A`, `$`, `"Sum"`.
func buildColRef(ctx grammar.IColRefContext) Col {
	letters := ctx.COL()
	switch {
	case letters != nil && ctx.DOLLAR() != nil:
		return ColLetters{Name: letters.GetText(), Abs: true}
	case letters != nil:
		return ColLetters{Name: letters.GetText()}
	case ctx.DOLLAR() != nil:
		return ColLast{}
	default:
		return ColNamed{Name: unquote(quoted(ctx.STRING().GetText()))}
	}
}

// buildRowWildcard converts `*`, `*$`, `*$+1`.
func buildRowWildcard(ctx grammar.IRowWildcardContext) (Endpoint, error) {
	row, err := optionalRow(ctx.RowRef())
	if err != nil {
		return nil, err
	}
	return RowSelector{Row: row}, nil
}

// optionalRow maps an absent row context to nil.
func optionalRow(ctx grammar.IRowRefContext) (RowRef, error) {
	if ctx == nil {
		return nil, nil
	}
	return buildRowRef(ctx)
}

// buildRowRef converts the §5.2 row forms: `1` (before), `+1` (after), `*`
// (all), and the `$` family (last/absolute).
func buildRowRef(ctx grammar.IRowRefContext) (RowRef, error) {
	switch {
	case ctx.STAR() != nil:
		return RowAll{}, nil
	case ctx.DOLLAR() != nil:
		return lastOrAbsRow(ctx)
	case ctx.PLUS() != nil:
		return afterRow(ctx.NUMBER())
	default:
		return beforeRow(ctx.NUMBER())
	}
}

// afterRow converts `+n`: n rows after the current row.
func afterRow(number antlr.TerminalNode) (RowRef, error) {
	n, err := intToken(number)
	if err != nil {
		return nil, err
	}
	return RowAfter{N: n}, nil
}

// beforeRow converts a bare `n`: n rows before the current row.
func beforeRow(number antlr.TerminalNode) (RowRef, error) {
	n, err := intToken(number)
	if err != nil {
		return nil, err
	}
	return RowBefore{N: n}, nil
}

// lastRowShape is the common surface of the two `$`-row contexts (rowRef and
// numRow), letting one conversion serve both.
type lastRowShape interface {
	NUMBER() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	DASH() antlr.TerminalNode
}

// lastOrAbsRow converts the `$` row family: `$` (last), `$+n`/`$-n` (offset
// from last), `$n` (absolute row n).
func lastOrAbsRow(ctx lastRowShape) (RowRef, error) {
	number := ctx.NUMBER()
	if number == nil {
		return RowLast{}, nil
	}
	n, err := intToken(number)
	if err != nil {
		return nil, err
	}
	switch {
	case ctx.PLUS() != nil:
		return RowLast{Offset: n}, nil
	case ctx.DASH() != nil:
		return RowLast{Offset: -n}, nil
	default:
		return RowAbs{N: n}, nil
	}
}

// buildNumericEndpoint converts `[col,row]` with either part elidable.
func buildNumericEndpoint(ctx grammar.INumericRefContext) (Endpoint, error) {
	col, err := numericCol(ctx.SignedInt())
	if err != nil {
		return nil, err
	}
	row, err := numericRow(ctx.NumRow())
	if err != nil {
		return nil, err
	}
	return CellEndpoint{Col: col, Row: row}, nil
}

// numericCol converts the optional 0-based column index; absent means elided
// (`[,$+1]`).
func numericCol(ctx grammar.ISignedIntContext) (Col, error) {
	if ctx == nil {
		return ColElided{}, nil
	}
	n, err := signedInt(ctx)
	if err != nil {
		return nil, err
	}
	return ColIndex{Index: n}, nil
}

// signedInt converts `-?[0-9]+`.
func signedInt(ctx grammar.ISignedIntContext) (int, error) {
	n, err := intToken(ctx.NUMBER())
	if err != nil {
		return 0, err
	}
	if ctx.DASH() != nil {
		return -n, nil
	}
	return n, nil
}

// numericRow converts the numeric row forms: `n` (before), `-n` (from end),
// and the `$` family.
func numericRow(ctx grammar.INumRowContext) (RowRef, error) {
	if ctx == nil {
		return nil, nil
	}
	if si := ctx.SignedInt(); si != nil {
		return numericRelRow(si)
	}
	return lastOrAbsRow(ctx)
}

// numericRelRow maps a signed row: n≥0 is n rows before current, n<0 is |n|
// rows from the bottom (§5.2).
func numericRelRow(ctx grammar.ISignedIntContext) (RowRef, error) {
	n, err := signedInt(ctx)
	if err != nil {
		return nil, err
	}
	if n < 0 {
		return RowFromEnd{N: -n}, nil
	}
	return RowBefore{N: n}, nil
}

// buildGroupedRange converts `(C:E)1` / `([3]:[5])1`.
func buildGroupedRange(ctx grammar.IGroupedRangeContext) (Reference, error) {
	row, err := optionalRow(ctx.RowRef())
	if err != nil {
		return nil, err
	}
	if cols := ctx.AllColRef(); len(cols) == 2 {
		return GroupedRange{FromCol: buildColRef(cols[0]), ToCol: buildColRef(cols[1]), Row: row}, nil
	}
	return buildNumericGroupedRange(ctx, row)
}

// buildNumericGroupedRange converts the numeric-column grouped form.
func buildNumericGroupedRange(ctx grammar.IGroupedRangeContext, row RowRef) (Reference, error) {
	numerics := ctx.AllNumericRef()
	from, err := groupedNumericCol(numerics[0])
	if err != nil {
		return nil, err
	}
	to, err := groupedNumericCol(numerics[1])
	if err != nil {
		return nil, err
	}
	return GroupedRange{FromCol: from, ToCol: to, Row: row}, nil
}

// groupedNumericCol admits only a plain `[n]` in grouped-range column
// position; a row inside the brackets is rejected.
func groupedNumericCol(ctx grammar.INumericRefContext) (Col, error) {
	if ctx.COMMA() != nil {
		start := ctx.GetStart()
		return nil, constants.ErrSyntax.With(nil, "line", start.GetLine(), "column", start.GetColumn(), "message", "row not allowed in a grouped-range column")
	}
	return numericCol(ctx.SignedInt())
}
