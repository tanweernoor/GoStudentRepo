// Lab 08-02 frame.go
package main

import "fmt"

// The frame structure
type frame struct {
	headings   []string // The colum headings
	rows, cols int
	data       []row // a slice of rows
	max, min   int
	mean       float32
}

func makeFrame(h []string, d []row) (f *frame) {
	// If there no headings or no rows, then return nil
	// Remember that the namesd return value is initialized to
	// its zero value which is nil
	if len(h) == 0 || len(d) == 0 {
		return
	}
	// Make the empty frame
	f = &frame{make([]string, len(h)), 0, 0, make([]row, len(d)), 0, 0, 0.0}
	// Copy the headings and data into the frame.
	copy(f.headings, h)
	copy(f.data, d)
	f.cols = len(h) // colums is the number of headings
	f.rows = len(d) // rows is the size of the slice of rows
	// compute stats like in rows but now across the slice of rows
	sum := float32(0)
	for i, v := range f.data { //v is a row
		if i == 0 {
			f.max, f.min = v.max, v.min
		} else {
			if v.max > f.max {
				f.max = v.max
			} else if v.min < f.min {
				f.min = v.min
			}
		}
		sum += v.mean
	}
	f.mean = sum / float32(f.rows)
	return
}

func printFrame(f *frame) {
	fmt.Println("Data Frame")
	fmt.Println("rows=", f.rows, "columns=", f.cols)
	fmt.Println("max=", f.max, "min=", f.min, "mean=", f.mean)
	fmt.Println(f.headings)
	for _, v := range f.data {
		printRow(&v)
	}
}
