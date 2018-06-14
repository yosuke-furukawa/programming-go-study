package ex18

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_fetch(t *testing.T) {
	local, _, err := fetch("https://golang.org/")
	if err != nil {
		t.Error(err)
	}
	content, err := ioutil.ReadFile(local)
	if err != nil {
		t.Error(err)
	}
	t.Log(local)
	t.Log(string(content))
	if err := os.RemoveAll(local); err != nil {
		t.Error(err)
	}
}
