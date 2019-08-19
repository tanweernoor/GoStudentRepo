// Example 03-06 Basic for loop

package main

import "fmt"

func main() {
	var total = 0
	for count := 0; count < 100; count++ {
		total += count
	}
	fmt.Println("total = ", total)

}
