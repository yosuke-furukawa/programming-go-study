package ex10

import "sort"

func IsPalindrome(s sort.Interface) bool {
	half := s.Len() / 2
	for i := 0; i < half; i++ {
		j := s.Len() - i - 1
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		}
		return false
	}
	return true
}
