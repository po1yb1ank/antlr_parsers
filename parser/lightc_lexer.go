// Code generated from lightc.g4 by ANTLR 4.9. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 26, 136,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 25, 3, 25, 3, 26, 6, 26, 128, 10, 26, 13, 26, 14, 26, 129, 3, 27,
	6, 27, 133, 10, 27, 13, 27, 14, 27, 134, 2, 2, 28, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15,
	29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24,
	47, 2, 49, 2, 51, 25, 53, 26, 3, 2, 5, 5, 2, 11, 12, 15, 15, 34, 34, 5,
	2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 2, 135, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2,
	2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2,
	2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3,
	2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 3, 55,
	3, 2, 2, 2, 5, 57, 3, 2, 2, 2, 7, 61, 3, 2, 2, 2, 9, 64, 3, 2, 2, 2, 11,
	69, 3, 2, 2, 2, 13, 73, 3, 2, 2, 2, 15, 80, 3, 2, 2, 2, 17, 85, 3, 2, 2,
	2, 19, 90, 3, 2, 2, 2, 21, 92, 3, 2, 2, 2, 23, 94, 3, 2, 2, 2, 25, 96,
	3, 2, 2, 2, 27, 98, 3, 2, 2, 2, 29, 100, 3, 2, 2, 2, 31, 102, 3, 2, 2,
	2, 33, 104, 3, 2, 2, 2, 35, 106, 3, 2, 2, 2, 37, 108, 3, 2, 2, 2, 39, 110,
	3, 2, 2, 2, 41, 112, 3, 2, 2, 2, 43, 115, 3, 2, 2, 2, 45, 118, 3, 2, 2,
	2, 47, 122, 3, 2, 2, 2, 49, 124, 3, 2, 2, 2, 51, 127, 3, 2, 2, 2, 53, 132,
	3, 2, 2, 2, 55, 56, 7, 34, 2, 2, 56, 4, 3, 2, 2, 2, 57, 58, 7, 104, 2,
	2, 58, 59, 7, 113, 2, 2, 59, 60, 7, 116, 2, 2, 60, 6, 3, 2, 2, 2, 61, 62,
	7, 107, 2, 2, 62, 63, 7, 104, 2, 2, 63, 8, 3, 2, 2, 2, 64, 65, 7, 100,
	2, 2, 65, 66, 7, 113, 2, 2, 66, 67, 7, 113, 2, 2, 67, 68, 7, 110, 2, 2,
	68, 10, 3, 2, 2, 2, 69, 70, 7, 107, 2, 2, 70, 71, 7, 112, 2, 2, 71, 72,
	7, 118, 2, 2, 72, 12, 3, 2, 2, 2, 73, 74, 7, 116, 2, 2, 74, 75, 7, 103,
	2, 2, 75, 76, 7, 118, 2, 2, 76, 77, 7, 119, 2, 2, 77, 78, 7, 116, 2, 2,
	78, 79, 7, 112, 2, 2, 79, 14, 3, 2, 2, 2, 80, 81, 7, 120, 2, 2, 81, 82,
	7, 113, 2, 2, 82, 83, 7, 107, 2, 2, 83, 84, 7, 102, 2, 2, 84, 16, 3, 2,
	2, 2, 85, 86, 7, 111, 2, 2, 86, 87, 7, 99, 2, 2, 87, 88, 7, 107, 2, 2,
	88, 89, 7, 112, 2, 2, 89, 18, 3, 2, 2, 2, 90, 91, 7, 42, 2, 2, 91, 20,
	3, 2, 2, 2, 92, 93, 7, 43, 2, 2, 93, 22, 3, 2, 2, 2, 94, 95, 7, 125, 2,
	2, 95, 24, 3, 2, 2, 2, 96, 97, 7, 127, 2, 2, 97, 26, 3, 2, 2, 2, 98, 99,
	7, 62, 2, 2, 99, 28, 3, 2, 2, 2, 100, 101, 7, 64, 2, 2, 101, 30, 3, 2,
	2, 2, 102, 103, 7, 65, 2, 2, 103, 32, 3, 2, 2, 2, 104, 105, 7, 60, 2, 2,
	105, 34, 3, 2, 2, 2, 106, 107, 7, 61, 2, 2, 107, 36, 3, 2, 2, 2, 108, 109,
	7, 46, 2, 2, 109, 38, 3, 2, 2, 2, 110, 111, 7, 63, 2, 2, 111, 40, 3, 2,
	2, 2, 112, 113, 7, 63, 2, 2, 113, 114, 7, 63, 2, 2, 114, 42, 3, 2, 2, 2,
	115, 116, 7, 35, 2, 2, 116, 117, 7, 63, 2, 2, 117, 44, 3, 2, 2, 2, 118,
	119, 9, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 121, 8, 23, 2, 2, 121, 46,
	3, 2, 2, 2, 122, 123, 9, 3, 2, 2, 123, 48, 3, 2, 2, 2, 124, 125, 9, 4,
	2, 2, 125, 50, 3, 2, 2, 2, 126, 128, 5, 49, 25, 2, 127, 126, 3, 2, 2, 2,
	128, 129, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130,
	52, 3, 2, 2, 2, 131, 133, 5, 47, 24, 2, 132, 131, 3, 2, 2, 2, 133, 134,
	3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 54, 3, 2,
	2, 2, 5, 2, 129, 134, 3, 2, 3, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "' '", "'for'", "'if'", "'bool'", "'int'", "'return'", "'void'", "'main'",
	"'('", "')'", "'{'", "'}'", "'<'", "'>'", "'?'", "':'", "';'", "','", "'='",
	"'=='", "'!='",
}

var lexerSymbolicNames = []string{
	"", "Space", "For", "If", "Bool", "Int", "Return", "Void", "Main", "LeftParen",
	"RightParen", "LeftBrace", "RightBrace", "Less", "Greater", "Question",
	"Colon", "Semi", "Comma", "Assign", "Equal", "NotEqual", "WS", "Number",
	"Identifier",
}

var lexerRuleNames = []string{
	"Space", "For", "If", "Bool", "Int", "Return", "Void", "Main", "LeftParen",
	"RightParen", "LeftBrace", "RightBrace", "Less", "Greater", "Question",
	"Colon", "Semi", "Comma", "Assign", "Equal", "NotEqual", "WS", "NONDIGIT",
	"DIGIT", "Number", "Identifier",
}

type lightcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewlightcLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *lightcLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewlightcLexer(input antlr.CharStream) *lightcLexer {
	l := new(lightcLexer)
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
	l.GrammarFileName = "lightc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// lightcLexer tokens.
const (
	lightcLexerSpace      = 1
	lightcLexerFor        = 2
	lightcLexerIf         = 3
	lightcLexerBool       = 4
	lightcLexerInt        = 5
	lightcLexerReturn     = 6
	lightcLexerVoid       = 7
	lightcLexerMain       = 8
	lightcLexerLeftParen  = 9
	lightcLexerRightParen = 10
	lightcLexerLeftBrace  = 11
	lightcLexerRightBrace = 12
	lightcLexerLess       = 13
	lightcLexerGreater    = 14
	lightcLexerQuestion   = 15
	lightcLexerColon      = 16
	lightcLexerSemi       = 17
	lightcLexerComma      = 18
	lightcLexerAssign     = 19
	lightcLexerEqual      = 20
	lightcLexerNotEqual   = 21
	lightcLexerWS         = 22
	lightcLexerNumber     = 23
	lightcLexerIdentifier = 24
)
