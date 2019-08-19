// lab11-01.go
package main

import "fmt"

func launch() {
	fmt.Println("**BOOM**")
}

func main() {
	fmt.Println("Commencing countdown.")
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
	}
	launch()
}
