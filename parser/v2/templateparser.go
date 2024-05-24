package parser

import (
	"fmt"

	"github.com/a-h/parse"
)

// TemplateExpression.
// t1 Func(p Parameter) {
// t1 (data Data) Func(p Parameter) {
// t1 (data []string) Func(p Parameter) {

// FallbackTemplateExpression.
// templ Func(p Parameter) {
// templ (data Data) Func(p Parameter) {
// templ (data []string) Func(p Parameter) {
type templateExpression struct {
	Expression Expression
}

<<<<<<< HEAD
var templateExpressionParser = parse.Func(func(pi *parse.Input) (r templateExpression, ok bool, err error) {
	start := pi.Index()

	if !peekPrefix(pi, "templ ") {
		return r, false, nil
=======
var templateExpressionStartParser = parse.String("t1 ")
var fallbackTemplateExpressionStartParser = parse.String("templ ")

var templateExpressionParser = parse.Func(func(pi *parse.Input) (r templateExpression, ok bool, err error) {
	// Check the prefix first.
	if _, ok, err = templateExpressionStartParser.Parse(pi); err != nil || !ok {
		if _, ok, err = fallbackTemplateExpressionStartParser.Parse(pi); err != nil || !ok {
			return
		}
>>>>>>> 0c99b15 (Big Bang: removed storybook - to be replaced by tndr and add support for Ontologies)
	}

	// Once we have the prefix, everything to the brace is Go.
	// e.g.
	// t1  (x []string) Test() {
	// or templ  (x []string) Test() {
	// becomes:
	// func (x []string) Test() templ.Component {
<<<<<<< HEAD
	if _, r.Expression, err = parseTemplFuncDecl(pi); err != nil {
		return r, false, err
=======

	// Once we've got a prefix, read until {\n.
	until := parse.All(openBraceWithOptionalPadding, parse.NewLine)
	msg := "thunderf1sh: malformed t1 expression, expected `t1 functionName() {` or `templ functionName() {`"
	if r.Expression, ok, err = ExpressionOf(parse.StringUntil(until)).Parse(pi); err != nil || !ok {
		err = parse.Error(msg, pi.Position())
		return
>>>>>>> 0c99b15 (Big Bang: removed storybook - to be replaced by tndr and add support for Ontologies)
	}

	// Eat " {\n".
	if _, ok, err = parse.All(openBraceWithOptionalPadding, parse.NewLine).Parse(pi); err != nil || !ok {
		err = parse.Error("templ: malformed templ expression, expected `templ functionName() {`", pi.PositionAt(start))
		return
	}

	return r, true, nil
})

const (
	unterminatedMissingCurly = `unterminated (missing closing '{\n') - https://templ.guide/syntax-and-usage/statements#incomplete-statements`
	unterminatedMissingEnd   = `missing end (expected '}') - https://templ.guide/syntax-and-usage/statements#incomplete-statements`
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
	docType,                    // <!DOCTYPE html>
	htmlComment,                // <!--
	goComment,                  // // or /*
	contextRetrievalExpression, // /- .... -/  for retrieval of context values as string
	rawElements,                // <text>, <>, or <style> element (special behaviour - contents are not parsed).
	element,                    // <a>, <br/> etc.
	ifExpression,               // if {}
	forExpression,              // for {}
	switchExpression,           // switch {}
	callTemplateExpression,     // {! TemplateName(a, b, c) }
	templElementExpression,     // @TemplateName(a, b, c) { <div>Children</div> }
	childrenExpression,         // { children... }
	stringExpression,           // { "abc" }
	whitespaceExpression,       // { " " }
	textParser,                 // anything &amp; everything accepted...
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
			if n, ok := node.(CallTemplateExpression); ok {
				op.Diagnostics = append(op.Diagnostics, Diagnostic{
					Message: "`{! foo }` syntax is deprecated. Use `@foo` syntax instead. Run `templ fmt .` to fix all instances.",
					Range:   n.Expression.Range,
				})
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

		err = fmt.Errorf("%v not found", p.untilName)
		return
	}

	return op, true, nil
}
