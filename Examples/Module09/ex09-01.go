// Example 09-01 Method for int32

package main

import "fmt"

type myint int32

func (x myint) negate() {
	x = -x
}

func main() {

	z := myint(89)
	fmt.Println(z)
	z.negate()
	fmt.Println(z)
}
