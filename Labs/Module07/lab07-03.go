// lab07-03.go  Map functions to value
// applying a slice of functions to a single argument

package main

import (
	"fmt"
)

func fmap(funcs []func(int) int, arg int) []int {
	results := make([]int, len(funcs))
	// iterate through the slice of functions
	for index, f := range funcs {
		results[index] = f(arg)
	}
	return results
}

func square(i int) int {
	return i * i
}
func neg(i int) int {
	return -i
}
func main() {
	var functions = []func(int) int{square, neg}
	res := fmap(functions, 2)
	fmt.Println(" res=", res)
}
