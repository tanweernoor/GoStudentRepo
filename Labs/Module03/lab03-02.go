// Lab 03-02 Count Vowels and Digits
package main

import (
	"fmt"
	"os"
)

func main() {

	var vowels, digits int
	for i := 1; i < len(os.Args); i++ {
		word := os.Args[i]
		for j := 0; j < len(word); j++ {
			switch word[j] {
			case 'a', 'e', 'i', 'o', 'u':
				vowels++
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				digits++
			}
		}
	}
	fmt.Println("Vowels=", vowels, "Digits=", digits)

}
