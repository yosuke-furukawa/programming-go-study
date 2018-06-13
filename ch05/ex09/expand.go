package ex09

import (
	"regexp"
)

func Expand(s string, f func(string) string) string {
	pattern := regexp.MustCompile(`\$\w+`)
	wrapper := func(arg string) string {
		return f(arg[1:])
	}
	return pattern.ReplaceAllStringFunc(s, wrapper)
}
