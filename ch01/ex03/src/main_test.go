package main

import "testing"

func TestEcho1(t *testing.T) {
	ans := echo1([]string{"a", "b", "c", "d", "e"})
	if ans != "a b c d e" {
		t.Errorf("unexpected result %s", ans)
	}
}

func TestEcho2(t *testing.T) {
	ans := echo2([]string{"a", "b", "c", "d", "e"})
	if ans != "a b c d e" {
		t.Errorf("unexpected result %s", ans)
	}
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1([]string{"a", "b", "c", "d", "e"})
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2([]string{"a", "b", "c", "d", "e"})
	}
}
