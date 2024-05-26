# Comments

# HTML comments

Inside t1 statements, use HTML comments.

```t1 title="template.t1"
t1 template() {
	<!-- Single line -->
	<!--
		Single or multiline.
	-->
}
```

Comments are rendered to the template output.

```html title="Output"
<!-- Single line -->
<!--
	Single or multiline.
-->
```

As per HTML, nested comments are not supported.

# Go comments

Outside of t1 statements, use Go comments.

```t1
package main

// Use standard Go comments outside t1 statements.
var greeting = "Hello!"

t1 hello(name string) {
	<p>{greeting} { name }</p>
}
```
