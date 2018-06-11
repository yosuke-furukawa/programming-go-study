package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	for k, v := range visit(map[string]int{}, doc) {
		fmt.Printf("%s : %v \n", k, v)
	}
}

func visit(m map[string]int, n *html.Node) map[string]int {
	switch n.Data {
	case "div":
		m["div"]++
	case "span":
		m["span"]++
	case "p":
		m["p"]++
	}
	if n.FirstChild != nil {
		m = visit(m, n.FirstChild)
	}
	if n.NextSibling != nil {
		m = visit(m, n.NextSibling)
	}
	return m
}
