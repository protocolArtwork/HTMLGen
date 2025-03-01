package main

import (
	"fmt"
	. "ssg"
)

func main() {

	h := HTMLComponent{
		Name:     "div",
		Props:    HTMLProps{"class": "container", "id": "myDiv"},
		Children: []Viewable{HTMLString("Hello, world!")},
	}

	fmt.Print(h.Render())
}
