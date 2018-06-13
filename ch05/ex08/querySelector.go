package ex08

import (
	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	done := false
	if pre != nil {
		done = pre(n)
	}

	if done {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, pre, post)
		if node != nil {
			return node
		}
	}

	if post != nil {
		done = post(n)
	}

	if done {
		return n
	}

	return nil
}

func startElement(id string) func(n *html.Node) (found bool) {
	return func(n *html.Node) (found bool) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					found = true
					return
				}
			}
		}
		return
	}
}

func endElement(id string) func(n *html.Node) (found bool) {
	return func(n *html.Node) (found bool) {
		return
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, startElement(id), endElement(id))
}
