// lab10-04.go
package main

import "fmt"

type shaper interface {
	area() float64
}
type quadder interface {
	diagonal() float64
}

func main() {
	shapes := []shaper{makeCircle(1.0), makeSquare(2.0), makeTriangle(2.0, 3.0), makeRectangle(2.0, 3.0)}
	for _, s := range shapes {
		switch t := s.(type) {
		case *circle:
			fmt.Println("Circle circumference is", t.circumference())
		case quadder:
			fmt.Println("Quadder diagonal is", t.diagonal())
		default:
			fmt.Printf("Object of type %T\n", t)
		}

	}
}
