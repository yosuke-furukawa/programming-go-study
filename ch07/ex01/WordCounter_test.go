package ex01

import (
	"fmt"
	"testing"
)

func TestWordCounter_Write(t *testing.T) {
	tests := []struct {
		words    string
		expected int
	}{
		{
			"foo\nbar\nbaz",
			9,
		},
		{
			"foo\nbar\nbaz\n",
			9,
		},
		{
			"foo\nbar\nbaz\nboo\ntest\ntoste\n",
			21,
		},
	}

	for _, test := range tests {
		var c WordCounter
		fmt.Fprint(&c, test.words)
		if int(c) != test.expected {
			t.Errorf("Counter is incorrect, c = %d", c)
		}
	}
}
