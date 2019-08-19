// Example 06-06 Iteration

package main

import "fmt"

func main() {

	dotw := map[string]int{
		"Sun": 1, "Mon": 2, "Tue": 3, "Wed": 4,
		"Thu": 5, "Fri": 6, "Sat": 7}
	for day, num := range dotw {
		fmt.Printf("(%s : %d) ", day, num)
	}
}
