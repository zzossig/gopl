/*
	Add additional handlers so that clients can create, read, update, and delete database entries.
	For example, a request of the form `/update?item=socks&price=6` will update the price of an item in the inventory and report an error if the item does not exist or if the price is invalid.
	(Warning: this change introduces concurrent variable update.)
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.RawQuery
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	q := req.URL.RawQuery
	kv := strings.Split(q, "=")
	if len(kv) != 2 {
		fmt.Fprintf(w, "wrong syntax: expected - /create?key=val, got - %s", kv)
		return
	}
	f, err := strconv.ParseFloat(kv[1], 32)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	db[kv[0]] = dollars(f)
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	key := req.URL.RawQuery
	if price, ok := db[key]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		fmt.Fprint(w, "no item found")
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	q := req.URL.RawQuery
	kv := strings.Split(q, "=")
	if len(kv) != 2 {
		fmt.Fprintf(w, "wrong syntax: expected - /create?key=val, got - %s", kv)
		return
	}
	f, err := strconv.ParseFloat(kv[1], 32)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	if _, ok := db[kv[0]]; ok {
		db[kv[0]] = dollars(f)
		fmt.Fprintf(w, "%s price updated to %.2f", kv[0], f)
	} else {
		fmt.Fprintf(w, "no item found")
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	key := req.URL.RawQuery
	if _, ok := db[key]; ok {
		delete(db, key)
		fmt.Fprintf(w, "item deleted: %s", key)
	} else {
		fmt.Fprint(w, "no item found")
	}
}
