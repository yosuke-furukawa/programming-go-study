package tempconv

import "fmt"

// Celsius is a type of degree
type Celsius float64
// Fahrenheit is a type of degree
type Fahrenheit float64
// Kelvin is a type of degree
type Kelvin float64

// Constants
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
	ZeroK Kelvin = 0
)

// CToF Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9/5 + 32)
}

// FToC Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5/9)
}

// CToK Celsius to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

// KToC Kelvin to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k + Kelvin(AbsoluteZeroC))
}

// FToK Fahrenheit to Kelvin
func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f + 459.67) * 5/9)
}

// KToF Kelvin to Fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(k * 9/5 - 459.67)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
