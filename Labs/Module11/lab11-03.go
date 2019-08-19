// lab11-03 .go
package main

import "fmt"
import "time"
import "os"

func launch() {
	fmt.Println("**BOOM**")
}

func main() {
	abort := make(chan int)
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- 0
	}()
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick

	}
	launch()
}
