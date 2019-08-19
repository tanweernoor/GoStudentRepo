// Example 08-05 Comparisons

package main

import "fmt"

type point struct{ x, y int }

func main() {
	p1 := point{2, 3}
	p2 := point{4, 5}
	p3 := point{2, 3}
	fmt.Println("p1 == p2? ", p1 == p2)
	fmt.Println("p1 == p3? ", p1 == p3)
	pp1 := &point{1, 1}
	pp2 := &point{1, 1}
	fmt.Println("pp1 == pp2? ", pp1 == pp2)
	fmt.Println("*pp1 == *pp? ", *pp1 == *pp2)

}
