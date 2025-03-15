package app

import (
	"context"
	"goat/app/c"
	"goat/app/template"
	"goat/goat"
	"syscall/js"
)

var todos = []c.Todo{
	{"Make a WASM Frontend Framework in Go", true},
}

func App() string {
	ctx := context.WithValue(context.Background(), "todos", todos)

	goat.MountFunc("toggleTodo", func(args []js.Value) {
		i := args[0].Int()
		todos[i].Completed = !todos[i].Completed
		ctx := context.WithValue(context.Background(), "todos", todos)
		goat.Render("TodoList", goat.HTML(template.TodoList(), ctx))
	})

	goat.MountFunc("addTodo", func(args []js.Value) {
		desc := args[0].String()
		todos = append(todos, c.Todo{
			Description: desc,
		})
		ctx := context.WithValue(context.Background(), "todos", todos)
		goat.Render("TodoList", goat.HTML(template.TodoList(), ctx))
	})

	goat.MountFunc("removeTodo", func(args []js.Value) {
		i := args[0].Int()
		var newTodos []c.Todo
		for j, todo := range todos {
			if i != j {
				newTodos = append(newTodos, todo)
			}
		}
		todos = newTodos
		ctx := context.WithValue(context.Background(), "todos", todos)
		goat.Render("TodoList", goat.HTML(template.TodoList(), ctx))
	})

	return goat.HTML(template.App(), ctx)
}
