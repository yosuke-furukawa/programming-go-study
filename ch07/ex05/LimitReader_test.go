package ex05

import (
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	tests := []struct {
		actual      string
		size        int64
		expected    string
		expectedErr error
	}{
		{
			"hello yosuke furukawa",
			5,
			"hello",
			nil,
		},
		{
			"hello",
			0,
			"",
			io.EOF,
		},
		{
			"hello",
			-1,
			"",
			io.EOF,
		},
		{
			"hello",
			2000,
			"hello",
			nil,
		},
	}

	for _, test := range tests {
		reader := LimitReader(strings.NewReader(test.actual), test.size)
		b := make([]byte, 1024)
		num, err := reader.Read(b)

		if err != test.expectedErr {
			t.Errorf("error is thrown %s", err)
		}
		actual := string(b[:num])

		if actual != test.expected {
			t.Errorf("expected %s but actual %s", test.expected, actual)
		}
	}
}
