package parser

import (
	"github.com/a-h/parse"
)

var contextRetrievalExpressionStart = parse.String("/-")
var contextRetrievalExpressionEnd = parse.String("-/")

type contextRetrievalExpressionParser struct {
}

var contextRetrievalExpressionP = contextRetrievalExpressionParser{}

func (p contextRetrievalExpressionParser) Parse(pi *parse.Input) (n Node, ok bool, err error) {
	// Comment start.
	var c HDTContextRetrievalExpression
	if _, ok, err = contextRetrievalExpressionStart.Parse(pi); err != nil || !ok {
		return
	}

	// Once we've got the comment start sequence, parse anything until the end
	// sequence as the comment contents.
	if c.Contents, ok, err = parse.StringUntil(contextRetrievalExpressionEnd).Parse(pi); err != nil || !ok {
		err = parse.Error("expected end comment literal '-/' not found", pi.Position())
		return
	}
	// Move past the end element.
	_, _, _ = contextRetrievalExpressionEnd.Parse(pi)
	return c, true, nil
}

var contextRetrievalExpression = parse.Any[Node](goSingleLineComment, contextRetrievalExpressionP)
