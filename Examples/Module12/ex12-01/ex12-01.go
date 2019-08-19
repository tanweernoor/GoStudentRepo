// Example 12-01 Buggy Program to Test
package main

import "fmt"

func count(s string) (vowels, cons int) {

	for _, letter := range s {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		default:
			cons++
		}
	}
	return
}
func main() {

	input := "This is a test"
	v, c := count(input)
	fmt.Printf("vowels=%d, consonants=%d\n", v, c)

}
