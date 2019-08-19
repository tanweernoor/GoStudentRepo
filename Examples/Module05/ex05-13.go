// Example 05-13 Odd behavior

package main

import "fmt"

func main() {

	a := [...]int{100, 200, 300}
	s := a[:2]
	fmt.Println("Initially a=", a, "s= ", s)
	s = append(s, -1) //result is a[2] == -1
	s[0] = 0          // result a[0] == 0
	fmt.Println("After first op a=", a, "s=", s)
	s = append(s, -2) // this would go in a[3]?
	fmt.Println("After second op a=", a, "s=", s)
	s[0] = 999 // a now remains unchanged
	fmt.Println("After third op a=", a, "s=", s)
}
