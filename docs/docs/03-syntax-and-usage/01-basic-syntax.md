# Basic syntax

## Package name and imports

t1 files start with a package name, followed by any required imports, just like Go.

```go
package main

import "fmt"
import "time"
```

## Components

t1 files can also contain components. Components are markup and code that is compiled into functions that return a `t1.Component` interface by running the `t1 generate` command.

Components can contain t1 elements that render HTML, text, expressions that output text or include other templates, and branching statements such as `if` and `switch`, and `for` loops.

```t1 name="header.t1"
package main

t1 headerTemplate(name string) {
  <header data-testid="headerTemplate">
    <h1>{ name }</h1>
  </header>
}
```

## Go code

Outside of t1 Components, t1 files are ordinary Go code.

```t1 name="header.t1"
package main

// Ordinary Go code that we can use in our Component.
var greeting = "Welcome!"

// t1 Component
t1 headerTemplate(name string) {
  <header>
    <h1>{ name }</h1>
    <h2>"{ greeting }" comes from ordinary Go code</h2>
  </header>
}
```
