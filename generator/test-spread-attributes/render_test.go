package testspreadattributes

import (
	_ "embed"
	"testing"

	"github.com/senforsce/t1"
	"github.com/senforsce/t1/generator/htmldiff"
)

//go:embed expected.html
var expected string

func Test(t *testing.T) {
	component := BasicTemplate(t1.Attributes{
		// Should render as `bool` as the value is true, and the conditional render is also true.
		"bool": t1.KV(true, true),
		// Should not render, as the conditional render value is false.
		"bool-disabled": t1.KV(true, false),
		// Should render as `dateId="my-custom-id"`.
		"dateId": "my-custom-id",
		// Should render as `hx-get="/page"`.
		"hx-get": "/page",
		// Should render as `id="test"`.
		"id": "test",
		// Should not render, as the attribute value, and the conditional render value is false.
		"no-bool": t1.KV(false, false),
		// Should not render, as the conditional render value is false.
		"no-text": t1.KV("empty", false),
		// Should render as `nonshare`, as the value is true.
		"nonshade": true,
		// Should not render, as the value is false.
		"shade": false,
		// Should render text="lorem" as the value is true.
		"text": t1.KV("lorem", true),
		// Optional attribute based on result of func() bool.
		"optional-from-func-false": func() bool { return false },
		// Optional attribute based on result of func() bool.
		"optional-from-func-true": func() bool { return true },
	})

	diff, err := htmldiff.Diff(component, expected)
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}
