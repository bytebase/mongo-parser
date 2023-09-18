package parser_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	mongoparser "github.com/bytebase/mongo-parser"
	"github.com/stretchr/testify/require"
)

type CustomErrorListener struct {
	errors []string
}

func NewCustomErrorListener() *CustomErrorListener {
	return new(CustomErrorListener)
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, fmt.Sprintf("%s (%d:%d)", msg, line, column))
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

			require.Empty(t, lexerErrors.errors)
			require.Empty(t, parserErrors.errors)
		})
	}
}

func TestObject(t *testing.T) {
	strings := []string{
		`{
			_id: ObjectId("64c0b8c4e65c51195e0584b2"),
			name: 'danny',
			age: 13,
			groups: [ 'basketball', 'swimming' ],
			tree: { a: 'a', b: 1 }
		}`,
		`[
			{
				_id: ObjectId("64c0b8c4e65c51195e0584b2"),
				name: 'danny',
				age: 13,
				groups: [ 'basketball', 'swimming' ],
				tree: { a: 'a', b: 1 }
			},
			{
				_id: ObjectId("64c0c59f43a63d0d400dfef6"),
				name: 'zp',
				age: 31,
				groups: [ 'hello', 'swimming' ]
			},
			{ _id: ObjectId("64c0c5b343a63d0d400dfef7"), name: 'tz', info: 66 }
		]`,
	}
	for _, s := range strings {
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

		_ = p.ArgumentList()

		require.Empty(t, lexerErrors.errors)
		require.Empty(t, parserErrors.errors)
	}
}
