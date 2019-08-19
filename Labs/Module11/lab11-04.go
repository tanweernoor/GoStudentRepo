// lab11-01.go
package main

import "fmt"
import "time"
import "os"

func launch() {
	fmt.Println("**BOOM**")
}

func countdown() {
	abort := make(chan int)
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- 0
	}()
	fmt.Println("Commencing countdown.\nPress Enter to Abort")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-tick:
			fmt.Println(countdown)
		case <-abort:
			fmt.Println("Aborted")
			return
		}
	}
	launch()

}

func main() {
	countdown()
}
