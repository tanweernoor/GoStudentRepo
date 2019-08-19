// Example 11-05 Anonymous goroutines
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go func() {
		for i := 0; ; i++ {
			fmt.Println("Anon:", i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
}
