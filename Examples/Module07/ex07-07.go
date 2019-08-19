// Example 07-07  Anonymous inner function

package main

import "fmt"

func main() {
	defer func(name string) {
		fmt.Println("Hello ", name)
		return
	}("World")
	fmt.Println("Main Function")
}
