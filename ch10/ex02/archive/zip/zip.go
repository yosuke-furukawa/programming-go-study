package zip

import (
	azip "archive/zip"
	"strings"

	"os"

	"github.com/yosuke-furukawa/programming-go-study/ch10/ex02/archive"
)

func init() {
	archive.RegisterFormat("zip", Search, Match)
}

func Search(filepath, query string) ([]string, error) {
	result := []string{}
	r, err := azip.OpenReader(filepath)
	defer r.Close()
	if err != nil {
		return nil, err
	}
	for _, f := range r.File {
		if strings.Contains(f.Name, query) {
			result = append(result, f.Name)
		}
	}
	return result, nil
}

func Match(filepath string) bool {
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		return false
	}
	buf := make([]byte, 2)
	f.Read(buf)
	return strings.HasPrefix(string(buf), "PK")
}
