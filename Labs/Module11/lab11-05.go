// lab11-05.go
package main

import (
	"fmt"
)

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}
func main() {
	source := make(chan int)
	go Generate(source)
	// for testing read from the generator until we hit 5
	for {
		i := <-source
		if 5 == i {
			break
		} else {
			fmt.Println(i)
		}
	}
}
