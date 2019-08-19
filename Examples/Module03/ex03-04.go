// Example 03-04 Local Variables

package main

import "fmt"

func main() {

	if x, y := 22, "hi"; x == 0 {
		fmt.Println("Value of x=", x, " y=", y)
	} else if x%2 == 0 {
		fmt.Println("Value of x=", x, " y=", y)
	} else {
		fmt.Println("Value of x=", x, " y=", y)
	}
}
