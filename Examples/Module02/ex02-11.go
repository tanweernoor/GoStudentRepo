// Example 02-11 iota and enums
package main

import "fmt"

const (
	_ = iota
	a
	b
	c
)

func main() {
	fmt.Printf("a=%T b=%T c=%T\n", a, b, c)
	fmt.Printf("a=%d b=%d c=%d\n", a, b, c)
}
