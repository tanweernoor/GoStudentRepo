// Example 02-08  Shadowing
package main

import "fmt"

var k int = 1

func main() {

	fmt.Printf("Package Scope k=%d \n", k)
	k := 2
	fmt.Printf("Function Scope k=%d \n", k)
	{
		k := 3
		fmt.Printf("Block Scope k=%d \n", k)
	}
	fmt.Printf("Function Scope k=%d \n", k)
}
