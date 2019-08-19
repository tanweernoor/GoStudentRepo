// Example 10-03  Interface Variable as parameter

package main

type Swimmer interface {
	swim()
}

func submerge(s Swimmer) {
	s.swim()
}
func main() {
	submerge(new(fish))
	submerge(new(human))

}
