// Example 04-04 Multiple return values

package main

import "fmt"

func divides(x, y int) (int, int) {
	return (y / x), (y % x)
}
func main() {
	x, y := 3, 23
	quot, rem := divides(x, y)
	fmt.Println(x, "/", y, " is ", quot, "R", rem)
}
