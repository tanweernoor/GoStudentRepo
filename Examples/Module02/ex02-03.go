// Example 02-03  Implicit Initialization
package main

import "fmt"

var x float32 = 32.0
var c = 5.0 + 3.1i
var b = true
var str = "Hi"

func main() {
	var message = "x=%T c=%T b=%T str=%T i=%T \n"
	var i = 3
	fmt.Printf(message, x, c, b, str, i)
}
