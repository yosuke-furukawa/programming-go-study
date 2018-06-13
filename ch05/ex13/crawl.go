package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"os"

	"strings"

	"path"

	"io"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func download(u string) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	item, _ := url.Parse(u)
	base := path.Base(item.Path)
	ext := path.Ext(item.Path)
	downloaddir := "./download/"
	downloadpath := path.Join(downloaddir, item.Path)

	if strings.HasSuffix(base, "/") {
		downloadpath = path.Join(downloadpath, "index.html")
	} else if ext == "" {
		downloadpath = path.Join(downloadpath, "index.html")
	}
	log.Println(downloadpath)

	err = os.MkdirAll(path.Dir(downloadpath), os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(downloadpath)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func crawl(url string) []string {
	if !strings.HasPrefix(url, hostname) {
		return nil
	}
	log.Println(url)
	err := download(url)
	if err != nil {
		log.Print(err)
	}
	list, err := extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

var hostname string

func main() {
	hostname = os.Args[1]
	breadthFirst(crawl, []string{hostname})
}
