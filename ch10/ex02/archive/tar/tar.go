package tar

import (
	atar "archive/tar"
	"strings"

	"os"

	"io"

	"github.com/yosuke-furukawa/programming-go-study/ch10/ex02/archive"
)

func init() {
	archive.RegisterFormat("tar", Search, Match)
}

func Search(filepath, query string) ([]string, error) {
	result := []string{}
	f, err := os.Open(filepath)
	if err != nil {
		return result, err
	}
	tr := atar.NewReader(f)
	if err != nil {
		return nil, err
	}
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return result, err
		}
		if strings.Contains(hdr.Name, query) {
			result = append(result, hdr.Name)
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
	buf := make([]byte, 262)
	f.Read(buf)
	return strings.HasSuffix(string(buf), "ustar")
}
