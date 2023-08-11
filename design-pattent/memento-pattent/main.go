package main

import (
	"example/memento"
	"fmt"
)

func main() {
	editor := memento.Editor{}
	history := memento.History{}

	fmt.Print("editor.Content: ", editor.Content)
	editor.SetContent("a")
	history.Push(editor.CreateState())

	editor.SetContent("b")
	history.Push(editor.CreateState())

	editor.SetContent("c")
	editor.Restore(history.Pop())
	fmt.Println("editor.Content: ", editor.Content)

}
