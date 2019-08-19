// Example 10-04  Method not found in interface

package main

func main() {
	var s Swimmer
	s = new(human)
	s.walk()

}
