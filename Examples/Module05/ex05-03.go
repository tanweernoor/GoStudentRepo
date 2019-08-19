// Example 05-03 Array comparison

package main

import "fmt"

func main() {
	var ar1 = [5]int{0, 1, 2, 3, 4}
	var ar2 = [5]int{0, 1, 2, 3, 4}
	fmt.Println("ar1 == ar2 is ", ar1 == ar2)
	fmt.Println("ar1 != ar2 is ", ar1 != ar2)

}
