// Example 04-15 Runtime panic

package main

import "fmt"

func rec() {
	r := recover()
	if r != nil {
		fmt.Println("recovered value = ", r)
	}
}
func main() {
	x, y := 1, 0
	defer rec()
	fmt.Println(x / y)
}
