// Code generated from mongo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // mongo
import "github.com/antlr4-go/antlr/v4"

// mongoListener is a complete listener for a parse tree produced by mongoParser.
type mongoListener interface {
	antlr.ParseTreeListener

	// EnterMongoCommands is called when entering the mongoCommands production.
	EnterMongoCommands(c *MongoCommandsContext)

	// EnterCommands is called when entering the commands production.
	EnterCommands(c *CommandsContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterEmptyCommand is called when entering the emptyCommand production.
	EnterEmptyCommand(c *EmptyCommandContext)

	// EnterCollection is called when entering the collection production.
	EnterCollection(c *CollectionContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterObjectLiteral is called when entering the objectLiteral production.
	EnterObjectLiteral(c *ObjectLiteralContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterElementList is called when entering the elementList production.
	EnterElementList(c *ElementListContext)

	// EnterPropertyNameAndValueList is called when entering the propertyNameAndValueList production.
	EnterPropertyNameAndValueList(c *PropertyNameAndValueListContext)

	// EnterPropertyAssignment is called when entering the propertyAssignment production.
	EnterPropertyAssignment(c *PropertyAssignmentContext)

	// EnterPropertyValue is called when entering the propertyValue production.
	EnterPropertyValue(c *PropertyValueContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterPropertyName is called when entering the propertyName production.
	EnterPropertyName(c *PropertyNameContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// ExitMongoCommands is called when exiting the mongoCommands production.
	ExitMongoCommands(c *MongoCommandsContext)

	// ExitCommands is called when exiting the commands production.
	ExitCommands(c *CommandsContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitEmptyCommand is called when exiting the emptyCommand production.
	ExitEmptyCommand(c *EmptyCommandContext)

	// ExitCollection is called when exiting the collection production.
	ExitCollection(c *CollectionContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitObjectLiteral is called when exiting the objectLiteral production.
	ExitObjectLiteral(c *ObjectLiteralContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitElementList is called when exiting the elementList production.
	ExitElementList(c *ElementListContext)

	// ExitPropertyNameAndValueList is called when exiting the propertyNameAndValueList production.
	ExitPropertyNameAndValueList(c *PropertyNameAndValueListContext)

	// ExitPropertyAssignment is called when exiting the propertyAssignment production.
	ExitPropertyAssignment(c *PropertyAssignmentContext)

	// ExitPropertyValue is called when exiting the propertyValue production.
	ExitPropertyValue(c *PropertyValueContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitPropertyName is called when exiting the propertyName production.
	ExitPropertyName(c *PropertyNameContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)
}
