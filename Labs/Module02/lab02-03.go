// Lab 02-03  Variable Scoping
package main

import "fmt"

var i = 1

func i_print() {
	fmt.Println("pkg level &i = ", &i)
}

func main() {
	i_print()
	i := 2
	fmt.Println("first assigment &i=", &i)
	i, j := 3, 1
	j += 1 // just so the compiler doesn't complain
	fmt.Println("second assigment &i=", &i)
	i_print()
}
