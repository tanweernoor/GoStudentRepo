// Example 08-09 Pointers as fields

package main

import "fmt"

type employee struct {
	fname, lname string
	id           int
	job          string
	salary       int
	boss         *employee
}

func main() {
	greta := employee{fname: "Greta", lname: "Smith"}
	anil := employee{fname: "Greta", lname: "Smith", boss: &greta}
	fmt.Println("Anil's boss is", anil.boss.fname)
}
