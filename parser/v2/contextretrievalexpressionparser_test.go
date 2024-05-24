package parser

import (
	"testing"

	"github.com/a-h/parse"
	"github.com/google/go-cmp/cmp"
)

func TestContextRetrievalExpressionParser(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected HDTContextRetrievalExpression
	}{

		{
			name:  "multiline triple can be on one line",
			input: `/- senforsce:Version -/`,
			expected: HDTContextRetrievalExpression{
				Contents: " senforsce:Version ",
			},
		},
		{
			name: "multiline triple can span lines",
			input: `/- subject:Test 
			rdf:has senforsce:HTMX -/`,
			expected: HDTContextRetrievalExpression{
				Contents: " subject:Test\nrdf:has senforsce:HTMX ",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			input := parse.NewInput(tt.input)
			result, ok, err := contextRetrievalExpression.Parse(input)
			if err != nil {
				t.Fatalf("parser error: %v", err)
			}
			if !ok {
				t.Fatalf("failed to parse at %d", input.Index())
			}
			if diff := cmp.Diff(tt.expected, result); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestContextRetrievalExpressionParserError(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected error
	}{
		{
			name:  "unclosed triple comments result in an error",
			input: `/- unclosed triple literal`,
			expected: parse.Error("expected end triple literal '-/' not found",
				parse.Position{
					Index: 24,
					Line:  0,
					Col:   24,
				}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			input := parse.NewInput(tt.input)
			_, _, err := contextRetrievalExpression.Parse(input)
			if diff := cmp.Diff(tt.expected, err); diff != "" {
				t.Error(diff)
			}
		})
	}
}
