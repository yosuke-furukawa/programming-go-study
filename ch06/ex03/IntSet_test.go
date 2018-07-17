package ex03

import "testing"

func TestIntSet_IntersectWith(t *testing.T) {

	tests := []struct {
		source   []int
		target   []int
		expected []int
	}{
		{
			[]int{1, 2},
			[]int{1, 3},
			[]int{1},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{3, 5, 6},
			[]int{3, 5, 6},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			[]int{},
		},
	}

	for _, test := range tests {
		set1 := &IntSet{}
		set1.AddAll(test.source...)

		set2 := &IntSet{}
		set2.AddAll(test.target...)

		set1.IntersectWith(set2)
		t.Logf("set %s", set1)
		for _, item := range test.expected {
			if !set1.Has(item) {
				t.Errorf("IntSet is not included this item, %d in %s", item, set1)
			}
		}

		if set1.Len() != len(test.expected) {
			t.Errorf("IntSet is not expected size. %d, %s", set1.Len(), set1)

		}

	}
}

func TestIntSet_DifferenceWith(t *testing.T) {

	tests := []struct {
		source   []int
		target   []int
		expected []int
	}{
		{
			[]int{1, 2},
			[]int{1, 3},
			[]int{2},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{3, 5, 6},
			[]int{1, 2, 4},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			[]int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		set1 := &IntSet{}
		set1.AddAll(test.source...)

		set2 := &IntSet{}
		set2.AddAll(test.target...)

		set1.DifferenceWith(set2)
		t.Logf("set %s", set1)
		for _, item := range test.expected {
			if !set1.Has(item) {
				t.Errorf("IntSet is not included this item, %d in %s", item, set1)
			}
		}

		if set1.Len() != len(test.expected) {
			t.Errorf("IntSet is not expected size. %d, %s", set1.Len(), set1)

		}

	}
}

func TestIntSet_SymmetricDifference(t *testing.T) {

	tests := []struct {
		source   []int
		target   []int
		expected []int
	}{
		{
			[]int{1, 2},
			[]int{1, 3},
			[]int{2, 3},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{3, 5, 6, 7},
			[]int{1, 2, 4, 7},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, test := range tests {
		set1 := &IntSet{}
		set1.AddAll(test.source...)

		set2 := &IntSet{}
		set2.AddAll(test.target...)

		set1.SymmetricDifference(set2)
		t.Logf("set %s", set1)
		for _, item := range test.expected {
			if !set1.Has(item) {
				t.Errorf("IntSet is not included this item, %d in %s", item, set1)
			}
		}

		if set1.Len() != len(test.expected) {
			t.Errorf("IntSet is not expected size. %d, %s", set1.Len(), set1)

		}

	}
}
