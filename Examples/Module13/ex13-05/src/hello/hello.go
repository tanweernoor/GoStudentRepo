// Example 13-07 Dependent init functions
package main

import (
	"fmt"
	"stringutil"
	_ "util/stringutil"
)

func init() {
	fmt.Println("Main init 1")
}
func init() {
	fmt.Println("Main init 2")
}

func main() {
	fmt.Println(stringutil.Reverse("Hello World!"))

}
