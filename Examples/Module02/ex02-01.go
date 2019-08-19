// Example 02-01  Declaring Variables
package main

import "fmt"

var x float32
var c complex64
var b bool
var str string

func main() {
	var message string = "x=%f c=%f b=%t str=|%s| i=%d \n"
	var i int
	fmt.Printf(message, x, c, b, str, i)
}
