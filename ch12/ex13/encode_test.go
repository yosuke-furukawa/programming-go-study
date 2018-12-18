package ex13

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	type Movie struct {
		Title     string            `sexpr:"title"`
		Subtitle  string            `sexpr:"subtitle"`
		Year      int               `sexpr:"year"`
		Actor     map[string]string `sexpr:"actor"`
		Oscars    []string          `sexpr:"year"`
		Sequel    *string           `sexpr:"sequel"`
		Published bool              `sexpr:"published"`
		Private   bool              `sexpr:"private"`
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		Published: true,
		Private:   false,
	}

	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	var m Movie
	decoder := NewDecoder(bytes.NewBuffer(data))
	decoder.Decode(&m)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if !reflect.DeepEqual(m, strangelove) {
		t.Errorf("marshal is not equal %v, %v", m, strangelove)
	}

}
