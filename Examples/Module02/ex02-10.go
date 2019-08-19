// Example 02-10 iota and enums
package main

import "fmt"

const (
	a = iota
	b
	c
)

func main() {
	fmt.Printf("a=%T b=%T c=%T\n", a, b, c)
	fmt.Printf("a=%d b=%d c=%d\n", a, b, c)
}
