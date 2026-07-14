package tsvt

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"

	"github.com/uplang/tsvsheet.go/internal/constants"
	grammar "github.com/uplang/tsvsheet.go/src/grammar/tsvsheet"
)

// quoted is a lexer-guaranteed double-quoted token text.
type quoted string

// unquote strips the enclosing quotes the STRING token guarantees.
func unquote(s quoted) string { return string(s[1 : len(s)-1]) }

// lineOf is the 1-based source line a rule begins on.
func lineOf(ctx antlr.ParserRuleContext) LineNumber {
	return LineNumber(ctx.GetStart().GetLine())
}

// intToken converts a NUMBER token to an int; the grammar admits fractional
// NUMBERs everywhere, so integer positions (row indexes, header counts) reject
// them here with a positioned syntax error.
func intToken(node antlr.TerminalNode) (int, error) {
	n, err := strconv.Atoi(node.GetText())
	if err != nil {
		sym := node.GetSymbol()
		return 0, constants.ErrSyntax.With(err, "line", sym.GetLine(), "column", sym.GetColumn(), "message", "expected an integer")
	}
	return n, nil
}

// buildTemplate converts the parse tree root into the typed AST.
func buildTemplate(ws grammar.IWorksheetContext) (Template, error) {
	contexts := ws.AllLine()
	lines := make([]Line, 0, len(contexts))
	for _, ctx := range contexts {
		built, err := buildLine(ctx)
		if err != nil {
			return Template{}, err
		}
		lines = append(lines, built)
	}
	return Template{Lines: lines}, nil
}

// buildLine dispatches on the three line shapes: section marker, structural
// command, or cell row.
func buildLine(ctx grammar.ILineContext) (Line, error) {
	if section := ctx.SectionCommand(); section != nil {
		return buildSection(section)
	}
	if structural := ctx.StructuralCommand(); structural != nil {
		return buildStructural(structural)
	}
	return buildRow(ctx.Cells())
}

// buildSection converts `=header(n)`, `=body`, or `=final`.
func buildSection(ctx grammar.ISectionCommandContext) (Line, error) {
	at := lineOf(ctx)
	if ctx.HEADER() != nil {
		count, err := intToken(ctx.NUMBER())
		if err != nil {
			return nil, err
		}
		return HeaderMarker{At: at, Count: count}, nil
	}
	if ctx.BODY() != nil {
		return BodyMarker{At: at}, nil
	}
	return FinalMarker{At: at}, nil
}

// buildStructural converts a standalone `=reference modifier` command.
func buildStructural(ctx grammar.IStructuralCommandContext) (Line, error) {
	ref, err := buildReference(ctx.Reference())
	if err != nil {
		return nil, err
	}
	return Structural{At: lineOf(ctx), Ref: ref, Mod: buildModifier(ctx.Modifier())}, nil
}

// buildRow walks the cells rule's children so consecutive TABs preserve empty
// cells at their column positions.
func buildRow(ctx grammar.ICellsContext) (Line, error) {
	children := ctx.GetChildren()
	cells := make([]Cell, 0, len(children))
	filled := false
	for _, child := range children {
		cell, ok := child.(grammar.ICellContext)
		if !ok { // a TAB terminal: close the current slot
			cells, filled = closeSlot(cells, filled)
			continue
		}
		built, err := buildCell(cell)
		if err != nil {
			return nil, err
		}
		cells, filled = append(cells, built), true
	}
	if !filled && len(children) > 0 { // trailing TAB leaves a final empty slot
		cells = append(cells, EmptyCell{})
	}
	return Row{At: lineOf(ctx), Cells: cells}, nil
}

// closeSlot ends the field before a TAB, inserting an EmptyCell when the field
// had no cell.
func closeSlot(cells []Cell, filled bool) ([]Cell, bool) {
	if !filled {
		return append(cells, EmptyCell{}), false
	}
	return cells, false
}

// buildCell dispatches on the three cell shapes: formula, placement, literal.
func buildCell(ctx grammar.ICellContext) (Cell, error) {
	if formula := ctx.Formula(); formula != nil {
		expr, err := buildExpr(formula.Expression())
		if err != nil {
			return nil, err
		}
		return FormulaCell{Expr: expr}, nil
	}
	if ctx.Reference() != nil {
		return buildPlacement(ctx)
	}
	return LiteralCell{Value: buildLiteral(ctx.Literal())}, nil
}

// buildPlacement converts a reference cell with its optional modifier and
// payload.
func buildPlacement(ctx grammar.ICellContext) (Cell, error) {
	ref, err := buildReference(ctx.Reference())
	if err != nil {
		return nil, err
	}
	payload, err := buildPayload(ctx.Payload())
	if err != nil {
		return nil, err
	}
	return PlacementCell{Ref: ref, Mod: optionalModifier(ctx.Modifier()), Payload: payload}, nil
}

// buildPayload converts a placed payload: a formula, or a literal whose
// leading `=` is the separator (§11.1). A nil context is no payload.
func buildPayload(ctx grammar.IPayloadContext) (Payload, error) {
	if ctx == nil {
		return nil, nil
	}
	if formula := ctx.Formula(); formula != nil {
		expr, err := buildExpr(formula.Expression())
		if err != nil {
			return nil, err
		}
		return FormulaPayload{Expr: expr}, nil
	}
	return LiteralPayload{Value: buildLiteral(ctx.Literal())}, nil
}

// buildLiteral converts a bareword or number literal token. The grammar's
// STRING literal alternative is unreachable — a quoted token always matches the
// reference/formula alternatives first (§11.4 / ADR 0003 rule 16) — so a literal
// datum is only ever a NAME (bareword) or a NUMBER.
func buildLiteral(ctx grammar.ILiteralContext) Literal {
	if name := ctx.NAME(); name != nil {
		return Literal{Kind: LiteralName, Text: name.GetText()}
	}
	return Literal{Kind: LiteralNumber, Text: ctx.NUMBER().GetText()}
}

// optionalModifier maps an absent modifier context to ModNone.
func optionalModifier(ctx grammar.IModifierContext) Modifier {
	if ctx == nil {
		return ModNone
	}
	return buildModifier(ctx)
}

// buildModifier converts `>`, `<`, or `!`.
func buildModifier(ctx grammar.IModifierContext) Modifier {
	switch {
	case ctx.GT() != nil:
		return ModShift
	case ctx.LT() != nil:
		return ModMove
	default:
		return ModDelete
	}
}
