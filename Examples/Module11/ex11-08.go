// Example 11-08 Buffered channels

package main

import (
	"fmt"
	"time"
)

func service(message string, c chan string) {
	for i := 0; ; i++ {
		fmt.Println("service: ", i, "buffered items ", len(c))
		c <- fmt.Sprintf("%s %d", message, i)
	}
}

func main() {
	c := make(chan string, 3)
	go service("Message: ", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("main %d Got back: %q\n", i, <-c)
		time.Sleep(time.Second)
	}
	fmt.Println("Done!")
}
