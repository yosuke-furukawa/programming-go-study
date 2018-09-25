package popcount

import (
	"testing"
)

var results []int

func BenchmarkPopCountLoop(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountLoop(uint64(i)))
	}
}

func BenchmarkPopCountLoopLazy(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountLoopLazy(uint64(i)))
	}
}
