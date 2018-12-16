package ex01

import (
	"bytes"
	"testing"
)

func TestDisplay(t *testing.T) {
	tests := []struct {
		label string
		obj   interface{}
		exp   string
	}{
		{
			"string",
			"abc",
			`Display string (string):
string = "abc"
`,
		},
		{
			"movie",
			struct {
				Title, Subtitle string
				Year            int
				Color           bool
				Actor           map[string]string
				Oscars          []string
				Sequel          *string
			}{
				"Dr. Strangelove",
				"How I Learned to Stop Worrying and Love the Bomb",
				1964,
				false,
				map[string]string{
					"Dr Strangelove": "Peter Sellers",
				},
				[]string{
					"Best Actor (Nomin.)",
				},
				nil,
			},
			`Display movie (struct { Title string; Subtitle string; Year int; Color bool; Actor map[string]string; Oscars []string; Sequel *string }):
movie.Title = "Dr. Strangelove"
movie.Subtitle = "How I Learned to Stop Worrying and Love the Bomb"
movie.Year = 1964
movie.Color = false
movie.Actor["Dr Strangelove"] = "Peter Sellers"
movie.Oscars[0] = "Best Actor (Nomin.)"
movie.Sequel = nil
`,
		},
	}

	for _, test := range tests {

		out = new(bytes.Buffer)
		Display(test.label, test.obj)
		got := out.(*bytes.Buffer).String()
		if got != test.exp {
			t.Errorf("%s = %q, want = %q", test.obj, got, test.exp)
		}
	}
}

func TestDisplayWithKeyStructOrArrayMap(t *testing.T) {
	tests := []struct {
		label string
		obj   interface{}
		exp   string
	}{
		{
			"maps",
			map[struct{ label string }]string{
				{"foo"}: "bar",
				{"bar"}: "baz",
			},
			`Display maps (map[struct { label string }]string):
maps[label: "foo", ] = "bar"
maps[label: "bar", ] = "baz"
`,
		},
		{
			"mapa",
			map[[3]string]string{
				{"foo", "bar", "baz"}:    "bar",
				{"foo1", "bar1", "baz1"}: "bar1",
			},
			`Display mapa (map[[3]string]string):
mapa[array[0] = foo array[1] = bar array[2] = baz ] = "bar"
mapa[array[0] = foo1 array[1] = bar1 array[2] = baz1 ] = "bar1"
`,
		},
	}

	for _, test := range tests {

		out = new(bytes.Buffer)
		Display(test.label, test.obj)
		got := out.(*bytes.Buffer).String()
		if got != test.exp {
			t.Errorf("%s = \n\n %q,\n\n %q", test.obj, got, test.exp)
		}
	}
}
