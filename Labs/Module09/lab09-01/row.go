// Lab09-01 Row Methods
package main

import "fmt"

type row struct {
	data     []int
	n        int
	mean     float32
	max, min int
}

func makeRow(d []int) (s *row) {
	if len(d) == 0 {
		s = nil
		return
	}
	s = &row{make([]int, len(d)), len(d), float32(0), d[0], d[0]}
	copy(s.data, d)
	s.computeStats()
	return
}

func (r *row) computeStats() {
	sum := 0
	r.max, r.min = r.data[0], r.data[0]
	for _, v := range r.data {
		sum += v
		if v > r.max {
			r.max = v
		} else if v < r.min {
			r.min = v
		}
	}
	r.n = len(r.data)
	r.mean = float32(sum) / float32(r.n)
	return
}

func (r *row) printRow() {
	fmt.Println(r.data, " n=", r.n, "max=", r.max, "min=", r.min, "mean=", r.mean)
}
