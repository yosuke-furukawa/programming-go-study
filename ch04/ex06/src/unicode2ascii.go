package src

import (
	"unicode"
	"unicode/utf8"
)

func UnicodeToAscii(bs []byte) []byte {
	var result []byte
	var size int
	beforeSpace := false
	for i := 0; i < len(bs); i += size {
		r, s := utf8.DecodeRune(bs[i:])
		if !unicode.IsSpace(r) {
			result = append(result, bs[i:i+s]...)
			beforeSpace = false
		} else {
			if !beforeSpace {
				result = append(result, byte(0x20))
				beforeSpace = true
			}
		}
		size = s
	}
	return result
}
