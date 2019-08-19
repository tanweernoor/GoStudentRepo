// Example 03-11 Simple Switch

package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			fmt.Println("Case 0")
		case 1:
			fmt.Println("Case 1")
			fallthrough
		default:
			fmt.Println("Default")
		}
	}
}
