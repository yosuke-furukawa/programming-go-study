package src

import (
	"bytes"
	"strings"
)

func comma(s string) string {
	if len(s)/3 <= 0 {
		return s
	}
	var result string
	var i int
	dot := strings.Index(s, ".")
	if dot < 0 {
		dot = len(s)
	}
	afterDot := s[dot:]
	beforeDot := s[0:dot]
	sig := ""
	if s[0] == '-' || s[0] == '+' {
		sig = string(s[0])
		beforeDot = s[1:dot]
	}
	for i = len(beforeDot); i/3 > 0; i -= 3 {
		var buf bytes.Buffer
		buf.WriteString(",")
		buf.WriteString(beforeDot[i-3 : i])
		result = string(append(buf.Bytes(), []byte(result)...))
	}
	result = sig + beforeDot[0:i] + result + afterDot

	return result
}
