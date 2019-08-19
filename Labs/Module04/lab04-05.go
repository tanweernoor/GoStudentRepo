// lab04-05.go
package main

import "fmt"

func first() {
	defer second()
	fmt.Println("first")
}
func second() {
	fmt.Println("second")
}

func main() {
	defer first()
	defer second()
}
