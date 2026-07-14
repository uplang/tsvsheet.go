package sheet

import "github.com/uplang/tsvsheet.go/internal/tsvt"

// phase groups a template's lines into the SPECIFICATION §9 sections.
type phase struct {
	header []tsvt.Row
	body   []tsvt.Line
	final  []tsvt.Line
}

// sectionID selects the active body/final section during partitioning.
type sectionID int

const (
	sectionBody sectionID = iota
	sectionFinal
)

// computeState is the mutable state of a single processing pass: the computed
// grid, the header-bound column names, and the logical column width. It is held
// behind a pointer by computation so the pass's value-receiver methods mutate a
// shared state.
type computeState struct {
	names map[string]int
	grid  Grid
	width int
}

// computation carries the mutable pass state (see computeState) through the
// SPECIFICATION §9 processing model.
type computation struct {
	state *computeState
}

// Compute runs the SPECIFICATION §9 processing model over a data grid and a
// parsed template, returning the computed output grid. Cell-level failures are
// error values in the grid (ADR 0003); the error return is reserved for a
// template the processor structurally rejects.
func Compute(t tsvt.Template, g Grid) (Grid, error) {
	for _, d := range Check(t) {
		if d.IsFatal {
			return nil, d.err()
		}
	}
	phases := partition(t)
	comp := computation{state: &computeState{grid: g.clone(), names: map[string]int{}, width: g.cols()}}
	comp.bindHeaders(phases.header)
	comp.runBody(phases.body)
	comp.runFinal(phases.final)
	return comp.state.grid, nil
}

// partition splits template lines into header / body / final by section markers
// (§4). With no markers every non-marker line is body (§4 minimal form).
func partition(t tsvt.Template) phase {
	var p phase
	active := sectionBody
	headers := headerCount(0)
	for _, line := range t.Lines {
		p, active, headers = route(p, line, active, headers)
	}
	return p
}

// route dispatches one template line to the active section, honoring a pending
// header-line count, and returns the updated phase, active section, and count.
func route(p phase, line tsvt.Line, active sectionID, headers headerCount) (phase, sectionID, headerCount) {
	if headers > 0 {
		if row, ok := line.(tsvt.Row); ok {
			p.header = append(p.header, row)
		}
		return p, active, headers - 1
	}
	switch m := line.(type) {
	case tsvt.HeaderMarker:
		return p, active, headerCount(m.Count)
	case tsvt.BodyMarker:
		return p, sectionBody, 0
	case tsvt.FinalMarker:
		return p, sectionFinal, 0
	default:
		return appendToSection(p, active, line), active, 0
	}
}

// appendToSection appends a line to the phase's active section.
func appendToSection(p phase, active sectionID, line tsvt.Line) phase {
	if active == sectionFinal {
		p.final = append(p.final, line)
		return p
	}
	p.body = append(p.body, line)
	return p
}

// bindHeaders names columns from the header rows: each row's cells name columns
// left to right by field index (§5.1).
func (c computation) bindHeaders(rows []tsvt.Row) {
	for _, row := range rows {
		for col, cell := range row.Cells {
			if name, ok := headerName(cell); ok {
				c.state.names[name] = col
			}
		}
	}
}

// headerName extracts a column name from a header cell: a bareword/number
// literal, or a named-column reference. Other cells (formulas, empties) do not
// name a column.
func headerName(cell tsvt.Cell) (string, bool) {
	switch c := cell.(type) {
	case tsvt.LiteralCell:
		return c.Value.Text, true
	case tsvt.PlacementCell:
		if named, ok := c.Ref.(tsvt.RangeRef); ok {
			return namedFromRange(named)
		}
	}
	return "", false
}

// namedFromRange returns the label of a single named-column reference used as a
// header cell (`"Sum"`); other reference shapes do not name a column.
func namedFromRange(ref tsvt.RangeRef) (string, bool) {
	cell, ok := ref.From.(tsvt.CellEndpoint)
	if !ok {
		return "", false
	}
	named, ok := cell.Col.(tsvt.ColNamed)
	if !ok {
		return "", false
	}
	return named.Name, true
}

// runBody applies each body line to every data row in order (§9.4).
func (c computation) runBody(lines []tsvt.Line) {
	for row := 0; row < c.state.grid.rows(); row++ {
		for _, line := range lines {
			c.applyLine(line, row)
		}
	}
}

// runFinal applies each final line once over the finished grid (§9.5), with no
// current row.
func (c computation) runFinal(lines []tsvt.Line) {
	for _, line := range lines {
		c.applyLine(line, noRow)
	}
}

// noRow marks the absence of a current row (final phase).
const noRow = -1
