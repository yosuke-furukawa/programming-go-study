package ex02

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	tests := []struct {
		data  []string
		count int64
	}{
		{
			[]string{"hello", "world", "world"},
			15,
		},
		{
			[]string{"hello", "world", "world", "foo", "bar", "baz"},
			24,
		},
	}

	for _, test := range tests {
		writer, writtenCount := CountingWriter(bytes.NewBufferString(""))
		for _, str := range test.data {
			writer.Write([]byte(str))
		}

		if *writtenCount != test.count {
			t.Errorf("CountingWriter is incorrect, count = %d, expected = %d", *writtenCount, test.count)
		}
	}
}
