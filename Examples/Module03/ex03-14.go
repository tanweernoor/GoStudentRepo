// Example 03-14 Switch with no test

package main

import "fmt"

func main() {
	x := 22

	switch {
	case x == 0:
		fmt.Println("zero")
	case x%2 == 0:
		fmt.Println("even")
	default:
		fmt.Println("odd")
	}

}
