package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Laugusti/gopl/ch7/exercise_7_14"
)

var calcTempl = template.Must(template.New("calculator.html").ParseFiles("calculator.html"))

type Result struct {
	Res string
	Err error
}

func getResult(s string) Result {
	expr, err := eval.Parse(s)
	if err != nil {
		return Result{Err: fmt.Errorf("error parsing expression: %v", err)}
	}
	vars := make(map[eval.Var]bool)
	err = expr.Check(vars)
	if err != nil {
		return Result{Err: fmt.Errorf("invalid expression: %v", err)}
	}
	var keys []string
	for k := range vars {
		keys = append(keys, string(k))
	}
	if len(keys) != 0 {
		return Result{Err: fmt.Errorf("unexpected variable(s): %v", keys)}
	}
	return Result{Res: fmt.Sprintf("%.6g", expr.Eval(eval.Env{}))}
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		message := fmt.Sprintf("error parsing form: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		return
	}
	s := r.URL.Query().Get("expr")
	var res Result
	if s != "" {
		res = getResult(s)
	}

	err = calcTempl.Execute(w, res)
	if err != nil {
		log.Printf("template execution failed: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8000", nil))
}
