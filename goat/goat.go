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

func HTML(j TemplJoint) string {
	var buf bytes.Buffer
	err := j.Render(context.Background(), &buf)
	if err != nil {
		js.Global().Get("console").Call("error", "Error rendering template:", err.Error())
		return ""
	}

	return buf.String()
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

func MountFunc(name string, f func() string) {
	js.Global().Set(name, js.FuncOf(func(this js.Value, args []js.Value) any {
		return f()
	}))
}
