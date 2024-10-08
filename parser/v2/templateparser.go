package parser

import (
	"fmt"

	"github.com/a-h/parse"
)

// TemplateExpression.

// TemplateExpression.
// t1 Func(p Parameter) {
// t1 (data Data) Func(p Parameter) {
// t1 (data []string) Func(p Parameter) {
type templateExpression struct {
	Expression Expression
}

var templateExpressionParser = parse.Func(func(pi *parse.Input) (r templateExpression, ok bool, err error) {
	start := pi.Index()

	if !peekPrefix(pi, "t1 ") {
		return r, false, nil
	}

	// Once we have the prefix, everything to the brace is Go.
	// e.g.
	// t1 (x []string) Test() {
	// becomes:
	// func (x []string) Test() t1.Component {
	if _, r.Expression, err = parseT1FuncDecl(pi); err != nil {
		return r, false, err
	}

	// Eat " {\n".
	if _, ok, err = parse.All(openBraceWithOptionalPadding, parse.StringFrom(parse.Optional(parse.NewLine))).Parse(pi); err != nil || !ok {
		err = parse.Error("templ: malformed t1 expression, expected `t1 functionName() {`", pi.PositionAt(start))
		return
	}

	return r, true, nil
})

const (
	unterminatedMissingCurly = `unterminated (missing closing '{\n') - https://t1.guide/syntax-and-usage/statements#incomplete-statements`
	unterminatedMissingEnd   = `missing end (expected '}') - https://t1.guide/syntax-and-usage/statements#incomplete-statements`
)

// Template node (element, call, if, switch, for, whitespace etc.)
func newTemplateNodeParser[TUntil any](until parse.Parser[TUntil], untilName string) templateNodeParser[TUntil] {
	return templateNodeParser[TUntil]{
		until:     until,
		untilName: untilName,
	}
}

type templateNodeParser[TUntil any] struct {
	until     parse.Parser[TUntil]
	untilName string
}

var rawElements = parse.Any[Node](styleElement, scriptElement)

var templateNodeParsers = []parse.Parser[Node]{
	docType,      // <!DOCTYPE html>
	htmlComment,  // <!--
	goComment,    // // or /*
	rawElements,  // <text>, <>, or <style> element (special behaviour - contents are not parsed).
	element,      // <a>, <br/> etc.
	ifExpression, // if {}
	contextRetrievalExpressionP,
	forExpression,          // for {}
	switchExpression,       // switch {}
	callTemplateExpression, // {! TemplateName(a, b, c) }
	t1ElementExpression,    // @TemplateName(a, b, c) { <div>Children</div> }
	childrenExpression,     // { children... }
	goCode,                 // {{ myval := x.myval }}
	stringExpression,       // { "abc" }
	whitespaceExpression,   // { " " }
	textParser,             // anything &amp; everything accepted...
}

func (p templateNodeParser[T]) Parse(pi *parse.Input) (op Nodes, ok bool, err error) {
	for {
		// Check if we've reached the end.
		if p.until != nil {
			start := pi.Index()
			_, ok, err = p.until.Parse(pi)
			if err != nil {
				return
			}
			if ok {
				pi.Seek(start)
				return op, true, nil
			}
		}

		// Attempt to parse a node.
		// Loop through the parsers and try to parse a node.
		var matched bool
		for _, p := range templateNodeParsers {
			var node Node
			node, matched, err = p.Parse(pi)
			if err != nil {
				return Nodes{}, false, err
			}
			if matched {
				op.Nodes = append(op.Nodes, node)
				break
			}
		}
		if matched {
			continue
		}

		if p.until == nil {
			// In this case, we're just reading as many nodes as we can until we can't read any more.
			// If we've reached here, we couldn't find a node.
			// The element parser checks the final node returned to make sure it's the expected close tag.
			break
		}

		err = UntilNotFoundError{
			ParseError: parse.Error(fmt.Sprintf("%v not found", p.untilName), pi.Position()),
		}
		return
	}

	return op, true, nil
}

type UntilNotFoundError struct {
	parse.ParseError
}
