package main

import (
	"bytes"
  "fmt"
  "net"
  "os"
	"testing"
)

func TestServerAndFetch(t *testing.T) {
	buffer := &bytes.Buffer{}
	ch := make(chan net.Listener)
	er := make(chan error)
  go server("localhost:0", ch, er)
	for {
		select {
		case listener := <-ch:
      addr := listener.Addr().String()
      fmt.Println("listen on ", addr)
      resp, err := fetch(addr, buffer)
      if err != nil {
        t.Errorf("error is thrown %v", err)
      }

      if resp.Status != "200 OK" {
        t.Errorf("status is not 200 OK")
      }

      if resp.Header["Content-Type"][0] != "image/gif" {
        t.Errorf("content type is not image/gif")
      }
      os.Exit(0)

    case err := <-er:
      t.Fatal("error ", err)
		}
	}
}

