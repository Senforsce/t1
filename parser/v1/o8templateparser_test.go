package parser

import (
	"testing"

	"github.com/a-h/lexical/input"
	"github.com/google/go-cmp/cmp"
)

func TestO8TemplateParser(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected O8Template
	}{
		{
			name: "o8: no parameters, no content",
			input: `{% o8 Name() %}
{% endo8 %}`,
			expected: O8Template{
				Name: Expression{
					Value: "Name",
					Range: Range{
						From: Position{
							Index: 10,
							Line:  1,
							Col:   10,
						},
						To: Position{
							Index: 14,
							Line:  1,
							Col:   14,
						},
					},
				},
				Parameters: Expression{
					Value: "",
					Range: Range{
						From: Position{
							Index: 15,
							Line:  1,
							Col:   15,
						},
						To: Position{
							Index: 15,
							Line:  1,
							Col:   15,
						},
					},
				},
			},
		},
		{
			name: "o8: no spaces",
			input: `{%o8 Name()%}
{% endo8 %}`,
			expected: O8Template{
				Name: Expression{
					Value: "Name",
					Range: Range{
						From: Position{
							Index: 9,
							Line:  1,
							Col:   9,
						},
						To: Position{
							Index: 13,
							Line:  1,
							Col:   13,
						},
					},
				},
				Parameters: Expression{
					Value: "",
					Range: Range{
						From: Position{
							Index: 14,
							Line:  1,
							Col:   14,
						},
						To: Position{
							Index: 14,
							Line:  1,
							Col:   14,
						},
					},
				},
			},
		},
		{
			name: "o8: containing a Turtle triple Sentence",
			input: `{% o8 Name() %}
sen:AbdoulSy rdfs:type schema:Person .
{% endo8 %}`,
			expected: O8Template{
				Name: Expression{
					Value: "Name",
					Range: Range{
						From: Position{
							Index: 10,
							Line:  1,
							Col:   10,
						},
						To: Position{
							Index: 14,
							Line:  1,
							Col:   14,
						},
					},
				},
				Parameters: Expression{
					Value: "",
					Range: Range{
						From: Position{
							Index: 15,
							Line:  1,
							Col:   15,
						},
						To: Position{
							Index: 15,
							Line:  1,
							Col:   15,
						},
					},
				},
				Value: `sen:AbdoulSy rdfs:type schema:Person .` + "\n",
			},
		},
		{
			name: "o8: single argument",
			input: `{% o8 Name(value string) %}
sen:AbdoulSy sen:hasName value .
{% o8 %}`,
			expected: O8Template{
				Name: Expression{
					Value: "Name",
					Range: Range{
						From: Position{
							Index: 6,
							Line:  1,
							Col:   6,
						},
						To: Position{
							Index: 10,
							Line:  1,
							Col:   10,
						},
					},
				},
				Parameters: Expression{
					Value: "value string",
					Range: Range{
						From: Position{
							Index: 11,
							Line:  1,
							Col:   11,
						},
						To: Position{
							Index: 23,
							Line:  1,
							Col:   23,
						},
					},
				},
				Value: `sen:AbdoulSy sen:hasName value .` + "\n",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			input := input.NewFromString(tt.input)
			result := newO8TemplateParser().Parse(input)
			if result.Error != nil {
				t.Fatalf("parser error: %v", result.Error)
			}
			if !result.Success {
				t.Fatalf("failed to parse at %d", input.Index())
			}
			if diff := cmp.Diff(tt.expected, result.Item); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
