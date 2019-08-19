// Example 11-04 Multiple goroutines
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
	go service("Alpha:")
	go service("Beta:")
	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
}
