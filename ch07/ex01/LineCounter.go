package ex01

import (
	"bufio"
	"strings"
)

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(strings.NewReader(string(p)))
	for input.Scan() {
		*c++
	}
	return len(p), nil
}
