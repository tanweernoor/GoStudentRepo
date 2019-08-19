// Example 03-01 Operators

package main

import "fmt"

func main() {

	a, b, c := uint8(1), uint16(1), int8(1)
	//fmt.Println("1. a + b =", a+b)
	//fmt.Println("2. a + c =", a+c)
	fmt.Println("3. a + uint8(b)=", a+uint8(b))
	fmt.Println("4. uint8(a) + c=", int8(a)+c)
}
