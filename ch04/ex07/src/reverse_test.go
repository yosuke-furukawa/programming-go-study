package src

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		actual string
		expect string
	}{
		{
			"a b c d e",
			"e d c b a",
		},
		{
			"こんにちは 世界",
			"界世 はちにんこ",
		},
		{
			"Hello World　こんにちは世界",
			"界世はちにんこ　dlroW olleH",
		},
	}

	for _, test := range tests {
		bs := []byte(test.actual)
		Reverse(bs)
		answer := string(bs)
		if answer != test.expect {
			t.Errorf("data is not correct %#v", answer)
		}
	}
}
