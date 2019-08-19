// Example 05-04 Array as parameter

package main

import "fmt"

func delta(prm [3]int) {
	prm[0] = -1
	fmt.Println("prm = ", prm)
}

func main() {
	var arg = [3]int{99, 98, 97}
	fmt.Println("arg = ", arg)
	delta(arg)
	fmt.Println("arg = ", arg)

}
