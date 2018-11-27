package ex05

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s      string
		sep    string
		expect int
	}{
		{
			"a:b:c", ":", 3,
		},
		{
			"a,b,c", ",", 3,
		},
		{
			"a,bc", ",", 2,
		},
	}

	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.expect {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, got, test.expect)
		}
	}
}
