// Example 09-04 Support code

package main

import "math"

type point struct{ x, y int }

func (p *point) swap() {
	p.x, p.y = p.y, p.x
}

type circle struct {
	center point
	radius float32
}

func (c *circle) area() float32 {
	return c.radius * c.radius * math.Pi
}
