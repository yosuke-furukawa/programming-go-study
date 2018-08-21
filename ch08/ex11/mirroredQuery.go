package main

import (
	"log"
	"net/http"
	"os"
)

var done = make(chan struct{})

func mirroredQuery(reqs []string) string {
	responses := make(chan string, len(reqs))

	for _, req := range reqs {
		go func(req string) {
			responses <- request(req, done)
		}(req)
	}
	return <-responses
}

func request(hostname string, done chan struct{}) (response string) {
	req, err := http.NewRequest("GET", "http://"+hostname, nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	log.Println(resp.Status)
	log.Println(resp.ContentLength)
	if err != nil {
		log.Println(err)
		return ""
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return ""
	}

	resp.Body.Close()
	close(done)
	return hostname
}

func main() {
	log.Println(mirroredQuery(os.Args[1:]))
}
