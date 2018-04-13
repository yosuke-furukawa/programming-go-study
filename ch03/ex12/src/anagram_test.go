package src

import "testing"

func TestAnagram(t *testing.T) {
	t.Parallel()
	s1 := "abcd"
	s2 := "dbac"
	result := anagram(s1, s2)
	if !result {
		t.Errorf("%s and %s are not anagram", s1, s2)
	}
	s1 = "こんにちは"
	s2 = "はんこちに"
	result = anagram(s1, s2)
	if !result {
		t.Errorf("%s and %s are not anagram", s1, s2)
	}
}

func TestFailAnagram(t *testing.T) {
	t.Parallel()
	s1 := "abcd"
	s2 := "aaaa"
	result := anagram(s1, s2)
	if result {
		t.Errorf("%s and %s are anagram", s1, s2)
	}
	s1 = "こんにちは"
	s2 = "はんここち"
	result = anagram(s1, s2)
	if result {
		t.Errorf("%s and %s are not anagram", s1, s2)
	}
}
