// ex12-04.go
package main

import (
	"fmt"

	"github.com/pkg/profile"
)

func fib(i int) (sum int) {
	if i <= 1 {
		sum = 1
	} else {
		sum = fib(i-1) + fib(i-2)
	}
	return
}

func main() {
	defer profile.Start().Stop()
	n := 30
	fmt.Println("Fibonacci num for", n, "is", fib(n))
}
