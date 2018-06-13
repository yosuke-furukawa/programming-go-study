package ex15

import "testing"

func TestMax1(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			8,
		},
		{
			[]int{8, 7, 6, 5, 5, 6, 7, 8},
			8,
		},
		{
			[]int{8},
			8,
		},
	}

	for _, test := range tests {
		max := Max1(test.data...)
		if max != test.expected {
			t.Errorf("expected %d, but actual %d", test.expected, max)
		}
	}
}
func TestMax2(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			8,
		},
		{
			[]int{8, 7, 6, 5, 5, 6, 7, 8},
			8,
		},
		{
			[]int{8},
			8,
		},
	}

	for _, test := range tests {
		max := Max2(test.data[0], test.data[1:]...)
		if max != test.expected {
			t.Errorf("expected %d, but actual %d", test.expected, max)
		}
	}
}

func TestMin1(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			1,
		},
		{
			[]int{8, 7, 6, 5, 5, 6, 7, 8},
			5,
		},
		{
			[]int{8},
			8,
		},
	}

	for _, test := range tests {
		min := Min1(test.data...)
		if min != test.expected {
			t.Errorf("expected %d, but actual %d", test.expected, min)
		}
	}
}
func TestMin2(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			1,
		},
		{
			[]int{8, 7, 6, 5, 5, 6, 7, 8},
			5,
		},
		{
			[]int{8},
			8,
		},
	}

	for _, test := range tests {
		min := Min2(test.data[0], test.data[1:]...)
		if min != test.expected {
			t.Errorf("expected %d, but actual %d", test.expected, min)
		}
	}
}
