// Code generated from emark.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // emark

import "github.com/antlr/antlr4/runtime/Go/antlr"

// emarkListener is a complete listener for a parse tree produced by emarkParser.
type emarkListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterRowop is called when entering the rowop production.
	EnterRowop(c *RowopContext)

	// EnterBlop is called when entering the blop production.
	EnterBlop(c *BlopContext)

	// EnterColop is called when entering the colop production.
	EnterColop(c *ColopContext)

	// EnterCtx is called when entering the ctx production.
	EnterCtx(c *CtxContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitRowop is called when exiting the rowop production.
	ExitRowop(c *RowopContext)

	// ExitBlop is called when exiting the blop production.
	ExitBlop(c *BlopContext)

	// ExitColop is called when exiting the colop production.
	ExitColop(c *ColopContext)

	// ExitCtx is called when exiting the ctx production.
	ExitCtx(c *CtxContext)
}
