// Example 04-06 Deferred execution

package main

import "fmt"

func f(message string) {
	fmt.Println("Value of param =", message)
	return
}
func main() {
	m := "before defer"
	defer f(m)
	m = "after defer"
	fmt.Println("Value of m when main() exits =", m)
}
