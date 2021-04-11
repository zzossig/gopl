/*
	Change the handler for `/list` to print its output as an HTML table, not text.
	You may find the `html/template` package (§4.6) useful.
*/

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32
type database map[string]dollars

var (
	table = `
<table>
	<tr>
		<th>name</th>
		<th>price</th>
	</tr>
	{{range .}}
		<tr>
			<td>{{.Name}}</td>
			<td>{{.Price}}</td>
		</tr>
	{{end}}
</table>
`
)

type T struct {
	Name  string
	Price dollars
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) list(w http.ResponseWriter, req *http.Request) {
	var data []T

	for name, price := range db {
		data = append(data, T{name, price})
	}

	t, _ := template.New("foo").Parse(table)
	_ = t.Execute(w, data)

}
