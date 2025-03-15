package main

import (
	"goat/app"
	"goat/goat"
	"syscall/js"
)

func main() {
	done := make(chan struct{}, 0)

	js.Global().Get("console").Call("log", "Hello from Go WebAssembly!")

	goat.RenderRoot(app.App())

	<-done
}
