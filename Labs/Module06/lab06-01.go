// Lab 06  Part 1
package main

import (
	"fmt"
	"io/ioutil"
)

// Read a file into a []byte slice
func myread(filename string) (text string) {
	rawtext, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("rawtext=", rawtext)
	return
}
func main() {
	myread("churchill.txt")

}
