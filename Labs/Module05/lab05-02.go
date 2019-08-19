// lab05-02.go
package main

import (
	"fmt"
)

func swap(a *[2]int) {
	a[0], a[1] = a[1], a[0]
	fmt.Println("Swapped ", a, "Address", &a[0])
}
func main() {
	ar := [2]int{1, 2}
	fmt.Println("Before ", ar, "Address", &ar[0])
	swap(&ar)
	fmt.Println("After ", ar, "Address", &ar[0])

}
