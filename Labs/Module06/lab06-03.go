// Lab 06  Part 3
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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

func strproc(instr, punctuation string) string {
	// take a slice to use as a work area
	str := instr[:]
	// iterate over the ";,." string
	for _, char := range punctuation {
		// convert the rune to a string of length 1
		str = strings.Replace(str, string(char), "", -1)
	}
	return str
}

func main() {
	s := myread("churchill.txt")
	s = strproc(s, ".;,")
	fmt.Println(s)
}
