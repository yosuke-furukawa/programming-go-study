package weightconv

import "fmt"

// Pound is a type of weight
type Pound float64
// KiloGram is a type of weight
type KiloGram float64

// Constants
const (
	coeffience = 0.45359237
)

// PToK Pound To Kilogram
func PToK(p Pound) KiloGram {
	return KiloGram(p * coeffience)
}

// KToP Kilogram To Pound
func KToP(k KiloGram) Pound {
	return Pound(k / coeffience)
}

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}

func (k KiloGram) String() string {
	return fmt.Sprintf("%gkg", k)
}
