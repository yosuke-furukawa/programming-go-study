package src

import (
	"unicode/utf8"
)

func reverse(bs []byte) {
	for i, j := 0, len(bs)-1; i < len(bs)/2; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
}

func Reverse(bs []byte) {
	var size int
	// 一旦文字だけ先に反転する
	for i := 0; i < len(bs); i += size {
		_, size = utf8.DecodeRune(bs[i:])
		reverse(bs[i : i+size])
	}
	// 全体を反転する
	reverse(bs)
}
