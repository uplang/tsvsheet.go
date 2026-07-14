// Package tsvt is the covered seam over the ANTLR-generated .tsvt parser: it
// turns template source into an immutable typed AST (or a sentinel syntax
// error) and hides every ANTLR type from the rest of the program.
package tsvt

// The eight AST interfaces below are sealed: each has an unexported marker
// method, so only the node types in this package can satisfy it. Rather than
// declare that method on all ~30 node types, each interface has one zero-size
// marker struct that carries the method, embedded in every node of that
// interface. The compute layer consumes the AST by type switch; the markers
// exist purely to bound each switch's variant set at compile time.
type (
	lineMarker      struct{}
	cellMarker      struct{}
	payloadMarker   struct{}
	referenceMarker struct{}
	endpointMarker  struct{}
	colMarker       struct{}
	rowMarker       struct{}
	exprMarker      struct{}
)

func (lineMarker) isLine()           {}
func (cellMarker) isCell()           {}
func (payloadMarker) isPayload()     {}
func (referenceMarker) isReference() {}
func (endpointMarker) isEndpoint()   {}
func (colMarker) isCol()             {}
func (rowMarker) isRow()             {}
func (exprMarker) isExpr()           {}

// Source is raw .tsvt template text.
type Source []byte

// LineNumber is a 1-based line position in the template source.
type LineNumber int

// Template is a parsed .tsvt file: one Line per source line.
type Template struct {
	Lines []Line
}

// Line is one template instruction: a section marker, a structural command, or
// a row of cells. The set is sealed.
type Line interface{ isLine() }

// HeaderMarker is `=header(n)`: the next n lines define the header rows (§4).
type HeaderMarker struct {
	lineMarker
	At    LineNumber
	Count int
}

// BodyMarker is `=body`: subsequent lines apply per data row (§4).
type BodyMarker struct {
	lineMarker
	At LineNumber
}

// FinalMarker is `=final`: subsequent lines apply once, after all rows (§4).
type FinalMarker struct {
	lineMarker
	At LineNumber
}

// Structural is a standalone structural command, e.g. `=A<` (§4/§6).
type Structural struct {
	lineMarker
	At  LineNumber
	Ref Reference
	Mod Modifier
}

// Row is a line of TAB-separated cells; empty cells preserve column position.
type Row struct {
	lineMarker
	At    LineNumber
	Cells []Cell
}

// Cell is one TAB-delimited field of a row. The set is sealed.
type Cell interface{ isCell() }

// EmptyCell is an empty field (consecutive TABs).
type EmptyCell struct{ cellMarker }

// FormulaCell is a positional formula: `=C + D`.
type FormulaCell struct {
	cellMarker
	Expr Expr
}

// LiteralCell is a bare literal datum or header label: `1`, `Total`.
type LiteralCell struct {
	cellMarker
	Value Literal
}

// PlacementCell is a reference with an optional modifier and payload:
// `E=…`, `C<`, `A$+1=Total`, or a bare reference used as a header label.
type PlacementCell struct {
	cellMarker
	Ref     Reference
	Mod     Modifier // ModNone when absent
	Payload Payload  // nil when absent
}

// Payload is what follows a placed reference: a formula or a literal datum
// whose leading `=` is the separator, not a formula marker (§11.1).
type Payload interface{ isPayload() }

// FormulaPayload places a computed expression: `E=C + D`.
type FormulaPayload struct {
	payloadMarker
	Expr Expr
}

// LiteralPayload places verbatim text: `A$+1=Total`.
type LiteralPayload struct {
	payloadMarker
	Value Literal
}

// LiteralKind distinguishes the literal token shapes a literal datum can take.
//
// There is no string kind: a double-quoted token is always a named-column
// reference (ColNamed), because in every position where a literal could appear
// — a bare cell and an addressed payload — the reference and formula
// alternatives outrank the literal alternative for a STRING (see build.go).
// The §11.4 "string literal" role is recovered semantically: a ColNamed with
// no matching header binding is a string literal (ADR 0003 rule 16).
type LiteralKind string

// The literal kinds: a bareword name and a number.
const (
	LiteralName   LiteralKind = "name"
	LiteralNumber LiteralKind = "number"
)

// Literal is a literal datum: a bareword name or a number, verbatim.
type Literal struct {
	Kind LiteralKind
	Text string
}

// Modifier is a structural operator (§6): shift, move, delete, or none.
type Modifier string

// The modifiers of §6. ModNone marks an absent modifier.
const (
	ModNone   Modifier = ""
	ModShift  Modifier = ">"
	ModMove   Modifier = "<"
	ModDelete Modifier = "!"
)

// Reference is the §5 reference algebra. The set is sealed.
type Reference interface{ isReference() }

// RangeRef is a single endpoint or a two-endpoint range/matrix: `C`, `C:E`,
// `C1:E3`, `[3,1]:[5,3]`, `$B$1:$F$-1`.
type RangeRef struct {
	referenceMarker
	From Endpoint
	To   Endpoint // nil for a single endpoint
}

// GroupedRange is a column range with one trailing row applied across it:
// `(C:E)1`, `([3]:[5])1` (§5.3).
type GroupedRange struct {
	referenceMarker
	FromCol Col
	ToCol   Col
	Row     RowRef // nil when absent
}

// Endpoint is one end of a range: a cell/column reference or a row selector.
type Endpoint interface{ isEndpoint() }

// CellEndpoint is a column with an optional row: `A`, `C1`, `$B`, `[3,1]`,
// `[,$+1]` (elided column), `"Sum"`.
type CellEndpoint struct {
	endpointMarker
	Col Col
	Row RowRef // nil when absent (whole column / current row per context)
}

// RowSelector is a whole-row reference with the column elided: `*`, `*$`,
// `*$+1` (§5.2).
type RowSelector struct {
	endpointMarker
	Row RowRef // nil for bare `*` (each row)
}

// Col is a column reference (§5.1). The set is sealed.
type Col interface{ isCol() }

// ColLetters is a spreadsheet-style column: `A`, `AA`; Abs marks `$B`.
type ColLetters struct {
	colMarker
	Name string
	Abs  bool
}

// ColLast is `$`: the last column.
type ColLast struct{ colMarker }

// ColNamed is a header-named column: `"Sum"`.
type ColNamed struct {
	colMarker
	Name string
}

// ColIndex is a 0-based numeric column, negatives from the end: `[3]`, `[-1]`.
type ColIndex struct {
	colMarker
	Index int
}

// ColElided is a numeric reference with the column omitted: `[,$+1]`.
type ColElided struct{ colMarker }

// RowRef is a row reference (§5.2). The set is sealed.
type RowRef interface{ isRow() }

// RowBefore is n rows before the current row: `C1` (n=1), `C0`/`C` (n=0),
// `[3,1]`.
type RowBefore struct {
	rowMarker
	N int
}

// RowAfter is n rows after the current row: `C+1`.
type RowAfter struct {
	rowMarker
	N int
}

// RowAll is every row: `E*`.
type RowAll struct{ rowMarker }

// RowLast is the last row with an offset: `C$` (0), `$F$-1` (-1), `*$+1` (+1),
// `[3,$]`.
type RowLast struct {
	rowMarker
	Offset int
}

// RowAbs is an absolute 1-based data row: `C$4`.
type RowAbs struct {
	rowMarker
	N int
}

// RowFromEnd is n rows from the bottom, from the numeric form: `[-3,-5]` (n=5).
type RowFromEnd struct {
	rowMarker
	N int
}

// Expr is a §11 formula expression. The set is sealed.
type Expr interface{ isExpr() }

// BinaryOp is a §11 binary operator.
type BinaryOp string

// The binary operators, tightest-binding tier first: multiplicative, additive,
// comparison (§11.2).
const (
	OpMul BinaryOp = "*"
	OpDiv BinaryOp = "/"
	OpMod BinaryOp = "%"
	OpAdd BinaryOp = "+"
	OpSub BinaryOp = "-"
	OpEq  BinaryOp = "="
	OpNe  BinaryOp = "<>"
	OpLt  BinaryOp = "<"
	OpLe  BinaryOp = "<="
	OpGt  BinaryOp = ">"
	OpGe  BinaryOp = ">="
)

// UnaryOp is a §11 unary sign operator.
type UnaryOp string

// The unary operators.
const (
	OpNeg UnaryOp = "-"
	OpPos UnaryOp = "+"
)

// Binary is a binary operation.
type Binary struct {
	exprMarker
	Op    BinaryOp
	Left  Expr
	Right Expr
}

// Unary is a unary sign operation.
type Unary struct {
	exprMarker
	Op UnaryOp
	X  Expr
}

// Call is a function call; Name is case-preserved (identity is resolved
// case-insensitively by the evaluator, §11.3).
type Call struct {
	exprMarker
	Name string
	Args []Expr
}

// RefOperand is a reference used as an expression operand.
type RefOperand struct {
	exprMarker
	Ref Reference
}

// Number is a numeric literal; Text preserves the source form.
type Number struct {
	exprMarker
	Text string
}
