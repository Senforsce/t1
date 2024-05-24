package parser

import (
	"testing"

	"github.com/a-h/parse"
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
			input: `o8 Name() {
}`,
			expected: O8Template{
				Name: Expression{
					Value: "Name",
					Range: Range{
						From: Position{
							Index: 3,
							Line:  0,
							Col:   3,
						},
						To: Position{
							Index: 7,
							Line:  0,
							Col:   7,
						},
					},
				},
				Parameters: Expression{
					Value: "",
					Range: Range{
						From: Position{
							Index: 8,
							Line:  0,
							Col:   8,
						},
						To: Position{
							Index: 8,
							Line:  0,
							Col:   8,
						},
					},
				},
			},
		},
		{
			name: "o8: no spaces",
			input: `o8 Name(){
}`,
			expected: O8Template{
				Name: Expression{
					Value: "Name",
					Range: Range{
						From: Position{
							Index: 7,
							Line:  0,
							Col:   7,
						},
						To: Position{
							Index: 11,
							Line:  0,
							Col:   11,
						},
					},
				},
				Parameters: Expression{
					Value: "",
					Range: Range{
						From: Position{
							Index: 8,
							Line:  0,
							Col:   8,
						},
						To: Position{
							Index: 8,
							Line:  0,
							Col:   8,
						},
					},
				},
			},
		},
		{
			name: "o8: containing a Turtle Triple",
			input: `o8 Name() {
sen:AbdoulSy rdfs:type schema:Person .
}`,
			expected: O8Template{
				Name: Expression{
					Value: "Name",
					Range: Range{
						From: Position{
							Index: 3,
							Line:  0,
							Col:   3,
						},
						To: Position{
							Index: 7,
							Line:  0,
							Col:   7,
						},
					},
				},
				Parameters: Expression{
					Value: "",
					Range: Range{
						From: Position{
							Index: 8,
							Line:  0,
							Col:   8,
						},
						To: Position{
							Index: 8,
							Line:  0,
							Col:   8,
						},
					},
				},
				Value: `sen:AbdoulSy rdfs:type schema:Person .` + "\n",
			},
		},
		{
			name: "o8: single argument",
			input: `o8 Name(value string) {
sen:AbdoulSy sen:hasNickname value .
}`,
			expected: O8Template{
				Name: Expression{
					Value: "Name",
					Range: Range{
						From: Position{
							Index: 3,
							Line:  0,
							Col:   3,
						},
						To: Position{
							Index: 7,
							Line:  0,
							Col:   7,
						},
					},
				},
				Parameters: Expression{
					Value: "value string",
					Range: Range{
						From: Position{
							Index: 12,
							Line:  0,
							Col:   12,
						},
						To: Position{
							Index: 24,
							Line:  0,
							Col:   24,
						},
					},
				},
				Value: `sen:AbdoulSy sen:hasNickname value .` + "\n",
			},
		},
		{
			name: "o8: comment with single quote",
			input: `o8 Name() {
	//'
} Trailing '`, // Without a single quote later, issue #360 isn't triggered.
			expected: O8Template{
				Name: Expression{
					Value: "Name",
					Range: Range{
						From: Position{
							Index: 3,
							Line:  0,
							Col:   3,
						},
						To: Position{
							Index: 7,
							Line:  0,
							Col:   7,
						},
					},
				},
				Parameters: Expression{
					Value: "",
					Range: Range{
						From: Position{
							Index: 8,
							Line:  0,
							Col:   8,
						},
						To: Position{
							Index: 8,
							Line:  0,
							Col:   8,
						},
					},
				},
				Value: `	//'` + "\n",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			input := parse.NewInput(tt.input)
			actual, ok, err := O8TemplateParser.Parse(input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !ok {
				t.Fatalf("unexpected failure for input %q", tt.input)
			}
			if diff := cmp.Diff(tt.expected, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}
