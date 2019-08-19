// Example 10-05 Amphib Type

package main

type Swimmer interface {
	swim()
}

type Walker interface {
	walk()
}

type Amphib interface {
	Walker
	Swimmer
}

func main() {
	var a Amphib
	a = new(human)
	a.walk()
	a.swim()

}
