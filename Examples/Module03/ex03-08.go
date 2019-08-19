// Example 03-08 Non Local Variables

package main

import "fmt"

func main() {
	var total, i int = 1000, 1000
	for i, total = 0, 0; i < 100; i++ {
		if total > 200 {
			continue
		} else {
			total += i
		}
	}
	fmt.Println("total = ", total)

}
