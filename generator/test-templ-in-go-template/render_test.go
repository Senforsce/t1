package testgotemplates

import (
	"context"
	_ "embed"
	"strings"
	"testing"

	"github.com/senforsce/tndr"
	"github.com/senforsce/tndr/generator/htmldiff"
)

//go:embed expected.html
var expected string

func TestExample(t *testing.T) {
	// Create the t1 component.
	t1Component := greeting()
	html, err := t1.ToGoHTML(context.Background(), t1Component)
	if err != nil {
		t.Fatalf("failed to convert to html: %v", err)
	}

	// Use it within the text/html template.
	b := new(strings.Builder)
	err = example.Execute(b, html)
	if err != nil {
		t.Fatalf("failed to execute template: %v", err)
	}

	// Compare the output with the expected.
	diff, err := htmldiff.DiffStrings(expected, b.String())
	if err != nil {
		t.Fatalf("failed to diff strings: %v", err)
	}
	if diff != "" {
		t.Error(diff)
	}
}
