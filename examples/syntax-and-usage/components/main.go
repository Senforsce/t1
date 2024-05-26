package main

import (
	"context"
	"fmt"
	"html"
	"io"
	"os"

	"github.com/senforsce/t1"
)

func main() {
	ctx := context.Background()
	list([]string{"a", "b", "c"}).Render(ctx, os.Stdout)
	codeList([]string{"A", "B", "C"}).Render(ctx, os.Stdout)
}

func codeList(items []string) t1.Component {
	return t1.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		if _, err = io.WriteString(w, "<ol>"); err != nil {
			return
		}
		for _, item := range items {
			if _, err = io.WriteString(w, fmt.Sprintf("<li>%s</li>", html.EscapeString(item))); err != nil {
				return
			}
		}
		if _, err = io.WriteString(w, "</ol>"); err != nil {
			return
		}
		return nil
	})
}
