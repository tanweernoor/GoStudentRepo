// Lab 02-07  Pointers
package main

import "fmt"

var i, j = 1, 2
var pi, pj = &i, &j

func list() {
	fmt.Println("i=", i, "j=", j)
}

func main() {
	list()
	i, j := 100, 101
	i, j = j, i
	*pi, *pj = *pj, *pi
	list()
}
