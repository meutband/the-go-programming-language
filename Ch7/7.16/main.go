// This function takes an expression from HTML, parses and evaluates the expression
package main

import (
	"fmt"
	eval "gobook/Ch7/7.14"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/eval", evaluateHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func renderCalculator(out io.Writer) {
	calculator := template.Must(template.ParseFiles("./index.html"))
	var data struct{}
	if err := calculator.Execute(out, data); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	renderCalculator(w)
}

func evaluateHandler(w http.ResponseWriter, r *http.Request) {

	e := r.URL.Query().Get("e")

	expr, err := eval.Parse(e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error parsing 'e'")
		return
	}

	vars := make(map[eval.Var]bool)
	err = expr.Check(vars)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error checking vars")
		return
	}

	env := make(map[eval.Var]float64)
	result := strconv.FormatFloat(expr.Eval(env), 'f', -1, 64)
	fmt.Fprintf(w, "%s = %v", e, result)
}
