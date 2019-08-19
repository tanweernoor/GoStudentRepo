// Example 05-05 Array assignment

package main

import "fmt"

func main() {
	var ar1 = [3]int{99, 98, 97}
	var ar2 [3]int
	ar2 = ar1
	ar2[0] = 0
	fmt.Println("ar1 =", ar1)
	fmt.Println("ar2 =", ar2)
}
