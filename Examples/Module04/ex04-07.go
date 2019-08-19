// Example 04-07 Stacked defer

package main

import "fmt"

func f(k int) {
	fmt.Printf("executed f(%d)\n", k)
	return
}

func main() {
	for i := 1; i < 4; i++ {
		fmt.Printf("called f(%d)\n", i)
		defer f(i)
	}
	fmt.Println("end main")
}
