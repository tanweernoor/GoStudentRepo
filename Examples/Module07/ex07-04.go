// Example 07-04  Function literal

package main

import "fmt"

func main() {
	f := func(i int) bool { return i == 0 }
	fmt.Println("2 is 0 is", f(2))
	fmt.Println("2 is 0 is", f(0))
}
