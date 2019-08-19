// Example 08-02 Struct initialization

package main

import "fmt"

type point struct{ x, y int }
type employee struct {
	fname, lname string
	id           int
	job          string
	salary       int
}

func main() {

	var p = point{2, 3}
	fmt.Println("p =", p)

	anil := employee{"Anil", "Patel", 9891, "Developer", 100000}
	greta := employee{id: 8897, fname: "Greta", lname: "Smith"}

	fmt.Println("anil =", anil)
	fmt.Println("grea =", greta)
}
