package testcomplexattributes

import (
	_ "embed"
	"testing"

	"github.com/senforsce/tndr/generator/htmldiff"
)

//go:embed expected.html
var expected string

func Test(t *testing.T) {
	component := ComplexAttributes()

	diff, err := htmldiff.Diff(component, expected)
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}
