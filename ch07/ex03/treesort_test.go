package ex03

import "testing"

func TestTree_String(t *testing.T) {
	tests := []struct {
		values   []int
		expected string
	}{
		{
			[]int{1, 2, 3, 4, 5},
			"1 -- 2 -- 3 -- 4 -- 5",
		},
		{
			[]int{5, 3, 2, 1, 4},
			"1 -- 2 -- 3 -- 4 -- 5",
		},
		{
			[]int{1, 2, 3},
			"1 -- 2 -- 3",
		},
		{
			[]int{2, 0, 4},
			"0 -- 2 -- 4",
		},
	}

	for _, test := range tests {
		root := &tree{
			test.values[0],
			nil,
			nil,
		}

		for _, value := range test.values[1:] {
			root = add(root, value)
		}

		if root.String() != test.expected {
			t.Errorf("root string is not match actual %s, expected %s", root.String(), test.expected)
		}
	}
}
