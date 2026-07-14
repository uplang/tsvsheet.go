// Code generated from TsvsheetParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tsvsheetgrammar // TsvsheetParser
import "github.com/antlr4-go/antlr/v4"


// TsvsheetParserListener is a complete listener for a parse tree produced by TsvsheetParser.
type TsvsheetParserListener interface {
	antlr.ParseTreeListener

	// EnterWorksheet is called when entering the worksheet production.
	EnterWorksheet(c *WorksheetContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterSectionCommand is called when entering the sectionCommand production.
	EnterSectionCommand(c *SectionCommandContext)

	// EnterStructuralCommand is called when entering the structuralCommand production.
	EnterStructuralCommand(c *StructuralCommandContext)

	// EnterCells is called when entering the cells production.
	EnterCells(c *CellsContext)

	// EnterCell is called when entering the cell production.
	EnterCell(c *CellContext)

	// EnterPayload is called when entering the payload production.
	EnterPayload(c *PayloadContext)

	// EnterFormula is called when entering the formula production.
	EnterFormula(c *FormulaContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterRangeRef is called when entering the rangeRef production.
	EnterRangeRef(c *RangeRefContext)

	// EnterEndpoint is called when entering the endpoint production.
	EnterEndpoint(c *EndpointContext)

	// EnterCellRef is called when entering the cellRef production.
	EnterCellRef(c *CellRefContext)

	// EnterColRef is called when entering the colRef production.
	EnterColRef(c *ColRefContext)

	// EnterRowWildcard is called when entering the rowWildcard production.
	EnterRowWildcard(c *RowWildcardContext)

	// EnterRowRef is called when entering the rowRef production.
	EnterRowRef(c *RowRefContext)

	// EnterNumericRef is called when entering the numericRef production.
	EnterNumericRef(c *NumericRefContext)

	// EnterSignedInt is called when entering the signedInt production.
	EnterSignedInt(c *SignedIntContext)

	// EnterNumRow is called when entering the numRow production.
	EnterNumRow(c *NumRowContext)

	// EnterGroupedRange is called when entering the groupedRange production.
	EnterGroupedRange(c *GroupedRangeContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterStringExpr is called when entering the stringExpr production.
	EnterStringExpr(c *StringExprContext)

	// EnterUnaryExpr is called when entering the unaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterRefExpr is called when entering the refExpr production.
	EnterRefExpr(c *RefExprContext)

	// EnterNumberExpr is called when entering the numberExpr production.
	EnterNumberExpr(c *NumberExprContext)

	// EnterMulExpr is called when entering the mulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterCallExpr is called when entering the callExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterCompareExpr is called when entering the compareExpr production.
	EnterCompareExpr(c *CompareExprContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// ExitWorksheet is called when exiting the worksheet production.
	ExitWorksheet(c *WorksheetContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitSectionCommand is called when exiting the sectionCommand production.
	ExitSectionCommand(c *SectionCommandContext)

	// ExitStructuralCommand is called when exiting the structuralCommand production.
	ExitStructuralCommand(c *StructuralCommandContext)

	// ExitCells is called when exiting the cells production.
	ExitCells(c *CellsContext)

	// ExitCell is called when exiting the cell production.
	ExitCell(c *CellContext)

	// ExitPayload is called when exiting the payload production.
	ExitPayload(c *PayloadContext)

	// ExitFormula is called when exiting the formula production.
	ExitFormula(c *FormulaContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitRangeRef is called when exiting the rangeRef production.
	ExitRangeRef(c *RangeRefContext)

	// ExitEndpoint is called when exiting the endpoint production.
	ExitEndpoint(c *EndpointContext)

	// ExitCellRef is called when exiting the cellRef production.
	ExitCellRef(c *CellRefContext)

	// ExitColRef is called when exiting the colRef production.
	ExitColRef(c *ColRefContext)

	// ExitRowWildcard is called when exiting the rowWildcard production.
	ExitRowWildcard(c *RowWildcardContext)

	// ExitRowRef is called when exiting the rowRef production.
	ExitRowRef(c *RowRefContext)

	// ExitNumericRef is called when exiting the numericRef production.
	ExitNumericRef(c *NumericRefContext)

	// ExitSignedInt is called when exiting the signedInt production.
	ExitSignedInt(c *SignedIntContext)

	// ExitNumRow is called when exiting the numRow production.
	ExitNumRow(c *NumRowContext)

	// ExitGroupedRange is called when exiting the groupedRange production.
	ExitGroupedRange(c *GroupedRangeContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitStringExpr is called when exiting the stringExpr production.
	ExitStringExpr(c *StringExprContext)

	// ExitUnaryExpr is called when exiting the unaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitRefExpr is called when exiting the refExpr production.
	ExitRefExpr(c *RefExprContext)

	// ExitNumberExpr is called when exiting the numberExpr production.
	ExitNumberExpr(c *NumberExprContext)

	// ExitMulExpr is called when exiting the mulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitCallExpr is called when exiting the callExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitCompareExpr is called when exiting the compareExpr production.
	ExitCompareExpr(c *CompareExprContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)
}
