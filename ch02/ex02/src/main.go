package main

import (
	"fmt"
	"os"
	"strconv"

	"bufio"
	"github.com/yosuke-furukawa/programming-go-study/ch02/ex02/src/distconv"
	"github.com/yosuke-furukawa/programming-go-study/ch02/ex02/src/tempconv"
	"github.com/yosuke-furukawa/programming-go-study/ch02/ex02/src/weightconv"
)

func transform(value float64, converter string) string {
	switch converter {
	case "w":
		fallthrough
	case "weight":
		p := weightconv.Pound(value)
		k := weightconv.KiloGram(value)
		return fmt.Sprintf("%s = %s, %s = %s\n", p, weightconv.PToK(p), k, weightconv.KToP(k))
	case "d":
		fallthrough
	case "distance":
		m := distconv.Meter(value)
		f := distconv.Feet(value)
		return fmt.Sprintf("%s = %s, %s = %s\n", m, distconv.MToF(m), f, distconv.FToM(f))
	case "t":
		fallthrough
	case "temperature":
		f := tempconv.Fahrenheit(value)
		c := tempconv.Celsius(value)
		return fmt.Sprintf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
	return "convert string should be weight, distance, temperature"
}

func main() {
	if len(os.Args[1:]) < 2 {
		stdin := bufio.NewScanner(os.Stdin)

		fmt.Print("input: converter type (weight, distance, temperature) > ")
		var c string
		if stdin.Scan() {
			c = stdin.Text()
		}

		fmt.Print("input: value >")
		var v string
		if stdin.Scan() {
			v = stdin.Text()
		}
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "convert error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(transform(value, c))
	} else {
		c := os.Args[1]
		v := os.Args[2]
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "convert error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(transform(value, c))
	}
}
