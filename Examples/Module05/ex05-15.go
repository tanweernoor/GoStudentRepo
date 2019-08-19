// Example 05-15 Slices as parameters

package main

import "fmt"

func f(p []int) {
	p[0] = -1
}

func main() {

	s := []int{1, 2, 3, 4}
	fmt.Println("Before call", s)
	f(s)
	fmt.Println("After call", s)
}
