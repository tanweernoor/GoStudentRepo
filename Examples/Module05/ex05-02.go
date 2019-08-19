// Example 05-02 Explicit array initialization

package main

import "fmt"

func main() {
	var ar1 = [5]int{0, 1, 2, 3, 4}
	var ar2 = [...]string{"Hello", "World"}
	ar3 := [2]bool{true, false}
	ar4 := [...]int{3: -1, 4: -1}
	fmt.Println("ar1=", ar1, "length=", len(ar1))
	fmt.Println("ar2=", ar2, "length=", len(ar2))
	fmt.Println("ar3=", ar3, "length=", len(ar3))
	fmt.Println("ar4=", ar4, "length=", len(ar4))
}
