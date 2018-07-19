package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"sort"

	"github.com/yosuke-furukawa/programming-go-study/ch07/ex09/querystring"
	"github.com/yosuke-furukawa/programming-go-study/ch07/ex09/sorting"
	"github.com/yosuke-furukawa/programming-go-study/ch07/ex09/types"
)

var tracks = []*sorting.Track{
	{
		"Go",
		"Delialah",
		"From the Roots up",
		2012,
		sorting.Length("3m38s"),
	},
	{
		"Go",
		"Moby",
		"Moby",
		1992,
		sorting.Length("3m38s"),
	},
	{
		"Go Ahead",
		"Alicia Keys",
		"As I Am",
		2007,
		sorting.Length("4m36s"),
	},
	{
		"Ready 2 Go",
		"Martin Solveig",
		"Smash",
		2011,
		sorting.Length("4m24s"),
	},
}

func sortHandler(tmpl *template.Template, table *sorting.SortTable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Error: do not parse form")
			return
		}

		query := types.Query{
			Key: "",
		}

		querystring.Decode(&query, r.Form)
		table.Select(query.Key)
		sort.Sort(table)
		if err := tmpl.Execute(w, table.T); err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "cannot exec templ")
			return
		}
	}
}

func server(host string, ch chan<- net.Listener, er chan<- error) {
	table := &sorting.SortTable{T: tracks, FirstKey: "", SecondKey: ""}
	templ, err := ioutil.ReadFile("./template/sorttable.html")
	if err != nil {
		er <- err
		return
	}
	sorttable := template.Must(template.New("sorttable").Parse(string(templ)))
	http.HandleFunc("/", sortHandler(sorttable, table))

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
