package main

import (
	"bytes"
  "net/http"
  "net/http/httptest"
  "io/ioutil"
	"testing"
)

func TestServerAndFetch(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(lissajousHandler))
  defer ts.Close()
  resp, err := fetch(ts.URL, ioutil.Discard)
  if err != nil {
    t.Errorf("error is thrown %v", err)
  }

  if resp.StatusCode != http.StatusOK {
    t.Errorf("status is not 200 OK")
  }

  if resp.Header.Get("Content-Type") != "image/gif" {
    t.Errorf("content type is not image/gif")
  }
}

func TestServerAndFetchWithCycle10(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(lissajousHandler))
  defer ts.Close()
  resp, err := fetch(ts.URL+"?cycles=10", ioutil.Discard)
  if err != nil {
    t.Errorf("error is thrown %v", err)
  }

  if resp.StatusCode != http.StatusOK {
    t.Errorf("status is not 200 OK")
  }

  if resp.Header.Get("Content-Type") != "image/gif" {
    t.Errorf("content type is not image/gif")
  }
}

func TestErrorLissajous(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(lissajousHandler))
  defer ts.Close()
	buffer := &bytes.Buffer{}
  resp, err := fetch(ts.URL+"?cycles=aaa", buffer)
  if err != nil {
    t.Errorf("error is thrown %v", err)
  }

  if resp.StatusCode != http.StatusBadRequest {
    t.Errorf("status is not 500 actual status: %v", resp.Status)
  }

  if buffer.String() != "Error: do not parse cycles" {
    t.Errorf("message is not parse error %v", buffer.String())
  }
}

