// Code generated from mongo.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type mongoLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MongoLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mongolexerLexerInit() {
	staticData := &MongoLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'{'", "','", "'}'", "'['", "']'", "':'", "", "",
		"", "'null'", "", "", "", "", "';'", "'.'", "'db'", "'\\n'", "'\\r\\n'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "SingleLineComment", "MultiLineComment",
		"StringLiteral", "NullLiteral", "BooleanLiteral", "NumericLiteral",
		"DecimalLiteral", "LineTerminator", "SEMICOLON", "DOT", "DB", "LF",
		"CRLF", "STRING_LITERAL", "DOUBLE_QUOTED_STRING_LITERAL", "SINGLE_QUOTED_STRING_LITERAL",
		"WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "SingleLineComment",
		"MultiLineComment", "StringLiteral", "NullLiteral", "BooleanLiteral",
		"NumericLiteral", "DecimalLiteral", "LineTerminator", "SEMICOLON", "DOT",
		"DB", "LF", "CRLF", "STRING_LITERAL", "DOUBLE_QUOTED_STRING_LITERAL",
		"SINGLE_QUOTED_STRING_LITERAL", "STRING_ESCAPE", "DecimalIntegerLiteral",
		"ExponentPart", "DecimalDigit", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 221, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 5, 8, 80, 8, 8, 10, 8, 12, 8, 83, 9, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 5, 9, 91, 8, 9, 10, 9, 12, 9, 94, 9, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 3, 10, 103, 8, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 3, 12, 119, 8, 12, 1, 13, 3, 13, 122, 8, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 5, 14, 129, 8, 14, 10, 14, 12, 14, 132, 9, 14, 1, 14, 3, 14,
		135, 8, 14, 1, 14, 1, 14, 4, 14, 139, 8, 14, 11, 14, 12, 14, 140, 1, 14,
		3, 14, 144, 8, 14, 1, 14, 1, 14, 3, 14, 148, 8, 14, 3, 14, 150, 8, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 4, 21, 170, 8, 21,
		11, 21, 12, 21, 171, 1, 22, 1, 22, 1, 22, 5, 22, 177, 8, 22, 10, 22, 12,
		22, 180, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 5, 23, 187, 8, 23, 10,
		23, 12, 23, 190, 9, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 5, 25, 200, 8, 25, 10, 25, 12, 25, 203, 9, 25, 3, 25, 205, 8, 25,
		1, 26, 1, 26, 3, 26, 209, 8, 26, 1, 26, 4, 26, 212, 8, 26, 11, 26, 12,
		26, 213, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 92, 0, 29, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23,
		12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41,
		21, 43, 22, 45, 23, 47, 24, 49, 0, 51, 0, 53, 0, 55, 0, 57, 25, 1, 0, 10,
		3, 0, 10, 10, 13, 13, 8232, 8233, 9, 0, 9, 10, 32, 32, 34, 34, 40, 41,
		44, 46, 58, 59, 92, 92, 123, 123, 125, 125, 2, 0, 34, 34, 92, 92, 2, 0,
		39, 39, 92, 92, 3, 0, 34, 34, 39, 39, 92, 92, 1, 0, 49, 57, 2, 0, 69, 69,
		101, 101, 2, 0, 43, 43, 45, 45, 1, 0, 48, 57, 2, 0, 9, 9, 32, 32, 238,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 1, 59, 1, 0, 0, 0, 3, 61, 1, 0, 0, 0,
		5, 63, 1, 0, 0, 0, 7, 65, 1, 0, 0, 0, 9, 67, 1, 0, 0, 0, 11, 69, 1, 0,
		0, 0, 13, 71, 1, 0, 0, 0, 15, 73, 1, 0, 0, 0, 17, 75, 1, 0, 0, 0, 19, 86,
		1, 0, 0, 0, 21, 102, 1, 0, 0, 0, 23, 104, 1, 0, 0, 0, 25, 118, 1, 0, 0,
		0, 27, 121, 1, 0, 0, 0, 29, 149, 1, 0, 0, 0, 31, 151, 1, 0, 0, 0, 33, 155,
		1, 0, 0, 0, 35, 157, 1, 0, 0, 0, 37, 159, 1, 0, 0, 0, 39, 162, 1, 0, 0,
		0, 41, 164, 1, 0, 0, 0, 43, 169, 1, 0, 0, 0, 45, 173, 1, 0, 0, 0, 47, 183,
		1, 0, 0, 0, 49, 193, 1, 0, 0, 0, 51, 204, 1, 0, 0, 0, 53, 206, 1, 0, 0,
		0, 55, 215, 1, 0, 0, 0, 57, 217, 1, 0, 0, 0, 59, 60, 5, 40, 0, 0, 60, 2,
		1, 0, 0, 0, 61, 62, 5, 41, 0, 0, 62, 4, 1, 0, 0, 0, 63, 64, 5, 123, 0,
		0, 64, 6, 1, 0, 0, 0, 65, 66, 5, 44, 0, 0, 66, 8, 1, 0, 0, 0, 67, 68, 5,
		125, 0, 0, 68, 10, 1, 0, 0, 0, 69, 70, 5, 91, 0, 0, 70, 12, 1, 0, 0, 0,
		71, 72, 5, 93, 0, 0, 72, 14, 1, 0, 0, 0, 73, 74, 5, 58, 0, 0, 74, 16, 1,
		0, 0, 0, 75, 76, 5, 47, 0, 0, 76, 77, 5, 47, 0, 0, 77, 81, 1, 0, 0, 0,
		78, 80, 8, 0, 0, 0, 79, 78, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0, 81, 79, 1,
		0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 84, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84,
		85, 6, 8, 0, 0, 85, 18, 1, 0, 0, 0, 86, 87, 5, 47, 0, 0, 87, 88, 5, 42,
		0, 0, 88, 92, 1, 0, 0, 0, 89, 91, 9, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 94,
		1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0,
		94, 92, 1, 0, 0, 0, 95, 96, 5, 42, 0, 0, 96, 97, 5, 47, 0, 0, 97, 98, 1,
		0, 0, 0, 98, 99, 6, 9, 0, 0, 99, 20, 1, 0, 0, 0, 100, 103, 3, 47, 23, 0,
		101, 103, 3, 45, 22, 0, 102, 100, 1, 0, 0, 0, 102, 101, 1, 0, 0, 0, 103,
		22, 1, 0, 0, 0, 104, 105, 5, 110, 0, 0, 105, 106, 5, 117, 0, 0, 106, 107,
		5, 108, 0, 0, 107, 108, 5, 108, 0, 0, 108, 24, 1, 0, 0, 0, 109, 110, 5,
		116, 0, 0, 110, 111, 5, 114, 0, 0, 111, 112, 5, 117, 0, 0, 112, 119, 5,
		101, 0, 0, 113, 114, 5, 102, 0, 0, 114, 115, 5, 97, 0, 0, 115, 116, 5,
		108, 0, 0, 116, 117, 5, 115, 0, 0, 117, 119, 5, 101, 0, 0, 118, 109, 1,
		0, 0, 0, 118, 113, 1, 0, 0, 0, 119, 26, 1, 0, 0, 0, 120, 122, 5, 45, 0,
		0, 121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123,
		124, 3, 29, 14, 0, 124, 28, 1, 0, 0, 0, 125, 126, 3, 51, 25, 0, 126, 130,
		5, 46, 0, 0, 127, 129, 3, 55, 27, 0, 128, 127, 1, 0, 0, 0, 129, 132, 1,
		0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 134, 1, 0, 0,
		0, 132, 130, 1, 0, 0, 0, 133, 135, 3, 53, 26, 0, 134, 133, 1, 0, 0, 0,
		134, 135, 1, 0, 0, 0, 135, 150, 1, 0, 0, 0, 136, 138, 5, 46, 0, 0, 137,
		139, 3, 55, 27, 0, 138, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 138,
		1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 143, 1, 0, 0, 0, 142, 144, 3, 53,
		26, 0, 143, 142, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 150, 1, 0, 0, 0,
		145, 147, 3, 51, 25, 0, 146, 148, 3, 53, 26, 0, 147, 146, 1, 0, 0, 0, 147,
		148, 1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 125, 1, 0, 0, 0, 149, 136,
		1, 0, 0, 0, 149, 145, 1, 0, 0, 0, 150, 30, 1, 0, 0, 0, 151, 152, 7, 0,
		0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 6, 15, 0, 0, 154, 32, 1, 0, 0, 0,
		155, 156, 5, 59, 0, 0, 156, 34, 1, 0, 0, 0, 157, 158, 5, 46, 0, 0, 158,
		36, 1, 0, 0, 0, 159, 160, 5, 100, 0, 0, 160, 161, 5, 98, 0, 0, 161, 38,
		1, 0, 0, 0, 162, 163, 5, 10, 0, 0, 163, 40, 1, 0, 0, 0, 164, 165, 5, 13,
		0, 0, 165, 166, 5, 10, 0, 0, 166, 42, 1, 0, 0, 0, 167, 170, 8, 1, 0, 0,
		168, 170, 3, 49, 24, 0, 169, 167, 1, 0, 0, 0, 169, 168, 1, 0, 0, 0, 170,
		171, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 44, 1,
		0, 0, 0, 173, 178, 5, 34, 0, 0, 174, 177, 8, 2, 0, 0, 175, 177, 3, 49,
		24, 0, 176, 174, 1, 0, 0, 0, 176, 175, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0,
		178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 181, 1, 0, 0, 0, 180,
		178, 1, 0, 0, 0, 181, 182, 5, 34, 0, 0, 182, 46, 1, 0, 0, 0, 183, 188,
		5, 39, 0, 0, 184, 187, 8, 3, 0, 0, 185, 187, 3, 49, 24, 0, 186, 184, 1,
		0, 0, 0, 186, 185, 1, 0, 0, 0, 187, 190, 1, 0, 0, 0, 188, 186, 1, 0, 0,
		0, 188, 189, 1, 0, 0, 0, 189, 191, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 191,
		192, 5, 39, 0, 0, 192, 48, 1, 0, 0, 0, 193, 194, 5, 92, 0, 0, 194, 195,
		7, 4, 0, 0, 195, 50, 1, 0, 0, 0, 196, 205, 5, 48, 0, 0, 197, 201, 7, 5,
		0, 0, 198, 200, 3, 55, 27, 0, 199, 198, 1, 0, 0, 0, 200, 203, 1, 0, 0,
		0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203,
		201, 1, 0, 0, 0, 204, 196, 1, 0, 0, 0, 204, 197, 1, 0, 0, 0, 205, 52, 1,
		0, 0, 0, 206, 208, 7, 6, 0, 0, 207, 209, 7, 7, 0, 0, 208, 207, 1, 0, 0,
		0, 208, 209, 1, 0, 0, 0, 209, 211, 1, 0, 0, 0, 210, 212, 3, 55, 27, 0,
		211, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213,
		214, 1, 0, 0, 0, 214, 54, 1, 0, 0, 0, 215, 216, 7, 8, 0, 0, 216, 56, 1,
		0, 0, 0, 217, 218, 7, 9, 0, 0, 218, 219, 1, 0, 0, 0, 219, 220, 6, 28, 1,
		0, 220, 58, 1, 0, 0, 0, 22, 0, 81, 92, 102, 118, 121, 130, 134, 140, 143,
		147, 149, 169, 171, 176, 178, 186, 188, 201, 204, 208, 213, 2, 0, 1, 0,
		6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// mongoLexerInit initializes any static state used to implement mongoLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewmongoLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MongoLexerInit() {
	staticData := &MongoLexerLexerStaticData
	staticData.once.Do(mongolexerLexerInit)
}

// NewmongoLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewmongoLexer(input antlr.CharStream) *mongoLexer {
	MongoLexerInit()
	l := new(mongoLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MongoLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "mongo.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// mongoLexer tokens.
const (
	mongoLexerT__0                         = 1
	mongoLexerT__1                         = 2
	mongoLexerT__2                         = 3
	mongoLexerT__3                         = 4
	mongoLexerT__4                         = 5
	mongoLexerT__5                         = 6
	mongoLexerT__6                         = 7
	mongoLexerT__7                         = 8
	mongoLexerSingleLineComment            = 9
	mongoLexerMultiLineComment             = 10
	mongoLexerStringLiteral                = 11
	mongoLexerNullLiteral                  = 12
	mongoLexerBooleanLiteral               = 13
	mongoLexerNumericLiteral               = 14
	mongoLexerDecimalLiteral               = 15
	mongoLexerLineTerminator               = 16
	mongoLexerSEMICOLON                    = 17
	mongoLexerDOT                          = 18
	mongoLexerDB                           = 19
	mongoLexerLF                           = 20
	mongoLexerCRLF                         = 21
	mongoLexerSTRING_LITERAL               = 22
	mongoLexerDOUBLE_QUOTED_STRING_LITERAL = 23
	mongoLexerSINGLE_QUOTED_STRING_LITERAL = 24
	mongoLexerWHITESPACE                   = 25
)

func isExternalIdentifierText(text string) bool {
	return text == "db"
}
