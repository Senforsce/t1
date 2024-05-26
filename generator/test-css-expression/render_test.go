package testcssexpression

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/senforsce/t1"
)

var expected = t1.ComponentCSSClass{
	ID:    "className_34fc",
	Class: t1.SafeCSS(`.className_34fc{background-color:#ffffff;max-height:calc(100vh - 170px);color:#ff0000;}`),
}

func TestCSSExpression(t *testing.T) {
	if diff := cmp.Diff(expected, className()); diff != "" {
		t.Error(diff)
	}
}
