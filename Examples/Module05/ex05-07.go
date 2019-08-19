// Example 05-07 Multidimensional array

package main

import "fmt"

func main() {
	var matrix [2][3]int
	value := 10
	for row, col := range matrix {
		for index, _ := range col {
			matrix[row][index] = value
			value++
		}
	}
	fmt.Println("Matrix: ", matrix)
}
