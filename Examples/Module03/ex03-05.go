// Example 03-05 Local and Non-Local Variables

package main

import "fmt"

var x int = 10

func main() {

	if x, y := 22, "hi"; x == 0 {
		fmt.Println("Value of x=", x, " y=", y)
	} else if x%2 == 0 {
		fmt.Println("Value of x=", x, " y=", y)
	} else {
		fmt.Println("Value of x=", x, " y=", y)
	}
	fmt.Println("Value of x=", x)
}
