// Example 05-06 Iteration with range

package main

import "fmt"

func main() {
	words := []string{"the", "best", "of", "times"}
	for index, value := range words {
		fmt.Println(index, " ", value)
	}
}
