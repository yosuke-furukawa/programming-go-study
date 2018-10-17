package archive

import (
	"fmt"
)

type format struct {
	name   string
	search func(filepath, query string) ([]string, error)
	match  func(filepath string) bool
}

var formats []format

func RegisterFormat(name string, search func(filepath, query string) ([]string, error), match func(string) bool) {
	formats = append(formats, format{name, search, match})
}

func Search(filepath, query string) ([]string, error) {
	for _, format := range formats {
		if format.match(filepath) {
			return format.search(filepath, query)
		}
	}
	return nil, fmt.Errorf("no such format")
}
