// Code generated from TsvsheetParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tsvsheetgrammar // TsvsheetParser
import "github.com/antlr4-go/antlr/v4"


// A complete Visitor for a parse tree produced by TsvsheetParser.
type TsvsheetParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TsvsheetParser#worksheet.
	VisitWorksheet(ctx *WorksheetContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#sectionCommand.
	VisitSectionCommand(ctx *SectionCommandContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#structuralCommand.
	VisitStructuralCommand(ctx *StructuralCommandContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#cells.
	VisitCells(ctx *CellsContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#cell.
	VisitCell(ctx *CellContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#payload.
	VisitPayload(ctx *PayloadContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#formula.
	VisitFormula(ctx *FormulaContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#reference.
	VisitReference(ctx *ReferenceContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#rangeRef.
	VisitRangeRef(ctx *RangeRefContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#endpoint.
	VisitEndpoint(ctx *EndpointContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#cellRef.
	VisitCellRef(ctx *CellRefContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#colRef.
	VisitColRef(ctx *ColRefContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#rowWildcard.
	VisitRowWildcard(ctx *RowWildcardContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#rowRef.
	VisitRowRef(ctx *RowRefContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#numericRef.
	VisitNumericRef(ctx *NumericRefContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#signedInt.
	VisitSignedInt(ctx *SignedIntContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#numRow.
	VisitNumRow(ctx *NumRowContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#groupedRange.
	VisitGroupedRange(ctx *GroupedRangeContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#modifier.
	VisitModifier(ctx *ModifierContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#stringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#addExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#refExpr.
	VisitRefExpr(ctx *RefExprContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#numberExpr.
	VisitNumberExpr(ctx *NumberExprContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#mulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#callExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#parenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#compareExpr.
	VisitCompareExpr(ctx *CompareExprContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by TsvsheetParser#argList.
	VisitArgList(ctx *ArgListContext) interface{}

}