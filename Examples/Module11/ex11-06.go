// Example 11-06 Basic channel
package main

import (
	"fmt"
	"time"
)

func service(message string, c chan string) {
	for i := 0; ; i++ {
		fmt.Println("service: ", i)
		c <- fmt.Sprintf("%s %d", message, i)
	}
}

func main() {
	c := make(chan string)
	go service("Message:", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("main %d Got back: %q\n", i, <-c)
		time.Sleep(time.Second)
	}
	fmt.Println("Done!")
}
