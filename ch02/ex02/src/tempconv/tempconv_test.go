package tempconv

import "testing"

func TestCToK(t *testing.T) {
	k := CToK(Celsius(AbsoluteZeroC))
	if k != Kelvin(0) {
		t.Errorf("Kelvin is not 0 actual: %f, expected: %f", k, Kelvin(0))
	}
}

func TestKToC(t *testing.T) {
	c := KToC(Kelvin(0))
	expected := Celsius(-273.15)
	if c != expected {
		t.Errorf("Kelvin is not correct actual: %f, expected: %f", c, expected)
	}
}

func TestFToK(t *testing.T) {
	f := FToK(Fahrenheit(100))
	expected := Kelvin(310.9277777777778)
	if f != expected {
		t.Errorf("Kelvin is not correct actual: %f, expected: %f", f, expected)
	}
}

func TestKToF(t *testing.T) {
	f := KToF(Kelvin(100))
	expected := Fahrenheit(-279.67)
	if f != expected {
		t.Errorf("Kelvin is not correct actual: %f, expected: %f", f, expected)
	}
}
