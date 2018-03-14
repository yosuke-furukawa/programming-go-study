// +build amd64,!gccgo,!appengine

TEXT ·PopCountCPU(SB), 4, $0-16
    MOVQ x+0(FP), R8
    POPCNTL R8, R8
    MOVQ R8, ret+8(FP)
    RET
