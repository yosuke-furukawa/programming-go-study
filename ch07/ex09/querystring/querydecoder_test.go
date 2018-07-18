package querystring

import "testing"

func TestDecode(t *testing.T) {
	type Query struct {
		Width  int    `url:"w"`
		Height int    `url:"h"`
		Top    string `url:"t"`
		Bottom string `url:"b"`
	}

	query := Query{}
	q := make(map[string][]string)
	q["w"] = []string{"10"}
	q["h"] = []string{"100"}
	q["t"] = []string{"green"}
	q["b"] = []string{"blue"}

	Decode(&query, q)

	if query.Width != 10 {
		t.Errorf("query width is not corrct %d", query.Width)
	}

	if query.Height != 100 {
		t.Errorf("query height is not corrct %d", query.Height)
	}

	if query.Top != "green" {
		t.Errorf("query top is not corrct %s", query.Top)
	}

	if query.Bottom != "blue" {
		t.Errorf("query bottom is not corrct %s", query.Bottom)
	}
}
