// Example 03-12 Switch Statement Break

package main

import "fmt"

func main() {

	for i := 0; i < 2; i++ {
		switch i {
		case 0:
			fmt.Println("Case 0")
		case 1:
			fmt.Println("Case 1")
			break
			fmt.Println("After break")
		default:

			fmt.Println("Default")
		}
	}
}
