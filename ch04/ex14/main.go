package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/yosuke-furukawa/programming-go-study/ch04/ex14/querystring"
	"github.com/yosuke-furukawa/programming-go-study/ch04/ex14/types"
	"github.com/yosuke-furukawa/programming-go-study/ch04/github"
)

func sortHandler(tmpl *template.Template) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Error: do not parse form")
			return
		}

		query := types.Query{
			Query: "node",
		}

		querystring.Decode(&query, r.Form)
		issues, err := github.SearchIssues([]string{query.Query})
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "cannot search")
			return
		}
		if err := tmpl.Execute(w, issues); err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "cannot exec templ")
			return
		}
	}
}

func server(host string, ch chan<- net.Listener, er chan<- error) {
	templ, err := ioutil.ReadFile("./template/issues.html")
	if err != nil {
		er <- err
		return
	}
	issueList := template.Must(template.New("issuelist").Parse(string(templ)))
	http.HandleFunc("/", issueHandler(issueList))

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
