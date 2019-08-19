// Example 08-03 Using new

package main

import "fmt"

type point struct{ x, y int }

func main() {

	var pp1 *point = new(point) // pp1 is a pointer
	fmt.Println("pp1 =", pp1, "*pp=", *pp1)

	p := point{3, 4} // p is a variable
	pp := &p         // pp is a pointer
	fmt.Println("pp =", pp, "*pp=", *pp, "p=", p)

	pp2 := &point{5, 6} //pp2 is a pointer
	fmt.Println("pp2 =", pp2, "*pp2=", *pp2)

}
