// Example 07-06  Executing a literal directly

package main

import "fmt"

func main() {
	z := func(x int) (y int) {
		y = x + 1
		return
	}(0)
	fmt.Println(z)
}
