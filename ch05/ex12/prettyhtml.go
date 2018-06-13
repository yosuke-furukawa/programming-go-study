package ex12

import (
	"fmt"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node) string) (result string) {
	if pre != nil {
		result += pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result += forEachNode(c, pre, post)
	}

	if post != nil {
		result += post(n)
	}
	return
}

func PrettyHTML(n *html.Node) string {
	depth := 0
	return forEachNode(n, func(n *html.Node) (result string) {
		if n.Type == html.ElementNode {
			attr := ""
			for _, a := range n.Attr {
				attr += fmt.Sprintf(" %s=\"%s\"", a.Key, a.Val)
			}
			if n.FirstChild == nil {
				result = fmt.Sprintf("%*s<%s%s />\n", depth*2, "", n.Data, attr)
			} else {
				result = fmt.Sprintf("%*s<%s%s>\n", depth*2, "", n.Data, attr)
			}
			depth++
		} else if n.Type == html.TextNode {
			result = fmt.Sprintf("%*s%s\n", depth*2, "", n.Data)
		}
		return
	}, func(n *html.Node) (result string) {
		if n.Type == html.ElementNode {
			depth--
			if n.FirstChild == nil {
				return
			}
			result = fmt.Sprintf("%*s</%s>\n", depth*2, "", n.Data)
		}
		return
	})
}
