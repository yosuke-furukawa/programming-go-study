package querystring

import (
	"net/url"
	"reflect"
	"strconv"
)

// Decode maps query params to struct using reflection
func Decode(dest interface{}, u url.Values) {
	t := reflect.TypeOf(dest).Elem()
	v := reflect.ValueOf(dest).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("url")
		switch field.Type.Kind() {
		case reflect.String:
			if len(u[tag]) > 0 {
				v.Field(i).SetString(string(u[tag][0]))
			}
		case reflect.Int:
			if len(u[tag]) > 0 {
				val, _ := strconv.ParseInt(u[tag][0], 10, 64)
				v.Field(i).SetInt(val)
			}
		}
	}
}
