// Example 13-05 Hello.go with aliases

package main

import (
	"fmt"
	s1 "stringutil"
	s2 "util/stringutil"
)

func main() {
	fmt.Println(s1.Reverse("Hello World!"))
	fmt.Println(s2.Reverse("Hello World!"))
}
