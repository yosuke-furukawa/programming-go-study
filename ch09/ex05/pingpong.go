package main

import (
	"log"
	"time"
)

type Ball struct {
	hits int
}

func main() {
	ping := make(chan struct{})
	pong := make(chan struct{})

	ball := &Ball{}
	go func() {
		for {
			<-ping
			ball.hits++
			log.Println(ball.hits)
			pong <- struct{}{}
		}
	}()

	go func() {
		for {
			<-pong
			ping <- struct{}{}
		}
	}()

	pong <- struct{}{}
	time.Sleep(1 * time.Second)
}
