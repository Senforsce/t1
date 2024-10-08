package testtextwhitespace

import (
	"context"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/senforsce/tndr"
)

func TestTextWhitespace(t *testing.T) {
	for _, test := range []struct {
		name     string
		input    t1.Component
		expected string
	}{
		{
			name:     "whitespace is added within t1 statements",
			input:    WhitespaceIsAddedWithinTemplStatements(),
			expected: WhitespaceIsAddedWithinTemplStatementsExpected,
		},
		{
			name:     "inline elements are not padded",
			input:    InlineElementsAreNotPadded(),
			expected: InlineElementsAreNotPaddedExpected,
		},
		{
			name:     "whitespace in HTML is normalised",
			input:    WhiteSpaceInHTMLIsNormalised(),
			expected: WhiteSpaceInHTMLIsNormalisedExpected,
		},
		{
			name:     "whitespace around values is maintained",
			input:    WhiteSpaceAroundValues(),
			expected: WhiteSpaceAroundValuesExpected,
		},
		{
			name:     "whitespace around templated values is maintained",
			input:    WhiteSpaceAroundTemplatedValues("t1", "allows whitespace around templated values."),
			expected: WhiteSpaceAroundTemplatedValuesExpected,
		},
	} {
		w := new(strings.Builder)
		err := test.input.Render(context.Background(), w)
		if err != nil {
			t.Errorf("failed to render: %v", err)
		}
		if diff := cmp.Diff(test.expected, w.String()); diff != "" {
			t.Error(diff)
		}
	}
}
