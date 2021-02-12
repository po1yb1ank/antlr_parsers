// Code generated from lightc.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // lightc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// lightcListener is a complete listener for a parse tree produced by lightcParser.
type lightcListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTypedef is called when entering the typedef production.
	EnterTypedef(c *TypedefContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterAssign_end is called when entering the assign_end production.
	EnterAssign_end(c *Assign_endContext)

	// EnterForloop is called when entering the forloop production.
	EnterForloop(c *ForloopContext)

	// EnterBool_expression is called when entering the bool_expression production.
	EnterBool_expression(c *Bool_expressionContext)

	// EnterRelop is called when entering the relop production.
	EnterRelop(c *RelopContext)

	// EnterIfstat is called when entering the ifstat production.
	EnterIfstat(c *IfstatContext)

	// EnterReturnval is called when entering the returnval production.
	EnterReturnval(c *ReturnvalContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTypedef is called when exiting the typedef production.
	ExitTypedef(c *TypedefContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitAssign_end is called when exiting the assign_end production.
	ExitAssign_end(c *Assign_endContext)

	// ExitForloop is called when exiting the forloop production.
	ExitForloop(c *ForloopContext)

	// ExitBool_expression is called when exiting the bool_expression production.
	ExitBool_expression(c *Bool_expressionContext)

	// ExitRelop is called when exiting the relop production.
	ExitRelop(c *RelopContext)

	// ExitIfstat is called when exiting the ifstat production.
	ExitIfstat(c *IfstatContext)

	// ExitReturnval is called when exiting the returnval production.
	ExitReturnval(c *ReturnvalContext)
}
