package db

import (
	"fmt"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type Database map[string]dollars

func (db Database) List(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db Database) Price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}

	fmt.Fprintf(w, "%s\n", price)
}

func (db Database) Update(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item := query.Get("item")
	price := query.Get("price")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "item is empty\n")
		return
	}
	if price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "price is empty\n")
		return
	}

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "cannot parse price %s\n", price)
		return
	}
	db[item] = dollars(p)
	w.WriteHeader(http.StatusCreated)
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db Database) Insert(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item := query.Get("item")
	price := query.Get("price")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "item is empty\n")
		return
	}
	if price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "price is empty\n")
		return
	}
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item is found: %q\n", item)
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "cannot parse price %s\n", price)
		return
	}
	db[item] = dollars(p)
	w.WriteHeader(http.StatusOK)
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db Database) Delete(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item := query.Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
	w.WriteHeader(http.StatusOK)
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
