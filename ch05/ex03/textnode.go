package main

import (
	"fmt"
	"os"

	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.TextNode && n.Parent != nil && (n.Parent.Data != "script" && n.Parent.Data != "style") {
		fmt.Print(strings.Replace(n.Data, "\n", " ", -1))
	}
	if n.FirstChild != nil {
		visit(n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(n.NextSibling)
	}
}
