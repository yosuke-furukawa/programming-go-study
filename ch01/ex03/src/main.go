package main

import (
	"fmt"
	"os"
	"strings"
)

func Echo1(args []string) string {
	s, sep := "", ""
	for _, arg := range args[0:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func Echo2(args []string) string {
	return strings.Join(args[0:], " ")
}

func main() {
	fmt.Println(Echo1(os.Args[1:]))
	fmt.Println(Echo2(os.Args[1:]))
}
