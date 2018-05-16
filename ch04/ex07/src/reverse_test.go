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

var answer string

func BenchmarkReverse(b *testing.B) {
	tests := []struct {
		actual     string
		expect     string
		actualbyte []byte
	}{
		{
			"a b c d e",
			"e d c b a",
			nil,
		},
		{
			"こんにちは 世界",
			"界世 はちにんこ",
			nil,
		},
		{
			"Hello World　こんにちは世界",
			"界世はちにんこ　dlroW olleH",
			nil,
		},
	}

	for _, test := range tests {
		test.actualbyte = []byte(test.actual)
	}

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Reverse(test.actualbyte)
			answer = string(test.actualbyte)
		}
	}
}

func BenchmarkReverseDoubleArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
			answer = string(ReverseDoubleArray(bs))
		}
	}
}

func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
			answer = ReverseString(test.actual)
		}
	}
}
