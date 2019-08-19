// Example 04-14 Runtime panic

package main

import "fmt"

func division(num, denom int) int {
	if denom == 0 {
		panic("Dividing by zero?!?")
	}
	return (num / denom)
}

func main() {
	res := division(56, 0)
	fmt.Println(res)

}
