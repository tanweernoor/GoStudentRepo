// lab05-06.go
package main

import (
	"fmt"
)

func main() {
	a1 := [2]int{1, 2}
	a2 := [2]int{1, 2}
	fmt.Println("a1 == a2?", a1 == a2)

	original := []int{1, 2, 3, 4, 5, 6, 7}
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := original

	fmt.Println("original and s1 ", &original[0] == &s1[0])
	fmt.Println("original and s2 ", &original[0] == &s2[0])

}
