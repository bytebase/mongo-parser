package parser_test

import (
	"os"
	"path"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	mongoparser "github.com/bytebase/mongo-parser"
	"github.com/stretchr/testify/require"
)

type CustomErrorListener struct {
	errors int
}

func NewCustomErrorListener() *CustomErrorListener {
	return new(CustomErrorListener)
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors += 1
	antlr.ConsoleErrorListenerINSTANCE.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
}

func (l *CustomErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (l *CustomErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (l *CustomErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

func TestMongoParser(t *testing.T) {
	examples, err := os.ReadDir("examples")
	require.NoError(t, err)

	for _, file := range examples {
		filePath := path.Join("examples", file.Name())
		t.Run(filePath, func(t *testing.T) {
			t.Parallel()
			input, err := antlr.NewFileStream(filePath)
			require.NoError(t, err)

			lexer := mongoparser.NewmongoLexer(input)

			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := mongoparser.NewmongoParser(stream)

			lexerErrors := &CustomErrorListener{}
			lexer.RemoveErrorListeners()
			lexer.AddErrorListener(lexerErrors)

			parserErrors := &CustomErrorListener{}
			p.RemoveErrorListeners()
			p.AddErrorListener(parserErrors)

			p.BuildParseTrees = true

			_ = p.MongoCommands()

			require.Equal(t, 0, lexerErrors.errors)
			require.Equal(t, 0, parserErrors.errors)
		})
	}
}

func TestObject(t *testing.T) {
	s := `{
		_id: ObjectId("64c0b8c4e65c51195e0584b2"),
		name: 'danny',
		age: 13,
		groups: [ 'basketball', 'swimming' ],
		tree: { a: 'a', b: 1 }
	  }`
	input := antlr.NewInputStream(s)

	lexer := mongoparser.NewmongoLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := mongoparser.NewmongoParser(stream)

	lexerErrors := &CustomErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	parserErrors := &CustomErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	p.BuildParseTrees = true

	_ = p.ObjectLiteral()

	require.Equal(t, 0, lexerErrors.errors)
	require.Equal(t, 0, parserErrors.errors)
}
