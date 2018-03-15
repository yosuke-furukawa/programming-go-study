package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/yosuke-furukawa/programming-go-study/ch03/ex04/src/querystring"
	"github.com/yosuke-furukawa/programming-go-study/ch03/ex04/src/surface"
	"github.com/yosuke-furukawa/programming-go-study/ch03/ex04/src/types"
)

func server(host string, ch chan<- net.Listener, er chan<- error) {
	http.HandleFunc("/", surfaceHandler)

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
			log.Println("open http://localhost:8080/?w=2000&h=2000&t=red&b=green&f=black&c=300")
		case err := <-er:
			log.Fatal("error ", err)
		}
	}
}

func surfaceHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error: do not parse form")
		return
	}

	query := types.Query{
		Stroke: "grey",
		Fill:   "white",
		Width:  600,
		Height: 800,
		Top:    "red",
		Bottom: "blue",
		Cells:  100,
	}

	querystring.Decode(&query, r.Form)

	w.Header().Set("Content-Type", "image/svg+xml")
	surface.Surface(w, query)
}
