// Lab 06  Part 2
package main

import (
	"fmt"
	"io/ioutil"
)

// Read a file into a []byte slice
// and return a string
func myread(filename string) (text string) {
	rawtext, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// convert the rawtext into a string
	text = string(rawtext)
	return
}
func main() {
	s := myread("churchill.txt")
	fmt.Println(s)
}
