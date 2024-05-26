# Components

t1 Components are markup and code that is compiled into functions that return a `t1.Component` interface by running the `t1 generate` command.

Components can contain t1 elements that render HTML, text, expressions that output text or include other templates, and branching statements such as `if` and `switch`, and `for` loops.

```t1 title="header.t1"
package main

t1 headerTemplate(name string) {
  <header data-testid="headerTemplate">
    <h1>{ name }</h1>
  </header>
}
```

:::tip
Since t1 produces Go code, you can share templates the same way that you share Go code - by sharing your Go module.

t1 follows the same rules as Go. If a `t1` block starts with an uppercase letter, then it is public, otherwise, it is private.
:::

## Code-only components

Since t1 Components ultimately implement the `t1.Component`, any code that implements the interface can be used in place of a t1 component generated from a `*.t1` file.

```go
package main

import (
	"context"
	"io"
	"os"

	"github.com/senforsce/t1"
)

func button(text string) t1.Component {
	return t1.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, "<button>"+text+"</button>")
		return err
	})
}

func main() {
	button("Click me").Render(context.Background(), os.Stdout)
}
```

```html title="Output"
<button>Click me</button>
```

:::warning
This code is unsafe! In code-only components, you're responsible for escaping the HTML content yourself, e.g. with the `t1.EscapeString` function.
:::

## Method components

t1 components can be returned from methods (functions attached to types).

Go code:

```t1
package main

type Data struct {
	message string
}

t1 (d Data) Method() {
	<div>{ d.message }</div>
}

func main() {
	d := Data{
		message: "You can implement methods on a type.",
	}
	d.Method().Render(context.Background(), os.Stdout)
}
```
