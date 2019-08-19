// Example 10-07 Switching on types

package main

import "fmt"

type Swimmer interface {
	swim()
}

func main() {

	var x = []Swimmer{&human{"bobby"}, &fish{}, &cat{}, &squid{}}

	for index, swimmer := range x {
		switch t := swimmer.(type) {
		case *human:
			fmt.Printf("Item %d is a human and is ", index)
			t.walk()
		case *squid:
			fmt.Println("Item ", index, "is a squid")
		case *fish:
			fmt.Println("Item ", index, "is a fish")
		default:
			fmt.Printf("Item %d is of type %T\n", index, t)
		}
	}

}
