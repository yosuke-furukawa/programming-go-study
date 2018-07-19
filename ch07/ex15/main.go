package main

import (
	"bufio"
	"os"

	"io"

	"fmt"

	"regexp"

	"strconv"

	"github.com/yosuke-furukawa/programming-go-study/ch07/ex15/eval"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	var expression eval.Expr
	variables := eval.Env{}
	// usage: env: 100.0
	re := regexp.MustCompile("(\\w+): (([1-9]\\d*|0)(\\.\\d+)?)")
	for {
		line, _, err := stdin.ReadLine()
		if err == io.EOF {
			break
		}

		if expression == nil {
			expr := string(line)
			exp, err := eval.Parse(expr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
			fmt.Printf("expression: %s\n", exp)
			expression = exp
		}

		fmt.Printf("variable? var: 100.0 \n\n")
		vars, _, err := stdin.ReadLine()
		if err == io.EOF {
			break
		}
		vs := re.FindStringSubmatch(string(vars))
		if len(vs) <= 2 {
			fmt.Fprintln(os.Stderr, "variable: 100.0 is valid")
		}
		v, err := strconv.ParseFloat(vs[2], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		variables[eval.Var(vs[1])] = v
	}
	result := expression.Eval(variables)
	fmt.Printf("result = %f", result)
}
