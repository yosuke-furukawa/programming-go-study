package main

import (
	"log"
	"os"

	"fmt"

	"github.com/yosuke-furukawa/programming-go-study/ch07/ex18/xmltree"
)

func main() {
	tree, err := xmltree.Build(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tree)

}
