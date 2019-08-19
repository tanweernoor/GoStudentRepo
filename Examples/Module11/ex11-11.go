// Example 11-11 Unidrectional Channels
package main

import (
	"fmt"
	"time"
)

func sender(output chan<- int) {
	for i := 0; i < 3; i++ {
		output <- i
	}
	close(output)
}

func receiver(input <-chan int) {
	for j := range input {
		fmt.Println("Received ", j)

	}
	fmt.Println("I'm done ")
}

func main() {

	c := make(chan int)
	go receiver(c)
	go sender(c)
	time.Sleep(1 * time.Second)

}
