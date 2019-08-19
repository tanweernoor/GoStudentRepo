// lab05-09.go
package main

import (
	"fmt"
)

func mkcopy(orig []int) (cp []int) {
	cp = make([]int, len(orig))
	for i, val := range orig {
		cp[i] = val
	}
	return
}

func main() {
	src := []int{1, 2, 3, 4}
	dest := mkcopy(src)
	fmt.Println(src, dest)
	fmt.Println(&src[0], &dest[0]) //confim different
}
