package ex01

import "testing"

func TestIntSet_Add(t *testing.T) {

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
		for _, item := range test.data {
			set.Add(item)
		}

		actual := true
		for _, item := range test.has {
			actual = actual && set.Has(item)
		}
		if actual != test.expected {
			t.Errorf("IntSet does not pass tests, %s", set)
		}
	}
}

func TestIntSet_Len(t *testing.T) {
	tests := []struct {
		data []int
		size int
	}{
		{
			[]int{1, 2, 3},
			3,
		},
		{
			[]int{1, 2, 3, 1, 2, 3},
			3,
		},
		{
			[]int{1, 2, 3, 4},
			4,
		},
		{
			[]int{},
			0,
		},
	}

	for _, test := range tests {
		set := &IntSet{}
		for _, item := range test.data {
			set.Add(item)
		}

		t.Logf("len %d", set.Len())

		if set.Len() != test.size {
			t.Errorf("IntSet does not pass tests, %s", set)
		}
	}
}

func TestIntSet_Remove(t *testing.T) {
	tests := []struct {
		data     []int
		removes  []int
		expected []int
		len      int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			[]int{},
			0,
		},
		{
			[]int{1, 2, 3, 1, 2, 3},
			[]int{1, 2, 3},
			[]int{},
			0,
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3, 5},
			[]int{4},
			1,
		},
		{
			[]int{1, 2, 3, 4},
			[]int{0},
			[]int{1, 2, 3, 4},
			4,
		},
	}

	for _, test := range tests {
		set := &IntSet{}
		for _, item := range test.data {
			set.Add(item)
		}

		for _, item := range test.removes {
			set.Remove(item)
		}
		t.Logf("set %s", set)
		has := true
		for _, item := range test.expected {
			has = has && set.Has(item)
		}

		if !has || set.Len() != test.len {
			t.Errorf("IntSet does not pass tests, %s", set)
		}
	}
}

func TestIntSet_ClearAll(t *testing.T) {
	tests := []struct {
		data []int
		len  int
	}{
		{
			[]int{1, 2, 3},
			0,
		},
		{
			[]int{1, 2, 3, 1, 2, 3},
			0,
		},
		{
			[]int{1, 2, 3, 4},
			0,
		},
	}

	for _, test := range tests {
		set := &IntSet{}
		for _, item := range test.data {
			set.Add(item)
		}
		t.Logf("set %s", set)
		set.ClearAll()
		t.Logf("set %s", set)

		if set.Len() != test.len {
			t.Errorf("IntSet does not pass tests, %s", set)
		}
	}
}

func TestIntSet_Copy(t *testing.T) {
	tests := []struct {
		data   []int
		add    int
		copied string
	}{
		{
			[]int{1, 2, 3},
			10,
			"{1 2 3}",
		},
		{
			[]int{1, 2, 3, 1, 2, 3},
			10,
			"{1 2 3}",
		},
		{
			[]int{1, 2, 3, 4},
			10,
			"{1 2 3 4}",
		},
	}

	for _, test := range tests {
		set := &IntSet{}
		for _, item := range test.data {
			set.Add(item)
		}
		copied := set.Copy()
		set.Add(test.add)
		t.Logf("set %s copied %s", set, copied)

		if copied.String() != test.copied {
			t.Errorf("copied and test data is incorrect %s, %s", set, copied)
		}
	}
}
