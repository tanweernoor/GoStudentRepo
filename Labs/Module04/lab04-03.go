// lab04-03.go
package main

import (
	"fmt"
)

func calc(x, y int) (int, int) {
	return x * y, x + y
}

func main() {
	i, j := 2, 4
	p, s := calc(i, j)
	fmt.Println(i, "*", j, "=", p)
	fmt.Println(i, "+", j, "=", s)
}
