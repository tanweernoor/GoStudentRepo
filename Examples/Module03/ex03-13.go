// Example 03-13 Non-integral test

package main

import "fmt"

func main() {

	os := "fedora"
	switch os {
	case "fedora", "redhat":
		fmt.Println("Open Source")
	case "Windows":
		fmt.Println("Proprietary")
	default:
		fmt.Println("unknown")
	}

}
