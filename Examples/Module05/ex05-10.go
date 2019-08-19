// Example 05-10 Basic Slice

package main

import "fmt"

func main() {

	var a = [6]int{0, 1, 2, 3, 4, 5}
	s := a[2:] // s is a slice of the array a
	fmt.Println("s= ", s)

	a[4] = -20 // changing underlying array
	fmt.Println("s= ", s)

	s[0] = 999 // change the array via the slice
	fmt.Println("a= ", a)
}
