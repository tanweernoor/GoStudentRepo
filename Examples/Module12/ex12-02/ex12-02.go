// Example 12-02 Fixed Program
package main

import "fmt"

func count(s string) (vowels, cons int) {

	for _, letter := range s {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		case 'A', 'E', 'I', 'O', 'U':
			vowels++
		case ' ':
			break
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
