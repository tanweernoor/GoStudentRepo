// lab12-01 Palindrome function

package main

import "fmt"

func isPal(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	t := []string{"radar", "gogo", "a", "level", "butter", ""}
	for _, s := range t {
		fmt.Println(s, isPal(s))
	}

}
