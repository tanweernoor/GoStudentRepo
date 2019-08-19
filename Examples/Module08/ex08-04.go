// Example 08-04 Field access

package main

import "fmt"

type point struct{ x, y int }

func main() {
	p := point{2, 3}
	pp := &point{4, 5}
	fmt.Println("x coord of p", p.x)
	fmt.Println("x coord of pp", pp.x)
	pp.y = -1
	p.y = 0
	fmt.Println("p=", p, " pp=", *pp)
}
