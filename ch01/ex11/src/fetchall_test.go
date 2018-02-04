package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestFetchAllTooManySites(t *testing.T) {
	buffer := &bytes.Buffer{}
	data, _ := ioutil.ReadFile("../alexasite.txt")
	sites := strings.Split(string(data), "\n")[0:1000]
	fetchAll(sites, buffer)
	fmt.Println(buffer)
}
