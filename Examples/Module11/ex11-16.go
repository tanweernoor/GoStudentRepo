// Example 11-16 Race Condition and Mutex
package main

import "fmt"
import "time"
import "sync"

var (
	bal_lock sync.Mutex
	balance  int
)

func update(amount int) {
	balance = amount
}
func query() int {
	return balance
}

func main() {
	go func() {
		bal_lock.Lock()
		bal := query()
		time.Sleep(time.Duration(time.Millisecond))
		update(bal + 100)
		bal_lock.Unlock()
	}()
	go func() {
		bal_lock.Lock()
		bal := query()
		update(bal + 1)
		bal_lock.Unlock()
	}()

	time.Sleep(time.Duration(time.Second))
	fmt.Println("Final Balance=", query())
}
