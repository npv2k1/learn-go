package main

import "fmt"

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func main() {

	rect1 := Rect{10, 20}
	fmt.Println(rect1.area())

}
