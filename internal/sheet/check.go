package sheet

import (
	"strings"

	"github.com/uplang/tsvsheet.go/internal/tsvt"
)

// Diagnostic is an advisory finding about a formula cell: an unknown function or
// a reference that is not a valid A1 address (both compute to error values).
type Diagnostic struct {
	Cell    string `json:"cell"`
	Message string `json:"message"`
	IsFatal bool   `json:"fatal"`
}

// Check reports the static diagnostics of a parsed sheet: unknown functions and
// non-A1 references in each formula. Syntax errors are already rejected by
// Parse, so Check never reports them.
func Check(s Sheet) []Diagnostic {
	var diags []Diagnostic
	for r, row := range s.cells {
		for c, cl := range row {
			if cl.isFormula() {
				diags = append(diags, checkFormula(cl.formula, Address{Row: r, Col: c})...)
			}
		}
	}
	return diags
}

// checkFormula collects the diagnostics of one formula.
func checkFormula(expr tsvt.Expr, at Address) []Diagnostic {
	diags := unknownFunctions(expr, at)
	return append(diags, nonA1References(expr, at)...)
}

// unknownFunctions flags each call to a name outside the builtin set.
func unknownFunctions(expr tsvt.Expr, at Address) []Diagnostic {
	label := at.String()
	var diags []Diagnostic
	walkCalls(expr, func(call tsvt.Call) {
		if !isKnownFunc(funcName(call.Name)) {
			diags = append(diags, Diagnostic{Cell: label, Message: "unknown function: " + call.Name})
		}
	})
	return diags
}

// nonA1References flags each reference that is not a valid A1 cell or range.
func nonA1References(expr tsvt.Expr, at Address) []Diagnostic {
	label := at.String()
	var diags []Diagnostic
	walkRefs(expr, func(ref tsvt.Reference) {
		if !isA1Reference(ref) {
			diags = append(diags, Diagnostic{Cell: label, Message: "not an A1 reference: " + RenderReference(ref)})
		}
	})
	return diags
}

// isKnownFunc reports whether name (case-insensitive) is a builtin, including
// the lazily-dispatched `if`.
func isKnownFunc(name funcName) boolResult {
	lower := strings.ToLower(string(name))
	if lower == "if" {
		return true
	}
	_, ok := functions[lower]
	return boolResult(ok)
}

// isA1Reference reports whether a reference is a single A1 cell or an A1 range.
func isA1Reference(ref tsvt.Reference) boolResult {
	rangeRef, ok := ref.(tsvt.RangeRef)
	if !ok {
		return false
	}
	if _, fromOK := a1Address(rangeRef.From); !fromOK {
		return false
	}
	if rangeRef.To == nil {
		return true
	}
	_, toOK := a1Address(rangeRef.To)
	return toOK
}

// walkCalls visits every function call in an expression tree.
func walkCalls(expr tsvt.Expr, visit func(tsvt.Call)) {
	switch e := expr.(type) {
	case tsvt.Unary:
		walkCalls(e.X, visit)
	case tsvt.Binary:
		walkCalls(e.Left, visit)
		walkCalls(e.Right, visit)
	case tsvt.Call:
		visit(e)
		for _, arg := range e.Args {
			walkCalls(arg, visit)
		}
	}
}

// walkRefs visits every reference operand in an expression tree.
func walkRefs(expr tsvt.Expr, visit func(tsvt.Reference)) {
	switch e := expr.(type) {
	case tsvt.RefOperand:
		visit(e.Ref)
	case tsvt.Unary:
		walkRefs(e.X, visit)
	case tsvt.Binary:
		walkRefs(e.Left, visit)
		walkRefs(e.Right, visit)
	case tsvt.Call:
		for _, arg := range e.Args {
			walkRefs(arg, visit)
		}
	}
}
