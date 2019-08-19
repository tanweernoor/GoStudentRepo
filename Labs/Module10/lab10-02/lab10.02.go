// lab10-02 Abstract type via Interface
package main

import "fmt"

type shaper interface {
	area() float64
}
type quadder interface {
	diagonal() float64
}

func main() {
	var shape shaper

	shape = makeCircle(1.0)
	fmt.Println("Area of Circle ", shape.area())

	shape = makeSquare(2.0)
	fmt.Println("Area of Square ", shape.area())

	var box quadder

	box = makeSquare(2.0)
	fmt.Println("Diagonal of Square ", box.diagonal())

	box = makeRectangle(2.0, 8.0)
	fmt.Println("Diagonal of Rectangle ", box.diagonal())
}
