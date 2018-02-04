package main

import (
	"bytes"
	"testing"
  "fmt"
)

func TestFetchAll(t *testing.T) {
	buffer1 := &bytes.Buffer{}
	buffer2 := &bytes.Buffer{}
	fetchAll([]string{"http://gopl.io", "http://www.4chan.org/"}, buffer1)
	fetchAll([]string{"http://gopl.io", "http://www.4chan.org/"}, buffer2)
  fmt.Println(buffer1)
  fmt.Println(buffer2)
}

