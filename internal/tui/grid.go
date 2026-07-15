package tui

// width is the number of columns to display, at least one so the cursor always
// has a column. The computed grid mirrors the source grid's shape, so either
// gives the column count.
func (m Model) width() int {
	if w := gridWidth(m.state.Computed); w > 0 {
		return w
	}
	return 1
}

// height is the number of rows to display, at least one.
func (m Model) height() int {
	if n := len(m.state.Computed); n > 0 {
		return n
	}
	return 1
}

// gridWidth is the widest row of a grid.
func gridWidth(g [][]string) int {
	w := 0
	for _, row := range g {
		if len(row) > w {
			w = len(row)
		}
	}
	return w
}

// computedAt returns the computed value at (row, col), or empty when absent.
func (m Model) computedAt(row, col int) string {
	return cellAt(m.state.Computed, cursorPos(row), cursorPos(col))
}

// sourceAt returns the source text at (row, col), or empty when absent.
func (m Model) sourceAt(row, col int) string {
	return cellAt(m.state.Source, cursorPos(row), cursorPos(col))
}

// cellAt reads a grid cell, returning empty for any out-of-bounds position.
func cellAt(g [][]string, row, col cursorPos) string {
	r, c := int(row), int(col)
	if r < 0 || r >= len(g) || c < 0 || c >= len(g[r]) {
		return ""
	}
	return g[r][c]
}
