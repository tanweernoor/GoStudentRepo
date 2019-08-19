// Example 06-04 Map capacity

package main

import "fmt"

var errs map[string]string

func main() {
	errs = make(map[string]string, 200)
	fmt.Println("Length of errs ", len(errs))
}
