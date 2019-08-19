// Lab 02-02  List of intitalizers

package main

import "fmt"

var i, b, str, x = 89, true, "hi", float32(45)

func main() {
	fmt.Printf("i=%T b=%T str=%T x=%T \n", i, b, str, x)
	fmt.Printf("i=%d b=%t str=%s x=%f \n", i, b, str, x)
}
