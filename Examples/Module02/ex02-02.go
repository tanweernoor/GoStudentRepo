// Example 02-02  Initializing Variables
package main

import "fmt"

var x float32 = 32.0
var c complex64 = complex64(0)
var b bool = true
var str string = "Hi"

func main() {
	var message string = "x=%f c=%f b=%t str=|%s| i=%d \n"
	var i int = 3
	fmt.Printf(message, x, c, b, str, i)
}
