// lab07-06.go  dispatch function
// executing anonymous functions
package main

import (
	"fmt"
)

var status int = 0

func oops() {
	switch {
	case status < 1:
		func() { status++; fmt.Println("It's all cool") }()
	case status < 3:
		func() { status++; fmt.Println("Things are heating up") }()
	default:
		func() { status++; panic("Too hot for me") }()
	}
}
func main() {
	for true {
		oops()
	}
}
