// Code generated from lightc.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // lightc

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 94, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 5, 4, 52, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 80, 10, 9, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	2, 2, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2, 5, 4, 2, 6, 7, 9,
	9, 3, 2, 25, 26, 4, 2, 15, 16, 22, 23, 2, 87, 2, 24, 3, 2, 2, 2, 4, 33,
	3, 2, 2, 2, 6, 51, 3, 2, 2, 2, 8, 53, 3, 2, 2, 2, 10, 58, 3, 2, 2, 2, 12,
	61, 3, 2, 2, 2, 14, 63, 3, 2, 2, 2, 16, 79, 3, 2, 2, 2, 18, 81, 3, 2, 2,
	2, 20, 83, 3, 2, 2, 2, 22, 88, 3, 2, 2, 2, 24, 25, 5, 4, 3, 2, 25, 26,
	7, 3, 2, 2, 26, 27, 7, 10, 2, 2, 27, 28, 7, 11, 2, 2, 28, 29, 7, 12, 2,
	2, 29, 30, 7, 13, 2, 2, 30, 31, 5, 6, 4, 2, 31, 32, 7, 14, 2, 2, 32, 3,
	3, 2, 2, 2, 33, 34, 9, 2, 2, 2, 34, 5, 3, 2, 2, 2, 35, 36, 5, 8, 5, 2,
	36, 37, 7, 19, 2, 2, 37, 38, 5, 6, 4, 2, 38, 52, 3, 2, 2, 2, 39, 40, 7,
	13, 2, 2, 40, 41, 5, 6, 4, 2, 41, 42, 7, 14, 2, 2, 42, 43, 5, 6, 4, 2,
	43, 52, 3, 2, 2, 2, 44, 45, 5, 14, 8, 2, 45, 46, 5, 6, 4, 2, 46, 52, 3,
	2, 2, 2, 47, 48, 5, 20, 11, 2, 48, 49, 5, 6, 4, 2, 49, 52, 3, 2, 2, 2,
	50, 52, 5, 22, 12, 2, 51, 35, 3, 2, 2, 2, 51, 39, 3, 2, 2, 2, 51, 44, 3,
	2, 2, 2, 51, 47, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 7, 3, 2, 2, 2, 53,
	54, 5, 4, 3, 2, 54, 55, 7, 3, 2, 2, 55, 56, 7, 26, 2, 2, 56, 57, 5, 10,
	6, 2, 57, 9, 3, 2, 2, 2, 58, 59, 7, 21, 2, 2, 59, 60, 5, 12, 7, 2, 60,
	11, 3, 2, 2, 2, 61, 62, 9, 3, 2, 2, 62, 13, 3, 2, 2, 2, 63, 64, 7, 4, 2,
	2, 64, 65, 7, 13, 2, 2, 65, 66, 5, 8, 5, 2, 66, 67, 7, 19, 2, 2, 67, 68,
	5, 16, 9, 2, 68, 69, 7, 19, 2, 2, 69, 70, 7, 14, 2, 2, 70, 15, 3, 2, 2,
	2, 71, 72, 7, 26, 2, 2, 72, 73, 5, 18, 10, 2, 73, 74, 7, 26, 2, 2, 74,
	80, 3, 2, 2, 2, 75, 76, 7, 25, 2, 2, 76, 77, 5, 18, 10, 2, 77, 78, 7, 26,
	2, 2, 78, 80, 3, 2, 2, 2, 79, 71, 3, 2, 2, 2, 79, 75, 3, 2, 2, 2, 80, 17,
	3, 2, 2, 2, 81, 82, 9, 4, 2, 2, 82, 19, 3, 2, 2, 2, 83, 84, 7, 5, 2, 2,
	84, 85, 7, 11, 2, 2, 85, 86, 5, 16, 9, 2, 86, 87, 7, 12, 2, 2, 87, 21,
	3, 2, 2, 2, 88, 89, 7, 8, 2, 2, 89, 90, 7, 3, 2, 2, 90, 91, 7, 25, 2, 2,
	91, 92, 7, 19, 2, 2, 92, 23, 3, 2, 2, 2, 4, 51, 79,
}
var literalNames = []string{
	"", "' '", "'for'", "'if'", "'bool'", "'int'", "'return'", "'void'", "'main'",
	"'('", "')'", "'{'", "'}'", "'<'", "'>'", "'?'", "':'", "';'", "','", "'='",
	"'=='", "'!='",
}
var symbolicNames = []string{
	"", "Space", "For", "If", "Bool", "Int", "Return", "Void", "Main", "LeftParen",
	"RightParen", "LeftBrace", "RightBrace", "Less", "Greater", "Question",
	"Colon", "Semi", "Comma", "Assign", "Equal", "NotEqual", "WS", "Number",
	"Identifier",
}

var ruleNames = []string{
	"program", "typedef", "statement", "declaration", "assign", "assign_end",
	"forloop", "bool_expression", "relop", "ifstat", "returnval",
}

type lightcParser struct {
	*antlr.BaseParser
}

// NewlightcParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *lightcParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewlightcParser(input antlr.TokenStream) *lightcParser {
	this := new(lightcParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "lightc.g4"

	return this
}

// lightcParser tokens.
const (
	lightcParserEOF        = antlr.TokenEOF
	lightcParserSpace      = 1
	lightcParserFor        = 2
	lightcParserIf         = 3
	lightcParserBool       = 4
	lightcParserInt        = 5
	lightcParserReturn     = 6
	lightcParserVoid       = 7
	lightcParserMain       = 8
	lightcParserLeftParen  = 9
	lightcParserRightParen = 10
	lightcParserLeftBrace  = 11
	lightcParserRightBrace = 12
	lightcParserLess       = 13
	lightcParserGreater    = 14
	lightcParserQuestion   = 15
	lightcParserColon      = 16
	lightcParserSemi       = 17
	lightcParserComma      = 18
	lightcParserAssign     = 19
	lightcParserEqual      = 20
	lightcParserNotEqual   = 21
	lightcParserWS         = 22
	lightcParserNumber     = 23
	lightcParserIdentifier = 24
)

// lightcParser rules.
const (
	lightcParserRULE_program         = 0
	lightcParserRULE_typedef         = 1
	lightcParserRULE_statement       = 2
	lightcParserRULE_declaration     = 3
	lightcParserRULE_assign          = 4
	lightcParserRULE_assign_end      = 5
	lightcParserRULE_forloop         = 6
	lightcParserRULE_bool_expression = 7
	lightcParserRULE_relop           = 8
	lightcParserRULE_ifstat          = 9
	lightcParserRULE_returnval       = 10
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lightcParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lightcParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Typedef() ITypedefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypedefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypedefContext)
}

func (s *ProgramContext) Space() antlr.TerminalNode {
	return s.GetToken(lightcParserSpace, 0)
}

func (s *ProgramContext) Main() antlr.TerminalNode {
	return s.GetToken(lightcParserMain, 0)
}

func (s *ProgramContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(lightcParserLeftParen, 0)
}

func (s *ProgramContext) RightParen() antlr.TerminalNode {
	return s.GetToken(lightcParserRightParen, 0)
}

func (s *ProgramContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(lightcParserLeftBrace, 0)
}

func (s *ProgramContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(lightcParserRightBrace, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *lightcParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, lightcParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Typedef()
	}
	{
		p.SetState(23)
		p.Match(lightcParserSpace)
	}
	{
		p.SetState(24)
		p.Match(lightcParserMain)
	}
	{
		p.SetState(25)
		p.Match(lightcParserLeftParen)
	}
	{
		p.SetState(26)
		p.Match(lightcParserRightParen)
	}
	{
		p.SetState(27)
		p.Match(lightcParserLeftBrace)
	}
	{
		p.SetState(28)
		p.Statement()
	}
	{
		p.SetState(29)
		p.Match(lightcParserRightBrace)
	}

	return localctx
}

// ITypedefContext is an interface to support dynamic dispatch.
type ITypedefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypedefContext differentiates from other interfaces.
	IsTypedefContext()
}

type TypedefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedefContext() *TypedefContext {
	var p = new(TypedefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lightcParserRULE_typedef
	return p
}

func (*TypedefContext) IsTypedefContext() {}

func NewTypedefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedefContext {
	var p = new(TypedefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lightcParserRULE_typedef

	return p
}

func (s *TypedefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedefContext) Int() antlr.TerminalNode {
	return s.GetToken(lightcParserInt, 0)
}

func (s *TypedefContext) Bool() antlr.TerminalNode {
	return s.GetToken(lightcParserBool, 0)
}

func (s *TypedefContext) Void() antlr.TerminalNode {
	return s.GetToken(lightcParserVoid, 0)
}

func (s *TypedefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.EnterTypedef(s)
	}
}

func (s *TypedefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.ExitTypedef(s)
	}
}

func (p *lightcParser) Typedef() (localctx ITypedefContext) {
	localctx = NewTypedefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, lightcParserRULE_typedef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<lightcParserBool)|(1<<lightcParserInt)|(1<<lightcParserVoid))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lightcParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lightcParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) Semi() antlr.TerminalNode {
	return s.GetToken(lightcParserSemi, 0)
}

func (s *StatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(lightcParserLeftBrace, 0)
}

func (s *StatementContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(lightcParserRightBrace, 0)
}

func (s *StatementContext) Forloop() IForloopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForloopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForloopContext)
}

func (s *StatementContext) Ifstat() IIfstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfstatContext)
}

func (s *StatementContext) Returnval() IReturnvalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnvalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnvalContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *lightcParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, lightcParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case lightcParserBool, lightcParserInt, lightcParserVoid:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Declaration()
		}
		{
			p.SetState(34)
			p.Match(lightcParserSemi)
		}
		{
			p.SetState(35)
			p.Statement()
		}

	case lightcParserLeftBrace:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(lightcParserLeftBrace)
		}
		{
			p.SetState(38)
			p.Statement()
		}
		{
			p.SetState(39)
			p.Match(lightcParserRightBrace)
		}
		{
			p.SetState(40)
			p.Statement()
		}

	case lightcParserFor:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.Forloop()
		}
		{
			p.SetState(43)
			p.Statement()
		}

	case lightcParserIf:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(45)
			p.Ifstat()
		}
		{
			p.SetState(46)
			p.Statement()
		}

	case lightcParserReturn:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(48)
			p.Returnval()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lightcParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lightcParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Typedef() ITypedefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypedefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypedefContext)
}

func (s *DeclarationContext) Space() antlr.TerminalNode {
	return s.GetToken(lightcParserSpace, 0)
}

func (s *DeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(lightcParserIdentifier, 0)
}

func (s *DeclarationContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *lightcParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, lightcParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Typedef()
	}
	{
		p.SetState(52)
		p.Match(lightcParserSpace)
	}
	{
		p.SetState(53)
		p.Match(lightcParserIdentifier)
	}
	{
		p.SetState(54)
		p.Assign()
	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lightcParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lightcParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) Assign() antlr.TerminalNode {
	return s.GetToken(lightcParserAssign, 0)
}

func (s *AssignContext) Assign_end() IAssign_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssign_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssign_endContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *lightcParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, lightcParserRULE_assign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(lightcParserAssign)
	}
	{
		p.SetState(57)
		p.Assign_end()
	}

	return localctx
}

// IAssign_endContext is an interface to support dynamic dispatch.
type IAssign_endContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssign_endContext differentiates from other interfaces.
	IsAssign_endContext()
}

type Assign_endContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_endContext() *Assign_endContext {
	var p = new(Assign_endContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lightcParserRULE_assign_end
	return p
}

func (*Assign_endContext) IsAssign_endContext() {}

func NewAssign_endContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_endContext {
	var p = new(Assign_endContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lightcParserRULE_assign_end

	return p
}

func (s *Assign_endContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_endContext) Identifier() antlr.TerminalNode {
	return s.GetToken(lightcParserIdentifier, 0)
}

func (s *Assign_endContext) Number() antlr.TerminalNode {
	return s.GetToken(lightcParserNumber, 0)
}

func (s *Assign_endContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_endContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_endContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.EnterAssign_end(s)
	}
}

func (s *Assign_endContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.ExitAssign_end(s)
	}
}

func (p *lightcParser) Assign_end() (localctx IAssign_endContext) {
	localctx = NewAssign_endContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, lightcParserRULE_assign_end)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		_la = p.GetTokenStream().LA(1)

		if !(_la == lightcParserNumber || _la == lightcParserIdentifier) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IForloopContext is an interface to support dynamic dispatch.
type IForloopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForloopContext differentiates from other interfaces.
	IsForloopContext()
}

type ForloopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForloopContext() *ForloopContext {
	var p = new(ForloopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lightcParserRULE_forloop
	return p
}

func (*ForloopContext) IsForloopContext() {}

func NewForloopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForloopContext {
	var p = new(ForloopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lightcParserRULE_forloop

	return p
}

func (s *ForloopContext) GetParser() antlr.Parser { return s.parser }

func (s *ForloopContext) For() antlr.TerminalNode {
	return s.GetToken(lightcParserFor, 0)
}

func (s *ForloopContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(lightcParserLeftBrace, 0)
}

func (s *ForloopContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ForloopContext) AllSemi() []antlr.TerminalNode {
	return s.GetTokens(lightcParserSemi)
}

func (s *ForloopContext) Semi(i int) antlr.TerminalNode {
	return s.GetToken(lightcParserSemi, i)
}

func (s *ForloopContext) Bool_expression() IBool_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_expressionContext)
}

func (s *ForloopContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(lightcParserRightBrace, 0)
}

func (s *ForloopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForloopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForloopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.EnterForloop(s)
	}
}

func (s *ForloopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.ExitForloop(s)
	}
}

func (p *lightcParser) Forloop() (localctx IForloopContext) {
	localctx = NewForloopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, lightcParserRULE_forloop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(lightcParserFor)
	}
	{
		p.SetState(62)
		p.Match(lightcParserLeftBrace)
	}
	{
		p.SetState(63)
		p.Declaration()
	}
	{
		p.SetState(64)
		p.Match(lightcParserSemi)
	}
	{
		p.SetState(65)
		p.Bool_expression()
	}
	{
		p.SetState(66)
		p.Match(lightcParserSemi)
	}
	{
		p.SetState(67)
		p.Match(lightcParserRightBrace)
	}

	return localctx
}

// IBool_expressionContext is an interface to support dynamic dispatch.
type IBool_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBool_expressionContext differentiates from other interfaces.
	IsBool_expressionContext()
}

type Bool_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBool_expressionContext() *Bool_expressionContext {
	var p = new(Bool_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lightcParserRULE_bool_expression
	return p
}

func (*Bool_expressionContext) IsBool_expressionContext() {}

func NewBool_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_expressionContext {
	var p = new(Bool_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lightcParserRULE_bool_expression

	return p
}

func (s *Bool_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Bool_expressionContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(lightcParserIdentifier)
}

func (s *Bool_expressionContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(lightcParserIdentifier, i)
}

func (s *Bool_expressionContext) Relop() IRelopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *Bool_expressionContext) Number() antlr.TerminalNode {
	return s.GetToken(lightcParserNumber, 0)
}

func (s *Bool_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.EnterBool_expression(s)
	}
}

func (s *Bool_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.ExitBool_expression(s)
	}
}

func (p *lightcParser) Bool_expression() (localctx IBool_expressionContext) {
	localctx = NewBool_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, lightcParserRULE_bool_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case lightcParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Match(lightcParserIdentifier)
		}
		{
			p.SetState(70)
			p.Relop()
		}
		{
			p.SetState(71)
			p.Match(lightcParserIdentifier)
		}

	case lightcParserNumber:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Match(lightcParserNumber)
		}
		{
			p.SetState(74)
			p.Relop()
		}
		{
			p.SetState(75)
			p.Match(lightcParserIdentifier)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRelopContext is an interface to support dynamic dispatch.
type IRelopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelopContext differentiates from other interfaces.
	IsRelopContext()
}

type RelopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelopContext() *RelopContext {
	var p = new(RelopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lightcParserRULE_relop
	return p
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lightcParserRULE_relop

	return p
}

func (s *RelopContext) GetParser() antlr.Parser { return s.parser }

func (s *RelopContext) Less() antlr.TerminalNode {
	return s.GetToken(lightcParserLess, 0)
}

func (s *RelopContext) Greater() antlr.TerminalNode {
	return s.GetToken(lightcParserGreater, 0)
}

func (s *RelopContext) Equal() antlr.TerminalNode {
	return s.GetToken(lightcParserEqual, 0)
}

func (s *RelopContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(lightcParserNotEqual, 0)
}

func (s *RelopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.EnterRelop(s)
	}
}

func (s *RelopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.ExitRelop(s)
	}
}

func (p *lightcParser) Relop() (localctx IRelopContext) {
	localctx = NewRelopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, lightcParserRULE_relop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<lightcParserLess)|(1<<lightcParserGreater)|(1<<lightcParserEqual)|(1<<lightcParserNotEqual))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIfstatContext is an interface to support dynamic dispatch.
type IIfstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfstatContext differentiates from other interfaces.
	IsIfstatContext()
}

type IfstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstatContext() *IfstatContext {
	var p = new(IfstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lightcParserRULE_ifstat
	return p
}

func (*IfstatContext) IsIfstatContext() {}

func NewIfstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstatContext {
	var p = new(IfstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lightcParserRULE_ifstat

	return p
}

func (s *IfstatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstatContext) If() antlr.TerminalNode {
	return s.GetToken(lightcParserIf, 0)
}

func (s *IfstatContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(lightcParserLeftParen, 0)
}

func (s *IfstatContext) Bool_expression() IBool_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_expressionContext)
}

func (s *IfstatContext) RightParen() antlr.TerminalNode {
	return s.GetToken(lightcParserRightParen, 0)
}

func (s *IfstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.EnterIfstat(s)
	}
}

func (s *IfstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.ExitIfstat(s)
	}
}

func (p *lightcParser) Ifstat() (localctx IIfstatContext) {
	localctx = NewIfstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, lightcParserRULE_ifstat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(lightcParserIf)
	}
	{
		p.SetState(82)
		p.Match(lightcParserLeftParen)
	}
	{
		p.SetState(83)
		p.Bool_expression()
	}
	{
		p.SetState(84)
		p.Match(lightcParserRightParen)
	}

	return localctx
}

// IReturnvalContext is an interface to support dynamic dispatch.
type IReturnvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnvalContext differentiates from other interfaces.
	IsReturnvalContext()
}

type ReturnvalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnvalContext() *ReturnvalContext {
	var p = new(ReturnvalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lightcParserRULE_returnval
	return p
}

func (*ReturnvalContext) IsReturnvalContext() {}

func NewReturnvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnvalContext {
	var p = new(ReturnvalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lightcParserRULE_returnval

	return p
}

func (s *ReturnvalContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnvalContext) Return() antlr.TerminalNode {
	return s.GetToken(lightcParserReturn, 0)
}

func (s *ReturnvalContext) Space() antlr.TerminalNode {
	return s.GetToken(lightcParserSpace, 0)
}

func (s *ReturnvalContext) Number() antlr.TerminalNode {
	return s.GetToken(lightcParserNumber, 0)
}

func (s *ReturnvalContext) Semi() antlr.TerminalNode {
	return s.GetToken(lightcParserSemi, 0)
}

func (s *ReturnvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.EnterReturnval(s)
	}
}

func (s *ReturnvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lightcListener); ok {
		listenerT.ExitReturnval(s)
	}
}

func (p *lightcParser) Returnval() (localctx IReturnvalContext) {
	localctx = NewReturnvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, lightcParserRULE_returnval)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(lightcParserReturn)
	}
	{
		p.SetState(87)
		p.Match(lightcParserSpace)
	}
	{
		p.SetState(88)
		p.Match(lightcParserNumber)
	}
	{
		p.SetState(89)
		p.Match(lightcParserSemi)
	}

	return localctx
}
