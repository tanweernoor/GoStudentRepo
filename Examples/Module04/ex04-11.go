// Example 04-11 Error Handling

package main

import "fmt"
import "errors"

func division(num, denom int) (int, error) {
	if denom == 0 {
		return 0, errors.New("Divide by zero")
	}
	return (num / denom), nil
}

func main() {
	res, e := division(56, 2)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(res)
	}
}
