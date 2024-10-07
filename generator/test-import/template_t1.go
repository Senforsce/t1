// Code generated by t1 - DO NOT EDIT.

package testimport

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/senforsce/tndr"
import "context"
import "io"
import "bytes"

func listItem() t1.Component {
	return t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
		t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
		if !t1_7745c5c3_IsBuffer {
			t1_7745c5c3_Buffer = t1.GetBuffer()
			defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
		}
		ctx = t1.InitializeContext(ctx)
		t1_7745c5c3_Var1 := t1.GetChildren(ctx)
		if t1_7745c5c3_Var1 == nil {
			t1_7745c5c3_Var1 = t1.NopComponent
		}
		ctx = t1.ClearChildren(ctx)
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<li>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		t1_7745c5c3_Err = t1_7745c5c3_Var1.Render(ctx, t1_7745c5c3_Buffer)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</li>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}

func list() t1.Component {
	return t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
		t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
		if !t1_7745c5c3_IsBuffer {
			t1_7745c5c3_Buffer = t1.GetBuffer()
			defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
		}
		ctx = t1.InitializeContext(ctx)
		t1_7745c5c3_Var2 := t1.GetChildren(ctx)
		if t1_7745c5c3_Var2 == nil {
			t1_7745c5c3_Var2 = t1.NopComponent
		}
		ctx = t1.ClearChildren(ctx)
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<ul>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		t1_7745c5c3_Err = t1_7745c5c3_Var2.Render(ctx, t1_7745c5c3_Buffer)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</ul>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}

func main() t1.Component {
	return t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
		t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
		if !t1_7745c5c3_IsBuffer {
			t1_7745c5c3_Buffer = t1.GetBuffer()
			defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
		}
		ctx = t1.InitializeContext(ctx)
		t1_7745c5c3_Var3 := t1.GetChildren(ctx)
		if t1_7745c5c3_Var3 == nil {
			t1_7745c5c3_Var3 = t1.NopComponent
		}
		ctx = t1.ClearChildren(ctx)
		t1_7745c5c3_Var4 := t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
			t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
			if !t1_7745c5c3_IsBuffer {
				t1_7745c5c3_Buffer = t1.GetBuffer()
				defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
			}
			t1_7745c5c3_Var5 := t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
				t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
				if !t1_7745c5c3_IsBuffer {
					t1_7745c5c3_Buffer = t1.GetBuffer()
					defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
				}
				_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<u>Item 1</u>")
				if t1_7745c5c3_Err != nil {
					return t1_7745c5c3_Err
				}
				if !t1_7745c5c3_IsBuffer {
					_, t1_7745c5c3_Err = io.Copy(t1_7745c5c3_W, t1_7745c5c3_Buffer)
				}
				return t1_7745c5c3_Err
			})
			t1_7745c5c3_Err = listItem().Render(t1.WithChildren(ctx, t1_7745c5c3_Var5), t1_7745c5c3_Buffer)
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(" ")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
			t1_7745c5c3_Var6 := t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
				t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
				if !t1_7745c5c3_IsBuffer {
					t1_7745c5c3_Buffer = t1.GetBuffer()
					defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
				}
				_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<u>Item 2</u>")
				if t1_7745c5c3_Err != nil {
					return t1_7745c5c3_Err
				}
				if !t1_7745c5c3_IsBuffer {
					_, t1_7745c5c3_Err = io.Copy(t1_7745c5c3_W, t1_7745c5c3_Buffer)
				}
				return t1_7745c5c3_Err
			})
			t1_7745c5c3_Err = listItem().Render(t1.WithChildren(ctx, t1_7745c5c3_Var6), t1_7745c5c3_Buffer)
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(" ")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
			t1_7745c5c3_Var7 := t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
				t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
				if !t1_7745c5c3_IsBuffer {
					t1_7745c5c3_Buffer = t1.GetBuffer()
					defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
				}
				_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<u>Item 3</u>")
				if t1_7745c5c3_Err != nil {
					return t1_7745c5c3_Err
				}
				if !t1_7745c5c3_IsBuffer {
					_, t1_7745c5c3_Err = io.Copy(t1_7745c5c3_W, t1_7745c5c3_Buffer)
				}
				return t1_7745c5c3_Err
			})
			t1_7745c5c3_Err = listItem().Render(t1.WithChildren(ctx, t1_7745c5c3_Var7), t1_7745c5c3_Buffer)
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
			if !t1_7745c5c3_IsBuffer {
				_, t1_7745c5c3_Err = io.Copy(t1_7745c5c3_W, t1_7745c5c3_Buffer)
			}
			return t1_7745c5c3_Err
		})
		t1_7745c5c3_Err = list().Render(t1.WithChildren(ctx, t1_7745c5c3_Var4), t1_7745c5c3_Buffer)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}