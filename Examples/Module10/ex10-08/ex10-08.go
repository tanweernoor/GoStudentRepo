// Example 10-08 The empty interface

package main

import "fmt"

type Swimmer interface {
	swim()
}

type myint int32
type point struct{ x, y int }

var i myint = 0

type whatever interface{}

func main() {

	mylist := []whatever{i, float64(45), &point{2, 3}, true, &fish{"tuna"}}
	for index, object := range mylist {
		switch v := object.(type) {
		case bool:
			fmt.Printf("item %d is a bool = %v\n", index, v)
		case myint:
			fmt.Printf("item %d is a myint = %v\n", index, v)
		case float64:
			fmt.Printf("item %d is a float64 = %v\n", index, v)
		case *point:
			fmt.Printf("item %d is a point = %v\n", index, *v)
		default:
			fmt.Printf("item %d is a %T = %v\n", index, v, v)
		}
	}

}
