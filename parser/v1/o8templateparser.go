package parser

import (
	"io"
	"strings"

	"github.com/a-h/lexical/parse"
)

func newO8TemplateParser() o8TemplateParser {
	return o8TemplateParser{}
}

type o8TemplateParser struct {
}

var endO8Parser = createEndParser("endo8") // {% endo8 %}

func (p o8TemplateParser) Parse(pi parse.Input) parse.Result {
	var r ScriptTemplate

	// Parse the name.
	pr := newO8ExpressionParser().Parse(pi)
	if !pr.Success {
		return pr
	}
	r.Name = pr.Item.(o8Expression).Name
	r.Parameters = pr.Item.(o8Expression).Parameters

	from := NewPositionFromInput(pi)
	// Read until {% endo8 %}
	sr := parse.StringUntil(endO8Parser)(pi)
	if sr.Error != nil {
		return sr
	}
	if sr.Success {
		r.Value = sr.Item.(string)
	}

	// Eat the final {% endo8 %}
	if endO8Parser(pi).Success {
		return parse.Success("script", r, nil)
	}
	return parse.Failure("script", newParseError("expected {% endo8 %} not found", from, NewPositionFromInput(pi)))
}

// {% o8 Func() %}
type o8Expression struct {
	Name       Expression
	Parameters Expression
}

func newO8ExpressionParser() o8ExpressionParser {
	return o8ExpressionParser{}
}

type o8ExpressionParser struct {
}

var o8ExpressionNameParser = parse.All(parse.WithStringConcatCombiner,
	parse.Letter,
	parse.Many(parse.WithStringConcatCombiner, 0, 1000, parse.Any(parse.Letter, parse.ZeroToNine)),
)

var o8ExpressionStartParser = createStartParser("o8")

func (p o8ExpressionParser) Parse(pi parse.Input) parse.Result {
	var r o8Expression

	// Check the prefix first.
	prefixResult := o8ExpressionStartParser(pi)
	if !prefixResult.Success {
		return prefixResult
	}

	// Once we have the prefix, we must have a name and parameters.
	// Read the name of the function.
	from := NewPositionFromInput(pi)
	pr := o8ExpressionNameParser(pi)
	if pr.Error != nil && pr.Error != io.EOF {
		return pr
	}
	// If there's no match, the name wasn't correctly terminated.
	if !pr.Success {
		return parse.Failure("o8ExpressionParser", newParseError("o8 expression: invalid name", from, NewPositionFromInput(pi)))
	}
	to := NewPositionFromInput(pi)
	r.Name = NewExpression(pr.Item.(string), from, to)
	from = to

	// Eat the open bracket.
	if lb := parse.Rune('(')(pi); !lb.Success {
		return parse.Failure("o8ExpressionParser", newParseError("o8 expression: parameters missing open bracket", from, NewPositionFromInput(pi)))
	}

	// Read the parameters.
	from = NewPositionFromInput(pi)
	pr = parse.StringUntil(parse.Rune(')'))(pi) // p Person, other Other, t thing.Thing)
	if pr.Error != nil && pr.Error != io.EOF {
		return pr
	}
	// If there's no match, the name wasn't correctly terminated.
	if !pr.Success {
		return parse.Failure("o8ExpressionParser", newParseError("o8 expression: parameters missing close bracket", from, NewPositionFromInput(pi)))
	}
	r.Parameters = NewExpression(strings.TrimSuffix(pr.Item.(string), ")"), from, NewPositionFromInput(pi))

	// Eat ") %}".
	from = NewPositionFromInput(pi)
	if lb := expressionFuncEnd(pi); !lb.Success {
		return parse.Failure("o8ExpressionParser", newParseError("o8 expression: unterminated (missing ') %}')", from, NewPositionFromInput(pi)))
	}

	// Expect a newline.
	from = NewPositionFromInput(pi)
	if lb := newLine(pi); !lb.Success {
		return parse.Failure("o8ExpressionParser", newParseError("o8 expression: missing terminating newline", from, NewPositionFromInput(pi)))
	}

	return parse.Success("o8ExpressionParser", r, nil)
}
