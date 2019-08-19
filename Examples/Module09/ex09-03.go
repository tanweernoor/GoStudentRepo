// Example 09-03 Method overloading

package main

import "fmt"

type myint int32
type myfloat float64

func (x *myint) negate() {
	*x = -(*x)
}
func (x *myfloat) negate() {
	*x = 0.0
}

func main() {

	i := myint(89)
	f := myfloat(189.9)
	i.negate()
	f.negate()
	fmt.Println(i, f)
}
