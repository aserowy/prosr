// Code generated from c:\Users\serow_000\source\repos\prosr\src\parser\Prosr1.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 25, 170,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 151, 10, 21, 3, 22, 3, 22,
	6, 22, 155, 10, 22, 13, 22, 14, 22, 156, 3, 23, 6, 23, 160, 10, 23, 13,
	23, 14, 23, 161, 3, 24, 6, 24, 165, 10, 24, 13, 24, 14, 24, 166, 3, 24,
	3, 24, 2, 2, 25, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 3, 2, 6, 4, 2, 67, 92, 99,
	124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12,
	15, 15, 34, 34, 2, 174, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23,
	3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2,
	31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2,
	2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2,
	2, 2, 47, 3, 2, 2, 2, 3, 49, 3, 2, 2, 2, 5, 56, 3, 2, 2, 2, 7, 58, 3, 2,
	2, 2, 9, 60, 3, 2, 2, 2, 11, 67, 3, 2, 2, 2, 13, 69, 3, 2, 2, 2, 15, 77,
	3, 2, 2, 2, 17, 81, 3, 2, 2, 2, 19, 83, 3, 2, 2, 2, 21, 85, 3, 2, 2, 2,
	23, 92, 3, 2, 2, 2, 25, 94, 3, 2, 2, 2, 27, 96, 3, 2, 2, 2, 29, 102, 3,
	2, 2, 2, 31, 105, 3, 2, 2, 2, 33, 112, 3, 2, 2, 2, 35, 116, 3, 2, 2, 2,
	37, 124, 3, 2, 2, 2, 39, 126, 3, 2, 2, 2, 41, 150, 3, 2, 2, 2, 43, 152,
	3, 2, 2, 2, 45, 159, 3, 2, 2, 2, 47, 164, 3, 2, 2, 2, 49, 50, 7, 117, 2,
	2, 50, 51, 7, 123, 2, 2, 51, 52, 7, 112, 2, 2, 52, 53, 7, 118, 2, 2, 53,
	54, 7, 99, 2, 2, 54, 55, 7, 122, 2, 2, 55, 4, 3, 2, 2, 2, 56, 57, 7, 63,
	2, 2, 57, 6, 3, 2, 2, 2, 58, 59, 7, 36, 2, 2, 59, 8, 3, 2, 2, 2, 60, 61,
	7, 114, 2, 2, 61, 62, 7, 116, 2, 2, 62, 63, 7, 113, 2, 2, 63, 64, 7, 117,
	2, 2, 64, 65, 7, 116, 2, 2, 65, 66, 7, 51, 2, 2, 66, 10, 3, 2, 2, 2, 67,
	68, 7, 61, 2, 2, 68, 12, 3, 2, 2, 2, 69, 70, 7, 114, 2, 2, 70, 71, 7, 99,
	2, 2, 71, 72, 7, 101, 2, 2, 72, 73, 7, 109, 2, 2, 73, 74, 7, 99, 2, 2,
	74, 75, 7, 105, 2, 2, 75, 76, 7, 103, 2, 2, 76, 14, 3, 2, 2, 2, 77, 78,
	7, 106, 2, 2, 78, 79, 7, 119, 2, 2, 79, 80, 7, 100, 2, 2, 80, 16, 3, 2,
	2, 2, 81, 82, 7, 125, 2, 2, 82, 18, 3, 2, 2, 2, 83, 84, 7, 127, 2, 2, 84,
	20, 3, 2, 2, 2, 85, 86, 7, 99, 2, 2, 86, 87, 7, 101, 2, 2, 87, 88, 7, 118,
	2, 2, 88, 89, 7, 107, 2, 2, 89, 90, 7, 113, 2, 2, 90, 91, 7, 112, 2, 2,
	91, 22, 3, 2, 2, 2, 92, 93, 7, 42, 2, 2, 93, 24, 3, 2, 2, 2, 94, 95, 7,
	43, 2, 2, 95, 26, 3, 2, 2, 2, 96, 97, 7, 101, 2, 2, 97, 98, 7, 99, 2, 2,
	98, 99, 7, 110, 2, 2, 99, 100, 7, 110, 2, 2, 100, 101, 7, 117, 2, 2, 101,
	28, 3, 2, 2, 2, 102, 103, 7, 113, 2, 2, 103, 104, 7, 112, 2, 2, 104, 30,
	3, 2, 2, 2, 105, 106, 7, 101, 2, 2, 106, 107, 7, 99, 2, 2, 107, 108, 7,
	110, 2, 2, 108, 109, 7, 110, 2, 2, 109, 110, 7, 103, 2, 2, 110, 111, 7,
	116, 2, 2, 111, 32, 3, 2, 2, 2, 112, 113, 7, 99, 2, 2, 113, 114, 7, 110,
	2, 2, 114, 115, 7, 110, 2, 2, 115, 34, 3, 2, 2, 2, 116, 117, 7, 111, 2,
	2, 117, 118, 7, 103, 2, 2, 118, 119, 7, 117, 2, 2, 119, 120, 7, 117, 2,
	2, 120, 121, 7, 99, 2, 2, 121, 122, 7, 105, 2, 2, 122, 123, 7, 103, 2,
	2, 123, 36, 3, 2, 2, 2, 124, 125, 7, 48, 2, 2, 125, 38, 3, 2, 2, 2, 126,
	127, 7, 116, 2, 2, 127, 128, 7, 103, 2, 2, 128, 129, 7, 114, 2, 2, 129,
	130, 7, 103, 2, 2, 130, 131, 7, 99, 2, 2, 131, 132, 7, 118, 2, 2, 132,
	133, 7, 103, 2, 2, 133, 134, 7, 102, 2, 2, 134, 40, 3, 2, 2, 2, 135, 136,
	7, 117, 2, 2, 136, 137, 7, 118, 2, 2, 137, 138, 7, 116, 2, 2, 138, 139,
	7, 107, 2, 2, 139, 140, 7, 112, 2, 2, 140, 151, 7, 105, 2, 2, 141, 142,
	7, 100, 2, 2, 142, 143, 7, 113, 2, 2, 143, 144, 7, 113, 2, 2, 144, 151,
	7, 110, 2, 2, 145, 146, 7, 107, 2, 2, 146, 147, 7, 112, 2, 2, 147, 148,
	7, 118, 2, 2, 148, 149, 7, 53, 2, 2, 149, 151, 7, 52, 2, 2, 150, 135, 3,
	2, 2, 2, 150, 141, 3, 2, 2, 2, 150, 145, 3, 2, 2, 2, 151, 42, 3, 2, 2,
	2, 152, 154, 9, 2, 2, 2, 153, 155, 9, 3, 2, 2, 154, 153, 3, 2, 2, 2, 155,
	156, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 44, 3,
	2, 2, 2, 158, 160, 9, 4, 2, 2, 159, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2,
	2, 161, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 46, 3, 2, 2, 2, 163,
	165, 9, 5, 2, 2, 164, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 164,
	3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 169, 8, 24,
	2, 2, 169, 48, 3, 2, 2, 2, 7, 2, 150, 156, 161, 166, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'syntax'", "'='", "'\"'", "'prosr1'", "';'", "'package'", "'hub'",
	"'{'", "'}'", "'action'", "'('", "')'", "'calls'", "'on'", "'caller'",
	"'all'", "'message'", "'.'", "'repeated'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "REPEATED", "TYPE", "IDENT", "NUMBER", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "REPEATED", "TYPE", "IDENT", "NUMBER", "WHITESPACE",
}

type Prosr1Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewProsr1Lexer(input antlr.CharStream) *Prosr1Lexer {

	l := new(Prosr1Lexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Prosr1.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Prosr1Lexer tokens.
const (
	Prosr1LexerT__0       = 1
	Prosr1LexerT__1       = 2
	Prosr1LexerT__2       = 3
	Prosr1LexerT__3       = 4
	Prosr1LexerT__4       = 5
	Prosr1LexerT__5       = 6
	Prosr1LexerT__6       = 7
	Prosr1LexerT__7       = 8
	Prosr1LexerT__8       = 9
	Prosr1LexerT__9       = 10
	Prosr1LexerT__10      = 11
	Prosr1LexerT__11      = 12
	Prosr1LexerT__12      = 13
	Prosr1LexerT__13      = 14
	Prosr1LexerT__14      = 15
	Prosr1LexerT__15      = 16
	Prosr1LexerT__16      = 17
	Prosr1LexerT__17      = 18
	Prosr1LexerREPEATED   = 19
	Prosr1LexerTYPE       = 20
	Prosr1LexerIDENT      = 21
	Prosr1LexerNUMBER     = 22
	Prosr1LexerWHITESPACE = 23
)
