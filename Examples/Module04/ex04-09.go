// Example 04-09 Varadic functions

package main

import "fmt"

func addup(nums ...int) (sum int) {
	for _, val := range nums {
		sum += val
	}
	return
}

func main() {
	fmt.Println(addup(1, 2, 3, 4, 5))
}
