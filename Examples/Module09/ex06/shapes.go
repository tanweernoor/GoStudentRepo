// Example 09-06 Support code

package main

import "math"

type point struct{ x, y int }

func (p *point) swap() {
	p.x, p.y = p.y, p.x
}

type circle struct {
	point
	radius float32
}

func (c *circle) swap() {
}

func (c *circle) area() float32 {
	return c.radius * c.radius * math.Pi
}
