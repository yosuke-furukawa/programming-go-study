package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	count := PopCount(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCount(16)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

func TestPopCountLoop(t *testing.T) {
	count := PopCountLoop(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCountLoop(16)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

func TestPopCountShift64(t *testing.T) {
	count := PopCountShift64(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCountShift64(16)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

func TestPopCountAndOne(t *testing.T) {
	count := PopCountAndOne(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCountAndOne(16)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

func TestPopCountHackersDelight(t *testing.T) {
	count := PopCountHackersDelight(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCountHackersDelight(16)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

func TestPopCountCPU(t *testing.T) {
	count := PopCountCPU(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCountCPU(16)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

var results []int

func BenchmarkPopCount(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCount(uint64(i)))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountLoop(uint64(i)))
	}
}

func BenchmarkPopCountShift64(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountShift64(uint64(i)))
	}
}

func BenchmarkPopCountAndOne(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountAndOne(uint64(i)))
	}
}

func BenchmarkPopCountHackersDelight(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountHackersDelight(uint64(i)))
	}
}

func BenchmarkPopCountOnesCount(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountOnesCount(uint64(i)))
	}
}

func BenchmarkPopCountCPU(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountCPU(uint64(i)))
	}
}

/*
BenchmarkPopCount-4                     50000000                25.4 ns/op
BenchmarkPopCountLoop-4                 50000000                26.7 ns/op
BenchmarkPopCountShift64-4              10000000               141 ns/op
BenchmarkPopCountAndOne-4               100000000               33.5 ns/op
BenchmarkPopCountHackersDelight-4       200000000               41.2 ns/op
BenchmarkPopCountOnesCount-4            100000000               12.6 ns/op
BenchmarkPopCountCPU-4                  200000000               24.5 ns/op
*/
