# Using with `html/template`

Templ components can be used with the Go standard library [`html/template`](https://pkg.go.dev/html/template) package.

## Using `html/template` in a t1 component

To use an existing `html/template` in a t1 component, use the `t1.FromGoHTML` function.

```t1 title="component.t1"
package testgotemplates

import "html/template"

var goTemplate = template.Must(template.New("example").Parse("<div>{{ . }}</div>"))

t1 Example() {
	<!DOCTYPE html>
	<html>
		<body>
			@t1.FromGoHTML(goTemplate, "Hello, World!")
		</body>
	</html>
}
```

```go title="main.go"
func main() {
	Example.Render(context.Background(), os.Stdout)
}
```

```html title="Output"
<!DOCTYPE html>
<html>
  <body>
    <div>Hello, World!</div>
  </body>
</html>
```

## Using a t1 component with `html/template`

To use a t1 component within a `html/template`, use the `t1.ToGoHTML` function to render the component into a `t1ate.HTML value`.

```t1 title="component.html"
package testgotemplates

import "html/template"

var example = template.Must(template.New("example").Parse(`<!DOCTYPE html>
<html>
	<body>
		{{ . }}
	</body>
</html>
`))

t1 greeting() {
	<div>Hello, World!</div>
}
```

```go title="main.go"
func main() {
	// Create the t1 component.
	t1Component := greeting()

	// Render the t1 component to a `t1ate.HTML` value.
	html, err := t1.ToGoHTML(context.Background(), t1Component)
	if err != nil {
		t.Fatalf("failed to convert to html: %v", err)
	}

	// Use the `t1ate.HTML` value within the text/html template.
	err = example.Execute(os.Stdout, html)
	if err != nil {
		t.Fatalf("failed to execute template: %v", err)
	}
}
```

```html title="Output"
<!DOCTYPE html>
<html>
  <body>
    <div>Hello, World!</div>
  </body>
</html>
```
