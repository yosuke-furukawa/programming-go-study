package ex02

import "testing"

func TestIntSet_AddAll(t *testing.T) {

	tests := []struct {
		data     []int
		has      []int
		expected bool
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			true,
		},
		{
			[]int{1, 2, 3, 1, 2, 3},
			[]int{1, 2, 3},
			true,
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3, 5},
			false,
		},
		{
			[]int{1, 2, 3, 4},
			[]int{0},
			false,
		},
	}

	for _, test := range tests {
		set := &IntSet{}
		set.AddAll(test.data...)

		actual := true
		for _, item := range test.has {
			actual = actual && set.Has(item)
		}
		if actual != test.expected {
			t.Errorf("IntSet does not pass tests, %s", set)
		}
	}
}
