# Expressions

## String expressions

Within a t1 element, expressions can be used to render strings. Content is automatically escaped using context-aware HTML encoding rules to protect against XSS and CSS injection attacks.

String literals, variables and functions that return a string can be used.

### Literals

You can use Go string literals.

```t1 title="component.t1"
package main

t1 component() {
  <div>{ "print this" }</div>
  <div>{ `and this` }</div>
}
```

```html title="Output"
<div>print this</div>
<div>and this</div>
```

### Variables

Any Go string variable can be used, for example:

- A string function parameter.
- A field on a struct.
- A variable or constant string that is in scope.

```t1 title="/main.t1"
package main

t1 greet(prefix string, p Person) {
  <div>{ prefix } { p.Name }{ exclamation }</div>
}
```

```t1 title="main.go"
package main

type Person struct {
  Name string
}

const exclamation = "!"

func main() {
  p := Person{ Name: "John" }
  component := greet("Hello", p)
  component.Render(context.Background(), os.Stdout)
}
```

```html title="Output"
<div>Hello John!</div>
```

### Functions

Functions that return `string` or `(string, error)` can be used.

```t1 title="component.t1"
package main

import "strings"
import "strconv"

func getString() (string, error) {
  return "DEF", nil
}

t1 component() {
  <div>{ strings.ToUpper("abc") }</div>
  <div>{ getString() }</div>
}
```

```html title="Output"
<div>ABC</div>
<div>DEF</div>
```

If the function returns an error, the `Render` function will return an error containing the location of the error and the underlying error.

### Escaping

t1 automatically escapes strings using HTML escaping rules.

```t1 title="component.t1"
package main

t1 component() {
  <div>{ `</div><script>alert('hello!')</script><div>` }</div>
}
```

```html title="Output"
<div>&lt;/div&gt;&lt;script&gt;alert(&#39;hello!&#39;)&lt;/script&gt;&lt;div&gt;</div>
```
