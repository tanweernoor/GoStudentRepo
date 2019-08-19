// lab07-04.go  Map functions to value
// using literals instead of function names
package main

import (
	"fmt"
)

func fmap(funcs []func(int) int, arg int) []int {
	results := make([]int, len(funcs))
	for index, f := range funcs {
		results[index] = f(arg)
	}
	return results
}

func main() {
	// this is the only change that has taken place from the
	// previous problem
	var functions = []func(int) int{func(i int) int { return i * i },
		func(i int) int { return -i }}
	res := fmap(functions, 2)
	fmt.Println(" res=", res)
}
