// Example 02-04  Implicit Initialization
package main

import "fmt"

var x = float32(0)
var c = complex64(0)
var b = true
var str = "Hi"

func main() {
	var message = "x=%T c=%T b=%T str=%T i=%T \n"
	var i = uint8(0)
	fmt.Printf(message, x, c, b, str, i)
}
