package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/yosuke-furukawa/programming-go-study/ch07/ex15/eval"
)

func expressionHandler(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	expr := query.Get("expr")
	log.Println(expr)
	exp, err := eval.Parse(expr)
	if err != nil {
		fmt.Fprintf(w, "parse error: %s", err)
		return
	}
	result := exp.Eval(eval.Env{})
	fmt.Fprintf(w, "%.6g", result)
}

func main() {
	http.HandleFunc("/", expressionHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
