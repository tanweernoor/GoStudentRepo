// Example 11-15 Shared Resource CSP Style
package main

import "fmt"
import "time"

var deposit chan int
var query chan int

func accountUpdater(d <-chan int, q chan<- int) {
	balance := 0
	for {
		select {
		case amt := <-d:
			balance = balance + amt
		case q <- balance:

		}
	}
}

func accountRequester(amt int, d chan<- int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(time.Millisecond))
		d <- amt
	}

}

func main() {
	deposit := make(chan int)
	query := make(chan int)
	go accountUpdater(deposit, query)
	go accountRequester(100, deposit)
	go accountRequester(10, deposit)
	time.Sleep(time.Duration(time.Second))
	fmt.Println("Final balance=", <-query)
}
