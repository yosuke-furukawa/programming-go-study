package ex01

import (
	"bufio"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	byteLen := len(p)

	readLen := 0
	for readLen < byteLen {
		advance, token, _ := bufio.ScanWords(p, true)
		p = p[advance:]
		if token != nil {
			tokenLen := len(token)
			*c = *c + WordCounter(tokenLen)
		}
		readLen += advance
	}

	return byteLen, nil
}
