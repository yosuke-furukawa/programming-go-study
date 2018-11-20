package ex05

import "testing"

func TestIntSet_IsSameBehaviorOnEmbedSet(t *testing.T) {

	tests := []struct {
		data   []int
		expect []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 2, 3, 1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 1, 1342, 123231, 111},
			[]int{1, 1342, 123231, 111},
		},
	}

	for _, test := range tests {
		iset := &IntSet{}
		set := make(map[int]struct{})
		for _, item := range test.data {
			iset.Add(item)
			set[item] = struct{}{}
		}

		if iset.Len() != len(set) {
			t.Errorf("IntSet is different length, compared by set")
		}

		for _, i := range iset.Elems() {
			_, ok := set[i]
			if !ok {
				t.Errorf("IntSet has wrong integer, %d", i)
			}
		}

		for _, i := range test.expect {
			has := iset.Has(i)
			_, ok := set[i]
			if !has {
				t.Errorf("IntSet has wrong integer, %d", i)
			}
			if !ok {
				t.Errorf("Set has wrong integer, %d", i)
			}

		}

	}
}
