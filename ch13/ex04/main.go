package main

import (
	"io"
	"log"
	"os"

	"github.com/yosuke-furukawa/programming-go-study/ch13/ex03/bzip"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close %v\n", err)
	}

}
