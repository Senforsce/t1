package generator

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/senforsce/t1/parser/v2"
)

func TestGeneratorSourceMap(t *testing.T) {
	w := new(bytes.Buffer)
	g := generator{
		w:         NewRangeWriter(w),
		sourceMap: parser.NewSourceMap(),
	}
	invalidExp := parser.TemplateFileGoExpression{
		Expression: parser.Expression{
			Value: "line1\nline2",
		},
	}
	if err := g.writeGoExpression(invalidExp); err != nil {
		t.Fatalf("failed to write Go expression: %v", err)
	}

	expected := parser.NewPosition(0, 0, 0)
	actual, ok := g.sourceMap.TargetPositionFromSource(0, 0)
	if !ok {
		t.Errorf("failed to get matching target")
	}
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("unexpected target:\n%v", diff)
	}

	withCommentExp := parser.TemplateFileGoExpression{
		Expression: parser.Expression{
			Value: `package main

// A comment.
t1 h1() {
	<h1></h1>
}
			`,
		},
	}
	if err := g.writeGoExpression(withCommentExp); err != nil {
		t.Fatalf("failed to write Go expression: %v", err)
	}
}
