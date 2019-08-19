// lab11-02.go
package main

import "fmt"
import "time"

func launch() {
	fmt.Println("**BOOM**")
}

func main() {
	fmt.Println("Commencing countdown.")
	// added the tick here
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		// blocking until the one second tick is received.
		<-tick
		fmt.Println(countdown)

	}
	launch()
}
