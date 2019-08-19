// lab04-01.go
package main

import (
	"fmt"
)

func product(x, y int) int {
	return x * y
}

func main() {
	i, j := 2, 4
	fmt.Println(i, "*", j, "=", product(i, j))
}
