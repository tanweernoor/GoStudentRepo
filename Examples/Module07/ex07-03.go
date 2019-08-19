// Example 07-03  Function as return value

package main

import "fmt"

func f1() string {
	return "I'm f1"
}
func f2() func() string {
	return f1
}

func main() {
	f := f2()
	fmt.Println(f())
}
