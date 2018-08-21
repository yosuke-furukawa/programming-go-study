package fetch

import (
	"log"
	"net/url"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) *html.Node {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, pre, post)
		if node != nil {
			return node
		}
	}

	if post != nil {
		post(n)
	}

	return nil
}

func startElement(hostname, saveDir string) func(n *html.Node) {
	return func(n *html.Node) {
		if n.Type == html.ElementNode {
			for i, a := range n.Attr {
				if a.Key == "href" {
					u, err := url.Parse(a.Val)
					if err != nil {
						log.Fatal(err)
					}
					if u.Hostname() == "" {
						a.Val = saveDir + u.Path
						n.Attr[i].Val = saveDir + u.Path
					}
					if u.Hostname() == hostname {
						a.Val = saveDir + u.Path
						n.Attr[i].Val = saveDir + u.Path
					}
					log.Printf("a href = %v ", a.Val)
					return
				}
			}
		}
		return
	}
}

func endElement(hostname, saveDir string) func(n *html.Node) {
	return func(n *html.Node) {
		return
	}
}

func ReplaceTag(doc *html.Node, hostname, saveDir string) *html.Node {
	return forEachNode(doc, startElement(hostname, saveDir), endElement(hostname, saveDir))
}
