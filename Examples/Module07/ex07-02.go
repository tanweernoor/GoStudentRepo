// Example 07-02  Functions as parameters

package main

import "fmt"

func f1() string {
	return "I'm f1"
}
func f2(fparam func() string) {
	fmt.Println(fparam())
}

func main() {
	f2(f1)
}
