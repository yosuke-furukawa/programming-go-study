package ex04

import "testing"

func TestIntSet_Elems(t *testing.T) {

	tests := []struct {
		source   []int
		expected []int
	}{
		{
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		set := &IntSet{}
		set.AddAll(test.source...)

		elems := set.Elems()
		for i, item := range test.expected {
			if elems[i] != item {
				t.Errorf("IntSet is not included this item, %d in %s", item, set)
			}
		}

		if set.Len() != len(test.expected) {
			t.Errorf("IntSet is not expected size. %d, %s", set.Len(), set)

		}

	}
}
