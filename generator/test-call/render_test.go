package testcall

import (
	_ "embed"
	"testing"

	"github.com/senforsce/t1/generator/htmldiff"
)

//go:embed expected.html
var expected string

func Test(t *testing.T) {
	component := showAll()

	diff, err := htmldiff.Diff(component, expected)
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}
