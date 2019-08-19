// Example 02-09  Constants
package main

import "fmt"

const a float32 = 1
const b = 4 / 3

func main() {
	const c string = "Hello Word"
	fmt.Printf("a=%T b=%T c=%T\n", a, b, c)
	fmt.Printf("a=%f b=%d c=%s\n", a, b, c)
}
