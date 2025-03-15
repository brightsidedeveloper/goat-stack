package app

import (
	"goat/app/template"
	"goat/goat"
)

var count int

func App() string {
	return goat.HTML(template.App())
}
