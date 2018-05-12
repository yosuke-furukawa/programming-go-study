package src

import (
	"testing"
)

func TestCheckBit(t *testing.T) {
	t.Parallel()
	diff := CheckBit([32]byte{byte(0)}, [32]byte{byte(5)})
	if diff != 2 {
		t.Errorf("diff is not 2 %d", diff)
	}
}

func TestCheckShaBit(t *testing.T) {
	t.Parallel()
	diff := CheckShaBit("x", "X")
	if diff != 125 {
		t.Errorf("diff is not 125 %d", diff)
	}
}
