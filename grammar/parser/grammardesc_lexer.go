// Code generated from GrammarDesc.g4 by ANTLR 4.9. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 15, 103,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 8, 3, 8, 7, 8, 52, 10, 8, 12, 8, 14, 8, 55, 11, 8, 3, 9, 3, 9, 3,
	9, 7, 9, 60, 10, 9, 12, 9, 14, 9, 63, 11, 9, 3, 10, 3, 10, 3, 10, 7, 10,
	68, 10, 10, 12, 10, 14, 10, 71, 11, 10, 3, 11, 3, 11, 7, 11, 75, 10, 11,
	12, 11, 14, 11, 78, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 7, 12, 84, 10,
	12, 12, 12, 14, 12, 87, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 7, 13, 93,
	10, 13, 12, 13, 14, 13, 96, 11, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14,
	3, 14, 5, 76, 85, 94, 2, 15, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 3, 2, 9, 4, 2, 67, 92,
	99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 3, 2, 67, 92, 4, 2, 50, 59, 67,
	92, 3, 2, 99, 124, 4, 2, 50, 59, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34,
	2, 108, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 3, 29, 3, 2, 2, 2, 5, 31, 3, 2, 2, 2,
	7, 33, 3, 2, 2, 2, 9, 38, 3, 2, 2, 2, 11, 44, 3, 2, 2, 2, 13, 47, 3, 2,
	2, 2, 15, 49, 3, 2, 2, 2, 17, 56, 3, 2, 2, 2, 19, 64, 3, 2, 2, 2, 21, 72,
	3, 2, 2, 2, 23, 81, 3, 2, 2, 2, 25, 90, 3, 2, 2, 2, 27, 99, 3, 2, 2, 2,
	29, 30, 7, 60, 2, 2, 30, 4, 3, 2, 2, 2, 31, 32, 7, 61, 2, 2, 32, 6, 3,
	2, 2, 2, 33, 34, 7, 117, 2, 2, 34, 35, 7, 109, 2, 2, 35, 36, 7, 107, 2,
	2, 36, 37, 7, 114, 2, 2, 37, 8, 3, 2, 2, 2, 38, 39, 7, 117, 2, 2, 39, 40,
	7, 118, 2, 2, 40, 41, 7, 99, 2, 2, 41, 42, 7, 116, 2, 2, 42, 43, 7, 118,
	2, 2, 43, 10, 3, 2, 2, 2, 44, 45, 7, 47, 2, 2, 45, 46, 7, 64, 2, 2, 46,
	12, 3, 2, 2, 2, 47, 48, 7, 63, 2, 2, 48, 14, 3, 2, 2, 2, 49, 53, 9, 2,
	2, 2, 50, 52, 9, 3, 2, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51,
	3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 16, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2,
	56, 57, 7, 38, 2, 2, 57, 61, 9, 4, 2, 2, 58, 60, 9, 5, 2, 2, 59, 58, 3,
	2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62,
	18, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 65, 7, 38, 2, 2, 65, 69, 9, 6,
	2, 2, 66, 68, 9, 7, 2, 2, 67, 66, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67,
	3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 20, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2,
	72, 76, 7, 36, 2, 2, 73, 75, 11, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 78, 3,
	2, 2, 2, 76, 77, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77, 79, 3, 2, 2, 2, 78,
	76, 3, 2, 2, 2, 79, 80, 7, 36, 2, 2, 80, 22, 3, 2, 2, 2, 81, 85, 7, 35,
	2, 2, 82, 84, 11, 2, 2, 2, 83, 82, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85,
	86, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 88, 3, 2, 2, 2, 87, 85, 3, 2, 2,
	2, 88, 89, 7, 35, 2, 2, 89, 24, 3, 2, 2, 2, 90, 94, 7, 125, 2, 2, 91, 93,
	11, 2, 2, 2, 92, 91, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2,
	94, 92, 3, 2, 2, 2, 95, 97, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 98, 7,
	127, 2, 2, 98, 26, 3, 2, 2, 2, 99, 100, 9, 8, 2, 2, 100, 101, 3, 2, 2,
	2, 101, 102, 8, 14, 2, 2, 102, 28, 3, 2, 2, 2, 9, 2, 53, 61, 69, 76, 85,
	94, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "':'", "';'", "'skip'", "'start'", "'->'", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "ID", "TERM_ID", "NON_TERM_ID", "REGEXP", "ATTR",
	"CODE", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "ID", "TERM_ID", "NON_TERM_ID",
	"REGEXP", "ATTR", "CODE", "WS",
}

type GrammarDescLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewGrammarDescLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *GrammarDescLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGrammarDescLexer(input antlr.CharStream) *GrammarDescLexer {
	l := new(GrammarDescLexer)
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
	l.GrammarFileName = "GrammarDesc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GrammarDescLexer tokens.
const (
	GrammarDescLexerT__0        = 1
	GrammarDescLexerT__1        = 2
	GrammarDescLexerT__2        = 3
	GrammarDescLexerT__3        = 4
	GrammarDescLexerT__4        = 5
	GrammarDescLexerT__5        = 6
	GrammarDescLexerID          = 7
	GrammarDescLexerTERM_ID     = 8
	GrammarDescLexerNON_TERM_ID = 9
	GrammarDescLexerREGEXP      = 10
	GrammarDescLexerATTR        = 11
	GrammarDescLexerCODE        = 12
	GrammarDescLexerWS          = 13
)
