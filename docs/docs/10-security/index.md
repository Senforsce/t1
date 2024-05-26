# Security

## Injection attacks

t1 is designed to prevent user-provided data from being used to inject vulnerabilities.

`<script>` and `<style>` tags could allow user data to inject vulnerabilities, so variables are not permitted in these sections.

```html
t1 Example() {
<script type="text/javascript">
  function showAlert() {
    alert("hello");
  }
</script>
<style type="text/css">
  /* Only CSS is allowed */
</style>
}
```

`onClick` attributes, and other `on*` attributes are used to execute JavaScript. To prevent user data from being unescaped, `on*` attributes accept a `t1.ComponentScript`.

```html
script onClickHandler(msg string) { alert(msg); } t1 Example(msg string) {
<div onClick="{" onClickHandler(msg) }>{ "will be HTML encoded using t1.Escape" }</div>
}
```

Style attributes cannot be expressions, only constants, to avoid escaping vulnerabilities. t1 style templates (`css className()`) should be used instead.

```html
t1 Example() {
  <div style={ "will throw an error" }></div>
}
```

Class names are sanitized by default. A failed class name is replaced by `--t1-css-class-safe-name`. The sanitization can be bypassed using the `t1.SafeClass` function, but the result is still subject to escaping.

```html
t1 Example() {
  <div class={ "unsafe</style&gt;-will-sanitized", t1.SafeClass("&sanitization bypassed") }></div>
}
```

Rendered output:

```html
<div class="--t1-css-class-safe-name &amp;sanitization bypassed"></div>
```

```html
t1 Example() {
<div>Node text is not modified at all.</div>
<div>{ "will be escaped using t1.EscapeString" }</div>
}
```

`href` attributes must be a `t1.SafeURL` and are sanitized to remove JavaScript URLs unless bypassed.

```html
t1 Example() {
  <a href="http://constants.example.com/are/not/sanitized">Text</a>
  <a href={ t1.URL("will be sanitized by t1.URL to remove potential attacks") }</a>
  <a href={ t1.SafeURL("will not be sanitized by t1.URL") }</a>
}
```

Within css blocks, property names, and constant CSS property values are not sanitized or escaped.

```css
css className() {
  background-color: #ffffff;
}
```

CSS property values based on expressions are passed through `t1.SanitizeCSS` to replace potentially unsafe values with placeholders.

```css
css className() {
	color: { red };
}
```

## Code signing

Binaries are created by https://github.com/a-h and signed with https://adrianhesketh.com/a-h.gpg
