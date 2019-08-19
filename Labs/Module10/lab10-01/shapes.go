// shapes.go
package main

import "math"

//type definitions

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type triangle struct {
	base, height float64
}

type rectangle struct {
	height, width float64
}

// Constructors

func makeCircle(r float64) *circle {
	return &circle{r}
}
func makeSquare(s float64) *square {
	return &square{s}
}
func makeRectangle(h, w float64) *rectangle {
	return &rectangle{h, w}
}
func makeTriangle(b, h float64) *triangle {
	return &triangle{b, h}
}

// Area and related functions
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c *circle) circumference() float64 {
	return float64(2) * math.Pi * c.radius
}
func (s *square) area() float64 {
	return s.side * s.side
}
func (r *rectangle) area() float64 {
	return r.height * r.width
}
func (r *rectangle) diagonal() float64 {
	return math.Sqrt((r.width * r.width) + (r.height * r.height))
}
func (r *square) diagonal() float64 {
	return math.Sqrt((r.side * r.side))
}
func (t *triangle) area() float64 {
	return t.height * t.base * float64(.5)
}
