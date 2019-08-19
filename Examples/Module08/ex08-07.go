// Example 08-07 Embedded structs

package main

import "fmt"

type point struct{ x, y int }

type circle struct {
	center point
	radius float32
}

func main() {
	c := circle{point{50, 32}, 13.0}
	fmt.Println("c=", c)
	c.center.x = 0
	fmt.Println("c=", c)
}
