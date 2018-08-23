package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yosuke-furukawa/programming-go-study/ch08/ex06/links"
)

type urlDepth struct {
	url   string
	depth int
}

func main() {
	depth := flag.Int("depth", 0, "depth")
	url := flag.String("url", "", "url")
	flag.Parse()

	worklist := make(chan []urlDepth)
	unseenLinks := make(chan urlDepth)

	go func() {
		worklist <- []urlDepth{
			{
				*url,
				0,
			},
		}
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link.url)
				go func(link urlDepth) {
					w := []urlDepth{}
					for _, flink := range foundLinks {
						w = append(w, urlDepth{
							flink,
							link.depth + 1,
						})
					}
					worklist <- w
				}(link)
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if link.depth >= *depth {
				return
			}
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
