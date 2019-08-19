// lab07-05.go  dispatch function
// function as a return value
package main

import (
	"fmt"
)

var status int = 0

func oops() (f func()) {
	switch {
	case status < 1:
		f = func() { status++; fmt.Println("It's all cool") }
	case status < 3:
		f = func() { status++; fmt.Println("Things are heating up") }
	default:
		f = func() { status++; panic("Too hot for me") }
	}

	return
}
func main() {
	// creates an instance of the returned function
	// with a copy of status set to 0
	closure := oops()
	for i := 0; i < 5; i++ {
		closure()
	}
}
