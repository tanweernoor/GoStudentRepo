// lab05-05.go
package main

import (
	"fmt"
)

func reverse(x []int) {
	for start, end := 0, len(x)-1; start < end; start, end = start+1, end-1 {
		x[start], x[end] = x[end], x[start]
	}
}
func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	s := a[:]
	fmt.Println("Before ", a)
	reverse(s)
	fmt.Println("After ", a)

}
