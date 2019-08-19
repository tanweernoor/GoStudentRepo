// Example 11-07 Deadlocks

package main

import "fmt"

func service(c chan string) {
	c <- "Hello"     //write
	fmt.Println(<-c) // read
}

func main() {
	c := make(chan string)
	go service(c)
	c <- "ho"        // write
	fmt.Println(<-c) //read
}
