package ex05

import "testing"

func Test_CountWordsAndImages(t *testing.T) {
	words, img, err := CountWordsAndImages("https://www.google.com")

	if err != nil {
		t.Fatal(err)
	}
	if img < 1 {
		t.Fatal("image not found in google")
	}
	if words < 1 {
		t.Fatal("word not found in google")
	}
	t.Logf("words %d, img %d", words, img)
}
