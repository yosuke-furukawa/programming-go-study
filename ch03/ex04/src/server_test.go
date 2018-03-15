package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServerAndFetch(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(surfaceHandler))
	defer ts.Close()
	resp, err := fetch(ts.URL, ioutil.Discard)
	if err != nil {
		t.Errorf("error is thrown %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status is not 200 OK")
	}

	if resp.Header.Get("Content-Type") != "image/svg+xml" {
		t.Errorf("content type is not image/svg+xml")
	}
}

func TestServerAndFetchWithQuery(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(surfaceHandler))
	defer ts.Close()
	buffer := &bytes.Buffer{}
	resp, err := fetch(ts.URL+"?f=white&s=black&w=1000&h=1000&c=1000&t=blue&b=green", buffer)
	if err != nil {
		t.Errorf("error is thrown %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status is not 200 OK")
	}

	if resp.Header.Get("Content-Type") != "image/svg+xml" {
		t.Errorf("content type is not image/svg+xml")
	}

	str := buffer.String()
	expects := "style='fill:white'"
	if !strings.Contains(str, expects) {
		t.Errorf("string does not contain expect string %s", expects)
	}
	expects = "style='fill:blue'"
	if !strings.Contains(str, expects) {
		t.Errorf("string does not contain expect string %s", expects)
	}
	expects = "style='fill:green'"
	if !strings.Contains(str, expects) {
		t.Errorf("string does not contain expect string %s", expects)
	}
	expects = "width='1000' height='1000'"
	if !strings.Contains(str, expects) {
		t.Errorf("string does not contain expect string %s", expects)
	}
}
