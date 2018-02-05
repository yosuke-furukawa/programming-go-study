package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"
)

func server(host string, ch chan<- net.Listener, er chan<- error) {
	http.HandleFunc("/", lissajousHandler)
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

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: do not parse form")
	}

	var cycles float64
	var err error
	c := r.Form["cycles"]
	if len(c) == 0 || c[0] == "" {
		cycles = 5
	} else {
		cycles, err = strconv.ParseFloat(c[0], 64)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Error: do not parse cycles")
			log.Print(err)
			return
		}
	}

	lissajous(w, cycles)
}
