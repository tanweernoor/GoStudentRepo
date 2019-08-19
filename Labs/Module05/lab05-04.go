// lab05-04.go
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
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Before ", s)
	reverse(s)
	fmt.Println("After ", s)

}
