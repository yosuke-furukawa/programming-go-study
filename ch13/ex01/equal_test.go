package ex01

import "testing"

func TestEqual(t *testing.T) {
	tests := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			1,
			1,
			true,
		},
		{
			1233432893298939893,
			1,
			false,
		},
		{
			1233432893298939893,
			1233432893298939893,
			true,
		},
		{
			123.456,
			123.456,
			true,
		},
		{
			123343289329893.923,
			123343289329893.922,
			true,
		},
		{
			123343289.3298939893,
			123343289.3298939893,
			true,
		},
	}

	for _, test := range tests {
		actual := Equal(test.x, test.y)
		if actual != test.expected {
			t.Errorf("not matched %v, %v, expected %b, but actual %b", test.x, test.y, test.expected, actual)
		}
	}
}
