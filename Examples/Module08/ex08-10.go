// Example 08-10 Struct factory

package main

import "fmt"

type point struct{ x, y int }

func makePoint(x, y int) *point {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return &point{x, y}
}
func main() {
	p1 := makePoint(1, 1)
	fmt.Println(*p1)
	p1 = makePoint(-4, -9)
	fmt.Println(*p1)
}
