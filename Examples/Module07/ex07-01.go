// Example 07-01 Function variables

package main

import "fmt"

func f1() string {
	return "I'm f1"
}
func f2() string {
	return "and I'm f2"
}

func main() {
	f := f1
	fmt.Printf("f is of type '%T' \n", f)
	fmt.Println(f())
	f = f2
	fmt.Println(f())
}
