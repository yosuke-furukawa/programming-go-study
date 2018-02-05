package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1(args []string) string {
	s, sep := "", ""
	for _, arg := range args[0:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo2(args []string) string {
	return strings.Join(args[0:], " ")
}

func main() {
	fmt.Println(echo1(os.Args[1:]))
	fmt.Println(echo2(os.Args[1:]))
}
