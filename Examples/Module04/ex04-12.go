// Example 04-12 Comma OK Idiom

package main

import "fmt"

func division(num, denom int) (int, bool) {
	if denom == 0 {
		return 0, false
	}
	return (num / denom), true
}

func main() {
	res, ok := division(56, 2)
	if ok {
		fmt.Println(res)
	} else {
		fmt.Println("division failed")
	}
}
