// Example 05-01 Basic array syntax

package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println("a =", a)
	a[0] = 1
	a[1] = a[0] + 1
	fmt.Println("a =", a)
}
