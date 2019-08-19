// Example 13-06 Hello.go with init functions
package main

import (
	"fmt"
	s1 "stringutil"
	s2 "util/stringutil"
)

func init() {
	fmt.Println("Main init 1")
}
func init() {
	fmt.Println("Main init 2")
}

func main() {
	fmt.Println(s1.Reverse("Hello World!"))
	fmt.Println(s2.Reverse("Hello World!"))
}
