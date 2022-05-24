package basic

import (
	"testing"
)

// To run benchmarks
// $ go test -bench=.

// To run single benchmark
// $ go test -bench ConcatenateBuffer
// or
// $ go test -bench ConcatenateJoin

// Flags

// -benchtime
// $ go test -bench BenchmarkConcatenateJoin -benchtime 5s

// -benchmem
// $ go test -bench . -benchmem

var result string

func BenchmarkConcatenateBuffer(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = ConcatenateBuffer("test1", "test2")
	}
	result = s
}

func BenchmarkConcatenateJoin(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = ConcatenateJoin("test1", "test2")
	}
	result = s
}

// Run benchmarks with code, without CLI
// func RunBench() {
// 	res := testing.Benchmark(BenchmarkConcatenateBuffer)
// 	fmt.Printf("Time taken: %s", res.T)
// 	fmt.Printf("Number of bytes allocated: %d", res.Bytes)
// 	fmt.Printf("Memory allocations: %d", res.MemAllocs)
// }
