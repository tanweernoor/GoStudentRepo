// Example 09-06 Method overloading

package main

import "fmt"

func main() {

	c := circle{point{1, 0}, 3.0}
	fmt.Println(c)
	c.swap()
	fmt.Println(c, c.area())
}
