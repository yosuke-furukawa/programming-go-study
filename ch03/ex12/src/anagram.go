package src

import "strings"

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, v := range s1 {
		b := strings.Contains(s2, string(v))
		if !b {
			return false
		}
	}
	return true
}
