# Rendering raw HTML

To render HTML that has come from a trusted source, bypassing all HTML escaping and security mechanisms that t1 includes, use the `t1.Raw` function.

:::info
Only include HTML that comes from a trusted source.
:::

:::warning
Use of this function may introduce security vulnerabilities to your program.
:::

```t1 title="component.t1"
t1 Example() {
	<!DOCTYPE html>
	<html>
		<body>
			@t1.Raw("<div>Hello, World!</div>")
		</body>
	</html>
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
