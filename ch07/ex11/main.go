package main

import (
	"log"
	"net/http"

	"github.com/yosuke-furukawa/programming-go-study/ch07/ex11/db"
)

func main() {
	db := db.Database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/insert", db.Insert)
	http.HandleFunc("/delete", db.Delete)
	http.HandleFunc("/price", db.Price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
