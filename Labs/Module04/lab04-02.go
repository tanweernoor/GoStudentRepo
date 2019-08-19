// lab04-02.go
package main

import (
	"fmt"
)

func product(x, y int) (p int) {
	p = x * y
	return
}

func main() {
	i, j := 2, 4
	fmt.Println(i, "*", j, "=", product(i, j))
}
