// Example 05-14 Re-sizing Slices

package main

import "fmt"

func main() {

	s1 := make([]int, 1, 3)
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
		fmt.Println("s1=", s1, "len=", len(s1), "Cap=", cap(s1))
	}

}
