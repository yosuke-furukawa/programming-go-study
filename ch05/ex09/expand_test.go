package ex09

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	tests := []struct {
		actual string
		expect string
		repl   func(s string) string
	}{
		{
			"hoge hoge $furukawa san",
			"hoge hoge FURUKAWA san",
			func(s string) string { return strings.ToUpper(s) },
		},
		{
			"hoge hoge $furukawa $san",
			"hoge hoge FURUKAWA SAN",
			func(s string) string { return strings.ToUpper(s) },
		},
		{
			"hoge $ho$ge $furukawa $san",
			"hoge HOGE FURUKAWA SAN",
			func(s string) string { return strings.ToUpper(s) },
		},
	}

	for _, test := range tests {
		result := Expand(test.actual, test.repl)
		if test.expect != result {
			t.Errorf("expect is %s, but actual is %s", test.expect, result)
		}
	}
}
