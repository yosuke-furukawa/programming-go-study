package popcount

import "math/bits"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountLoop bit count using loop
func PopCountLoop(x uint64) int {
	var b byte
	for i := 0; i < 8; i++ {
		b += pc[byte(x>>uint(i*8))]
	}
	return int(b)
}

// PopCountShift64 bit count using shift operation
func PopCountShift64(x uint64) int {
	result := 0
	for i := 0; i < 64; i++ {
		if x>>uint(i)&1 == 1 {
			result++
		}
	}
	return result

}

// PopCount bit count using map
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountAndOne bit count using And Operation
func PopCountAndOne(x uint64) int {
	result := 0
	for ; x > 0; x &= x - 1 {
		result++
	}
	return result
}

// PopCountHackersDelight version
func PopCountHackersDelight(x uint64) int {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

// PopCountOnesCount using bits.OnesCount
func PopCountOnesCount(x uint64) int {
	return int(bits.OnesCount(uint(x)))
}

// PopCountCPU bit count using POPCNTL
func PopCountCPU(x uint64) int
