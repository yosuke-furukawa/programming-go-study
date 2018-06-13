package strings

import "testing"

func TestJoin(t *testing.T) {
	tests := []struct {
		data     []string
		sep      string
		expected string
	}{
		{
			[]string{"aaa", "bbb", "ccc"},
			"/",
			"aaa/bbb/ccc",
		},
		{
			[]string{"aaa", "bbb", "ccc", "ddd"},
			"----",
			"aaa----bbb----ccc----ddd",
		},
	}

	for _, test := range tests {
		result := Join(test.sep, test.data...)
		if test.expected != result {
			t.Errorf("expected %s, but actual %s", test.expected, result)
		}
	}
}
