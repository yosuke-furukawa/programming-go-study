package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	buffer := &bytes.Buffer{}
	err := fetch("http://gopl.io", buffer)
	if err != nil {
		t.Errorf("error is thrown %v", err)
	}
	if !strings.Contains(buffer.String(), "Status: 200 OK") {
		t.Errorf("Status code is not 200")
	}
	if !strings.Contains(buffer.String(), "<title>The Go Programming Language</title>") {
		t.Errorf("title is not matched")
	}
}

func TestErrorFetch(t *testing.T) {
	buffer := &bytes.Buffer{}
	err := fetch("http://bad.gopl.io", buffer)
	if err == nil {
		t.Errorf("error is not thrown")
	}
}

func TestFetchWithoutHttp(t *testing.T) {
	buffer := &bytes.Buffer{}
	err := fetch("gopl.io", buffer)
	if err != nil {
		t.Errorf("error is thrown %v", err)
	}
	if !strings.Contains(buffer.String(), "Status: 200 OK") {
		t.Errorf("Status code is not 200")
	}
	if !strings.Contains(buffer.String(), "<title>The Go Programming Language</title>") {
		t.Errorf("title is not matched")
	}
}
