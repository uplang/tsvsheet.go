// Code generated from TsvsheetParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tsvsheetgrammar // TsvsheetParser
import "github.com/antlr4-go/antlr/v4"

type BaseTsvsheetParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTsvsheetParserVisitor) VisitWorksheet(ctx *WorksheetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitSectionCommand(ctx *SectionCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitStructuralCommand(ctx *StructuralCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitCells(ctx *CellsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitCell(ctx *CellContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitPayload(ctx *PayloadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitFormula(ctx *FormulaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitReference(ctx *ReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitRangeRef(ctx *RangeRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitEndpoint(ctx *EndpointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitCellRef(ctx *CellRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitColRef(ctx *ColRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitRowWildcard(ctx *RowWildcardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitRowRef(ctx *RowRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitNumericRef(ctx *NumericRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitSignedInt(ctx *SignedIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitNumRow(ctx *NumRowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitGroupedRange(ctx *GroupedRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitModifier(ctx *ModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitRefExpr(ctx *RefExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitNumberExpr(ctx *NumberExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitMulExpr(ctx *MulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitCompareExpr(ctx *CompareExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTsvsheetParserVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}
