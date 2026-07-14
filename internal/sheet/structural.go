package sheet

import "github.com/uplang/tsvsheet.go/internal/tsvt"

// applyStructural applies a standalone column structural command to the whole
// grid at the point its phase reaches it (ADR 0003 rule 18): `<` inserts an
// empty column before the target, `>` after, `!` deletes it. Check has already
// rejected range-scoped and non-column forms, so the reference is a single
// column here.
func (c computation) applyStructural(cmd tsvt.Structural) {
	col, ok := c.structuralColumn(cmd.Ref)
	if !ok {
		return
	}
	switch cmd.Mod {
	case tsvt.ModMove:
		c.insertColumn(col)
	case tsvt.ModShift:
		c.insertColumn(col + 1)
	default: // tsvt.ModDelete
		c.deleteColumn(col)
	}
}

// structuralColumn resolves the target column index of a structural command.
// Check has guaranteed the reference is a single bare column (isBareColumn), so
// placementCell always succeeds; only an unbound named column fails to resolve.
func (c computation) structuralColumn(ref tsvt.Reference) (int, bool) {
	cell, _ := placementCell(ref)
	cr := c.resolverAt(noRow).resolveColumn(cell.Col)
	return cr.index, cr.isOK
}

// insertColumn inserts an empty column at index k in every row and shifts the
// header bindings at or past k up by one.
func (c computation) insertColumn(k int) {
	for row, line := range c.state.grid {
		c.state.grid[row] = insertAt(line, clampCol(colIndex(k), colIndex(len(line))), "")
	}
	c.state.width++
	c.shiftNames(k, +1)
}

// deleteColumn removes column k from every row and shifts the header bindings
// past k down by one, dropping the binding at k.
func (c computation) deleteColumn(k int) {
	for row, line := range c.state.grid {
		if k < len(line) {
			c.state.grid[row] = append(line[:k], line[k+1:]...)
		}
	}
	c.state.width--
	c.dropName(k)
	c.shiftNames(k+1, -1)
}

// clampCol bounds an insertion index to a row's current width so an insert past
// the end appends.
func clampCol(k, width colIndex) colIndex {
	if k > width {
		return width
	}
	return k
}

// insertAt inserts v at index i of a string slice.
func insertAt(line []string, i colIndex, v textVal) []string {
	out := make([]string, 0, len(line)+1)
	out = append(out, line[:i]...)
	out = append(out, string(v))
	return append(out, line[i:]...)
}

// shiftNames adjusts header bindings at or past k by delta.
func (c computation) shiftNames(k, delta int) {
	for name, col := range c.state.names {
		if col >= k {
			c.state.names[name] = col + delta
		}
	}
}

// dropName removes any header binding exactly at column k.
func (c computation) dropName(k int) {
	for name, col := range c.state.names {
		if col == k {
			delete(c.state.names, name)
		}
	}
}
