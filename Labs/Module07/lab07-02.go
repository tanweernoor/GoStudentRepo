// Lab 07-02  Apply mapf_r in place
// The in place version of mapf

package main

import (
	"fmt"
)

func mapf_r(f func(int) int, arg []int) {
	for index, value := range arg {
		// replace each entry i with f(i)
		arg[index] = f(value)
	}
	return
}

func square(i int) int {
	return i * i
}

func neg(i int) int {
	return -i
}
func main() {
	args1 := []int{1, 2, 3, 4, 5}
	args2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Before call", args1, args2)
	mapf_r(square, args1)
	mapf_r(neg, args2)
	fmt.Println("After call", args1, args2)
}
