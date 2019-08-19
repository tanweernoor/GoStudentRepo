// Example 06-01 Creating maps with literals

package main

import "fmt"

var errs map[int]string

func main() {
	severity := map[string]string{
		"Blue":   "normal",
		"Orange": "moderate",
		"Red":    "severe"}

	severity["Black"] = "apocalyptic"

	fmt.Println(severity, " size =", len(severity))
	fmt.Println(errs, " size =", len(errs))
}
