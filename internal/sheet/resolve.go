package sheet

import "github.com/uplang/tsvsheet.go/internal/tsvt"

// resolver resolves references against a grid at a current body row. row is -1
// in a no-row context (final phase or a template with no data rows), where only
// absolute rows resolve. width is the logical column width (data width adjusted
// by structural ops, ADR 0003 rule 19), so columns computed to the right do not
// change what `$` (last column) and negative indices mean.
type resolver struct {
	names map[string]int
	grid  Grid
	row   int
	width int
}

// colRes is the outcome of resolving a column: a real index, or an unbound
// named column carried as name for the string-literal fallback (ADR 0003
// rule 16).
type colRes struct {
	name  string
	index int
	isOK  bool
}

// cellset is a resolved reference as an operand: the resolved cell values and
// whether it was a single cell (a range/matrix is not single).
type cellset struct {
	values   []Value
	isSingle bool
}

// scalar reduces a cellset to one value: a single cell is its value; a
// multi-cell range used where a scalar is required is #VALUE! (ADR 0003 rule 8).
func (c cellset) scalar() Value {
	if c.isSingle && len(c.values) == 1 {
		return c.values[0]
	}
	return errorValue(ErrValue)
}

// resolveColumn resolves a column reference to an index against the grid width.
func (r resolver) resolveColumn(col tsvt.Col) colRes {
	switch c := col.(type) {
	case tsvt.ColLetters:
		return colRes{index: lettersToIndex(columnLetters(c.Name)), isOK: true}
	case tsvt.ColIndex:
		return colRes{index: normalizeIndex(colIndex(c.Index), colIndex(r.width)), isOK: true}
	case tsvt.ColLast:
		return colRes{index: r.width - 1, isOK: true}
	case tsvt.ColNamed:
		return r.resolveNamed(c.Name)
	default: // tsvt.ColElided — invalid as an operand column
		return colRes{isOK: false}
	}
}

// resolveNamed looks up a header-bound column; an unbound name is carried for
// the string-literal fallback.
func (r resolver) resolveNamed(name string) colRes {
	if index, found := r.names[name]; found {
		return colRes{index: index, isOK: true}
	}
	return colRes{name: name}
}

// normalizeIndex maps a possibly-negative 0-based index to a concrete column,
// counting negatives from the end (§5.1).
func normalizeIndex(index, width colIndex) int {
	if index < 0 {
		return int(width + index)
	}
	return int(index)
}

// resolveRow resolves a row reference to a 0-based grid row. ok is false when
// the reference cannot resolve in the current context (e.g. a relative row with
// no current row), which the caller renders as #REF!.
func (r resolver) resolveRow(row tsvt.RowRef) (int, bool) {
	last := r.grid.rows() - 1
	switch ref := row.(type) {
	case nil:
		return r.row, r.row >= 0
	case tsvt.RowBefore:
		return relativeRow(rowIndex(r.row), rowIndex(-ref.N))
	case tsvt.RowAfter:
		return relativeRow(rowIndex(r.row), rowIndex(ref.N))
	case tsvt.RowAbs:
		return ref.N - 1, true
	case tsvt.RowLast:
		return last + ref.Offset, true
	case tsvt.RowFromEnd:
		return r.grid.rows() - ref.N, true
	default: // tsvt.RowAll — not a single row; invalid in scalar context
		return 0, false
	}
}

// relativeRow offsets the current row, reporting ok=false when there is no
// current row (final phase).
func relativeRow(current, delta rowIndex) (int, bool) {
	if current < 0 {
		return 0, false
	}
	return int(current + delta), true
}

// resolveOperand resolves a reference for use as an expression operand.
func (r resolver) resolveOperand(ref tsvt.Reference) cellset {
	switch t := ref.(type) {
	case tsvt.RangeRef:
		return r.resolveRange(t)
	default: // tsvt.GroupedRange
		return r.resolveGrouped(t.(tsvt.GroupedRange))
	}
}

// resolveRange resolves a single endpoint or a two-endpoint matrix.
func (r resolver) resolveRange(ref tsvt.RangeRef) cellset {
	if ref.To == nil {
		return r.resolveEndpoint(ref.From)
	}
	return r.resolveMatrix(ref.From, ref.To)
}

// resolveEndpoint resolves one endpoint used alone: a single cell, or a whole
// row when the column is elided (a row selector).
func (r resolver) resolveEndpoint(ep tsvt.Endpoint) cellset {
	switch e := ep.(type) {
	case tsvt.CellEndpoint:
		return cellset{values: []Value{r.resolveCell(e.Col, e.Row)}, isSingle: true}
	default: // tsvt.RowSelector
		return r.resolveRowSelector(ep.(tsvt.RowSelector))
	}
}

// resolveCell resolves a single cell to a Value; an unbound named column is the
// string literal of its name (ADR 0003 rule 16), an out-of-grid position is
// #REF!.
func (r resolver) resolveCell(col tsvt.Col, row tsvt.RowRef) Value {
	cr := r.resolveColumn(col)
	if !cr.isOK {
		if cr.name != "" {
			return stringValue(textVal(cr.name))
		}
		return errorValue(ErrRef)
	}
	rowIdx, ok := r.resolveRow(row)
	if !ok {
		return errorValue(ErrRef)
	}
	return r.read(rowIdx, cr.index)
}

// read returns the value at a resolved position, or #REF! when out of grid.
func (r resolver) read(row, col int) Value {
	raw, ok := r.grid.at(row, col)
	if !ok {
		return errorValue(ErrRef)
	}
	return value(textVal(raw))
}

// resolveRowSelector resolves a whole-row reference to that row's cells.
func (r resolver) resolveRowSelector(sel tsvt.RowSelector) cellset {
	rowIdx, ok := r.resolveRow(sel.Row)
	if !ok {
		return cellset{values: []Value{errorValue(ErrRef)}}
	}
	values := make([]Value, r.width)
	for col := range values {
		values[col] = r.read(rowIdx, col)
	}
	return cellset{values: values}
}

// resolveMatrix resolves the rectangular hull of two cell endpoints (§5.3); a
// non-cell endpoint or an unresolvable corner is #REF! (ADR 0003 rule 4).
func (r resolver) resolveMatrix(from, to tsvt.Endpoint) cellset {
	a, aok := r.corner(from)
	b, bok := r.corner(to)
	if !aok || !bok {
		return cellset{values: []Value{errorValue(ErrRef)}}
	}
	return cellset{values: r.hull(a, b)}
}

// corner resolves a matrix endpoint to a concrete (row, col); ok is false for a
// non-cell endpoint or an unresolvable column/row.
func (r resolver) corner(ep tsvt.Endpoint) (Address, bool) {
	cell, isCell := ep.(tsvt.CellEndpoint)
	if !isCell {
		return Address{}, false
	}
	cr := r.resolveColumn(cell.Col)
	rowIdx, rowOK := r.resolveRow(cell.Row)
	if !cr.isOK || !rowOK {
		return Address{}, false
	}
	return Address{Row: rowIdx, Col: cr.index}, true
}

// hull reads every cell in the inclusive rectangle spanned by a and b.
func (r resolver) hull(a, b Address) []Value {
	r0, r1 := ordered(gridPos(a.Row), gridPos(b.Row))
	c0, c1 := ordered(gridPos(a.Col), gridPos(b.Col))
	values := make([]Value, 0, (r1-r0+1)*(c1-c0+1))
	for row := r0; row <= r1; row++ {
		for col := c0; col <= c1; col++ {
			values = append(values, r.read(row, col))
		}
	}
	return values
}

// ordered returns its two arguments low-first.
func ordered(x, y gridPos) (int, int) {
	if x <= y {
		return int(x), int(y)
	}
	return int(y), int(x)
}

// resolveGrouped resolves a grouped column range with one trailing row applied
// across it: `(C:E)1` (§5.3).
func (r resolver) resolveGrouped(ref tsvt.GroupedRange) cellset {
	from := r.resolveColumn(ref.FromCol)
	to := r.resolveColumn(ref.ToCol)
	rowIdx, rowOK := r.resolveRow(ref.Row)
	if !from.isOK || !to.isOK || !rowOK {
		return cellset{values: []Value{errorValue(ErrRef)}}
	}
	return cellset{values: r.groupedCells(from.index, to.index, rowIdx)}
}

// groupedCells reads one row across an inclusive column span.
func (r resolver) groupedCells(fromCol, toCol, row int) []Value {
	c0, c1 := ordered(gridPos(fromCol), gridPos(toCol))
	values := make([]Value, 0, c1-c0+1)
	for col := c0; col <= c1; col++ {
		values = append(values, r.read(row, col))
	}
	return values
}
