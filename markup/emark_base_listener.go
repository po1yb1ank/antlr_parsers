// Code generated from emark.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // emark

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseemarkListener is a complete listener for a parse tree produced by emarkParser.
type BaseemarkListener struct{}

var _ emarkListener = &BaseemarkListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseemarkListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseemarkListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseemarkListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseemarkListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseemarkListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseemarkListener) ExitProgram(ctx *ProgramContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseemarkListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseemarkListener) ExitTag(ctx *TagContext) {}

// EnterRowop is called when production rowop is entered.
func (s *BaseemarkListener) EnterRowop(ctx *RowopContext) {}

// ExitRowop is called when production rowop is exited.
func (s *BaseemarkListener) ExitRowop(ctx *RowopContext) {}

// EnterBlop is called when production blop is entered.
func (s *BaseemarkListener) EnterBlop(ctx *BlopContext) {}

// ExitBlop is called when production blop is exited.
func (s *BaseemarkListener) ExitBlop(ctx *BlopContext) {}

// EnterColop is called when production colop is entered.
func (s *BaseemarkListener) EnterColop(ctx *ColopContext) {}

// ExitColop is called when production colop is exited.
func (s *BaseemarkListener) ExitColop(ctx *ColopContext) {}

// EnterCtx is called when production ctx is entered.
func (s *BaseemarkListener) EnterCtx(ctx *CtxContext) {}

// ExitCtx is called when production ctx is exited.
func (s *BaseemarkListener) ExitCtx(ctx *CtxContext) {}
