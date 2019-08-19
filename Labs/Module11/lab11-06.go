// lab11-06.go
package main

import (
	"fmt"
	"time"
)

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {

			out <- i // Send 'i' to 'out'.
		}
	}
}
func main() {

	source := make(chan int)
	filtered := make(chan int)
	go Generate(source)
	go Filter(source, filtered, 2)
	for {
		i := <-source
		if i > 40 {
			break
		} else {
			fmt.Println(<-filtered)
			time.Sleep(time.Duration(time.Millisecond))

		}
	}
}
