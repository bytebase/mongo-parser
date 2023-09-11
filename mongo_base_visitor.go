// Code generated from mongo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // mongo
import "github.com/antlr4-go/antlr/v4"

type BasemongoVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasemongoVisitor) VisitMongoCommands(ctx *MongoCommandsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitCommands(ctx *CommandsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitCommand(ctx *CommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitEmptyCommand(ctx *EmptyCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitCollection(ctx *CollectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitElementList(ctx *ElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitPropertyNameAndValueList(ctx *PropertyNameAndValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitPropertyAssignment(ctx *PropertyAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitPropertyValue(ctx *PropertyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitPropertyName(ctx *PropertyNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemongoVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}
