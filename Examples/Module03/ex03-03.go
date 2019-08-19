// Example 03-03 Basic If Statement

package main

import "fmt"

func main() {
	x := 22
	if x == 0 {
		fmt.Printf("%d is zero\n", x)
	} else if x%2 == 0 {
		fmt.Printf("%d is even\n", x)
	} else {
		fmt.Printf("%d is odd\n", x)
	}
}
