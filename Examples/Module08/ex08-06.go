// Example 08-06 Struct pointers

package main

import "fmt"

type point struct{ x, y int }

func swap1(p point) {
	p.x, p.y = p.y, p.x
	fmt.Println("After executing swap1 p=", p)
}
func swap2(p *point) {
	p.x, p.y = p.y, p.x
}

func main() {
	a := point{1, 2}
	fmt.Println("Original a =", a)
	swap1(a)
	fmt.Println("After swap1 a =", a)
	swap2(&a)
	fmt.Println("After swap2 a =", a)

}
