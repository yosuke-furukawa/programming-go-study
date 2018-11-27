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

var input uint64 = 0x1234567890ABCDEF
var results []int
var benchmarkNum = 100000000

func benchmarkPopCount(b *testing.B, num int) {
	results = []int{}
	for i := 0; i < num; i++ {
		results = append(results, PopCount(input))
	}
}

func BenchmarkPopCount(b *testing.B) {
	benchmarkPopCount(b, benchmarkNum)
}

func benchmarkPopCountLoop(b *testing.B, num int) {
	results = []int{}
	for i := 0; i < num; i++ {
		results = append(results, PopCountLoop(input))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	benchmarkPopCountLoop(b, benchmarkNum)
}

func benchmarkPopCountShift64(b *testing.B, num int) {
	results = []int{}
	for i := 0; i < num; i++ {
		results = append(results, PopCountShift64(input))
	}
}

func BenchmarkPopCountShift64(b *testing.B) {
	benchmarkPopCountShift64(b, benchmarkNum)
}

func benchmarkPopCountAndOne(b *testing.B, num int) {
	results = []int{}
	for i := 0; i < num; i++ {
		results = append(results, PopCountAndOne(input))
	}
}

func BenchmarkPopCountAndOne(b *testing.B) {
	benchmarkPopCountAndOne(b, benchmarkNum)
}

func benchmarkPopCountHackersDelight(b *testing.B, num int) {
	results = []int{}
	for i := 0; i < num; i++ {
		results = append(results, PopCountHackersDelight(input))
	}
}
func BenchmarkPopCountHackersDelight(b *testing.B) {
	benchmarkPopCountHackersDelight(b, benchmarkNum)
}

func benchmarkPopCountOnesCount(b *testing.B, num int) {
	results = []int{}
	for i := 0; i < num; i++ {
		results = append(results, PopCountOnesCount(input))
	}
}
func BenchmarkPopCountOnesCount(b *testing.B) {
	benchmarkPopCountOnesCount(b, benchmarkNum)
}

func benchmarkPopCountCPU(b *testing.B, num int) {
	results = []int{}
	for i := 0; i < num; i++ {
		results = append(results, PopCountCPU(input))
	}
}
func BenchmarkPopCountCPU(b *testing.B) {
	benchmarkPopCountCPU(b, benchmarkNum)
}

/*
BenchmarkPopCount-4                     50000000                34.1 ns/op
BenchmarkPopCountLoop-4                 50000000                23.1 ns/op
BenchmarkPopCountShift64-4              10000000               121 ns/op
BenchmarkPopCountAndOne-4               50000000                32.5 ns/op
BenchmarkPopCountHackersDelight-4       100000000               38.9 ns/op
BenchmarkPopCountOnesCount-4            100000000               10.4 ns/op
BenchmarkPopCountCPU-4                  200000000               29.5 ns/op
*/
