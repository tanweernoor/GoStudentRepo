// lab05-07.go
package main

import (
	"fmt"
)

func equiv(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for index, val := range s1 {
		if s2[index] != val {
			return false
		}
	}
	return true
}
func main() {
	a1 := []int{1, 1}
	a2 := []int{1, 1}
	b1 := []int{1, 2, 3}
	b2 := []int{1, 2}
	a3 := a1
	fmt.Println("a1 == a2?", equiv(a1, a2))
	fmt.Println("a3 == a2?", equiv(a3, a2))
	fmt.Println("a1 == b1?", equiv(a1, b1))
	fmt.Println("a1 == b2?", equiv(a1, b2))
}
