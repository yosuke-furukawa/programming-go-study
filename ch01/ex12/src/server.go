package main

import (
	"log"
	"math/rand"
  "net"
	"net/http"
	"strconv"
	"time"
)

func server(host string, ch chan<- net.Listener, er chan<- error) {
	http.HandleFunc("/", handler)
	rand.Seed(time.Now().UTC().UnixNano())

  listener, err := net.Listen("tcp", host)

  if err != nil {
    er <- err
    return
  }
  ch <- listener

  err = http.Serve(listener, nil)
  if err != nil {
    er <- err
    return
  }
}

func main() {
	ch := make(chan net.Listener)
	er := make(chan error)
	go server("localhost:8080", ch, er)
	for {
		select {
		case listener := <-ch:
      log.Println("listen on ", listener.Addr().String())
    case err := <-er:
      log.Fatal("error ", err)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	var cycles float64
	var err error
	c := r.Form["cycles"]
	if len(c) == 0 || c[0] == "" {
		cycles = 5
	} else {
		cycles, err = strconv.ParseFloat(c[0], 64)
		if err != nil {
			log.Print(err)
		}
	}

	lissajous(w, cycles)
}
