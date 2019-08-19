// lab04-04.go
package main

import (
	"fmt"
)

func calc(x, y int) (prod, sum int) {
	prod, sum = x*y, x+y
	return
}

func main() {
	i, j := 2, 4
	p, s := calc(i, j)
	fmt.Println(i, "*", j, "=", p)
	fmt.Println(i, "+", j, "=", s)
}
