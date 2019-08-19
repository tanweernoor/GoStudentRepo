// lab07-05-a.go  dispatch function
// function as a return value
// cleaned up
package main

import (
	"fmt"
)

func normal() {
	status++
	fmt.Println("It's all cool")
}
func warning() {
	status++
	fmt.Println("Things are heating up")
}

func trouble() {
	status++
	panic("Too hot for me")
}

var status int = 0

func oops() (f func()) {
	switch {
	case status < 1:
		return normal
	case status < 3:
		return warning
	default:
		return trouble
	}

}
func main() {
	for true {
		oops()()
	}
}
