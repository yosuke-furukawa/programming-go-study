package ex08

import (
	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) (elements []*html.Node) {
	if pre != nil {
		if pre(n) {
			elements = append(elements, n)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes := forEachNode(c, pre, post)
		elements = append(elements, nodes...)
	}

	if post != nil {
		if post(n) {
			elements = append(elements, n)
		}
	}

	return
}

func startElement(tags ...string) func(n *html.Node) (found bool) {
	return func(n *html.Node) (found bool) {
		if n.Type == html.ElementNode {
			for _, tag := range tags {
				if tag == n.Data {
					found = true
				}
			}
		}
		return
	}
}

func endElement(tags ...string) func(n *html.Node) (found bool) {
	return func(n *html.Node) (found bool) {
		return
	}
}

func ElementsByName(doc *html.Node, names ...string) []*html.Node {
	return forEachNode(doc, startElement(names...), endElement(names...))
}
