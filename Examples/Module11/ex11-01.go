// Example 11-01 Normal function call
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
	service("Message:")
}
