package goat

import (
	"bytes"
	"context"
	"io"
	"syscall/js"
)

type TemplJoint interface {
	Render(context.Context, io.Writer) error
}

func HTML(j TemplJoint, c context.Context) string {
	var buf bytes.Buffer
	err := j.Render(c, &buf)
	if err != nil {
		js.Global().Get("console").Call("error", "Error rendering template:", err.Error())
		return ""
	}

	return buf.String()
}

func Log(args ...any) {
	js.Global().Get("console").Call("log", args...)
}

func RenderRoot(html string) {
	doc := js.Global().Get("document")
	output := doc.Call("getElementById", "root")
	output.Set("innerHTML", html)
}

func Render(id string, html string) {
	doc := js.Global().Get("document")
	output := doc.Call("getElementById", id)
	output.Set("outerHTML", html)
}

func MountFunc(name string, f func(args []js.Value)) {
	js.Global().Set(name, js.FuncOf(func(this js.Value, args []js.Value) any {
		f(args)
		return nil
	}))
}
