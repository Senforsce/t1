# Creating an HTTP server with t1

### Static pages

To use a t1 component as a HTTP handler, the `t1.Handler` function can be used.

This is suitable for use when the component is not used to display dynamic data.

```go title="components.t1"
package main

t1 hello() {
	<div>Hello</div>
}
```

```go title="main.go"
package main

import (
	"net/http"

	"github.com/senforsce/t1"
)

func main() {
	http.Handle("/", t1.Handler(hello()))

	http.ListenAndServe(":8080", nil)
}
```

### Displaying fixed data

In the previous example, the `hello` component does not take any parameters. Let's display the time when the server was started instead.

```go title="components.t1"
package main

import "time"

t1 timeComponent(d time.Time) {
	<div>{ d.String() }</div>
}

t1 notFoundComponent() {
	<div>404 - Not found</div>
}
```

```go title="main.go"
package main

import (
	"net/http"
	"time"

	"github.com/senforsce/t1"
)

func main() {
	http.Handle("/", t1.Handler(timeComponent(time.Now())))
	http.Handle("/404", t1.Handler(notFoundComponent(), t1.WithStatus(http.StatusNotFound)))

	http.ListenAndServe(":8080", nil)
}
```

:::tip
The `t1.WithStatus`, `t1.WithContentType`, and `t1.WithErrorHandler` functions can be passed as parameters to the `t1.Handler` function to control how content is rendered.
:::

The output will always be the date and time that the web server was started up, not the current time.

```
2023-04-26 08:40:03.421358 +0100 BST m=+0.000779501
```

To display the current time, we could update the component to use the `time.Now()` function itself, but this would limit the reusability of the component. It's better when components take parameters for their display values.

:::tip
Good t1 components are idempotent, pure functions - they don't rely on data that is not passed in through parameters. As long as the parameters are the same, they always return the same HTML - they don't rely on any network calls or disk access.
:::

## Displaying dynamic data

Let's update the previous example to display dynamic content.

t1 components implement the `t1.Component` interface, which provides a `Render` method.

The `Render` method can be used within HTTP handlers to write HTML to the `http.ResponseWriter`.

```go title="main.go"
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hello().Render(r.Context(), w)
	})

	http.ListenAndServe(":8080", nil)
}
```

Building on that example, we can implement the Go HTTP handler interface and use the component within our HTTP handler. In this case, displaying the latest date and time, instead of the date and time when the server started up.

```go title="main.go"
package main

import (
	"net/http"
	"time"
)

func NewNowHandler(now func() time.Time) NowHandler {
	return NowHandler{Now: now}
}

type NowHandler struct {
	Now func() time.Time
}

func (nh NowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	timeComponent(nh.Now()).Render(r.Context(), w)
}

func main() {
	http.Handle("/", NewNowHandler(time.Now))

	http.ListenAndServe(":8080", nil)
}
```
