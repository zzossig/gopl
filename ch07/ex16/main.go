// Write a web-based calculator program.

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"example.com/ch07/ex16/eval"
)

var mainPage = `
<html>
<body>
	<form action="/calc" method="GET">
		<div>
			<label for="expr">input a expression</label>
			<input name="expr" id="expr" type="text">
		</div>
		<button>Calculate</button>
	</form>
</body>
</html>
`
var calcPage = `
<html>
<body>
	<form action="/calc" method="GET">
		<div>
			<label for="expr">input a expression</label>
			<input name="expr" id="expr" type="text">
		</div>
		<button>Calculate</button>
	</form>
	<div>
		result:&nbsp;{{.}}
	</div>
</body>
</html>
`

var errPage = `
<div>invalid expression: {{.}}</div>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.Must(template.New("main").Parse(mainPage)).Execute(w, nil)
	})
	http.HandleFunc("/calc", func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.URL.String())
		expr := u.Query().Get("expr")
		parsed, err := eval.Parse(expr)
		if err != nil {
			template.Must(template.New("err").Parse(errPage)).Execute(w, expr)
		} else {
			got := fmt.Sprintf("%.6g", parsed.Eval(eval.Env{}))
			template.Must(template.New("calc").Parse(calcPage)).Execute(w, got)
		}
	})
	log.Fatal(http.ListenAndServe(":8001", nil))
}
