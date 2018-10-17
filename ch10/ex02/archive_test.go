package ex02

import (
	"testing"

	"github.com/yosuke-furukawa/programming-go-study/ch10/ex02/archive"
	_ "github.com/yosuke-furukawa/programming-go-study/ch10/ex02/archive/tar"
	_ "github.com/yosuke-furukawa/programming-go-study/ch10/ex02/archive/zip"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		filepath string
		query    string
		expected []string
	}{
		{
			"./fixtures/test.zip",
			"yosuke",
			[]string{"test/yosuke.txt"},
		},
		{
			"./fixtures/test.zip",
			"test",
			[]string{"test/yosuke.txt", "test/test.txt"},
		},
		{
			"./fixtures/test.tar",
			"yosuke",
			[]string{"./test/yosuke.txt"},
		},
		{
			"./fixtures/test.tar",
			"test",
			[]string{"./test/yosuke.txt", "./test/test.txt"},
		},
	}

	for _, test := range tests {
		results, err := archive.Search(test.filepath, test.query)
		if err != nil {
			t.Errorf("error is thrown, %v", err)
		}

		if len(results) == 0 {
			t.Errorf("results list is empty", len(results))
		}
		for _, expected := range test.expected {
			if !contains(expected, results) {
				t.Errorf("results list is not correct %s, expected %s", results, expected)
			}
		}
	}
}

func contains(expected string, results []string) bool {
	for _, result := range results {
		if result == expected {
			return true
		}
	}
	return false
}
