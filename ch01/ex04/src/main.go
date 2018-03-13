package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Dup usage: args []string
func Dup(args []string) (map[string]map[string]int, error) {
	counts := make(map[string]map[string]int)
	for _, filename := range args {
		counts[filename] = make(map[string]int)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[filename][line]++
		}
	}
	return counts, nil
}

func main() {
	counts, err := Dup(os.Args[1:])

	if err != nil {
		fmt.Fprintf(os.Stderr, "dup: %v\n", err)
	}

	for filename, count := range counts {
		for line, n := range count {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, filename)
			}
		}
	}
}
