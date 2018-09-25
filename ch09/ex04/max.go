package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

type pipeline struct {
	nextChan chan struct{}
	next     *pipeline
}

func main() {
	max := os.Args[1]
	m, err := strconv.Atoi(max)
	if err != nil {
		log.Fatal(err)
	}
	p := &pipeline{
		nil,
		nil,
	}
	for i := 0; i < m; i++ {
		p = &pipeline{
			make(chan struct{}),
			p,
		}
	}

	start := time.Now()

	for {
		go func(p *pipeline) {
			if p.nextChan == nil {
				return
			}
			close(p.nextChan)
		}(p)
		if p.nextChan == nil {
			break
		}
		<-p.nextChan
		if p.next == nil {
			break
		}
		p = p.next
	}
	log.Println("end %d msec, %d goroutines", time.Since(start)*time.Millisecond, m)

}
