// Lab 06  Part 4
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

// Process the string into a list of lowercase words
// with no punctuation
func strproc(instr, punctuation string) []string {
	// take a slice to use as a work area
	str := instr[:]
	// iterate over the ";,." string
	for _, char := range punctuation {
		// convert the rune to a string of length 1
		str = strings.Replace(str, string(char), "", -1)
	}
	// split the string into words and lower case letters
	words := strings.Split(strings.ToLower(str), " ")
	return words
}

func main() {

	wordlist := strproc(myread("churchill.txt"), ".;,")
	fmt.Println(wordlist)
}
