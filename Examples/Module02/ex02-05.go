// Example 02-05  Variations
package main

import "fmt"

var i, j int8
var b, str, x = true, "hi", float32(45)

func main() {
	fmt.Printf("i=%T j=%T b=%T str=%T x=%T \n",
		i, j, b, str, x)
}
