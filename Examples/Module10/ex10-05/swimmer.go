// Ex 10-05 Support Code

package main

import "fmt"

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
