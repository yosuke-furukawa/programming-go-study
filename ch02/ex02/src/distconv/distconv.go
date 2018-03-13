package distconv

import "fmt"

// Meter is a type of distance
type Meter float64

// Feet is a type of distance
type Feet float64

// Constants
const (
	coeffience = 3.2808
)

// MToF Meter To Feet
func MToF(m Meter) Feet {
	return Feet(m * coeffience)
}

// FToM Feet To Meter
func FToM(f Feet) Meter {
	return Meter(f / coeffience)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gM", m)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gF", f)
}
