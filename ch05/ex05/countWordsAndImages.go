package ex05

import (
	"bufio"
	"fmt"
	"net/http"

	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "img":
			images++
		}
	}
	if n.Type == html.TextNode {
		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			words++
		}
	}
	if n.FirstChild != nil {
		w, i := countWordsAndImages(n.FirstChild)
		words += w
		images += i
	}
	if n.NextSibling != nil {
		w, i := countWordsAndImages(n.NextSibling)
		words += w
		images += i
	}
	return
}
