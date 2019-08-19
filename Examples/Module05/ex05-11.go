// Example 05-11 Slices

package main

import "fmt"

func main() {

	var a = [...]int{0, 1, 2, 3, 4, 5}

	s0 := a[1:4]
	s1 := s0[1:3]
	s2 := s1[0:4] // ???

	fmt.Println("s0 length=", len(s0), " s1 length=", len(s1),
		" s2 length=", len(s2))
	fmt.Println("s0=", s0, " s1=", s1, " s2=", s2)
}
