package ex10

import (
	"testing"
)

type runes []rune

func (rs runes) Len() int {
	return len(rs)
}

func (rs runes) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func (rs runes) Less(i, j int) bool {
	return rs[i] < rs[j]
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		data         string
		isPalindrome bool
	}{
		{
			"abcdcba",
			true,
		},
		{
			"abcdcbaaa",
			false,
		},
		{
			"a",
			true,
		},
		{
			"abbba",
			true,
		},
	}

	for _, test := range tests {
		isPalindrome := IsPalindrome(runes([]rune(test.data)))
		if isPalindrome != test.isPalindrome {
			t.Errorf("Palindrome error, %s, %v but actual %v", test.data, test.isPalindrome, isPalindrome)
		}
	}
}
