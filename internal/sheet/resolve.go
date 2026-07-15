package sheet

import "github.com/uplang/tsvsheet.go/internal/tsvt"

// cellPhase tracks a cell's evaluation state for memoization and cycle
// detection.
type cellPhase int

const (
	phaseUnvisited cellPhase = iota
	phaseVisiting            // on the current evaluation stack → a cycle
	phaseDone
)

// computer memoizes cell values as they are evaluated in dependency order. Its
// cache and phase slices are allocated once and shared, so value-receiver
// methods mutate them in place (no reassignment) and every recursive read sees
// the same state.
type computer struct {
	sheet Sheet
	cache [][]Value
	phase [][]cellPhase
}

// newComputer builds a computer sized to the sheet.
func newComputer(s Sheet) computer {
	cache := make([][]Value, len(s.cells))
	phase := make([][]cellPhase, len(s.cells))
	for r, row := range s.cells {
		cache[r] = make([]Value, len(row))
		phase[r] = make([]cellPhase, len(row))
	}
	return computer{sheet: s, cache: cache, phase: phase}
}

// output is a cell's rendered value: a literal verbatim, a formula computed.
func (c computer) output(row rowIndex, col colIndex, cl cell) textVal {
	if !cl.isFormula() {
		return textVal(cl.text)
	}
	return textVal(c.read(row, col).String())
}

// read returns the value at (row, col), evaluating and memoizing it on first
// visit. A cell already on the evaluation stack is a circular reference; an
// out-of-grid position is #REF!.
func (c computer) read(row rowIndex, col colIndex) Value {
	cl, inGrid := c.sheet.at(row, col)
	if !inGrid {
		return errorValue(ErrRef)
	}
	switch c.phase[row][col] {
	case phaseDone:
		return c.cache[row][col]
	case phaseVisiting:
		return errorValue(ErrCirc)
	}
	c.phase[row][col] = phaseVisiting
	result := c.evalCell(cl)
	c.cache[row][col] = result
	c.phase[row][col] = phaseDone
	return result
}

// evalCell evaluates one cell: a literal parses to its value, a formula
// evaluates its expression (which reads other cells).
func (c computer) evalCell(cl cell) Value {
	if !cl.isFormula() {
		return value(textVal(cl.text))
	}
	return resolver{comp: c}.eval(cl.formula)
}

// resolver evaluates expressions against the computer, resolving A1 references.
type resolver struct {
	comp computer
}

// cellset is a resolved reference: the referenced cells' values, and whether it
// was a single cell (a range is not single).
type cellset struct {
	values   []Value
	isSingle boolResult
}

// scalar reduces a cellset to one value: a single cell is its value; a
// multi-cell range used where a scalar is required is #VALUE!.
func (c cellset) scalar() Value {
	if c.isSingle && len(c.values) == 1 {
		return c.values[0]
	}
	return errorValue(ErrValue)
}

// resolveOperand resolves a reference operand: a single A1 cell or an A1 range.
// A grouped range is not part of the A1 model and resolves to a single #REF!.
func (r resolver) resolveOperand(ref tsvt.Reference) cellset {
	rangeRef, ok := ref.(tsvt.RangeRef)
	if !ok {
		return refError()
	}
	if rangeRef.To == nil {
		return r.resolveSingle(rangeRef.From)
	}
	return r.resolveMatrix(rangeRef.From, rangeRef.To)
}

// refError is the #REF! result of an invalid single reference; isSingle so the
// error propagates through scalar() rather than being masked as #VALUE!.
func refError() cellset {
	return cellset{values: []Value{errorValue(ErrRef)}, isSingle: true}
}

// resolveSingle resolves a single-cell reference.
func (r resolver) resolveSingle(ep tsvt.Endpoint) cellset {
	at, ok := a1Address(ep)
	if !ok {
		return refError()
	}
	return cellset{values: []Value{r.comp.read(rowIndex(at.Row), colIndex(at.Col))}, isSingle: true}
}

// resolveMatrix resolves the rectangular hull of two A1 corners (`A1:B3`).
func (r resolver) resolveMatrix(from, to tsvt.Endpoint) cellset {
	a, aok := a1Address(from)
	b, bok := a1Address(to)
	if !aok || !bok {
		return cellset{values: []Value{errorValue(ErrRef)}}
	}
	return cellset{values: r.hull(a, b)}
}

// hull reads every cell in the inclusive rectangle spanned by a and b.
func (r resolver) hull(a, b Address) []Value {
	r0, r1 := ordered(gridPos(a.Row), gridPos(b.Row))
	c0, c1 := ordered(gridPos(a.Col), gridPos(b.Col))
	values := make([]Value, 0, int(r1-r0+1)*int(c1-c0+1))
	for row := r0; row <= r1; row++ {
		for col := c0; col <= c1; col++ {
			values = append(values, r.comp.read(rowIndex(row), colIndex(col)))
		}
	}
	return values
}

// a1Address converts an endpoint to an absolute (row, col); ok is false for any
// non-A1 form (a row selector, a `$`/named/numeric column, or a relative row).
func a1Address(ep tsvt.Endpoint) (Address, boolResult) {
	cellRef, isCell := ep.(tsvt.CellEndpoint)
	if !isCell {
		return Address{}, false
	}
	col, colOK := a1Column(cellRef.Col)
	row, rowOK := a1Row(cellRef.Row)
	if !colOK || !rowOK {
		return Address{}, false
	}
	return Address{Row: int(row), Col: int(col)}, true
}

// a1Column resolves a column to a 0-based index; only plain/absolute letters are
// A1 columns (a named, numeric, last, or elided column is not).
func a1Column(col tsvt.Col) (colIndex, boolResult) {
	letters, ok := col.(tsvt.ColLetters)
	if !ok {
		return 0, false
	}
	return colIndex(lettersToIndex(columnLetters(letters.Name))), true
}

// a1Row resolves a row to a 0-based index. A1 rows are 1-based absolute: `A1`
// parses as RowBefore{1}, `A$1` as RowAbs{1}; relative and wildcard forms are
// not A1.
func a1Row(row tsvt.RowRef) (rowIndex, boolResult) {
	switch r := row.(type) {
	case tsvt.RowBefore:
		return absoluteRow(rowNumber(r.N))
	case tsvt.RowAbs:
		return absoluteRow(rowNumber(r.N))
	default:
		return 0, false
	}
}

// absoluteRow maps a 1-based row number to a 0-based index, rejecting row 0.
func absoluteRow(n rowNumber) (rowIndex, boolResult) {
	if n < 1 {
		return 0, false
	}
	return rowIndex(n - 1), true
}

// ordered returns its two coordinates low-first.
func ordered(x, y gridPos) (gridPos, gridPos) {
	if x <= y {
		return x, y
	}
	return y, x
}
