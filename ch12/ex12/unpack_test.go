package ex12

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestUnpack(t *testing.T) {
	type Data struct {
		Emails []string `http:"mail" validate:"^[a-zA-Z0-9]+@[a-zA-Z0-9]+.[a-zA-Z0-9]+$"`
		Zip    string   `http:"zip" validate:"^[0-9]{3}-[0-9]{4}$"`
	}
	for _, test := range []struct {
		url  string
		err  string
		data Data
	}{
		{`http://localhost:12345/search?mail=test@example.com&zip=111-2222`, "", Data{[]string{"test@example.com"}, "111-2222"}},
		{`http://localhost:12345/search?mail=testexamplecom&zip=111-2222`, "validator is unmatched mail testexamplecom", Data{[]string{"test@example.com"}, "111-2222"}},
		{`http://localhost:12345/search?mail=test@example.com&zip=1112222`, "validator is unmatched zip 1112222", Data{[]string{"test@example.com"}, "111-2222"}},
	} {
		var data Data

		req, err := newRequest(test.url)
		if err != nil {
			t.Errorf("Parse: %v\n", err)
			continue
		}

		if err := Unpack(req, &data); err != nil {
			if test.err != fmt.Sprintf("%v", err) {
				t.Errorf("Unpack: %v\n", err)
			}
			continue
		}

		if !reflect.DeepEqual(data, test.data) {
			t.Errorf("%q => \n%+v but want %+v\n", test.url, data, test.data)
		}
	}
}

func newRequest(rawurl string) (*http.Request, error) {
	var req http.Request
	url, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	req.URL = url
	return &req, nil
}
