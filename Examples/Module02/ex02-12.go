// Example 02-12 Enums with Generator
package main

import "fmt"

const (
	a = iota * 2
	b
	c
)

func main() {
	fmt.Printf("a=%T b=%T c=%T\n", a, b, c)
	fmt.Printf("a=%d b=%d c=%d\n", a, b, c)
}
