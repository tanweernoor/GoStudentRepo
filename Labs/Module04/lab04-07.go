// Lab 04-07.go
package main

import "fmt"

func fib(i int) (sum int) {
	if i <= 1 {
		sum = 1
	} else {
		sum = fib(i-1) + fib(i-2)
	}
	return
}

func main() {
	n := 10
	fmt.Println("Fibonacci num for", n, "is", fib(n))
}
