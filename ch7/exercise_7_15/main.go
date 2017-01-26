package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Laugusti/gopl/ch7/exercise_7_14"
)

func main() {
	fmt.Print("Enter expression: ")
	r := bufio.NewReader(os.Stdin)
	line, _, err := r.ReadLine()
	if err != nil {
		log.Fatalf("error reading input: %v", err)
	}
	expr, err := eval.Parse(string(line))
	if err != nil {
		log.Fatalf("error parsing expression: %v", err)
	}
	vars := make(map[eval.Var]bool)
	err = expr.Check(vars)
	if err != nil {
		log.Fatalf("invalid expression: %v", err)
	}
	env := eval.Env{}
	for k := range vars {
		fmt.Printf("Provide value for %q: ", k)
		var v float64
		_, err = fmt.Scanf("%g", &v)
		if err != nil {
			log.Fatalf("error reading input: %v", err)
		}
		env[k] = v
	}

	result := fmt.Sprintf("%.6g", expr.Eval(env))

	fmt.Println()
	fmt.Printf("%s with %v = %q\n", expr, env, result)
}
