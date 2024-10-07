package testcontext

import (
	"context"
	_ "embed"
	"testing"

	"github.com/senforsce/tndr/generator/htmldiff"
)

//go:embed expected.html
var expected string

func Test(t *testing.T) {
	component := render()

	ctx := context.WithValue(context.Background(), contextKeyName, "test")

	diff, err := htmldiff.DiffCtx(ctx, component, expected)
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}
