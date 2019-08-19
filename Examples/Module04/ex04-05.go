// Example 04-05 Multiple named return values

package main

import "fmt"

func divides(x, y int) (q int, r int) {
	q = y / x
	r = y % x
	return
}
func main() {
	x, y := 3, 23
	quot, rem := divides(x, y)
	fmt.Println(x, "/", y, " is ", quot, "R", rem)
}
