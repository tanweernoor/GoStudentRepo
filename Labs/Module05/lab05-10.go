// lab05-10.go
package main

import (
	"fmt"
)

var text = [][][]string{
	{
		{"This", "is", "line", "one", "para", "one"},
		{"This", "is", "line", "two", "para", "one"},
		{"This", "is", "line", "three", "para", "one"}},

	{
		{"This", "is", "line", "one", "para", "two"},
		{"This", "is", "line", "two", "para", "two"},
		{"This", "is", "line", "three", "para", "two"}}}

func main() {
	words := 0
	for _, line := range text {
		for _, sentence := range line {
			words = words + len(sentence)
		}
	}
	fmt.Println("Number of words is", words)
}
