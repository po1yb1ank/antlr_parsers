// Code generated from lightc.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // lightc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaselightcListener is a complete listener for a parse tree produced by lightcParser.
type BaselightcListener struct{}

var _ lightcListener = &BaselightcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaselightcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaselightcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaselightcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaselightcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaselightcListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaselightcListener) ExitProgram(ctx *ProgramContext) {}

// EnterTypedef is called when production typedef is entered.
func (s *BaselightcListener) EnterTypedef(ctx *TypedefContext) {}

// ExitTypedef is called when production typedef is exited.
func (s *BaselightcListener) ExitTypedef(ctx *TypedefContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaselightcListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaselightcListener) ExitStatement(ctx *StatementContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaselightcListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaselightcListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaselightcListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaselightcListener) ExitAssign(ctx *AssignContext) {}

// EnterAssign_end is called when production assign_end is entered.
func (s *BaselightcListener) EnterAssign_end(ctx *Assign_endContext) {}

// ExitAssign_end is called when production assign_end is exited.
func (s *BaselightcListener) ExitAssign_end(ctx *Assign_endContext) {}

// EnterForloop is called when production forloop is entered.
func (s *BaselightcListener) EnterForloop(ctx *ForloopContext) {}

// ExitForloop is called when production forloop is exited.
func (s *BaselightcListener) ExitForloop(ctx *ForloopContext) {}

// EnterBool_expression is called when production bool_expression is entered.
func (s *BaselightcListener) EnterBool_expression(ctx *Bool_expressionContext) {}

// ExitBool_expression is called when production bool_expression is exited.
func (s *BaselightcListener) ExitBool_expression(ctx *Bool_expressionContext) {}

// EnterRelop is called when production relop is entered.
func (s *BaselightcListener) EnterRelop(ctx *RelopContext) {}

// ExitRelop is called when production relop is exited.
func (s *BaselightcListener) ExitRelop(ctx *RelopContext) {}

// EnterIfstat is called when production ifstat is entered.
func (s *BaselightcListener) EnterIfstat(ctx *IfstatContext) {}

// ExitIfstat is called when production ifstat is exited.
func (s *BaselightcListener) ExitIfstat(ctx *IfstatContext) {}

// EnterReturnval is called when production returnval is entered.
func (s *BaselightcListener) EnterReturnval(ctx *ReturnvalContext) {}

// ExitReturnval is called when production returnval is exited.
func (s *BaselightcListener) ExitReturnval(ctx *ReturnvalContext) {}
