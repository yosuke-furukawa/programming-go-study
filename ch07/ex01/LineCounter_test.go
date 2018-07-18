package ex01

import (
	"fmt"
	"testing"
)

func TestLineCounter_Write(t *testing.T) {
	tests := []struct {
		words    string
		expected int
	}{
		{
			"foo\nbar\nbaz",
			3,
		},
		{
			"foo\nbar\nbaz\n",
			3,
		},
		{
			"foo\nbar\nbaz\nboo\ntest\ntoste\n",
			6,
		},
	}

	for _, test := range tests {
		var c LineCounter
		fmt.Fprint(&c, test.words)
		if int(c) != test.expected {
			t.Errorf("Counter is incorrect, c = %d", c)
		}
	}
}
