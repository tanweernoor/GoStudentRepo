// lab11-07.go
package main

import (
	"fmt"
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
	go Generate(source)
	for i := 0; i < 10; i++ {
		prime := <-source
		fmt.Println(prime)
		filtered := make(chan int)
		go Filter(source, filtered, prime)
		source = filtered
	}
}
