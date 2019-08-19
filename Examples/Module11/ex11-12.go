// Example 11-12 Modification to receiver and sender
package main

import (
	"fmt"
	"time"
)

func sender(output chan<- int) {
	defer close(output)
	for i := 0; i < 3; i++ {
		output <- i
	}
}

func receiver(input <-chan int) {
	for {
		j, ok := <-input
		fmt.Println(ok)
		if !ok {
			break
		}
		fmt.Println("Received ", j)
	}
	fmt.Println("I'm done ")

}

func main() {

	c := make(chan int)
	go receiver(c)
	go sender(c)
	time.Sleep(1 * time.Millisecond)

}
