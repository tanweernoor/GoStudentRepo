// Example 11-03 Basic goroutines
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func service(message string) {
	for i := 0; ; i++ {
		fmt.Println(message, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}

}
func main() {
	go service("Message:")
	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
}
