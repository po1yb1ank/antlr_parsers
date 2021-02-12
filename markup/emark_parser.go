// Code generated from emark.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // emark

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 21, 110,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	3, 2, 3, 2, 3, 2, 6, 2, 18, 10, 2, 13, 2, 14, 2, 19, 3, 2, 3, 2, 3, 2,
	6, 2, 25, 10, 2, 13, 2, 14, 2, 26, 3, 2, 3, 2, 7, 2, 31, 10, 2, 12, 2,
	14, 2, 34, 11, 2, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 40, 10, 3, 12, 3, 14, 3,
	43, 11, 3, 3, 3, 3, 3, 7, 3, 47, 10, 3, 12, 3, 14, 3, 50, 11, 3, 3, 3,
	3, 3, 7, 3, 54, 10, 3, 12, 3, 14, 3, 57, 11, 3, 3, 3, 3, 3, 6, 3, 61, 10,
	3, 13, 3, 14, 3, 62, 5, 3, 65, 10, 3, 3, 4, 3, 4, 7, 4, 69, 10, 4, 12,
	4, 14, 4, 72, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 6, 5, 79, 10, 5, 13,
	5, 14, 5, 80, 3, 5, 3, 5, 3, 5, 6, 5, 86, 10, 5, 13, 5, 14, 5, 87, 3, 5,
	3, 5, 7, 5, 92, 10, 5, 12, 5, 14, 5, 95, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	7, 6, 101, 10, 6, 12, 6, 14, 6, 104, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2, 3, 4, 2, 15, 16, 18, 18, 2, 118, 2,
	14, 3, 2, 2, 2, 4, 64, 3, 2, 2, 2, 6, 66, 3, 2, 2, 2, 8, 75, 3, 2, 2, 2,
	10, 98, 3, 2, 2, 2, 12, 107, 3, 2, 2, 2, 14, 15, 7, 3, 2, 2, 15, 17, 7,
	5, 2, 2, 16, 18, 7, 16, 2, 2, 17, 16, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19,
	17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 22, 7, 18,
	2, 2, 22, 24, 7, 6, 2, 2, 23, 25, 7, 16, 2, 2, 24, 23, 3, 2, 2, 2, 25,
	26, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 28, 3, 2, 2,
	2, 28, 32, 7, 19, 2, 2, 29, 31, 5, 4, 3, 2, 30, 29, 3, 2, 2, 2, 31, 34,
	3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 35, 3, 2, 2, 2,
	34, 32, 3, 2, 2, 2, 35, 36, 7, 4, 2, 2, 36, 37, 7, 2, 2, 3, 37, 3, 3, 2,
	2, 2, 38, 40, 5, 12, 7, 2, 39, 38, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41,
	39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 44, 3, 2, 2, 2, 43, 41, 3, 2, 2,
	2, 44, 65, 5, 6, 4, 2, 45, 47, 5, 12, 7, 2, 46, 45, 3, 2, 2, 2, 47, 50,
	3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 51, 3, 2, 2, 2,
	50, 48, 3, 2, 2, 2, 51, 65, 5, 8, 5, 2, 52, 54, 5, 12, 7, 2, 53, 52, 3,
	2, 2, 2, 54, 57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56,
	58, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 58, 65, 5, 10, 6, 2, 59, 61, 5, 12,
	7, 2, 60, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63,
	3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64, 41, 3, 2, 2, 2, 64, 48, 3, 2, 2, 2,
	64, 55, 3, 2, 2, 2, 64, 60, 3, 2, 2, 2, 65, 5, 3, 2, 2, 2, 66, 70, 7, 7,
	2, 2, 67, 69, 5, 4, 3, 2, 68, 67, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68,
	3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 73, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2,
	73, 74, 7, 8, 2, 2, 74, 7, 3, 2, 2, 2, 75, 76, 7, 3, 2, 2, 76, 78, 7, 5,
	2, 2, 77, 79, 7, 16, 2, 2, 78, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80,
	78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 7, 18,
	2, 2, 83, 85, 7, 6, 2, 2, 84, 86, 7, 16, 2, 2, 85, 84, 3, 2, 2, 2, 86,
	87, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 89, 3, 2, 2,
	2, 89, 93, 7, 19, 2, 2, 90, 92, 5, 4, 3, 2, 91, 90, 3, 2, 2, 2, 92, 95,
	3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 96, 3, 2, 2, 2,
	95, 93, 3, 2, 2, 2, 96, 97, 7, 4, 2, 2, 97, 9, 3, 2, 2, 2, 98, 102, 7,
	9, 2, 2, 99, 101, 5, 4, 3, 2, 100, 99, 3, 2, 2, 2, 101, 104, 3, 2, 2, 2,
	102, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 105, 3, 2, 2, 2, 104,
	102, 3, 2, 2, 2, 105, 106, 7, 10, 2, 2, 106, 11, 3, 2, 2, 2, 107, 108,
	9, 2, 2, 2, 108, 13, 3, 2, 2, 2, 15, 19, 26, 32, 41, 48, 55, 62, 64, 70,
	80, 87, 93, 102,
}
var literalNames = []string{
	"", "'<block '", "'</block>'", "'rows='", "'columns='", "'<row>'", "'</row>'",
	"'<column'", "'</column>'", "'walign'", "'halign'", "'bgcolor'", "'width'",
	"", "", "'textcolor'", "' '", "'>'", "'<'",
}
var symbolicNames = []string{
	"", "Block", "BlockEnd", "Rows", "Columns", "Row", "RowEnd", "Column",
	"ColumnEnd", "Walign", "Halign", "BgColor", "Width", "Ctx", "Number", "TextColor",
	"Spacebar", "CloseTag", "OpenTag", "WS",
}

var ruleNames = []string{
	"program", "tag", "rowop", "blop", "colop", "ctx",
}

type emarkParser struct {
	*antlr.BaseParser
}

// NewemarkParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *emarkParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewemarkParser(input antlr.TokenStream) *emarkParser {
	this := new(emarkParser)
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
	this.GrammarFileName = "emark.g4"

	return this
}

// emarkParser tokens.
const (
	emarkParserEOF       = antlr.TokenEOF
	emarkParserBlock     = 1
	emarkParserBlockEnd  = 2
	emarkParserRows      = 3
	emarkParserColumns   = 4
	emarkParserRow       = 5
	emarkParserRowEnd    = 6
	emarkParserColumn    = 7
	emarkParserColumnEnd = 8
	emarkParserWalign    = 9
	emarkParserHalign    = 10
	emarkParserBgColor   = 11
	emarkParserWidth     = 12
	emarkParserCtx       = 13
	emarkParserNumber    = 14
	emarkParserTextColor = 15
	emarkParserSpacebar  = 16
	emarkParserCloseTag  = 17
	emarkParserOpenTag   = 18
	emarkParserWS        = 19
)

// emarkParser rules.
const (
	emarkParserRULE_program = 0
	emarkParserRULE_tag     = 1
	emarkParserRULE_rowop   = 2
	emarkParserRULE_blop    = 3
	emarkParserRULE_colop   = 4
	emarkParserRULE_ctx     = 5
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
	p.RuleIndex = emarkParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emarkParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Block() antlr.TerminalNode {
	return s.GetToken(emarkParserBlock, 0)
}

func (s *ProgramContext) Rows() antlr.TerminalNode {
	return s.GetToken(emarkParserRows, 0)
}

func (s *ProgramContext) Spacebar() antlr.TerminalNode {
	return s.GetToken(emarkParserSpacebar, 0)
}

func (s *ProgramContext) Columns() antlr.TerminalNode {
	return s.GetToken(emarkParserColumns, 0)
}

func (s *ProgramContext) CloseTag() antlr.TerminalNode {
	return s.GetToken(emarkParserCloseTag, 0)
}

func (s *ProgramContext) BlockEnd() antlr.TerminalNode {
	return s.GetToken(emarkParserBlockEnd, 0)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(emarkParserEOF, 0)
}

func (s *ProgramContext) AllNumber() []antlr.TerminalNode {
	return s.GetTokens(emarkParserNumber)
}

func (s *ProgramContext) Number(i int) antlr.TerminalNode {
	return s.GetToken(emarkParserNumber, i)
}

func (s *ProgramContext) AllTag() []ITagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagContext)(nil)).Elem())
	var tst = make([]ITagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagContext)
		}
	}

	return tst
}

func (s *ProgramContext) Tag(i int) ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emarkListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emarkListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *emarkParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, emarkParserRULE_program)
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
		p.SetState(12)
		p.Match(emarkParserBlock)
	}
	{
		p.SetState(13)
		p.Match(emarkParserRows)
	}
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == emarkParserNumber {
		{
			p.SetState(14)
			p.Match(emarkParserNumber)
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(19)
		p.Match(emarkParserSpacebar)
	}
	{
		p.SetState(20)
		p.Match(emarkParserColumns)
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == emarkParserNumber {
		{
			p.SetState(21)
			p.Match(emarkParserNumber)
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(26)
		p.Match(emarkParserCloseTag)
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<emarkParserBlock)|(1<<emarkParserRow)|(1<<emarkParserColumn)|(1<<emarkParserCtx)|(1<<emarkParserNumber)|(1<<emarkParserSpacebar))) != 0 {
		{
			p.SetState(27)
			p.Tag()
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(33)
		p.Match(emarkParserBlockEnd)
	}
	{
		p.SetState(34)
		p.Match(emarkParserEOF)
	}

	return localctx
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emarkParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emarkParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) Rowop() IRowopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRowopContext)
}

func (s *TagContext) AllCtx() []ICtxContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICtxContext)(nil)).Elem())
	var tst = make([]ICtxContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICtxContext)
		}
	}

	return tst
}

func (s *TagContext) Ctx(i int) ICtxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICtxContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICtxContext)
}

func (s *TagContext) Blop() IBlopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlopContext)
}

func (s *TagContext) Colop() IColopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColopContext)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emarkListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emarkListener); ok {
		listenerT.ExitTag(s)
	}
}

func (p *emarkParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, emarkParserRULE_tag)
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

	var _alt int

	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<emarkParserCtx)|(1<<emarkParserNumber)|(1<<emarkParserSpacebar))) != 0 {
			{
				p.SetState(36)
				p.Ctx()
			}

			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(42)
			p.Rowop()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<emarkParserCtx)|(1<<emarkParserNumber)|(1<<emarkParserSpacebar))) != 0 {
			{
				p.SetState(43)
				p.Ctx()
			}

			p.SetState(48)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(49)
			p.Blop()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<emarkParserCtx)|(1<<emarkParserNumber)|(1<<emarkParserSpacebar))) != 0 {
			{
				p.SetState(50)
				p.Ctx()
			}

			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(56)
			p.Colop()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(57)
					p.Ctx()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(60)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IRowopContext is an interface to support dynamic dispatch.
type IRowopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRowopContext differentiates from other interfaces.
	IsRowopContext()
}

type RowopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowopContext() *RowopContext {
	var p = new(RowopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emarkParserRULE_rowop
	return p
}

func (*RowopContext) IsRowopContext() {}

func NewRowopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowopContext {
	var p = new(RowopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emarkParserRULE_rowop

	return p
}

func (s *RowopContext) GetParser() antlr.Parser { return s.parser }

func (s *RowopContext) Row() antlr.TerminalNode {
	return s.GetToken(emarkParserRow, 0)
}

func (s *RowopContext) RowEnd() antlr.TerminalNode {
	return s.GetToken(emarkParserRowEnd, 0)
}

func (s *RowopContext) AllTag() []ITagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagContext)(nil)).Elem())
	var tst = make([]ITagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagContext)
		}
	}

	return tst
}

func (s *RowopContext) Tag(i int) ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *RowopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emarkListener); ok {
		listenerT.EnterRowop(s)
	}
}

func (s *RowopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emarkListener); ok {
		listenerT.ExitRowop(s)
	}
}

func (p *emarkParser) Rowop() (localctx IRowopContext) {
	localctx = NewRowopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, emarkParserRULE_rowop)
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
		p.SetState(64)
		p.Match(emarkParserRow)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<emarkParserBlock)|(1<<emarkParserRow)|(1<<emarkParserColumn)|(1<<emarkParserCtx)|(1<<emarkParserNumber)|(1<<emarkParserSpacebar))) != 0 {
		{
			p.SetState(65)
			p.Tag()
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(71)
		p.Match(emarkParserRowEnd)
	}

	return localctx
}

// IBlopContext is an interface to support dynamic dispatch.
type IBlopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlopContext differentiates from other interfaces.
	IsBlopContext()
}

type BlopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlopContext() *BlopContext {
	var p = new(BlopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emarkParserRULE_blop
	return p
}

func (*BlopContext) IsBlopContext() {}

func NewBlopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlopContext {
	var p = new(BlopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emarkParserRULE_blop

	return p
}

func (s *BlopContext) GetParser() antlr.Parser { return s.parser }

func (s *BlopContext) Block() antlr.TerminalNode {
	return s.GetToken(emarkParserBlock, 0)
}

func (s *BlopContext) Rows() antlr.TerminalNode {
	return s.GetToken(emarkParserRows, 0)
}

func (s *BlopContext) Spacebar() antlr.TerminalNode {
	return s.GetToken(emarkParserSpacebar, 0)
}

func (s *BlopContext) Columns() antlr.TerminalNode {
	return s.GetToken(emarkParserColumns, 0)
}

func (s *BlopContext) CloseTag() antlr.TerminalNode {
	return s.GetToken(emarkParserCloseTag, 0)
}

func (s *BlopContext) BlockEnd() antlr.TerminalNode {
	return s.GetToken(emarkParserBlockEnd, 0)
}

func (s *BlopContext) AllNumber() []antlr.TerminalNode {
	return s.GetTokens(emarkParserNumber)
}

func (s *BlopContext) Number(i int) antlr.TerminalNode {
	return s.GetToken(emarkParserNumber, i)
}

func (s *BlopContext) AllTag() []ITagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagContext)(nil)).Elem())
	var tst = make([]ITagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagContext)
		}
	}

	return tst
}

func (s *BlopContext) Tag(i int) ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *BlopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emarkListener); ok {
		listenerT.EnterBlop(s)
	}
}

func (s *BlopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emarkListener); ok {
		listenerT.ExitBlop(s)
	}
}

func (p *emarkParser) Blop() (localctx IBlopContext) {
	localctx = NewBlopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, emarkParserRULE_blop)
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
		p.SetState(73)
		p.Match(emarkParserBlock)
	}
	{
		p.SetState(74)
		p.Match(emarkParserRows)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == emarkParserNumber {
		{
			p.SetState(75)
			p.Match(emarkParserNumber)
		}

		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.Match(emarkParserSpacebar)
	}
	{
		p.SetState(81)
		p.Match(emarkParserColumns)
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == emarkParserNumber {
		{
			p.SetState(82)
			p.Match(emarkParserNumber)
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(emarkParserCloseTag)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<emarkParserBlock)|(1<<emarkParserRow)|(1<<emarkParserColumn)|(1<<emarkParserCtx)|(1<<emarkParserNumber)|(1<<emarkParserSpacebar))) != 0 {
		{
			p.SetState(88)
			p.Tag()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
		p.Match(emarkParserBlockEnd)
	}

	return localctx
}

// IColopContext is an interface to support dynamic dispatch.
type IColopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColopContext differentiates from other interfaces.
	IsColopContext()
}

type ColopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColopContext() *ColopContext {
	var p = new(ColopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emarkParserRULE_colop
	return p
}

func (*ColopContext) IsColopContext() {}

func NewColopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColopContext {
	var p = new(ColopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emarkParserRULE_colop

	return p
}

func (s *ColopContext) GetParser() antlr.Parser { return s.parser }

func (s *ColopContext) Column() antlr.TerminalNode {
	return s.GetToken(emarkParserColumn, 0)
}

func (s *ColopContext) ColumnEnd() antlr.TerminalNode {
	return s.GetToken(emarkParserColumnEnd, 0)
}

func (s *ColopContext) AllTag() []ITagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagContext)(nil)).Elem())
	var tst = make([]ITagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagContext)
		}
	}

	return tst
}

func (s *ColopContext) Tag(i int) ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *ColopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emarkListener); ok {
		listenerT.EnterColop(s)
	}
}

func (s *ColopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emarkListener); ok {
		listenerT.ExitColop(s)
	}
}

func (p *emarkParser) Colop() (localctx IColopContext) {
	localctx = NewColopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, emarkParserRULE_colop)
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
		p.SetState(96)
		p.Match(emarkParserColumn)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<emarkParserBlock)|(1<<emarkParserRow)|(1<<emarkParserColumn)|(1<<emarkParserCtx)|(1<<emarkParserNumber)|(1<<emarkParserSpacebar))) != 0 {
		{
			p.SetState(97)
			p.Tag()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(emarkParserColumnEnd)
	}

	return localctx
}

// ICtxContext is an interface to support dynamic dispatch.
type ICtxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCtxContext differentiates from other interfaces.
	IsCtxContext()
}

type CtxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCtxContext() *CtxContext {
	var p = new(CtxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emarkParserRULE_ctx
	return p
}

func (*CtxContext) IsCtxContext() {}

func NewCtxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CtxContext {
	var p = new(CtxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emarkParserRULE_ctx

	return p
}

func (s *CtxContext) GetParser() antlr.Parser { return s.parser }

func (s *CtxContext) Spacebar() antlr.TerminalNode {
	return s.GetToken(emarkParserSpacebar, 0)
}

func (s *CtxContext) Number() antlr.TerminalNode {
	return s.GetToken(emarkParserNumber, 0)
}

func (s *CtxContext) Ctx() antlr.TerminalNode {
	return s.GetToken(emarkParserCtx, 0)
}

func (s *CtxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CtxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CtxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emarkListener); ok {
		listenerT.EnterCtx(s)
	}
}

func (s *CtxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emarkListener); ok {
		listenerT.ExitCtx(s)
	}
}

func (p *emarkParser) Ctx() (localctx ICtxContext) {
	localctx = NewCtxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, emarkParserRULE_ctx)
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
		p.SetState(105)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<emarkParserCtx)|(1<<emarkParserNumber)|(1<<emarkParserSpacebar))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
