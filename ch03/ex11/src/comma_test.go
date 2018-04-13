package src

import (
	"testing"
)

func TestComma(t *testing.T) {
	t.Parallel()
	result := comma("12345678")
	if result != "12,345,678" {
		t.Errorf("comma should have proper position %s", result)
	}
	result = comma("10")
	if result != "10" {
		t.Errorf("comma should have proper position %s", result)
	}
}

func TestCommaWithDot(t *testing.T) {
	t.Parallel()
	result := comma("12345678.00000")
	if result != "12,345,678.00000" {
		t.Errorf("comma should have proper position %s", result)
	}
	result = comma("-12345678.00000")
	if result != "-12,345,678.00000" {
		t.Errorf("comma should have proper position %s", result)
	}
}
