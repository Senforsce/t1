# Template composition

Templates can be composed using the import expression.

```t1
t1 showAll() {
	@left()
	@middle()
	@right()
}

t1 left() {
	<div>Left</div>
}

t1 middle() {
	<div>Middle</div>
}

t1 right() {
	<div>Right</div>
}
```

```html title="Output"
<div>Left</div>
<div>Middle</div>
<div>Right</div>
```

# Children

Children can be passed to a component for it to wrap.

```t1
t1 showAll() {
	@wrapChildren() {
		<div>Inserted from the top</div>
	}
}

t1 wrapChildren() {
	<div id="wrapper">
		{ children... }
	</div>
}
```

:::note
The use of the `{ children... }` expression in the child component.
:::

```html title="output"
<div id="wrapper">
  <div>Inserted from the top</div>
</div>
```

# Components as parameters

Components can also be passed as parameters and rendered using the `@component` expression.

```t1
package main

t1 heading() {
    <h1>Heading</h1>
}

t1 layout(contents t1.Component) {
	<div id="heading">
		@heading()
	</div>
	<div id="contents">
		@contents
	</div>
}

t1 paragraph(contents string) {
	<p>{ contents }</p>
}
```

```go title="main.go"
package main

import (
	"context"
	"os"
)

func main() {
	c := paragraph("Dynamic contents")
	layout(c).Render(context.Background(), os.Stdout)
}
```

```html title="output"
<div id="heading">
  <h1>Heading</h1>
</div>
<div id="contents">
  <p>Dynamic contents</p>
</div>
```

You can pass `t1` components as parameters to other components within templates using standard Go function call syntax.

```t1
package main

t1 heading() {
    <h1>Heading</h1>
}

t1 layout(contents t1.Component) {
	<div id="heading">
		@heading()
	</div>
	<div id="contents">
		@contents
	</div>
}

t1 paragraph(contents string) {
	<p>{ contents }</p>
}

t1 root() {
	@layout(paragraph("Dynamic contents"))
}
```

```go title="main.go"
package main

import (
	"context"
	"os"
)

func main() {
	root().Render(context.Background(), os.Stdout)
}
```

```html title="output"
<div id="heading">
  <h1>Heading</h1>
</div>
<div id="contents">
  <p>Dynamic contents</p>
</div>
```
