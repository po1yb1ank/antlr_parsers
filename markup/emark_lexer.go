// Code generated from emark.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 21, 164,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 2, 2, 23, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 2, 29,
	2, 31, 15, 33, 16, 35, 17, 37, 18, 39, 19, 41, 20, 43, 21, 3, 2, 5, 3,
	2, 50, 59, 4, 2, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 161,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3,
	2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 3, 45,
	3, 2, 2, 2, 5, 53, 3, 2, 2, 2, 7, 62, 3, 2, 2, 2, 9, 68, 3, 2, 2, 2, 11,
	77, 3, 2, 2, 2, 13, 83, 3, 2, 2, 2, 15, 90, 3, 2, 2, 2, 17, 98, 3, 2, 2,
	2, 19, 108, 3, 2, 2, 2, 21, 115, 3, 2, 2, 2, 23, 122, 3, 2, 2, 2, 25, 130,
	3, 2, 2, 2, 27, 136, 3, 2, 2, 2, 29, 138, 3, 2, 2, 2, 31, 140, 3, 2, 2,
	2, 33, 142, 3, 2, 2, 2, 35, 144, 3, 2, 2, 2, 37, 154, 3, 2, 2, 2, 39, 156,
	3, 2, 2, 2, 41, 158, 3, 2, 2, 2, 43, 160, 3, 2, 2, 2, 45, 46, 7, 62, 2,
	2, 46, 47, 7, 100, 2, 2, 47, 48, 7, 110, 2, 2, 48, 49, 7, 113, 2, 2, 49,
	50, 7, 101, 2, 2, 50, 51, 7, 109, 2, 2, 51, 52, 7, 34, 2, 2, 52, 4, 3,
	2, 2, 2, 53, 54, 7, 62, 2, 2, 54, 55, 7, 49, 2, 2, 55, 56, 7, 100, 2, 2,
	56, 57, 7, 110, 2, 2, 57, 58, 7, 113, 2, 2, 58, 59, 7, 101, 2, 2, 59, 60,
	7, 109, 2, 2, 60, 61, 7, 64, 2, 2, 61, 6, 3, 2, 2, 2, 62, 63, 7, 116, 2,
	2, 63, 64, 7, 113, 2, 2, 64, 65, 7, 121, 2, 2, 65, 66, 7, 117, 2, 2, 66,
	67, 7, 63, 2, 2, 67, 8, 3, 2, 2, 2, 68, 69, 7, 101, 2, 2, 69, 70, 7, 113,
	2, 2, 70, 71, 7, 110, 2, 2, 71, 72, 7, 119, 2, 2, 72, 73, 7, 111, 2, 2,
	73, 74, 7, 112, 2, 2, 74, 75, 7, 117, 2, 2, 75, 76, 7, 63, 2, 2, 76, 10,
	3, 2, 2, 2, 77, 78, 7, 62, 2, 2, 78, 79, 7, 116, 2, 2, 79, 80, 7, 113,
	2, 2, 80, 81, 7, 121, 2, 2, 81, 82, 7, 64, 2, 2, 82, 12, 3, 2, 2, 2, 83,
	84, 7, 62, 2, 2, 84, 85, 7, 49, 2, 2, 85, 86, 7, 116, 2, 2, 86, 87, 7,
	113, 2, 2, 87, 88, 7, 121, 2, 2, 88, 89, 7, 64, 2, 2, 89, 14, 3, 2, 2,
	2, 90, 91, 7, 62, 2, 2, 91, 92, 7, 101, 2, 2, 92, 93, 7, 113, 2, 2, 93,
	94, 7, 110, 2, 2, 94, 95, 7, 119, 2, 2, 95, 96, 7, 111, 2, 2, 96, 97, 7,
	112, 2, 2, 97, 16, 3, 2, 2, 2, 98, 99, 7, 62, 2, 2, 99, 100, 7, 49, 2,
	2, 100, 101, 7, 101, 2, 2, 101, 102, 7, 113, 2, 2, 102, 103, 7, 110, 2,
	2, 103, 104, 7, 119, 2, 2, 104, 105, 7, 111, 2, 2, 105, 106, 7, 112, 2,
	2, 106, 107, 7, 64, 2, 2, 107, 18, 3, 2, 2, 2, 108, 109, 7, 121, 2, 2,
	109, 110, 7, 99, 2, 2, 110, 111, 7, 110, 2, 2, 111, 112, 7, 107, 2, 2,
	112, 113, 7, 105, 2, 2, 113, 114, 7, 112, 2, 2, 114, 20, 3, 2, 2, 2, 115,
	116, 7, 106, 2, 2, 116, 117, 7, 99, 2, 2, 117, 118, 7, 110, 2, 2, 118,
	119, 7, 107, 2, 2, 119, 120, 7, 105, 2, 2, 120, 121, 7, 112, 2, 2, 121,
	22, 3, 2, 2, 2, 122, 123, 7, 100, 2, 2, 123, 124, 7, 105, 2, 2, 124, 125,
	7, 101, 2, 2, 125, 126, 7, 113, 2, 2, 126, 127, 7, 110, 2, 2, 127, 128,
	7, 113, 2, 2, 128, 129, 7, 116, 2, 2, 129, 24, 3, 2, 2, 2, 130, 131, 7,
	121, 2, 2, 131, 132, 7, 107, 2, 2, 132, 133, 7, 102, 2, 2, 133, 134, 7,
	118, 2, 2, 134, 135, 7, 106, 2, 2, 135, 26, 3, 2, 2, 2, 136, 137, 9, 2,
	2, 2, 137, 28, 3, 2, 2, 2, 138, 139, 9, 3, 2, 2, 139, 30, 3, 2, 2, 2, 140,
	141, 5, 29, 15, 2, 141, 32, 3, 2, 2, 2, 142, 143, 5, 27, 14, 2, 143, 34,
	3, 2, 2, 2, 144, 145, 7, 118, 2, 2, 145, 146, 7, 103, 2, 2, 146, 147, 7,
	122, 2, 2, 147, 148, 7, 118, 2, 2, 148, 149, 7, 101, 2, 2, 149, 150, 7,
	113, 2, 2, 150, 151, 7, 110, 2, 2, 151, 152, 7, 113, 2, 2, 152, 153, 7,
	116, 2, 2, 153, 36, 3, 2, 2, 2, 154, 155, 7, 34, 2, 2, 155, 38, 3, 2, 2,
	2, 156, 157, 7, 64, 2, 2, 157, 40, 3, 2, 2, 2, 158, 159, 7, 62, 2, 2, 159,
	42, 3, 2, 2, 2, 160, 161, 9, 4, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163, 8,
	22, 2, 2, 163, 44, 3, 2, 2, 2, 3, 2, 3, 2, 3, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'<block '", "'</block>'", "'rows='", "'columns='", "'<row>'", "'</row>'",
	"'<column'", "'</column>'", "'walign'", "'halign'", "'bgcolor'", "'width'",
	"", "", "'textcolor'", "' '", "'>'", "'<'",
}

var lexerSymbolicNames = []string{
	"", "Block", "BlockEnd", "Rows", "Columns", "Row", "RowEnd", "Column",
	"ColumnEnd", "Walign", "Halign", "BgColor", "Width", "Ctx", "Number", "TextColor",
	"Spacebar", "CloseTag", "OpenTag", "WS",
}

var lexerRuleNames = []string{
	"Block", "BlockEnd", "Rows", "Columns", "Row", "RowEnd", "Column", "ColumnEnd",
	"Walign", "Halign", "BgColor", "Width", "DIGIT", "TXT", "Ctx", "Number",
	"TextColor", "Spacebar", "CloseTag", "OpenTag", "WS",
}

type emarkLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewemarkLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *emarkLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewemarkLexer(input antlr.CharStream) *emarkLexer {
	l := new(emarkLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "emark.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// emarkLexer tokens.
const (
	emarkLexerBlock     = 1
	emarkLexerBlockEnd  = 2
	emarkLexerRows      = 3
	emarkLexerColumns   = 4
	emarkLexerRow       = 5
	emarkLexerRowEnd    = 6
	emarkLexerColumn    = 7
	emarkLexerColumnEnd = 8
	emarkLexerWalign    = 9
	emarkLexerHalign    = 10
	emarkLexerBgColor   = 11
	emarkLexerWidth     = 12
	emarkLexerCtx       = 13
	emarkLexerNumber    = 14
	emarkLexerTextColor = 15
	emarkLexerSpacebar  = 16
	emarkLexerCloseTag  = 17
	emarkLexerOpenTag   = 18
	emarkLexerWS        = 19
)
