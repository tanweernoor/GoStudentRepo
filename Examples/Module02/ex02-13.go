// Example 02-13 Pointers

package main

import "fmt"

func main() {
	var i int = 187
	var p *int

	p = &i
	fmt.Println("i=", i, " &i or p =", p, " *p =", *p)
	*p = -12
	fmt.Println("i=", i, " &i or p =", p, " *p =", *p)

}
