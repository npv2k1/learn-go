package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/npv2k1/todo-cli"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	flag.Parse()

	todos := &todo.Todos{}
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)

	}

	switch {
	case *add:
		todos.Add("Sample task")
		if err := todos.Store(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
	}
}
