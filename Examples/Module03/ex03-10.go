// Example 03-10 Looping with range

package main

import "fmt"

func main() {

	test := "Hi!"

	for index, letter := range test {
		fmt.Printf("Letter %d is %#U\n", index, letter)
	}
}
