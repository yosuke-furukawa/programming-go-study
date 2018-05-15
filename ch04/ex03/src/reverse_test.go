package src

import (
	"testing"
)

func TestReverse(t *testing.T) {
	a := [6]int{0, 1, 2, 3, 4, 5}
	Reverse(&a)
	t.Logf("%#v", a)
	if a != [6]int{5, 4, 3, 2, 1, 0} {
		t.Errorf("reverse is not working correctly, %#v", a)
	}
}
