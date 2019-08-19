// Ex 10-08 Support Code

package main

import "fmt"

type fish struct{ species string }
type human struct{ name string }
type squid struct{ inkyness int }
type cat struct{ breed string }

func (f *fish) swim() {
	fmt.Println("Underwater")
}

func (f *squid) swim() {
	fmt.Println("Jetting along")
}
func (f *cat) swim() {
	fmt.Println("If I must...")
}
func (f *human) swim() {
	fmt.Println("Dog Paddle")
}

func (f *human) walk() {
	fmt.Println("Strolling around")
}
