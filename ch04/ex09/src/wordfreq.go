package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for {
		if !input.Scan() {
			break
		}
		words[input.Text()]++
	}
	fmt.Printf("\nword\tcount\n")
	for w, n := range words {
		fmt.Printf("%q\t%d\n", w, n)
	}
}
