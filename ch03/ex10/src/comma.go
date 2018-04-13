package src

import (
	"bytes"
)

func comma(s string) string {
	if len(s)/3 <= 0 {
		return s
	}
	var result string
	var i int
	for i = len(s); i/3 > 0; i -= 3 {
		var buf bytes.Buffer
		buf.WriteString(s[i-3 : i])
		result = string(append(buf.Bytes(), []byte(result)...))
	}
	result = s[0:i] + result

	return result
}
