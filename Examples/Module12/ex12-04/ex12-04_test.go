// Example 12-04 Benchmarking Test with Profiling
package main

import "testing"

// from fib_test.go
func BenchmarkFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(20)
	}
}
