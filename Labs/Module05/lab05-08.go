// lab05-08.go
package main

import (
	"fmt"
)

func main() {
	src := []int{1, 2, 3, 4}
	dest := []int{8, 9, 10, 11, 12, 13}
	copy(dest, src)
	fmt.Println(src, dest)
}
