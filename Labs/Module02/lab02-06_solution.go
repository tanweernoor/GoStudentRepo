// Lab 02-06  iota
package main

import "fmt"

const (
	_ = iota * iota
	_
	two_sq
	three_sq
	four_sq
)

const (
	_ = iota * iota * iota
	_
	two_cu
	three_cu
	four_cu
)

func main() {
	fmt.Println(two_sq, three_sq, four_sq)
	fmt.Println(two_cu, three_cu, four_cu)
}
