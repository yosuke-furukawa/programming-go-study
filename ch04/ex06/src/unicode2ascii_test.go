package src

import "testing"

func TestUnicodeToAscii(t *testing.T) {
	tests := []struct {
		actual string
		expect string
	}{
		{
			"a b c d\t e",
			"a b c d e",
		},
		{
			"   a b c d\t e",
			" a b c d e",
		},
		{
			"   \na b c d\t e",
			" a b c d e",
		},
		{
			"a b c　d\t e",
			"a b c d e",
		},
		{
			"a\rb　　　c\nd\t\ve",
			"a b c d e",
		},
	}

	for _, test := range tests {
		answer := string(UnicodeToAscii([]byte(test.actual)))
		if answer != test.expect {
			t.Errorf("data is not correct %#v", answer)
		}
	}
}
