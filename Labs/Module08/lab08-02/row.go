// Lab08-02 Rows
package main

import "fmt"

// The row data structure
type row struct {
	data     []int // this is a slice of ints
	n        int   // the number of rows
	mean     float32
	max, min int
}

func makeRow(d []int) (s *row) {
	// If there is no data, then we return nil
	if len(d) == 0 {
		s = nil
		return
	}
	// Make a row structure.
	s = &row{make([]int, len(d)), len(d), float32(0), 0, 0}
	// Copy the data otherwise a reference to a slice could result
	// in the underlying data being changed elsewhere and corrupting
	// out stats
	// s.data = d  does not create a copy of the data
	copy(s.data, d)
	// Use the procRow routine to compute the stats
	s.max, s.min, s.n, s.mean = rowProc(d)
	return
}

func rowProc(d []int) (max, min, n int, mean float32) {
	sum := 0
	// Initialize the min and max to the first element
	// in the row otherwise we may get a wrong result
	// if we initialize to 0.
	max, min = d[0], d[0]
	// iterate adding up the sum and resetting the max
	// and min as necessary
	for _, v := range d {
		sum += v
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}
	// Compute stats and done.
	n = len(d)
	mean = float32(sum) / float32(n)
	return
}

func printRow(r *row) {
	fmt.Println(r.data, " n=", r.n, "max=", r.max, "min=", r.min, "mean=", r.mean)
}
