// Code generated from mongo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // mongo
import "github.com/antlr4-go/antlr/v4"

// BasemongoListener is a complete listener for a parse tree produced by mongoParser.
type BasemongoListener struct{}

var _ mongoListener = &BasemongoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasemongoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasemongoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasemongoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasemongoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMongoCommands is called when production mongoCommands is entered.
func (s *BasemongoListener) EnterMongoCommands(ctx *MongoCommandsContext) {}

// ExitMongoCommands is called when production mongoCommands is exited.
func (s *BasemongoListener) ExitMongoCommands(ctx *MongoCommandsContext) {}

// EnterCommands is called when production commands is entered.
func (s *BasemongoListener) EnterCommands(ctx *CommandsContext) {}

// ExitCommands is called when production commands is exited.
func (s *BasemongoListener) ExitCommands(ctx *CommandsContext) {}

// EnterCommand is called when production command is entered.
func (s *BasemongoListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BasemongoListener) ExitCommand(ctx *CommandContext) {}

// EnterEmptyCommand is called when production emptyCommand is entered.
func (s *BasemongoListener) EnterEmptyCommand(ctx *EmptyCommandContext) {}

// ExitEmptyCommand is called when production emptyCommand is exited.
func (s *BasemongoListener) ExitEmptyCommand(ctx *EmptyCommandContext) {}

// EnterCollection is called when production collection is entered.
func (s *BasemongoListener) EnterCollection(ctx *CollectionContext) {}

// ExitCollection is called when production collection is exited.
func (s *BasemongoListener) ExitCollection(ctx *CollectionContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasemongoListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasemongoListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BasemongoListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BasemongoListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BasemongoListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BasemongoListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *BasemongoListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *BasemongoListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BasemongoListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BasemongoListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterElementList is called when production elementList is entered.
func (s *BasemongoListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *BasemongoListener) ExitElementList(ctx *ElementListContext) {}

// EnterPropertyNameAndValueList is called when production propertyNameAndValueList is entered.
func (s *BasemongoListener) EnterPropertyNameAndValueList(ctx *PropertyNameAndValueListContext) {}

// ExitPropertyNameAndValueList is called when production propertyNameAndValueList is exited.
func (s *BasemongoListener) ExitPropertyNameAndValueList(ctx *PropertyNameAndValueListContext) {}

// EnterPropertyAssignment is called when production propertyAssignment is entered.
func (s *BasemongoListener) EnterPropertyAssignment(ctx *PropertyAssignmentContext) {}

// ExitPropertyAssignment is called when production propertyAssignment is exited.
func (s *BasemongoListener) ExitPropertyAssignment(ctx *PropertyAssignmentContext) {}

// EnterPropertyValue is called when production propertyValue is entered.
func (s *BasemongoListener) EnterPropertyValue(ctx *PropertyValueContext) {}

// ExitPropertyValue is called when production propertyValue is exited.
func (s *BasemongoListener) ExitPropertyValue(ctx *PropertyValueContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasemongoListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasemongoListener) ExitLiteral(ctx *LiteralContext) {}

// EnterPropertyName is called when production propertyName is entered.
func (s *BasemongoListener) EnterPropertyName(ctx *PropertyNameContext) {}

// ExitPropertyName is called when production propertyName is exited.
func (s *BasemongoListener) ExitPropertyName(ctx *PropertyNameContext) {}

// EnterComment is called when production comment is entered.
func (s *BasemongoListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BasemongoListener) ExitComment(ctx *CommentContext) {}
