package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	price := r.URL.Query().Get("price")
	p64, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}
	db[item] = dollars(p64)
	fmt.Fprintf(w, "item updated\n%s: %s\n", item, db[item])
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "item deleted\n")
}

func (db database) create(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "missing item\n")
		return
	}
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "item exists: %q\n", item)
		return
	}
	price := r.URL.Query().Get("price")
	p64, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}
	db[item] = dollars(p64)
	fmt.Fprintf(w, "item created\n%s: %s\n", item, db[item])
}
