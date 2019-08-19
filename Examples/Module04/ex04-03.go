// Example 04-03 Multiple parameters

package main

import "fmt"

func divides(x, y int) (div bool) {
	div = (y % x) == 0
	return true
}
func main() {
	x, y := 2, 21
	fmt.Println(x, "|", y, " is ", divides(x, y))
}
