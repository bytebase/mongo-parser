// Code generated from mongo.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // mongo
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type mongoParser struct {
	*antlr.BaseParser
}

var MongoParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mongoParserInit() {
	staticData := &MongoParserStaticData
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
		"mongoCommands", "commands", "command", "emptyCommand", "collection",
		"functionCall", "arguments", "argumentList", "objectLiteral", "arrayLiteral",
		"elementList", "propertyNameAndValueList", "propertyAssignment", "propertyValue",
		"literal", "propertyName", "comment",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 124, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 4, 1, 41, 8, 1, 11, 1,
		12, 1, 42, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 52, 8, 2, 1,
		2, 3, 2, 55, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		3, 6, 66, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 73, 8, 7, 1, 8, 1,
		8, 3, 8, 77, 8, 8, 1, 8, 3, 8, 80, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9,
		86, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 93, 8, 10, 10, 10, 12,
		10, 96, 9, 10, 1, 11, 1, 11, 1, 11, 5, 11, 101, 8, 11, 10, 11, 12, 11,
		104, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 3,
		13, 114, 8, 13, 1, 14, 1, 14, 3, 14, 118, 8, 14, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 0, 0, 17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 0, 2, 1, 0, 11, 13, 1, 0, 9, 10, 123, 0, 34, 1, 0, 0, 0, 2,
		40, 1, 0, 0, 0, 4, 44, 1, 0, 0, 0, 6, 56, 1, 0, 0, 0, 8, 58, 1, 0, 0, 0,
		10, 60, 1, 0, 0, 0, 12, 63, 1, 0, 0, 0, 14, 72, 1, 0, 0, 0, 16, 74, 1,
		0, 0, 0, 18, 83, 1, 0, 0, 0, 20, 89, 1, 0, 0, 0, 22, 97, 1, 0, 0, 0, 24,
		105, 1, 0, 0, 0, 26, 113, 1, 0, 0, 0, 28, 117, 1, 0, 0, 0, 30, 119, 1,
		0, 0, 0, 32, 121, 1, 0, 0, 0, 34, 35, 3, 2, 1, 0, 35, 36, 5, 0, 0, 1, 36,
		1, 1, 0, 0, 0, 37, 41, 3, 4, 2, 0, 38, 41, 3, 6, 3, 0, 39, 41, 3, 32, 16,
		0, 40, 37, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 39, 1, 0, 0, 0, 41, 42,
		1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 3, 1, 0, 0, 0,
		44, 45, 5, 19, 0, 0, 45, 51, 5, 18, 0, 0, 46, 52, 3, 10, 5, 0, 47, 48,
		3, 8, 4, 0, 48, 49, 5, 18, 0, 0, 49, 50, 3, 10, 5, 0, 50, 52, 1, 0, 0,
		0, 51, 46, 1, 0, 0, 0, 51, 47, 1, 0, 0, 0, 52, 54, 1, 0, 0, 0, 53, 55,
		5, 17, 0, 0, 54, 53, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 5, 1, 0, 0, 0,
		56, 57, 5, 17, 0, 0, 57, 7, 1, 0, 0, 0, 58, 59, 5, 22, 0, 0, 59, 9, 1,
		0, 0, 0, 60, 61, 5, 22, 0, 0, 61, 62, 3, 12, 6, 0, 62, 11, 1, 0, 0, 0,
		63, 65, 5, 1, 0, 0, 64, 66, 3, 14, 7, 0, 65, 64, 1, 0, 0, 0, 65, 66, 1,
		0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 5, 2, 0, 0, 68, 13, 1, 0, 0, 0, 69,
		73, 3, 28, 14, 0, 70, 73, 3, 16, 8, 0, 71, 73, 3, 18, 9, 0, 72, 69, 1,
		0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 71, 1, 0, 0, 0, 73, 15, 1, 0, 0, 0, 74,
		76, 5, 3, 0, 0, 75, 77, 3, 22, 11, 0, 76, 75, 1, 0, 0, 0, 76, 77, 1, 0,
		0, 0, 77, 79, 1, 0, 0, 0, 78, 80, 5, 4, 0, 0, 79, 78, 1, 0, 0, 0, 79, 80,
		1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 82, 5, 5, 0, 0, 82, 17, 1, 0, 0, 0,
		83, 85, 5, 6, 0, 0, 84, 86, 3, 20, 10, 0, 85, 84, 1, 0, 0, 0, 85, 86, 1,
		0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 5, 7, 0, 0, 88, 19, 1, 0, 0, 0, 89,
		94, 3, 26, 13, 0, 90, 91, 5, 4, 0, 0, 91, 93, 3, 26, 13, 0, 92, 90, 1,
		0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95,
		21, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 102, 3, 24, 12, 0, 98, 99, 5, 4,
		0, 0, 99, 101, 3, 24, 12, 0, 100, 98, 1, 0, 0, 0, 101, 104, 1, 0, 0, 0,
		102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 23, 1, 0, 0, 0, 104, 102,
		1, 0, 0, 0, 105, 106, 3, 30, 15, 0, 106, 107, 5, 8, 0, 0, 107, 108, 3,
		26, 13, 0, 108, 25, 1, 0, 0, 0, 109, 114, 3, 28, 14, 0, 110, 114, 3, 16,
		8, 0, 111, 114, 3, 18, 9, 0, 112, 114, 3, 10, 5, 0, 113, 109, 1, 0, 0,
		0, 113, 110, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 112, 1, 0, 0, 0, 114,
		27, 1, 0, 0, 0, 115, 118, 7, 0, 0, 0, 116, 118, 5, 14, 0, 0, 117, 115,
		1, 0, 0, 0, 117, 116, 1, 0, 0, 0, 118, 29, 1, 0, 0, 0, 119, 120, 5, 22,
		0, 0, 120, 31, 1, 0, 0, 0, 121, 122, 7, 1, 0, 0, 122, 33, 1, 0, 0, 0, 13,
		40, 42, 51, 54, 65, 72, 76, 79, 85, 94, 102, 113, 117,
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

// mongoParserInit initializes any static state used to implement mongoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewmongoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MongoParserInit() {
	staticData := &MongoParserStaticData
	staticData.once.Do(mongoParserInit)
}

// NewmongoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewmongoParser(input antlr.TokenStream) *mongoParser {
	MongoParserInit()
	this := new(mongoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MongoParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "mongo.g4"

	return this
}

// mongoParser tokens.
const (
	mongoParserEOF                          = antlr.TokenEOF
	mongoParserT__0                         = 1
	mongoParserT__1                         = 2
	mongoParserT__2                         = 3
	mongoParserT__3                         = 4
	mongoParserT__4                         = 5
	mongoParserT__5                         = 6
	mongoParserT__6                         = 7
	mongoParserT__7                         = 8
	mongoParserSingleLineComment            = 9
	mongoParserMultiLineComment             = 10
	mongoParserStringLiteral                = 11
	mongoParserNullLiteral                  = 12
	mongoParserBooleanLiteral               = 13
	mongoParserNumericLiteral               = 14
	mongoParserDecimalLiteral               = 15
	mongoParserLineTerminator               = 16
	mongoParserSEMICOLON                    = 17
	mongoParserDOT                          = 18
	mongoParserDB                           = 19
	mongoParserLF                           = 20
	mongoParserCRLF                         = 21
	mongoParserSTRING_LITERAL               = 22
	mongoParserDOUBLE_QUOTED_STRING_LITERAL = 23
	mongoParserSINGLE_QUOTED_STRING_LITERAL = 24
	mongoParserWHITESPACE                   = 25
)

// mongoParser rules.
const (
	mongoParserRULE_mongoCommands            = 0
	mongoParserRULE_commands                 = 1
	mongoParserRULE_command                  = 2
	mongoParserRULE_emptyCommand             = 3
	mongoParserRULE_collection               = 4
	mongoParserRULE_functionCall             = 5
	mongoParserRULE_arguments                = 6
	mongoParserRULE_argumentList             = 7
	mongoParserRULE_objectLiteral            = 8
	mongoParserRULE_arrayLiteral             = 9
	mongoParserRULE_elementList              = 10
	mongoParserRULE_propertyNameAndValueList = 11
	mongoParserRULE_propertyAssignment       = 12
	mongoParserRULE_propertyValue            = 13
	mongoParserRULE_literal                  = 14
	mongoParserRULE_propertyName             = 15
	mongoParserRULE_comment                  = 16
)

// IMongoCommandsContext is an interface to support dynamic dispatch.
type IMongoCommandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Commands() ICommandsContext
	EOF() antlr.TerminalNode

	// IsMongoCommandsContext differentiates from other interfaces.
	IsMongoCommandsContext()
}

type MongoCommandsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMongoCommandsContext() *MongoCommandsContext {
	var p = new(MongoCommandsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_mongoCommands
	return p
}

func InitEmptyMongoCommandsContext(p *MongoCommandsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_mongoCommands
}

func (*MongoCommandsContext) IsMongoCommandsContext() {}

func NewMongoCommandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MongoCommandsContext {
	var p = new(MongoCommandsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_mongoCommands

	return p
}

func (s *MongoCommandsContext) GetParser() antlr.Parser { return s.parser }

func (s *MongoCommandsContext) Commands() ICommandsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandsContext)
}

func (s *MongoCommandsContext) EOF() antlr.TerminalNode {
	return s.GetToken(mongoParserEOF, 0)
}

func (s *MongoCommandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MongoCommandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MongoCommandsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterMongoCommands(s)
	}
}

func (s *MongoCommandsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitMongoCommands(s)
	}
}

func (s *MongoCommandsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitMongoCommands(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) MongoCommands() (localctx IMongoCommandsContext) {
	localctx = NewMongoCommandsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, mongoParserRULE_mongoCommands)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Commands()
	}
	{
		p.SetState(35)
		p.Match(mongoParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandsContext is an interface to support dynamic dispatch.
type ICommandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCommand() []ICommandContext
	Command(i int) ICommandContext
	AllEmptyCommand() []IEmptyCommandContext
	EmptyCommand(i int) IEmptyCommandContext
	AllComment() []ICommentContext
	Comment(i int) ICommentContext

	// IsCommandsContext differentiates from other interfaces.
	IsCommandsContext()
}

type CommandsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandsContext() *CommandsContext {
	var p = new(CommandsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_commands
	return p
}

func InitEmptyCommandsContext(p *CommandsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_commands
}

func (*CommandsContext) IsCommandsContext() {}

func NewCommandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandsContext {
	var p = new(CommandsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_commands

	return p
}

func (s *CommandsContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandsContext) AllCommand() []ICommandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommandContext); ok {
			len++
		}
	}

	tst := make([]ICommandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommandContext); ok {
			tst[i] = t.(ICommandContext)
			i++
		}
	}

	return tst
}

func (s *CommandsContext) Command(i int) ICommandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *CommandsContext) AllEmptyCommand() []IEmptyCommandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyCommandContext); ok {
			len++
		}
	}

	tst := make([]IEmptyCommandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyCommandContext); ok {
			tst[i] = t.(IEmptyCommandContext)
			i++
		}
	}

	return tst
}

func (s *CommandsContext) EmptyCommand(i int) IEmptyCommandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyCommandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyCommandContext)
}

func (s *CommandsContext) AllComment() []ICommentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommentContext); ok {
			len++
		}
	}

	tst := make([]ICommentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommentContext); ok {
			tst[i] = t.(ICommentContext)
			i++
		}
	}

	return tst
}

func (s *CommandsContext) Comment(i int) ICommentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *CommandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterCommands(s)
	}
}

func (s *CommandsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitCommands(s)
	}
}

func (s *CommandsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitCommands(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) Commands() (localctx ICommandsContext) {
	localctx = NewCommandsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, mongoParserRULE_commands)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&656896) != 0) {
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case mongoParserDB:
			{
				p.SetState(37)
				p.Command()
			}

		case mongoParserSEMICOLON:
			{
				p.SetState(38)
				p.EmptyCommand()
			}

		case mongoParserSingleLineComment, mongoParserMultiLineComment:
			{
				p.SetState(39)
				p.Comment()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DB() antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	FunctionCall() IFunctionCallContext
	SEMICOLON() antlr.TerminalNode
	Collection() ICollectionContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) DB() antlr.TerminalNode {
	return s.GetToken(mongoParserDB, 0)
}

func (s *CommandContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(mongoParserDOT)
}

func (s *CommandContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(mongoParserDOT, i)
}

func (s *CommandContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *CommandContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(mongoParserSEMICOLON, 0)
}

func (s *CommandContext) Collection() ICollectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (s *CommandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitCommand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, mongoParserRULE_command)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(mongoParserDB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(mongoParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(46)
			p.FunctionCall()
		}

	case 2:
		{
			p.SetState(47)
			p.Collection()
		}
		{
			p.SetState(48)
			p.Match(mongoParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.FunctionCall()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(53)
			p.Match(mongoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEmptyCommandContext is an interface to support dynamic dispatch.
type IEmptyCommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMICOLON() antlr.TerminalNode

	// IsEmptyCommandContext differentiates from other interfaces.
	IsEmptyCommandContext()
}

type EmptyCommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyCommandContext() *EmptyCommandContext {
	var p = new(EmptyCommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_emptyCommand
	return p
}

func InitEmptyEmptyCommandContext(p *EmptyCommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_emptyCommand
}

func (*EmptyCommandContext) IsEmptyCommandContext() {}

func NewEmptyCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyCommandContext {
	var p = new(EmptyCommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_emptyCommand

	return p
}

func (s *EmptyCommandContext) GetParser() antlr.Parser { return s.parser }

func (s *EmptyCommandContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(mongoParserSEMICOLON, 0)
}

func (s *EmptyCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyCommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterEmptyCommand(s)
	}
}

func (s *EmptyCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitEmptyCommand(s)
	}
}

func (s *EmptyCommandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitEmptyCommand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) EmptyCommand() (localctx IEmptyCommandContext) {
	localctx = NewEmptyCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, mongoParserRULE_emptyCommand)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(mongoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICollectionContext is an interface to support dynamic dispatch.
type ICollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode

	// IsCollectionContext differentiates from other interfaces.
	IsCollectionContext()
}

type CollectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionContext() *CollectionContext {
	var p = new(CollectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_collection
	return p
}

func InitEmptyCollectionContext(p *CollectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_collection
}

func (*CollectionContext) IsCollectionContext() {}

func NewCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionContext {
	var p = new(CollectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_collection

	return p
}

func (s *CollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(mongoParserSTRING_LITERAL, 0)
}

func (s *CollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterCollection(s)
	}
}

func (s *CollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitCollection(s)
	}
}

func (s *CollectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitCollection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) Collection() (localctx ICollectionContext) {
	localctx = NewCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, mongoParserRULE_collection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(mongoParserSTRING_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFUNCTION_NAME returns the FUNCTION_NAME token.
	GetFUNCTION_NAME() antlr.Token

	// SetFUNCTION_NAME sets the FUNCTION_NAME token.
	SetFUNCTION_NAME(antlr.Token)

	// Getter signatures
	Arguments() IArgumentsContext
	STRING_LITERAL() antlr.TerminalNode

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	FUNCTION_NAME antlr.Token
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) GetFUNCTION_NAME() antlr.Token { return s.FUNCTION_NAME }

func (s *FunctionCallContext) SetFUNCTION_NAME(v antlr.Token) { s.FUNCTION_NAME = v }

func (s *FunctionCallContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FunctionCallContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(mongoParserSTRING_LITERAL, 0)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, mongoParserRULE_functionCall)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)

		var _m = p.Match(mongoParserSTRING_LITERAL)

		localctx.(*FunctionCallContext).FUNCTION_NAME = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Arguments()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOPEN_PARENTHESIS returns the OPEN_PARENTHESIS token.
	GetOPEN_PARENTHESIS() antlr.Token

	// GetCLOSED_PARENTHESIS returns the CLOSED_PARENTHESIS token.
	GetCLOSED_PARENTHESIS() antlr.Token

	// SetOPEN_PARENTHESIS sets the OPEN_PARENTHESIS token.
	SetOPEN_PARENTHESIS(antlr.Token)

	// SetCLOSED_PARENTHESIS sets the CLOSED_PARENTHESIS token.
	SetCLOSED_PARENTHESIS(antlr.Token)

	// Getter signatures
	ArgumentList() IArgumentListContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser             antlr.Parser
	OPEN_PARENTHESIS   antlr.Token
	CLOSED_PARENTHESIS antlr.Token
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) GetOPEN_PARENTHESIS() antlr.Token { return s.OPEN_PARENTHESIS }

func (s *ArgumentsContext) GetCLOSED_PARENTHESIS() antlr.Token { return s.CLOSED_PARENTHESIS }

func (s *ArgumentsContext) SetOPEN_PARENTHESIS(v antlr.Token) { s.OPEN_PARENTHESIS = v }

func (s *ArgumentsContext) SetCLOSED_PARENTHESIS(v antlr.Token) { s.CLOSED_PARENTHESIS = v }

func (s *ArgumentsContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, mongoParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)

		var _m = p.Match(mongoParserT__0)

		localctx.(*ArgumentsContext).OPEN_PARENTHESIS = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30792) != 0 {
		{
			p.SetState(64)
			p.ArgumentList()
		}

	}
	{
		p.SetState(67)

		var _m = p.Match(mongoParserT__1)

		localctx.(*ArgumentsContext).CLOSED_PARENTHESIS = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	ObjectLiteral() IObjectLiteralContext
	ArrayLiteral() IArrayLiteralContext

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ArgumentListContext) ObjectLiteral() IObjectLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectLiteralContext)
}

func (s *ArgumentListContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, mongoParserRULE_argumentList)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case mongoParserStringLiteral, mongoParserNullLiteral, mongoParserBooleanLiteral, mongoParserNumericLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Literal()
		}

	case mongoParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.ObjectLiteral()
		}

	case mongoParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(71)
			p.ArrayLiteral()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectLiteralContext is an interface to support dynamic dispatch.
type IObjectLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PropertyNameAndValueList() IPropertyNameAndValueListContext

	// IsObjectLiteralContext differentiates from other interfaces.
	IsObjectLiteralContext()
}

type ObjectLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectLiteralContext() *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_objectLiteral
	return p
}

func InitEmptyObjectLiteralContext(p *ObjectLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_objectLiteral
}

func (*ObjectLiteralContext) IsObjectLiteralContext() {}

func NewObjectLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_objectLiteral

	return p
}

func (s *ObjectLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectLiteralContext) PropertyNameAndValueList() IPropertyNameAndValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyNameAndValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyNameAndValueListContext)
}

func (s *ObjectLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitObjectLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) ObjectLiteral() (localctx IObjectLiteralContext) {
	localctx = NewObjectLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, mongoParserRULE_objectLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(mongoParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == mongoParserSTRING_LITERAL {
		{
			p.SetState(75)
			p.PropertyNameAndValueList()
		}

	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == mongoParserT__3 {
		{
			p.SetState(78)
			p.Match(mongoParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(81)
		p.Match(mongoParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ElementList() IElementListContext

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) ElementList() IElementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementListContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, mongoParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(mongoParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4225096) != 0 {
		{
			p.SetState(84)
			p.ElementList()
		}

	}
	{
		p.SetState(87)
		p.Match(mongoParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElementListContext is an interface to support dynamic dispatch.
type IElementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPropertyValue() []IPropertyValueContext
	PropertyValue(i int) IPropertyValueContext

	// IsElementListContext differentiates from other interfaces.
	IsElementListContext()
}

type ElementListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementListContext() *ElementListContext {
	var p = new(ElementListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_elementList
	return p
}

func InitEmptyElementListContext(p *ElementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_elementList
}

func (*ElementListContext) IsElementListContext() {}

func NewElementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementListContext {
	var p = new(ElementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_elementList

	return p
}

func (s *ElementListContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementListContext) AllPropertyValue() []IPropertyValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropertyValueContext); ok {
			len++
		}
	}

	tst := make([]IPropertyValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropertyValueContext); ok {
			tst[i] = t.(IPropertyValueContext)
			i++
		}
	}

	return tst
}

func (s *ElementListContext) PropertyValue(i int) IPropertyValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyValueContext)
}

func (s *ElementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterElementList(s)
	}
}

func (s *ElementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitElementList(s)
	}
}

func (s *ElementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitElementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) ElementList() (localctx IElementListContext) {
	localctx = NewElementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, mongoParserRULE_elementList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.PropertyValue()
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == mongoParserT__3 {
		{
			p.SetState(90)
			p.Match(mongoParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.PropertyValue()
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyNameAndValueListContext is an interface to support dynamic dispatch.
type IPropertyNameAndValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPropertyAssignment() []IPropertyAssignmentContext
	PropertyAssignment(i int) IPropertyAssignmentContext

	// IsPropertyNameAndValueListContext differentiates from other interfaces.
	IsPropertyNameAndValueListContext()
}

type PropertyNameAndValueListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyNameAndValueListContext() *PropertyNameAndValueListContext {
	var p = new(PropertyNameAndValueListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_propertyNameAndValueList
	return p
}

func InitEmptyPropertyNameAndValueListContext(p *PropertyNameAndValueListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_propertyNameAndValueList
}

func (*PropertyNameAndValueListContext) IsPropertyNameAndValueListContext() {}

func NewPropertyNameAndValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameAndValueListContext {
	var p = new(PropertyNameAndValueListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_propertyNameAndValueList

	return p
}

func (s *PropertyNameAndValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameAndValueListContext) AllPropertyAssignment() []IPropertyAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropertyAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IPropertyAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropertyAssignmentContext); ok {
			tst[i] = t.(IPropertyAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *PropertyNameAndValueListContext) PropertyAssignment(i int) IPropertyAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyAssignmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyAssignmentContext)
}

func (s *PropertyNameAndValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameAndValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameAndValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterPropertyNameAndValueList(s)
	}
}

func (s *PropertyNameAndValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitPropertyNameAndValueList(s)
	}
}

func (s *PropertyNameAndValueListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitPropertyNameAndValueList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) PropertyNameAndValueList() (localctx IPropertyNameAndValueListContext) {
	localctx = NewPropertyNameAndValueListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, mongoParserRULE_propertyNameAndValueList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.PropertyAssignment()
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(98)
				p.Match(mongoParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(99)
				p.PropertyAssignment()
			}

		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyAssignmentContext is an interface to support dynamic dispatch.
type IPropertyAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PropertyName() IPropertyNameContext
	PropertyValue() IPropertyValueContext

	// IsPropertyAssignmentContext differentiates from other interfaces.
	IsPropertyAssignmentContext()
}

type PropertyAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyAssignmentContext() *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_propertyAssignment
	return p
}

func InitEmptyPropertyAssignmentContext(p *PropertyAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_propertyAssignment
}

func (*PropertyAssignmentContext) IsPropertyAssignmentContext() {}

func NewPropertyAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_propertyAssignment

	return p
}

func (s *PropertyAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyAssignmentContext) PropertyName() IPropertyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *PropertyAssignmentContext) PropertyValue() IPropertyValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyValueContext)
}

func (s *PropertyAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterPropertyAssignment(s)
	}
}

func (s *PropertyAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitPropertyAssignment(s)
	}
}

func (s *PropertyAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitPropertyAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) PropertyAssignment() (localctx IPropertyAssignmentContext) {
	localctx = NewPropertyAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, mongoParserRULE_propertyAssignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.PropertyName()
	}
	{
		p.SetState(106)
		p.Match(mongoParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.PropertyValue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyValueContext is an interface to support dynamic dispatch.
type IPropertyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	ObjectLiteral() IObjectLiteralContext
	ArrayLiteral() IArrayLiteralContext
	FunctionCall() IFunctionCallContext

	// IsPropertyValueContext differentiates from other interfaces.
	IsPropertyValueContext()
}

type PropertyValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyValueContext() *PropertyValueContext {
	var p = new(PropertyValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_propertyValue
	return p
}

func InitEmptyPropertyValueContext(p *PropertyValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_propertyValue
}

func (*PropertyValueContext) IsPropertyValueContext() {}

func NewPropertyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyValueContext {
	var p = new(PropertyValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_propertyValue

	return p
}

func (s *PropertyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyValueContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PropertyValueContext) ObjectLiteral() IObjectLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectLiteralContext)
}

func (s *PropertyValueContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *PropertyValueContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *PropertyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterPropertyValue(s)
	}
}

func (s *PropertyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitPropertyValue(s)
	}
}

func (s *PropertyValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitPropertyValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) PropertyValue() (localctx IPropertyValueContext) {
	localctx = NewPropertyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, mongoParserRULE_propertyValue)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case mongoParserStringLiteral, mongoParserNullLiteral, mongoParserBooleanLiteral, mongoParserNumericLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Literal()
		}

	case mongoParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.ObjectLiteral()
		}

	case mongoParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			p.ArrayLiteral()
		}

	case mongoParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(112)
			p.FunctionCall()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NullLiteral() antlr.TerminalNode
	BooleanLiteral() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode
	NumericLiteral() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(mongoParserNullLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(mongoParserBooleanLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(mongoParserStringLiteral, 0)
}

func (s *LiteralContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(mongoParserNumericLiteral, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, mongoParserRULE_literal)
	var _la int

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case mongoParserStringLiteral, mongoParserNullLiteral, mongoParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14336) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case mongoParserNumericLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Match(mongoParserNumericLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyNameContext is an interface to support dynamic dispatch.
type IPropertyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode

	// IsPropertyNameContext differentiates from other interfaces.
	IsPropertyNameContext()
}

type PropertyNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyNameContext() *PropertyNameContext {
	var p = new(PropertyNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_propertyName
	return p
}

func InitEmptyPropertyNameContext(p *PropertyNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_propertyName
}

func (*PropertyNameContext) IsPropertyNameContext() {}

func NewPropertyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameContext {
	var p = new(PropertyNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_propertyName

	return p
}

func (s *PropertyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(mongoParserSTRING_LITERAL, 0)
}

func (s *PropertyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterPropertyName(s)
	}
}

func (s *PropertyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitPropertyName(s)
	}
}

func (s *PropertyNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitPropertyName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) PropertyName() (localctx IPropertyNameContext) {
	localctx = NewPropertyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, mongoParserRULE_propertyName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(mongoParserSTRING_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SingleLineComment() antlr.TerminalNode
	MultiLineComment() antlr.TerminalNode

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_comment
	return p
}

func InitEmptyCommentContext(p *CommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = mongoParserRULE_comment
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = mongoParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) SingleLineComment() antlr.TerminalNode {
	return s.GetToken(mongoParserSingleLineComment, 0)
}

func (s *CommentContext) MultiLineComment() antlr.TerminalNode {
	return s.GetToken(mongoParserMultiLineComment, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mongoListener); ok {
		listenerT.ExitComment(s)
	}
}

func (s *CommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case mongoVisitor:
		return t.VisitComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *mongoParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, mongoParserRULE_comment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		_la = p.GetTokenStream().LA(1)

		if !(_la == mongoParserSingleLineComment || _la == mongoParserMultiLineComment) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
