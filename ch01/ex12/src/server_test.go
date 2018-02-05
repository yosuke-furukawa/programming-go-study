package main

import (
	"bytes"
  "net/http"
  "net/http/httptest"
	"testing"
)

func TestServerAndFetch(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(lissajousHandler))
  defer ts.Close()
	buffer := &bytes.Buffer{}
  resp, err := fetch(ts.URL, buffer)
  if err != nil {
    t.Errorf("error is thrown %v", err)
  }

  if resp.Status != "200 OK" {
    t.Errorf("status is not 200 OK")
  }

  if resp.Header["Content-Type"][0] != "image/gif" {
    t.Errorf("content type is not image/gif")
  }
}

func TestServerAndFetchWithCycle10(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(lissajousHandler))
  defer ts.Close()
	buffer := &bytes.Buffer{}
  resp, err := fetch(ts.URL+"?cycles=10", buffer)
  if err != nil {
    t.Errorf("error is thrown %v", err)
  }

  if resp.Status != "200 OK" {
    t.Errorf("status is not 200 OK")
  }

  if resp.Header["Content-Type"][0] != "image/gif" {
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

  if resp.Status != "500 Internal Server Error" {
    t.Errorf("status is not 500 actual status: %v", resp.Status)
  }
}

