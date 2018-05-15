package src

import "testing"

func isSameSlice(a, b []string) bool {
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

func TestReduct(t *testing.T) {
	d := []string{"a", "b", "b", "c", "d", "e", "e"}
	exp := []string{"a", "b", "c", "d", "e"}

	ans := Reduct(d)
	if !isSameSlice(exp, ans) {
		t.Errorf("answer is not correct %#v", ans)
	}

	d = []string{"a", "b", "b", "c", "d", "e", "e", "e"}
	exp = []string{"a", "b", "c", "d", "e"}

	ans = Reduct(d)
	if !isSameSlice(exp, ans) {
		t.Errorf("answer is not correct %#v", ans)
	}

	d = []string{"b", "b", "b", "c", "d", "e", "e", "e"}
	exp = []string{"b", "c", "d", "e"}

	ans = Reduct(d)
	if !isSameSlice(exp, ans) {
		t.Errorf("answer is not correct %#v", ans)
	}
}
