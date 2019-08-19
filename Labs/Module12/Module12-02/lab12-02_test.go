// lab12-02_test Palindrome function
package main

import "testing"

func TestPal1(t *testing.T) {
	if !isPal("radar") {
		// single quotes allow us to create a string with the double quote
		// marks embedded
		t.Error(`IsPalindrome("radar") = false`)
	}
	if !isPal("a") {
		t.Error(`IsPalindrome("a") = false`)
	}
	if !isPal("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
	if !isPal("Adamada") {
		t.Error(`IsPalindrome("Adamada") = false`)
	}
}

func TestNonPal1(t *testing.T) {
	if isPal("butter") {
		t.Error(`IsPalindrome("butter") = true`)
	}
}
