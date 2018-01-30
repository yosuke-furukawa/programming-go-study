package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestDup(t *testing.T) {
	counts, err := Dup([]string{"../fixtures/example1.txt"})
	if err != nil {
		t.Errorf("error is thrown %v", err)
	}
	for filename, count := range counts {
		if filename != "../fixtures/example1.txt" {
			t.Errorf("unexpected filename %s", filename)
		}
		for line, n := range count {
			if n < 1 {
				t.Errorf("unexpected linecount %f", n)
			}
			data, _ := ioutil.ReadFile(filename)
			if !strings.Contains(string(data), line) {
				t.Errorf("unexpected line %s", line)
			}
		}
	}
}
