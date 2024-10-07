package testtemplelement

import (
	"testing"

	_ "embed"

	"github.com/senforsce/tndr/generator/htmldiff"
)

//go:embed expected.html
var expected string

func Test(t *testing.T) {
	component := template()

	diff, err := htmldiff.Diff(component, expected)
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}
