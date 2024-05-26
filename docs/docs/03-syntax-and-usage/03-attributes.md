# Attributes

## Constant attributes

t1 elements can have HTML attributes that use the double quote character `"`.

```t1
t1 component() {
  <p data-testid="paragraph">Text</p>
}
```

```html title="Output"
<p data-testid="paragraph">Text</p>
```

## Boolean attributes

Boolean attributes (see https://html.spec.whatwg.org/multipage/common-microsyntaxes.html#boolean-attributes) where the presence of an attribute name without a value means true, and the attribute name not being present means false are supported.

```t1
t1 component() {
  <hr noshade/>
}
```

```html title="Output"
<hr noshade />
```

:::note
t1 is aware that `<hr/>` is a void element, and renders `<hr>` instead.
:::

To set boolean attributes using variables or template parameters, a question mark after the attribute name is used to denote that the attribute is boolean.

```t1
t1 component() {
  <hr noshade?={ false } />
}
```

```html title="Output"
<hr />
```

## Conditional attributes

Use an `if` statement within a t1 element to optionally add attributes to elements.

```t1
t1 component() {
  <hr style="padding: 10px"
    if true {
      class="itIsTrue"
    }
  />
}
```

```html title="Output"
<hr style="padding: 10px" class="itIsTrue" />
```

## Spread attributes

Use the `{ attrMap... }` syntax in the open tag of an element to append a dynamic map of attributes to the element's attributes.

It's possible to spread any variable of type `t1.Attributes`. `t1.Attributes` is a `map[string]any` type definition.

- If the value is a `string`, the attribute is added with the string value, e.g. `<div name="value">`.
- If the value is a `bool`, the attribute is added as a boolean attribute if the value is true, e.g. `<div name>`.
- If the value is a `t1.KeyValue[string, bool]`, the attribute is added if the boolean is true, e.g. `<div name="value">`.
- If the value is a `t1.KeyValue[bool, bool]`, the attribute is added if both boolean values are true, as `<div name>`.

```t1
t1 component(shouldBeUsed bool, attrs t1.Attributes) {
  <p { attrs... }></p>
  <hr
    if shouldBeUsed {
      { attrs... }
    }
  />
}

t1 usage() {
  @component(false, t1.Attributes{"data-testid": "paragraph"})
}
```

```html title="Output"
<p data-testid="paragraph">Text</p>
<hr />
```

## URL attributes

The `<a>` element's `href` attribute is treated differently. t1 expects you to provide a `t1.SafeURL` instead of a `string`.

Typically, you would do this by using the `t1.URL` function.

The `t1.URL` function sanitizes input URLs and checks that the protocol is `http`/`https`/`mailto` rather than `javascript` or another unexpected protocol.

```t1
t1 component(p Person) {
  <a href={ t1.URL(p.URL) }>{ strings.ToUpper(p.Name) }</a>
}
```

The `t1.URL` function only supports standard HTML elements and attributes (`<a href=""` and `<form action=""`).

For use on non-standard HTML elements (e.g. HTMX's `hx-*` attributes), convert the `t1.URL` to a `string` after sanitization.

```t1
t1 component(contact model.Contact) {
  <div hx-get={ string(t1.URL(fmt.Sprintf("/contacts/%s/email", contact.ID)))}>
    { contact.Name }
  </div>
}
```

:::caution
If you need to bypass this sanitization, you can use `t1.SafeURL(myURL)` to mark that your string is safe to use.

This may introduce security vulnerabilities to your program.
:::

## JavaScript attributes

`onClick` and other `on*` handlers have special behaviour, they expect a reference to a `script` template.

:::info
This ensures that any client-side JavaScript that is required for a component to function is only emitted once, that script name collisions are not possible, and that script input parameters are properly sanitized.
:::

```t1
script withParameters(a string, b string, c int) {
	console.log(a, b, c);
}

script withoutParameters() {
	alert("hello");
}

t1 Button(text string) {
	<button onClick={ withParameters("test", text, 123) } onMouseover={ withoutParameters() } type="button">{ text }</button>
}
```

```html title="Output"
<script type="text/javascript">
 function __t1_withParameters_1056(a, b, c){console.log(a, b, c);}function __t1_withoutParameters_6bbf(){alert("hello");}
</script>
<button onclick="__t1_withParameters_1056("test","Say hello",123)" onmouseover="__t1_withoutParameters_6bbf()" type="button">
 Say hello
</button>
```

## CSS attributes

CSS handling is discussed in detail in [CSS style management](css-style-management).
