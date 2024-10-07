// Code generated by t1 - DO NOT EDIT.

package testwhitespacearoundgokeywords

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/senforsce/tndr"
import "context"
import "io"
import "bytes"

import "fmt"

func WhitespaceIsConsistentInIf(firstIf, secondIf bool) t1.Component {
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
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>Start</button> ")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if firstIf {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>If</button> ")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
		} else if secondIf {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>ElseIf</button> ")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
		} else {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>Else</button> ")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>End</button>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}

const WhitespaceIsConsistentInTrueIfExpected = `<button>Start</button> <button>If</button> <button>End</button>`
const WhitespaceIsConsistentInTrueElseIfExpected = `<button>Start</button> <button>ElseIf</button> <button>End</button>`
const WhitespaceIsConsistentInTrueElseExpected = `<button>Start</button> <button>Else</button> <button>End</button>`

func WhitespaceIsConsistentInFalseIf() t1.Component {
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
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>Start</button> ")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if false {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>Will Not Render</button> ")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>End</button>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}

const WhitespaceIsConsistentInFalseIfExpected = `<button>Start</button> <button>End</button>`

func WhitespaceIsConsistentInSwitch(i int) t1.Component {
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
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>Start</button> ")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		switch i {
		case 1:
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>1</button> ")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
		default:
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>default</button> ")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>End</button>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}

const WhitespaceIsConsistentInOneSwitchExpected = `<button>Start</button> <button>1</button> <button>End</button>`
const WhitespaceIsConsistentInDefaultSwitchExpected = `<button>Start</button> <button>default</button> <button>End</button>`

func WhitespaceIsConsistentInSwitchNoDefault() t1.Component {
	return t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
		t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
		if !t1_7745c5c3_IsBuffer {
			t1_7745c5c3_Buffer = t1.GetBuffer()
			defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
		}
		ctx = t1.InitializeContext(ctx)
		t1_7745c5c3_Var4 := t1.GetChildren(ctx)
		if t1_7745c5c3_Var4 == nil {
			t1_7745c5c3_Var4 = t1.NopComponent
		}
		ctx = t1.ClearChildren(ctx)
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>Start</button> ")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		switch false {
		case true:
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>Will Not Render</button> ")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>End</button>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}

const WhitespaceIsConsistentInSwitchNoDefaultExpected = `<button>Start</button> <button>End</button>`

func WhitespaceIsConsistentInFor(i int) t1.Component {
	return t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
		t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
		if !t1_7745c5c3_IsBuffer {
			t1_7745c5c3_Buffer = t1.GetBuffer()
			defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
		}
		ctx = t1.InitializeContext(ctx)
		t1_7745c5c3_Var5 := t1.GetChildren(ctx)
		if t1_7745c5c3_Var5 == nil {
			t1_7745c5c3_Var5 = t1.NopComponent
		}
		ctx = t1.ClearChildren(ctx)
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>Start</button> ")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		for j := 0; j < i; j++ {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
			var t1_7745c5c3_Var6 string
			t1_7745c5c3_Var6, t1_7745c5c3_Err = t1.JoinStringErrs(fmt.Sprint(j))
			if t1_7745c5c3_Err != nil {
				return t1.Error{Err: t1_7745c5c3_Err, FileName: `generator/test-whitespace-around-go-keywords/template.t1`, Line: 58, Col: 25}
			}
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1_7745c5c3_Var6))
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</button> ")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<button>End</button>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}

const WhitespaceIsConsistentInForZeroExpected = `<button>Start</button> <button>End</button>`
const WhitespaceIsConsistentInForOneExpected = `<button>Start</button> <button>0</button> <button>End</button>`
const WhitespaceIsConsistentInForThreeExpected = `<button>Start</button> <button>0</button> <button>1</button> <button>2</button> <button>End</button>`