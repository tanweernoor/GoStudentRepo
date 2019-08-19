// Example 06-05  comma ok update

package main

import "fmt"

func update(m map[int]int, key int, val int) {
	_, ok := m[key]
	if ok {
		m[key] = val
	}
}
func main() {

	data := map[int]int{1: 100, 3: 300}
	update(data, 1, -1)
	update(data, 2, 200)
	fmt.Println(data)

}
