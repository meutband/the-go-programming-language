// This files manages a database of items and their costs. The output is written to HTML template
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price) //read
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var itemTable = template.Must(template.New("Items").Parse(`
<h1>Items</h1>
<table>
	<tr>
		<th> Item </th>
		<th> Price </th>
	</tr>
	{{ range $item, $price := . }}
		<tr>
			<td>{{ $item }}</td>
			<td>{{ $price }}</td>
		</tr>
	{{end}}
</table>
`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	itemTable.Execute(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {

	item := req.URL.Query().Get("item")

	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	_, ok := db[item]
	if !ok {

		p, err := strconv.ParseFloat(price, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "item: %q, price: %q\n", item, price)
			return
		}

		db[item] = dollars(p)
		w.WriteHeader(http.StatusCreated)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "item already created: %q\n", item)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	_, ok := db[item]
	if !ok && price != "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item: %q, price: %q\n", item, price)
		return
	}

	db[item] = dollars(p)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "item: %q, price: %q\n", item, price)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {

	item := req.URL.Query().Get("item")

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	delete(db, item)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "item deleted: %q\n", item)
}
