// Example 12-03 Benchmarking Test
package main

import "testing"

// from fib_test.go
func BenchmarkFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(20)
	}
}
