package src

import "testing"

func isSameSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	exp := []int{4, 5, 6, 1, 2, 3}
	ans := Rotate(arr, 3)

	if !isSameSlice(ans, exp) {
		t.Errorf("answer is not correct %v", ans)
	}

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	exp = []int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}
	ans = Rotate(arr, 3)

	if !isSameSlice(ans, exp) {
		t.Errorf("answer is not correct %v", ans)
	}

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	exp = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	ans = Rotate(arr, 1)

	if !isSameSlice(ans, exp) {
		t.Errorf("answer is not correct %v", ans)
	}

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	exp = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	ans = Rotate(arr, 13)

	if !isSameSlice(ans, exp) {
		t.Errorf("answer is not correct %v", ans)
	}
}
