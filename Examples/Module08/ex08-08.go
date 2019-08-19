// Example 08-08 Anonymous fields

package main

import "fmt"

type point struct{ x, y int }

type circle struct {
	point
	bool
	radius float32
}

func main() {
	c := circle{point{50, 32}, false, 3.0}
	fmt.Println("c=", c)
	c.x = 0
	c.bool = true
	fmt.Println("c=", c)
}
