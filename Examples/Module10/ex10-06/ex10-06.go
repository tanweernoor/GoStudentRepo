// Example 10-06 Type Testing

package main

type Swimmer interface {
	swim()
}

func main() {

	var x = []Swimmer{&human{"bobby"}, &fish{}}

	for _, swimmer := range x {
		if h, ok := swimmer.(*human); ok {
			h.walk()
		}

	}

}
