// Example 03-07 Multiple Declarations

package main

import "fmt"

func main() {
	var total = 0
	for i, m := 0, "abort"; i < 100; i++ {
		total += i
		if total > 100 {
			fmt.Println(m)
			break
		}
	}
	fmt.Println("total = ", total)

}
