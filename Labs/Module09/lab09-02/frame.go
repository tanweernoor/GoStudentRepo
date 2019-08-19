// Lab 09-02 frame.go
package main

import "fmt"

type frame struct {
	headings   []string
	rows, cols int
	data       []row
	max, min   int
	mean       float32
}

func (f *frame) addHeadings(h []string) {
	copy(f.headings, h)
}

func (f *frame) addRow(r *row) {
	if r.n != f.cols {
		return
	}
	f.data = append(f.data, *r)
	f.rows += 1
}

func (f *frame) computeStats() {
	sum := float32(0)
	for i, v := range f.data {
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
func makeFrame(h []string, d []row) (f *frame) {
	if len(h) == 0 || len(d) == 0 {
		return
	}
	f = &frame{make([]string, len(h)), 0, 0, make([]row, len(d)), 0, 0, 0.0}
	f.addHeadings(h)
	copy(f.data, d)
	f.cols = len(h)
	f.rows = len(d)
	f.computeStats()
	return
}

func (f *frame) printFrame() {
	fmt.Println("Data Frame")
	fmt.Println("rows=", f.rows, "columns=", f.cols)
	fmt.Println("max=", f.max, "min=", f.min, "mean=", f.mean)
	fmt.Println(f.headings)
	for _, v := range f.data {
		v.printRow()
	}
}
