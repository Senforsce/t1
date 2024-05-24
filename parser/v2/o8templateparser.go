package parser

import (
	"github.com/a-h/parse"
)

var O8TemplateParser = parse.Func(func(pi *parse.Input) (r O8Template, ok bool, err error) {
	start := pi.Index()

	// Parse the name.
	var se o8Expression
	if se, ok, err = o8ExpressionParser.Parse(pi); err != nil || !ok {
		pi.Seek(start)
		return
	}
	r.Name = se.Name
	r.Parameters = se.Parameters

	// Read code expression.
	var e Expression
	if e, ok, err = exp.Parse(pi); err != nil || !ok {
		pi.Seek(start)
		return
	}
	r.Value = e.Value

	// Try for }
	if _, ok, err = closeBraceWithOptionalPadding.Parse(pi); err != nil || !ok {
		err = parse.Error("o8 template: missing closing brace", pi.Position())
		return
	}

	return r, true, nil
})

// o8 Func() {
type o8Expression struct {
	Name       Expression
	Parameters Expression
}

var o8ExpressionNameParser = ExpressionOf(parse.StringFrom(
	parse.Letter,
	parse.StringFrom(parse.AtMost(1000, parse.Any(parse.Letter, parse.ZeroToNine))),
))

var o8ExpressionParser = parse.Func(func(pi *parse.Input) (r o8Expression, ok bool, err error) {
	// Check the prefix first.
	if _, ok, err = parse.String("o8 ").Parse(pi); err != nil || !ok {
		return
	}

	// Once we have the prefix, we must have a name and parameters.
	// Read the name of the function.
	if r.Name, ok, err = o8ExpressionNameParser.Parse(pi); err != nil || !ok {
		err = parse.Error("o8 expression: invalid name", pi.Position())
		return
	}

	// Eat the open bracket.
	if _, ok, err = openBracket.Parse(pi); err != nil || !ok {
		err = parse.Error("o8 expression: parameters missing open bracket", pi.Position())
		return
	}

	// Read the parameters.
	// p Person, other Other, t thing.Thing)
	if r.Parameters, ok, err = ExpressionOf(parse.StringUntil(closeBracket)).Parse(pi); err != nil || !ok {
		err = parse.Error("o8 expression: parameters missing close bracket", pi.Position())
		return
	}

	// Eat ") {".
	if _, ok, err = expressionFuncEnd.Parse(pi); err != nil || !ok {
		err = parse.Error("o8 expression: unterminated (missing ') {')", pi.Position())
		return
	}

	// Expect a newline.
	if _, ok, err = parse.NewLine.Parse(pi); err != nil || !ok {
		err = parse.Error("o8 expression: missing terminating newline", pi.Position())
		return
	}

	return r, true, nil
})
