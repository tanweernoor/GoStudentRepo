// lab03-01.go
package main

import (
	"fmt"
	"os"
)

func main() {

	var chars int
	for i := 1; i < len(os.Args); i++ {
		chars += len(os.Args[i])

	}
	fmt.Println("Words=", len(os.Args)-1, "Characters=", chars)
}
