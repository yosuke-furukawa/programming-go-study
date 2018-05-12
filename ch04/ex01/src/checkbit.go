package src

import (
	"crypto/sha256"
)

func CheckShaBit(a, b string) int {
	c1 := sha256.Sum256([]byte(a))
	c2 := sha256.Sum256([]byte(b))
	return CheckBit(c1, c2)
}

func CheckBit(c1, c2 [32]byte) int {
	diff := 0

	for i, _ := range c1 {
		b1, b2 := c1[i], c2[i]
		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))
			if b1&mask != b2&mask {
				diff++
			}
		}
	}

	return diff
}
