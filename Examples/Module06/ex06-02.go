// Example 06-02 Creating maps with make()

package main

import "fmt"

var errs map[int]string

func main() {
	errs = make(map[int]string)
	errs[0] = "Hardware"
	errs[1] = "Segmentation fault"
	fmt.Println(errs, " size =", len(errs))
	errs[0] = "Firmware fault"
	fmt.Println(errs, " size =", len(errs))
	fmt.Println("The errorcode '0' is a ", errs[0])
}
