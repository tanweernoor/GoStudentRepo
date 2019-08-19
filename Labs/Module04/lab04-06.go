// lab04-06.go
package main

import "fmt"

func limits(x ...int) (min, max int, empty bool) {
	if len(x) == 0 {
		empty = true
		return
	} else {
		min, max = x[0], x[0]
	}
	for _, val := range x {
		switch {
		case val > max:
			max = val
		case val < min:
			min = val
		}
	}
	return
}

func main() {
	min, max, e := limits(2, 2, 2, 4, 3, 2)
	fmt.Println("min=", min, "max=", max, "empty=", e)
	min, max, e = limits()
	fmt.Println("min=", min, "max=", max, "empty=", e)
}
