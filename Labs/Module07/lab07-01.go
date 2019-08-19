// lab07-01.go  Map function

package main

import (
	"fmt"
)

func mapf(f func(int) int, arg []int) []int {
	res := make([]int, len(arg))
	for index, value := range arg {
		res[index] = f(value)
	}
	return res
}

func square(i int) int {
	return i * i
}

func neg(i int) int {
	return -i
}

func main() {
	args := []int{1, 2, 3, 4, 5}
	res := mapf(square, args)
	fmt.Println("square: args=", args, " res=", res)
	res = mapf(neg, args)
	fmt.Println("neg: args=", args, " res=", res)
}
