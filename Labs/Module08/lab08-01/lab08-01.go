// Lab08-01 Making a Row
package main

func main() {
	//make three rows and headings
	r1 := makeRow([]int{1, 1, 3})
	r2 := makeRow([]int{9, 102, 5})
	r3 := makeRow([]int{43, -3, 715})
	printRow(r1)
	printRow(r2)
	printRow(r3)
}
