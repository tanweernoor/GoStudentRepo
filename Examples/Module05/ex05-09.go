// Example 05-09 Multidimensional initialization

package main

import "fmt"

func main() {
	var matrix = [4][4]int{
		{1, 2, 3, 4},
		{2, 4, 8, 16},
		{3, 9, 27, 81},
		{4, 16, 64, 256},
	}
	fmt.Println("Matrix: ", matrix)
}
