package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetch(url string, dst io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	_, err = io.Copy(dst, resp.Body)

	resp.Body.Close()
	return err
}

func main() {
	for _, url := range os.Args[1:] {
		err := fetch(url, os.Stdout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
	}
}
