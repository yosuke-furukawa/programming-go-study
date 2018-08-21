package main

import (
	"flag"
	"fmt"
	"log"

	"os"

	"github.com/yosuke-furukawa/programming-go-study/ch08/ex10/links"
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
	cancelChan := make(chan struct{})

	go func() {
		e := make([]byte, 1)
		os.Stdin.Read(e)
		log.Println(e)
		close(cancelChan)
	}()

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
				foundLinks := crawl(link.url, cancelChan)
				go func() {
					w := []urlDepth{}
					for _, flink := range foundLinks {
						w = append(w, urlDepth{
							flink,
							link.depth + 1,
						})
					}
					select {
					case <-cancelChan:
						return
					case worklist <- w:
					}
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
			if !seen[link.url] {
				select {
				case <-cancelChan:
					continue
				case unseenLinks <- link:
					seen[link.url] = true
				}
			}
		}
	}
}

func crawl(url string, cancelChan <-chan struct{}) []string {
	fmt.Println(url)
	list, err := links.Extract(url, cancelChan)
	if err != nil {
		log.Print(err)
	}
	return list
}
