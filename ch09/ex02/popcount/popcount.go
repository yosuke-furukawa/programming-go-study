package popcount

import "sync"

var initializeOnce sync.Once
var pc [256]byte
var pc2 [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func initialize() {
	for i := range pc2 {
		pc2[i] = pc2[i/2] + byte(i&1)
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

// PopCountLoop bit count using loop but lazy initialize
func PopCountLoopLazy(x uint64) int {
	initializeOnce.Do(initialize)
	var b byte
	for i := 0; i < 8; i++ {
		b += pc2[byte(x>>uint(i*8))]
	}
	return int(b)
}
