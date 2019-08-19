// Example 09-04 Inner struct methods

package main

import "fmt"

func main() {

	c := circle{point{1, 0}, 3.0}
	fmt.Println(c)
	c.center.swap()
	fmt.Println(c, c.area())
}
