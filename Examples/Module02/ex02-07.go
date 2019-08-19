// Example 02-07  Variable Scope
package main

import "fmt"

var k int = 1

func main() {

	i := 1
	{
		x := 1.2
		fmt.Printf("i=%d x=%f k=%d \n", i, x, k)
	}
}
