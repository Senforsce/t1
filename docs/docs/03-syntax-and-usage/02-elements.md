# Elements

t1 elements are used to render HTML within t1 components.

```t1 title="button.t1"
package main

t1 button(text string) {
	<button class="button">{ text }</button>
}
```

```go title="main.go"
package main

import (
	"context"
	"os"
)

func main() {
	button("Click me").Render(context.Background(), os.Stdout)
}
```

```html title="Output"
<button class="button">Click me</button>
```

:::info
t1 automatically minifies HTML responses, output is shown formatted for readability.
:::

## Tags must be closed

t1 requires that all HTML elements are closed with either a closing tag (`</a>`), or by using a self-closing element (`<hr/>`).

t1 is aware of which HTML elements are "void", and will omit the closing `/` from the element.

```t1 title="button.t1"
package main

t1 component() {
	<div>Test</div>
	<img src="images/test.png"/>
	<br/>
}
```

```t1 title="Output"
<div>Test</div>
<img src="images/test.png">
<br>
```

## Attributes and elements can contain expressions

t1 elements can contain placeholder expressions for attributes and content.

```t1 title="button.t1"
package main

t1 button(name string, content string) {
	<button value={ name }>{ content }</button>
}
```

Rendering the component to stdout, we can see the results.

```go title="main.go"
func main() {
	component := button("John", "Say Hello")
	component.Render(context.Background(), os.Stdout)
}
```

```html title="Output"
<button value="John">Say Hello</button>
```
