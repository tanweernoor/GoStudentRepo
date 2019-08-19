// Example 08-01 Basic struct definiton

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

	var anil employee
	var p point
	fmt.Println("anil=", anil, "p=", p)

}
