package main

import (
	"flag"
	"fmt"
	"log"

	"net/url"

	"os"

	"path/filepath"

	"github.com/yosuke-furukawa/programming-go-study/ch08/ex06/links"
	"github.com/yosuke-furukawa/programming-go-study/ch08/ex07/fetch"
)

type urlDepth struct {
	url   string
	depth int
}

func main() {
	depth := flag.Int("depth", 0, "depth")
	savePath := *(flag.String("path", "/tmp", "path"))
	if !filepath.IsAbs(savePath) {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		savePath = cwd + savePath
	}

	u := flag.String("url", "", "url")
	requestUrl, err := url.Parse(*u)
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()

	worklist := make(chan []urlDepth)
	unseenLinks := make(chan urlDepth)

	go func() {
		worklist <- []urlDepth{
			{
				*u,
				0,
			},
		}
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link.url, savePath)
				go func() {
					w := []urlDepth{}
					for _, flink := range foundLinks {
						w = append(w, urlDepth{
							flink,
							link.depth + 1,
						})
					}
					worklist <- w
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if link.depth >= *depth {
				return
			}
			parsedUrl, err := url.Parse(link.url)
			if err != nil {
				// url 不正はスキップ
				continue
			}

			if parsedUrl.Host != requestUrl.Host {
				// url domain が一致しなかったらskip
				continue
			}
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(url, savePath string) []string {
	fmt.Println(url)
	fetch.Fetch(url, savePath)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
