// Lab08-02 Adding the frame
package main

func main() {
	//make three rows and headings
	r1 := makeRow([]int{1, 1, 3})
	r2 := makeRow([]int{9, 102, 5})
	r3 := makeRow([]int{43, -3, 715})
	// create the slice of rows, notice we have to dereference
	d := []row{*r1, *r2, *r3}
	h := []string{"One", "Two", "Three"}
	f := makeFrame(h, d)
	printFrame(f)
}
