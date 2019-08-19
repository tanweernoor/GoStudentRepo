// Example 12-02 Fixed Program tests -- same as 12-01
package main

import "testing"

func TestOne(t *testing.T) {
	v, c := count("a test case")
	if c != 5 {
		t.Error("Test One: Expected 5 cons, got ", c)
	}
	if v != 4 {
		t.Error("Test One: Expected 4 vowels, got ", v)
	}
}
func TestTwo(t *testing.T) {
	v, c := count("And more stuff")
	if c != 8 {
		t.Error("Test Two: Expected 8 cons, got ", c)
	}
	if v != 4 {
		t.Error("Test Two: Expected 4 vowels, got ", v)
	}
}
