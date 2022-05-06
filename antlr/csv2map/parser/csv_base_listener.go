// Code generated from CSV.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CSV

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCSVListener is a complete listener for a parse tree produced by CSVParser.
type BaseCSVListener struct{}

var _ CSVListener = &BaseCSVListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCSVListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCSVListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCSVListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCSVListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCsvFile is called when production csvFile is entered.
func (s *BaseCSVListener) EnterCsvFile(ctx *CsvFileContext) {}

// ExitCsvFile is called when production csvFile is exited.
func (s *BaseCSVListener) ExitCsvFile(ctx *CsvFileContext) {}

// EnterHdr is called when production hdr is entered.
func (s *BaseCSVListener) EnterHdr(ctx *HdrContext) {}

// ExitHdr is called when production hdr is exited.
func (s *BaseCSVListener) ExitHdr(ctx *HdrContext) {}

// EnterRow is called when production row is entered.
func (s *BaseCSVListener) EnterRow(ctx *RowContext) {}

// ExitRow is called when production row is exited.
func (s *BaseCSVListener) ExitRow(ctx *RowContext) {}

// EnterField is called when production field is entered.
func (s *BaseCSVListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseCSVListener) ExitField(ctx *FieldContext) {}
