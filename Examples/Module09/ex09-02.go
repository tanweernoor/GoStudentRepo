// Example 09-02 Method for *int32

package main

import "fmt"

type myint int32

func (x *myint) negate() {
	*x = -*x
}

func main() {

	z := myint(89)
	fmt.Println("Before call", z)
	z.negate()
	fmt.Println("After call", z)
}
