package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}

	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item := query.Get("item")
	price := query.Get("price")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "item is empty\n")
	}
	if price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "price is empty\n")
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "cannot parse price %s\n", price)
	}
	db[item] = dollars(p)
	w.WriteHeader(http.StatusCreated)
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item := query.Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	delete(db, item)
	w.WriteHeader(http.StatusOK)
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
