// lab12-01 Palindrome function

package main

import "fmt"
import "unicode"

func isPal(s string) bool {
	var codepoints []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			codepoints = append(codepoints, unicode.ToLower(r))
		}
	}
	for i := range codepoints {
		if codepoints[i] != codepoints[len(codepoints)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	t := []string{"radar", "gogo", "a", "level", "butter", "été", "Abba"}
	for _, s := range t {
		fmt.Println(s, isPal(s))
	}

}
