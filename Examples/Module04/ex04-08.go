// Example 04-08 Recursion

package main

import "fmt"

// sums numbers from 1 - k

func sum(k int) int {
	fmt.Printf("executed sum(%d)\n", k)
	if k > 0 {
		return (k + sum(k-1))
	} else {
		return 0
	}
}

func main() {
	fmt.Println(sum(3))
}
