package ex02

import (
	"testing"
)

func TestIsCyclic(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}

	a, b, c, d := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}, &link{value: "d"}
	a.tail, b.tail, c.tail = b, a, c

	tests := []struct {
		x      interface{}
		cyclic bool
	}{
		{
			a,
			true,
		},
		{
			b,
			true,
		},
		{
			c,
			true,
		},
		{
			d,
			false,
		},
	}

	for _, test := range tests {
		if IsCyclic(test.x) != test.cyclic {
			t.Errorf("not expected, IsCyclic(%v) = %s", test.x, IsCyclic(test.x))
		}
	}
}
