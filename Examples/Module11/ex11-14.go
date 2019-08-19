// Example 11-14 Race Conditoin
package main

import "fmt"
import "time"

var balance int

func update(amount int) {
	balance = amount
}
func query() int {
	return balance
}

func main() {
	go func() {
		bal := query()
		time.Sleep(time.Duration(time.Millisecond))
		update(bal + 100)
	}()
	go func() {
		bal := query()
		update(bal + 1)
	}()

	time.Sleep(time.Duration(time.Second))
	// balance should be 101
	fmt.Println("Final Balance=", query())
}
