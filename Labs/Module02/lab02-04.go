// Lab 02-04  Strings
package main

import "fmt"

func main() {
	Cyrillic := "中 Я"
	Latin := "a b"
	Hanzi := "亲 仇"
	fmt.Println("Length of", Latin, " is", len(Latin))
	fmt.Println("Length of", Hanzi, " is", len(Hanzi))
	fmt.Println("Length of", Cyrillic, " is", len(Cyrillic))
}
