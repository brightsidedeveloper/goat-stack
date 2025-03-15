package app

import (
	"goat/app/c"
	"goat/app/template"
	"goat/goat"
	"syscall/js"
)

var todos = []c.Todo{
	{"Make a WASM Frontend Framework in Go", true},
}

type TodoList struct {
	key   string
	todos []c.Todo
}

func (t *TodoList) toggle(i int) {
	t.todos[i].Completed = !t.todos[i].Completed
}

func (t *TodoList) addTodo(desc string) {
	t.todos = append(t.todos, c.Todo{
		Description: desc,
	})
}

func (t *TodoList) removeTodo(i int) {
	var newTodos []c.Todo
	for j, todo := range todos {
		if i != j {
			newTodos = append(newTodos, todo)
		}
	}
	todos = newTodos
}

func (t *TodoList) render() string {
	return goat.HTML(template.TodoList(t.todos))
}

func (t *TodoList) rerender() {
	goat.Render(t.key, t.render())
}

func (t *TodoList) callbacks() {
	goat.MountFunc("toggleTodo", func(args []js.Value) {
		i := args[0].Int()
		t.toggle(i)
		t.rerender()
	})

	goat.MountFunc("addTodo", func(args []js.Value) {
		desc := args[0].String()
		t.addTodo(desc)
		t.rerender()
	})

	goat.MountFunc("removeTodo", func(args []js.Value) {
		i := args[0].Int()
		t.removeTodo(i)
		t.rerender()
	})
}

func App() string {

	todoList := &TodoList{
		key:   "TodoList",
		todos: todos,
	}

	todoList.callbacks()

	return todoList.render()
}
