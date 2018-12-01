package main

import (
	"fmt"
)

const state = "Maryland"

var name string

func main() {
	name = "Sameet"
	from := `Baltimore`
	var n int

	fmt.Printf("Hello, fellow %s Gophers!\n", state)
	fmt.Printf("my name is %s and I'm from %s.\n", name, from)
	fmt.Printf("By the time %d o'clock comes around, we'll know how to code in Go!\n", n)
	fmt.Println("Let's get started!")
}
