// lab10-03.go
package main

import "fmt"

type line struct{ length float64 }

type shaper interface {
	area() float64
}
type quadder interface {
	diagonal() float64
}

func main() {
	shapes := []shaper{makeCircle(1.0), makeSquare(2.0), makeTriangle(2.0, 3.0), makeRectangle(2.0, 3.0)}
	for _, s := range shapes {
		fmt.Println(s.area())
	}
}
