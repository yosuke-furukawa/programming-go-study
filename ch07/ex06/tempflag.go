package main

import (
	"flag"
	"fmt"

	"github.com/yosuke-furukawa/programming-go-study/ch07/ex06/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
