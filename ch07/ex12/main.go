package main

import (
	"html/template"
	"log"
	"net/http"

	"io/ioutil"

	"github.com/yosuke-furukawa/programming-go-study/ch07/ex12/db"
)

func main() {
	tmpl, err := ioutil.ReadFile("./template/database.html")
	if err != nil {
		panic(err)
	}
	dbtmpl := template.Must(template.New("database").Parse(string(tmpl)))
	db := db.Database{db.Prices{"shoes": 50, "socks": 5}, dbtmpl}
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/insert", db.Insert)
	http.HandleFunc("/delete", db.Delete)
	http.HandleFunc("/price", db.Price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
