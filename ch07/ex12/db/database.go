package db

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type Prices map[string]dollars

type Database struct {
	Prices Prices
	Tmpl   *template.Template
}

func (db Database) render(w http.ResponseWriter, status int) error {
	w.WriteHeader(status)
	if err := db.Tmpl.Execute(w, db); err != nil {
		return err
	}
	return nil
}

func (db Database) List(w http.ResponseWriter, req *http.Request) {
	err := db.render(w, http.StatusOK)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "template error")
	}
}

func (db Database) Price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db.Prices[item]

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

	_, ok := db.Prices[item]
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
	db.Prices[item] = dollars(p)

	if err := db.render(w, http.StatusOK); err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "template error")
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
	_, ok := db.Prices[item]
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
	db.Prices[item] = dollars(p)
	if err := db.render(w, http.StatusOK); err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "template error")
	}
}

func (db Database) Delete(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item := query.Get("item")
	_, ok := db.Prices[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db.Prices, item)
	if err := db.render(w, http.StatusOK); err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "template error")
	}
}
