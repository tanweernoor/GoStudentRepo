// Example 10-02  Interface Variable

package main

type Swimmer interface {
	swim()
}

func main() {
	var s Swimmer
	s = new(fish)
	s.swim()
	s = new(human)
	s.swim()

}
