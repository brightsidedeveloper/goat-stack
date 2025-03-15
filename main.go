package main

import (
	"goat/app"
	"goat/goat"
)

func main() {
	done := make(chan struct{}, 0)

	goat.Log("Hi from WASM!")

	goat.RenderRoot(app.App())

	<-done
}
