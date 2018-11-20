package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func count(reader io.Reader) (map[rune]int, map[string]int, [utf8.UTFMax + 1]int, int) {
	counts := make(map[rune]int)
	kinds := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(reader)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		if unicode.IsLetter(r) {
			kinds["letter"]++
		} else if unicode.IsNumber(r) {
			kinds["number"]++
		} else if unicode.IsControl(r) {
			kinds["control"]++
		} else if unicode.IsSpace(r) {
			kinds["space"]++
		}
	}
	return counts, kinds, utflen, invalid
}

func main() {
	counts, kinds, utflen, invalid := count(os.Stdin)

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	fmt.Printf("\nkind\tcount\n")
	for kind, n := range kinds {
		fmt.Printf("%s\t%d\n", kind, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
