package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCount(t *testing.T) {

	var tests = []struct {
		input  string
		counts map[rune]int
	}{
		{
			"test",
			map[rune]int{'t': 2, 'e': 1, 's': 1},
		},
		{
			"ブンブンハローユーチューブ",
			map[rune]int{'ブ': 3, 'ン': 2, 'ハ': 1, 'ロ': 1, 'ー': 3, 'ユ': 1, 'チ': 1, 'ュ': 1},
		},
	}

	for _, test := range tests {
		result, _, _, _ := count(strings.NewReader(test.input))
		same := reflect.DeepEqual(test.counts, result)

		if !same {
			t.Errorf("expect count is not same expect %v, actual %v", test.counts, result)
		}
	}
}
