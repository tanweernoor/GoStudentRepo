// Example 11-13 Select Statement Multiplexer

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sourceEven(outchan chan<- int) {
	for i := 0; ; i++ {
		outchan <- i * 2
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
func sourceOdd(outchan chan<- int) {
	for i := 0; ; i++ {
		outchan <- (i * 2) + 1
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
func counter(ceven <-chan int, codd <-chan int) {

	for {
		select {
		case e := <-ceven:
			fmt.Println("Even:", e)
		case o := <-codd:
			fmt.Println("Odd:", o)
		default:
			fmt.Println("no one is ready")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	codd := make(chan int)
	ceven := make(chan int)
	go sourceOdd(codd)
	go sourceEven(ceven)
	go counter(ceven, codd)
	time.Sleep(2 * time.Second)
}
