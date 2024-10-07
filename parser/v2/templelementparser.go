package parser

import (
	"github.com/a-h/parse"
	"github.com/senforsce/tndr/parser/v2/goexpression"
)

type t1ElementExpressionParser struct{}

func (p t1ElementExpressionParser) Parse(pi *parse.Input) (n Node, ok bool, err error) {
	// Check the prefix first.
	if _, ok, err = parse.Rune('@').Parse(pi); err != nil || !ok {
		return
	}

	var r TemplElementExpression
	// Parse the Go expression.
	if r.Expression, err = parseGo("t1 element", pi, goexpression.TemplExpression); err != nil {
		return r, false, err
	}

	// Once we've got a start expression, check to see if there's an open brace for children. {\n.
	var hasOpenBrace bool
	_, hasOpenBrace, err = openBraceWithOptionalPadding.Parse(pi)
	if err != nil {
		return
	}
	if !hasOpenBrace {
		return r, true, nil
	}

	// Once we've had the start of an element's children, we must conclude the block.

	// Node contents.
	np := newTemplateNodeParser(closeBraceWithOptionalPadding, "t1 element closing brace")
	var nodes Nodes
	if nodes, ok, err = np.Parse(pi); err != nil || !ok {
		err = parse.Error("@"+r.Expression.Value+": expected nodes, but none were found", pi.Position())
		return
	}
	r.Children = nodes.Nodes

	// Read the required closing brace.
	if _, ok, err = closeBraceWithOptionalPadding.Parse(pi); err != nil || !ok {
		err = parse.Error("@"+r.Expression.Value+": missing end (expected '}')", pi.Position())
		return
	}

	return r, true, nil
}

var t1ElementExpression t1ElementExpressionParser
