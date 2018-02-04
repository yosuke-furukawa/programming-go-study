package main

import (
	"io"
	"net/http"
	"strings"
)

func fetch(url string, dst io.Writer) (*http.Response, error) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		return resp, err
	}
	_, err = io.Copy(dst, resp.Body)
	resp.Body.Close()
	return resp, err
}
