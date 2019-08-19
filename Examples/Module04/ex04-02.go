// Example 04-02 Named return value

package main

import "fmt"

func square(x int) (result int) {
	result = x * x
	return
}
func main() {
	x := 4
	retval := square(x)
	fmt.Println("x = ", x, "square = ", retval)
}
