// Code generated from mongo.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // mongo
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by mongoParser.
type mongoVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by mongoParser#mongoCommands.
	VisitMongoCommands(ctx *MongoCommandsContext) interface{}

	// Visit a parse tree produced by mongoParser#commands.
	VisitCommands(ctx *CommandsContext) interface{}

	// Visit a parse tree produced by mongoParser#command.
	VisitCommand(ctx *CommandContext) interface{}

	// Visit a parse tree produced by mongoParser#emptyCommand.
	VisitEmptyCommand(ctx *EmptyCommandContext) interface{}

	// Visit a parse tree produced by mongoParser#collection.
	VisitCollection(ctx *CollectionContext) interface{}

	// Visit a parse tree produced by mongoParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by mongoParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by mongoParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by mongoParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by mongoParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by mongoParser#elementList.
	VisitElementList(ctx *ElementListContext) interface{}

	// Visit a parse tree produced by mongoParser#propertyNameAndValueList.
	VisitPropertyNameAndValueList(ctx *PropertyNameAndValueListContext) interface{}

	// Visit a parse tree produced by mongoParser#propertyAssignment.
	VisitPropertyAssignment(ctx *PropertyAssignmentContext) interface{}

	// Visit a parse tree produced by mongoParser#propertyValue.
	VisitPropertyValue(ctx *PropertyValueContext) interface{}

	// Visit a parse tree produced by mongoParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by mongoParser#propertyName.
	VisitPropertyName(ctx *PropertyNameContext) interface{}

	// Visit a parse tree produced by mongoParser#comment.
	VisitComment(ctx *CommentContext) interface{}
}
