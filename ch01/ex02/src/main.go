package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[0:] {
		fmt.Printf("%d: %s\n", i, arg)
	}
}
