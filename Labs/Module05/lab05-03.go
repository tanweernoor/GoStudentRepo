// lab05-03.go
package main

import (
	"fmt"
)

func swap(a []int) {
	a[0], a[1] = a[1], a[0]
	fmt.Println("Swapped ", a, "Address", &a[0])
}
func main() {
	s := []int{1, 2}
	fmt.Println("Before ", s, "Address", &s[0])
	swap(s)
	fmt.Println("After ", s, "Address", &s[0])

}
