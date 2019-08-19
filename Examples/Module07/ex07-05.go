// Example 07-05  Function parameter

package main

import "fmt"

func f2(p func(int) bool) {
	fmt.Println("2 is 0 is", p(2))
	fmt.Println("2 is 0 is", p(0))
}

func main() {
	f2(func(i int) bool { return i == 0 })
}
