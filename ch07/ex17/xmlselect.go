package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", stringify(stack), tok)
			}
		}
	}
}

func stringify(x []xml.StartElement) string {
	result := ""
	for _, elem := range x {
		result += fmt.Sprintf("  %s", elem.Name.Local)
	}
	return result
}

func containsAll(x []xml.StartElement, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0].Name.Local == y[0] {
			y = y[1:]
		}
		for _, attr := range x[0].Attr {
			if attr.Value == y[0] {
				y = y[1:]
				break
			}
		}
		x = x[1:]
	}
	return false
}
