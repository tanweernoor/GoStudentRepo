// Example 04-01 Basic function sytnax

package main

import "fmt"

func square(x int) int {
	y := x * x
	return y
}
func main() {
	x := 4
	retval := square(x)
	fmt.Println("x = ", x, "square = ", retval)
}
