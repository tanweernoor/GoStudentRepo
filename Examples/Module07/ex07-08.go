// Example 07-08  Simple Closure

package main

import "fmt"

func f(p func(string)) {
	p("Mars")
}

func main() {
	a := "again "
	z := func(name string) {
		fmt.Println("Hello ", a, name)
		return
	}
	f(z)
}
