// Example 03-02 parallel Assignment

package main

import "fmt"

func main() {

	first, last := "York", "New"
	fmt.Println(first, last)
	first, last = last, first
	fmt.Println(first, last)

}
