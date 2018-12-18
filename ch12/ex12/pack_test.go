package ex12

import "testing"

func TestPack(t *testing.T) {
	tests := []struct {
		Labels    []string `http:"l"`
		MaxResult int      `http:"max"`
		Exact     bool     `http:"x"`
	}{
		{
			[]string{"foo", "bar", "baz"},
			100,
			true,
		},
		{
			[]string{"foo"},
			0,
			false,
		},
	}

	expects := []string{
		"l=foo&l=bar&l=baz&max=100&x=true",
		"l=foo&max=0&x=false",
	}

	for i, test := range tests {
		url := Pack(&test)
		t.Logf(url)
		if url != expects[i] {
			t.Errorf("url is not match actual %s, expect %s", url, expects[i])
		}
	}
}
