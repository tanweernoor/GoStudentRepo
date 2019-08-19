// Example 10-01  Interface Definiton

package main

import "fmt"

type Swimmer interface {
	swim()
}

type fish struct{ species string }
type human struct{ name string }

func (f *fish) swim() {
	fmt.Println("Underwater")
}

func (f *human) swim() {
	fmt.Println("Dog Paddle")
}

func (f *human) walk() {
	fmt.Println("Strolling around")
}

func main() {
	tuna := new(fish)
	tuna.swim()
	knuth := new(human)
	knuth.swim()
	knuth.walk()

}
