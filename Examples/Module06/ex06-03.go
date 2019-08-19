// Example 06-03 Empty versus nil maps

package main

import "fmt"

var errs map[string]string

func main() {
	fmt.Println("Check for nil before make() ", errs == nil)
	fmt.Println("Length of errs ", len(errs))

	errs = make(map[string]string)

	fmt.Println("Check for nil after make() ", errs == nil)
	fmt.Println("Length of errs ", len(errs))
}
