// Lab09-01 Row Methods
package main

func main() {
	//make three rows and headings
	r1 := makeRow([]int{1, 1, 3})
	r2 := makeRow([]int{9, 102, 5})
	r3 := makeRow([]int{33, -3, 715})
	r1.printRow()
	r2.printRow()
	r3.printRow()
}
