package ex07

import (
	"math/rand"
	"testing"
	"time"
)

var result1 IntSet
var result2 Set

func gen(count, randNum int) []uint {
	list := []uint{}
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < count; i++ {
		n := rng.Intn(randNum)
		list = append(list, uint(n))
	}
	return list
}

func BenchmarkIntSet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iset := IntSet{
			words: []uint{},
		}
		for _, i := range gen(100000, 10000000000) {
			iset.Add(int(i))
		}
		result1 = iset
	}
}

func BenchmarkSet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iset := Set{
			data: make(map[int]struct{}),
		}
		for _, i := range gen(100000, 10000000000) {
			iset.Add(int(i))
		}
		result2 = iset
	}
}

func BenchmarkIntSet_UnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iset1 := IntSet{
			words: gen(100000, 10000000000),
		}
		iset2 := IntSet{
			words: gen(100000, 10000000000),
		}
		iset1.UnionWith(&iset2)

		result1 = iset1
	}
}

func BenchmarkSet_UnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iset1 := Set{
			data: make(map[int]struct{}),
		}
		for _, i := range gen(100000, 10000000000) {
			iset1.Add(int(i))
		}
		iset2 := Set{
			data: make(map[int]struct{}),
		}
		for _, i := range gen(100000, 10000000000) {
			iset2.Add(int(i))
		}
		iset1.UnionWith(&iset2)

		result2 = iset1
	}
}
