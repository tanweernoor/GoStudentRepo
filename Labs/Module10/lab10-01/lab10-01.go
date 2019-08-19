// lab10-01 Creating Objects
package main

import (
	"fmt"
)

func main() {
	c := makeCircle(1.0)
	fmt.Println(c.area())

	s := makeSquare(2.0)
	fmt.Println(s.area())

	r := makeRectangle(3.0, 4.0)
	fmt.Println(r.diagonal())
}
